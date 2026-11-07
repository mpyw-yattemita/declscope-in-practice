package gateway

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/provider"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (p *Provider) loadHTTPRoute(ctx context.Context, gateways []gatewayWithListeners, route *gatev1.HTTPRoute, conf *dynamic.Configuration, attachedRoutes attachedRoutes, servedRules servedRules, statusReport *statusReport) {
	logger := log.Ctx(ctx).With().
		Str("http_route", route.Name).
		Str("namespace", route.Namespace).
		Logger()

	routeParentRefs := matchingGatewayListenersForParentRef(gateways, route.Namespace, route.Spec.ParentRefs)
	if len(routeParentRefs) == 0 {
		return
	}

	for _, match := range routeParentRefs {
		acceptedCondition := metav1.Condition{
			Type:               string(gatev1.RouteConditionAccepted),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonNoMatchingParent),
		}

		var resolvedRefCondition *metav1.Condition
		for _, listener := range match.Listeners {
			// A parentRef can target specific listeners through its SectionName or Port.
			accepted := matchListener(listener, match.ParentRef)

			if accepted && !allowRoute(listener, route.Namespace, kindHTTPRoute) {
				if acceptedCondition.Status == metav1.ConditionFalse {
					acceptedCondition.Reason = string(gatev1.RouteReasonNotAllowedByListeners)
				}
				accepted = false
			}

			hostnames, ok := findMatchingHostnames(listener.Hostname, route.Spec.Hostnames)
			if accepted && !ok {
				if acceptedCondition.Status == metav1.ConditionFalse {
					acceptedCondition.Reason = string(gatev1.RouteReasonNoMatchingListenerHostname)
				}
				accepted = false
			}

			ref := listenerRef{
				GatewayNamespace: match.GatewayNamespace,
				GatewayName:      match.GatewayName,
				SectionName:      listener.SectionName,
			}
			if accepted && attachedRoutes.Conflicts(ref, kindHTTPRoute, hostnames) {
				if acceptedCondition.Status == metav1.ConditionFalse {
					acceptedCondition.Reason = string(routeReasonHostnameConflict)
				}
				accepted = false
			}

			if accepted {
				// Gateway listener should have AttachedRoutes set even when Gateway has unresolved refs.
				listener.Status.AttachedRoutes++

				attachedRoutes.Record(ref, kindHTTPRoute, hostnames)
			}

			// The ResolvedRefs condition must be reported for every parentRef,
			// even when the route does not attach to the listener.
			loader := httpRouteLoader{
				p:                p,
				gatewayName:      match.GatewayName,
				gatewayNamespace: match.GatewayNamespace,
				listener:         listener,
				route:            route,
				statusReport:     statusReport,
			}
			routerConfs, condition := loader.loadHTTPRouteConfiguration(logger.WithContext(ctx), hostnames)
			if resolvedRefCondition == nil || resolvedRefCondition.Status == metav1.ConditionTrue {
				resolvedRefCondition = new(condition)
			}

			if accepted && listener.Attached {
				for _, rc := range routerConfs {
					router := rc.Conf.HTTP.Routers[rc.Name]
					if servedBy, alreadyServed := servedRules.register(rc.Name, router.ParentRefs, router.Rule); alreadyServed {
						logger.Warn().Msgf("Traefik does not create router %q, because router %q serves the rule %q under the same parent routers", rc.Name, servedBy, router.Rule)
						continue
					}

					mergeHTTPConfiguration(rc.Conf, conf)
				}

				// Only consider the route attached if the listener is in an "attached" state.
				acceptedCondition.Reason = string(gatev1.RouteReasonAccepted)
				acceptedCondition.Status = metav1.ConditionTrue
			}
		}

		parentStatusConditions := []metav1.Condition{acceptedCondition}
		if resolvedRefCondition != nil {
			parentStatusConditions = append(parentStatusConditions, *resolvedRefCondition)
		}

		statusReport.RecordHTTPRouteStatus(ktypes.NamespacedName{Namespace: route.Namespace, Name: route.Name}, gatev1.RouteParentStatus{
			ParentRef:      match.ParentRef,
			ControllerName: controllerName,
			Conditions:     parentStatusConditions,
		})
	}
}

// httpRouteRouter is a router of an HTTPRoute, with the configuration of its resources.
type httpRouteRouter struct {
	Name string
	Conf *dynamic.Configuration
}

// httpRouteLoader loads the configuration of an HTTPRoute attached to a Gateway listener.
// It holds what every backend and filter of the route needs to be loaded,
// as a RequestMirror filter loads a backend like a rule does.
type httpRouteLoader struct {
	p                *Provider
	gatewayName      string
	gatewayNamespace string
	listener         gatewayListener
	route            *gatev1.HTTPRoute
	statusReport     *statusReport
}

func (l *httpRouteLoader) loadHTTPRouteConfiguration(ctx context.Context, hostnames []gatev1.Hostname) ([]httpRouteRouter, metav1.Condition) {
	route := l.route

	var routers []httpRouteRouter

	condition := metav1.Condition{
		Type:               string(gatev1.RouteConditionResolvedRefs),
		Status:             metav1.ConditionTrue,
		ObservedGeneration: route.Generation,
		LastTransitionTime: metav1.Now(),
		Reason:             string(gatev1.RouteConditionResolvedRefs),
	}

	for ri, routeRule := range route.Spec.Rules {
		respondingTimeouts := routerRespondingTimeouts(ctx, routeRule.Timeouts)
		backendRequestTimeout := parseBackendRequestTimeout(ctx, routeRule.Timeouts)
		retry := l.p.buildRetry(ctx, routeRule.Retry)

		for _, match := range routeRule.Matches {
			rule, priority := buildMatchRule(hostnames, match)
			router := dynamic.Router{
				// "default" stands for the default rule syntax in Traefik v3, i.e. the v3 syntax.
				RuleSyntax:         "default",
				Rule:               rule,
				Priority:           priority + len(route.Spec.Rules) - ri,
				ParentRefs:         l.listener.RouterNames,
				RespondingTimeouts: respondingTimeouts,
			}

			routerName := makeRouterName(strings.ToLower(kindHTTPRoute), rule, route.Namespace, route.Name, l.gatewayNamespace, l.gatewayName, l.listener.EPName, ri)

			routerConf := &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           make(map[string]*dynamic.Router),
					Middlewares:       make(map[string]*dynamic.Middleware),
					Services:          make(map[string]*dynamic.Service),
					ServersTransports: make(map[string]*dynamic.ServersTransport),
				},
			}

			// TODO loadMiddlewares errors could change the condition.
			middlewares, err := l.loadMiddlewares(ctx, routerConf, routerName, routeRule.Filters, match.Path)
			if err != nil {
				log.Ctx(ctx).Error().Err(err).Msg("Unable to load HTTPRoute filters")

				errWrrName := routerName + "-err-wrr"
				routerConf.HTTP.Services[errWrrName] = &dynamic.Service{
					Weighted: &dynamic.WeightedRoundRobin{
						Services: []dynamic.WRRService{
							{
								Name:   "invalid-httproute-filter",
								Status: new(500),
								Weight: new(1),
							},
						},
					},
				}
				router.Service = errWrrName
			} else {
				router.Middlewares = middlewares

				// The retry middleware comes after the filters, so that only the call to the backend is retried.
				if retry != nil {
					retryName := routerName + "-retry"
					routerConf.HTTP.Middlewares[retryName] = &dynamic.Middleware{Retry: retry}
					router.Middlewares = append(router.Middlewares, retryName)
				}

				var serviceCondition *metav1.Condition
				router.Service, serviceCondition = l.loadHTTPRouteRuleService(ctx, routerConf, routerName, routeRule, backendRequestTimeout, match.Path)
				if serviceCondition != nil {
					condition = *serviceCondition
				}
			}

			l.p.applyRouterTransform(ctx, &router, route)

			routerConf.HTTP.Routers[routerName] = &router
			routers = append(routers, httpRouteRouter{Name: routerName, Conf: routerConf})
		}
	}

	return routers, condition
}

// routerRespondingTimeouts returns the router responding timeouts for the request timeout of an HTTPRoute rule,
// or nil when it has none.
func routerRespondingTimeouts(ctx context.Context, timeouts *gatev1.HTTPRouteTimeouts) *dynamic.RouterRespondingTimeouts {
	if timeouts == nil || timeouts.Request == nil {
		return nil
	}

	// The CRD CEL validation guarantees a valid GEP-2257 duration string.
	// Both timeouts count from the same instant, so they share a single deadline bounding the whole request.
	// A zero duration disables the timeout, lifting the entrypoint read and write deadlines, as GEP-1742 requires.
	d, err := time.ParseDuration(string(*timeouts.Request))
	if err != nil {
		// TODO: Maybe it would be better to handle the error and put the routers in error.
		log.Ctx(ctx).Error().Err(err).Msg("Ignoring HTTPRoute timeouts.request: invalid duration")
		return nil
	}

	return &dynamic.RouterRespondingTimeouts{
		ReadTimeout:  new(ptypes.Duration(d)),
		WriteTimeout: new(ptypes.Duration(d)),
	}
}

// parseBackendRequestTimeout returns the backend request timeout of an HTTPRoute rule, or zero when it has none.
func parseBackendRequestTimeout(ctx context.Context, timeouts *gatev1.HTTPRouteTimeouts) time.Duration {
	if timeouts == nil || timeouts.BackendRequest == nil {
		return 0
	}

	// The CRD CEL validation guarantees a valid GEP-2257 duration string.
	d, err := time.ParseDuration(string(*timeouts.BackendRequest))
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Ignoring HTTPRoute timeouts.backendRequest: invalid duration")
		return 0
	}

	return d
}

// buildRetry returns the retry middleware configuration for the retry settings of an HTTPRoute rule,
// or nil when it has none.
func (p *Provider) buildRetry(ctx context.Context, retry *gatev1.HTTPRouteRetry) *dynamic.Retry {
	if retry == nil {
		return nil
	}

	// Traefik counts the initial request as an attempt, whereas Gateway API only counts the retries.
	// The CRD defaults attempts to 1.
	conf := &dynamic.Retry{
		Attempts:            ptr.Deref(retry.Attempts, 1) + 1,
		MaxRequestBodyBytes: new(p.RetryMaxRequestBodyBytes),
	}

	for _, code := range retry.Codes {
		conf.Status = append(conf.Status, strconv.Itoa(int(code)))
	}

	if retry.Backoff != nil {
		// The CRD CEL validation guarantees a valid GEP-2257 duration string.
		d, err := time.ParseDuration(string(*retry.Backoff))
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("Ignoring HTTPRoute retry.backoff: invalid duration")
		} else {
			conf.InitialInterval = ptypes.Duration(d)
		}
	}

	return conf
}

// loadHTTPRouteRuleService returns the name of the service the routers of the rule forward to,
// and a non-nil condition when a backend reference cannot be resolved.
func (l *httpRouteLoader) loadHTTPRouteRuleService(ctx context.Context, conf *dynamic.Configuration, routerName string, routeRule gatev1.HTTPRouteRule, backendRequestTimeout time.Duration, pathMatch *gatev1.HTTPPathMatch) (string, *metav1.Condition) {
	if len(routeRule.BackendRefs) != 1 || !isInternalService(routeRule.BackendRefs[0].BackendRef) {
		return l.loadWRRService(ctx, conf, routerName, routeRule, backendRequestTimeout, pathMatch)
	}

	route := l.route
	if !isCrossProviderNamespaceAllowed(l.p.CrossProviderNamespaces, route.Namespace) {
		return "", &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot load HTTPRoute BackendRef %s: internal service reference is not allowed: HTTPRoute namespace %q is not in crossProviderNamespaces", routeRule.BackendRefs[0].Name, route.Namespace),
		}
	}

	return string(routeRule.BackendRefs[0].Name), nil
}

func (l *httpRouteLoader) loadWRRService(ctx context.Context, conf *dynamic.Configuration, routerName string, routeRule gatev1.HTTPRouteRule, backendRequestTimeout time.Duration, pathMatch *gatev1.HTTPPathMatch) (string, *metav1.Condition) {
	name := routerName + "-wrr"
	if _, ok := conf.HTTP.Services[name]; ok {
		return name, nil
	}

	// A rule with omitted or empty backendRefs, and without any filter that
	// could make the request receive a response, has no backend to forward to,
	// and must explicitly respond with a 500 status code.
	if len(routeRule.BackendRefs) == 0 && len(routeRule.Filters) == 0 {
		conf.HTTP.Services[name] = &dynamic.Service{
			Weighted: &dynamic.WeightedRoundRobin{
				Services: []dynamic.WRRService{
					{
						Name:   "no-backend-refs",
						Status: new(500),
						Weight: new(1),
					},
				},
			},
		}

		return name, nil
	}

	var wrr dynamic.WeightedRoundRobin
	var condition *metav1.Condition
	for bi, backendRef := range routeRule.BackendRefs {
		// TODO in loadService we need to always return a non-nil serviceName even when there is an error which is not the
		// usual defacto.
		svcName, errCondition := l.loadService(ctx, conf, routerName, bi, backendRef, backendRequestTimeout, pathMatch)
		weight := new(int(ptr.Deref(backendRef.Weight, 1)))
		if errCondition != nil {
			log.Ctx(ctx).Error().
				Msgf("Unable to load HTTPRoute backend: %s", errCondition.Message)

			condition = errCondition
			wrr.Services = append(wrr.Services, dynamic.WRRService{
				Name:   svcName,
				Status: new(500),
				Weight: weight,
			})
			continue
		}

		wrr.Services = append(wrr.Services, dynamic.WRRService{
			Name:   svcName,
			Weight: weight,
		})
	}

	conf.HTTP.Services[name] = &dynamic.Service{Weighted: &wrr}
	return name, condition
}

// loadService returns a dynamic.Service config corresponding to the given gatev1.HTTPBackendRef.
// Note that the returned dynamic.Service config can be nil (for cross-provider, internal services, and backendFunc).
// A non-zero backendRequestTimeout bounds the wait for the response headers of a Kubernetes Service backend.
func (l *httpRouteLoader) loadService(ctx context.Context, conf *dynamic.Configuration, routerName string, backendIndex int, backendRef gatev1.HTTPBackendRef, backendRequestTimeout time.Duration, pathMatch *gatev1.HTTPPathMatch) (string, *metav1.Condition) {
	route := l.route
	kind := ptr.Deref(backendRef.Kind, kindService)

	group := groupCore
	if backendRef.Group != nil && *backendRef.Group != "" {
		group = string(*backendRef.Group)
	}

	namespace := route.Namespace
	if backendRef.Namespace != nil && *backendRef.Namespace != "" {
		namespace = string(*backendRef.Namespace)

		if strings.Contains(string(backendRef.Name), "@") {
			return provider.Normalize(fmt.Sprintf("%s-svc-%s-%s-%d", routerName, namespace, backendRef.Name, backendIndex)), &metav1.Condition{
				Type:               string(gatev1.RouteConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: route.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.RouteReasonRefNotPermitted),
				Message:            fmt.Sprintf("Cannot load HTTPBackendRef %s/%s/%s/%s: namespace is not allowed with a cross-provider reference", group, kind, namespace, backendRef.Name),
			}
		}
	}

	serviceName := fmt.Sprintf("%s-svc-%s-%s-%d", routerName, namespace, backendRef.Name, backendIndex)

	if err := l.p.isReferenceGranted(kindHTTPRoute, route.Namespace, group, string(kind), string(backendRef.Name), namespace); err != nil {
		return serviceName, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot load HTTPBackendRef %s/%s/%s/%s: %s", group, kind, namespace, backendRef.Name, err),
		}
	}

	middlewares, err := l.loadMiddlewares(ctx, conf, serviceName, backendRef.Filters, pathMatch)
	if err != nil {
		return serviceName, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonInvalidKind),
			Message:            fmt.Sprintf("Cannot load filters on HTTPBackendRef %s/%s/%s/%s: %s", group, kind, namespace, backendRef.Name, err),
		}
	}

	// TODO may be we could incorporate this "ignored" case into the loadHTTPBackendRef.
	if group != groupCore || kind != kindService {
		name, service, err := l.p.loadHTTPBackendRef(namespace, backendRef)
		if err != nil {
			return serviceName, &metav1.Condition{
				Type:               string(gatev1.RouteConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: route.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.RouteReasonInvalidKind),
				Message:            fmt.Sprintf("Cannot load HTTPBackendRef %s/%s/%s/%s: %s", group, kind, namespace, backendRef.Name, err),
			}
		}

		if service != nil {
			service.Middlewares = middlewares
			conf.HTTP.Services[name] = service
		}

		return name, nil
	}

	port := ptr.Deref(backendRef.Port, gatev1.PortNumber(0))
	if port == 0 {
		return serviceName, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonUnsupportedProtocol),
			Message:            fmt.Sprintf("Cannot load HTTPBackendRef %s/%s/%s/%s: port is required", group, kind, namespace, backendRef.Name),
		}
	}

	lb, st, errCondition := l.loadHTTPServers(namespace, backendRef)
	if errCondition != nil {
		return serviceName, errCondition
	}

	// GEP-1742 backendRequest bounds a single backend round trip. Waiting for the response headers is the part of
	// it a ServersTransport can bound without also cutting long-lived response bodies (e.g. streaming).
	// A zero duration means no timeout, which needs no dedicated ServersTransport.
	if backendRequestTimeout > 0 {
		if st == nil {
			st = &dynamic.ServersTransport{}
		}

		st.ForwardingTimeouts = &dynamic.ForwardingTimeouts{}
		st.ForwardingTimeouts.SetDefaults()
		st.ForwardingTimeouts.ResponseHeaderTimeout = ptypes.Duration(backendRequestTimeout)
	}

	if st != nil {
		lb.ServersTransport = serviceName
		conf.HTTP.ServersTransports[serviceName] = st
	}

	conf.HTTP.Services[serviceName] = &dynamic.Service{LoadBalancer: lb, Middlewares: middlewares}

	return serviceName, nil
}

func (p *Provider) loadHTTPBackendRef(namespace string, backendRef gatev1.HTTPBackendRef) (string, *dynamic.Service, error) {
	// Support for cross-provider references (e.g: api@internal).
	// This provides the same behavior as for IngressRoutes.
	if *backendRef.Kind == "TraefikService" && strings.Contains(string(backendRef.Name), "@") {
		if !isCrossProviderNamespaceAllowed(p.CrossProviderNamespaces, namespace) {
			return "", nil, fmt.Errorf("TraefikService %q reference is not allowed: namespace %q is not in crossProviderNamespaces", string(backendRef.Name), namespace)
		}

		return string(backendRef.Name), nil, nil
	}

	backendFunc, ok := p.groupKindBackendFuncs[string(*backendRef.Group)][string(*backendRef.Kind)]
	if !ok {
		return "", nil, fmt.Errorf("unsupported HTTPBackendRef %s/%s/%s", *backendRef.Group, *backendRef.Kind, backendRef.Name)
	}
	if backendFunc == nil {
		return "", nil, fmt.Errorf("undefined backendFunc for HTTPBackendRef %s/%s/%s", *backendRef.Group, *backendRef.Kind, backendRef.Name)
	}

	return backendFunc(string(backendRef.Name), namespace)
}

func (l *httpRouteLoader) loadMiddlewares(ctx context.Context, conf *dynamic.Configuration, parentName string, filters []gatev1.HTTPRouteFilter, pathMatch *gatev1.HTTPPathMatch) ([]string, error) {
	type namedMiddleware struct {
		Name   string
		Config *dynamic.Middleware
	}

	pm := ptr.Deref(pathMatch, gatev1.HTTPPathMatch{
		Type:  new(gatev1.PathMatchPathPrefix),
		Value: new("/"),
	})

	var middlewares []namedMiddleware
	for i, filter := range filters {
		name := fmt.Sprintf("%s-%s-%d", parentName, strings.ToLower(string(filter.Type)), i)

		switch filter.Type {
		case gatev1.HTTPRouteFilterRequestRedirect:
			middlewares = append(middlewares, namedMiddleware{
				name,
				createRequestRedirect(filter.RequestRedirect, pm),
			})

		case gatev1.HTTPRouteFilterRequestHeaderModifier:
			middlewares = append(middlewares, namedMiddleware{
				name,
				createRequestHeaderModifier(filter.RequestHeaderModifier),
			})

		case gatev1.HTTPRouteFilterResponseHeaderModifier:
			middlewares = append(middlewares, namedMiddleware{
				name,
				createResponseHeaderModifier(filter.ResponseHeaderModifier),
			})

		case gatev1.HTTPRouteFilterExtensionRef:
			name, middleware, err := l.p.loadHTTPRouteFilterExtensionRef(l.route.Namespace, filter.ExtensionRef)
			if err != nil {
				return nil, fmt.Errorf("loading ExtensionRef filter %s: %w", filter.Type, err)
			}
			middlewares = append(middlewares, namedMiddleware{
				name,
				middleware,
			})

		case gatev1.HTTPRouteFilterURLRewrite:
			middleware, err := createURLRewrite(filter.URLRewrite, pm)
			if err != nil {
				return nil, fmt.Errorf("invalid filter %s: %w", filter.Type, err)
			}
			middlewares = append(middlewares, namedMiddleware{
				name,
				middleware,
			})

		case gatev1.HTTPRouteFilterCORS:
			middlewares = append(middlewares, namedMiddleware{
				name,
				createCORS(filter.CORS),
			})

		case gatev1.HTTPRouteFilterRequestMirror:
			middleware, err := l.loadRequestMirror(ctx, conf, name, filter.RequestMirror)
			if err != nil {
				// As per the spec, an invalid mirror must be dropped without failing the request.
				log.Ctx(ctx).Debug().Err(err).Msgf("Dropping HTTPRoute RequestMirror filter %s", name)
				continue
			}
			middlewares = append(middlewares, namedMiddleware{
				name,
				middleware,
			})

		default:
			// As per the spec: https://gateway-api.sigs.k8s.io/api-types/httproute/#filters-optional
			// In all cases where incompatible or unsupported filters are
			// specified, implementations MUST add a warning condition to
			// status.
			return nil, fmt.Errorf("unsupported filter %s", filter.Type)
		}
	}

	var middlewareNames []string
	for _, m := range middlewares {
		if m.Config != nil {
			conf.HTTP.Middlewares[m.Name] = m.Config
		}
		middlewareNames = append(middlewareNames, m.Name)
	}

	return middlewareNames, nil
}

func (p *Provider) loadHTTPRouteFilterExtensionRef(namespace string, extensionRef *gatev1.LocalObjectReference) (string, *dynamic.Middleware, error) {
	if extensionRef == nil {
		return "", nil, errors.New("filter extension ref undefined")
	}

	filterFunc, ok := p.groupKindFilterFuncs[string(extensionRef.Group)][string(extensionRef.Kind)]
	if !ok {
		return "", nil, fmt.Errorf("unsupported filter extension ref %s/%s/%s", extensionRef.Group, extensionRef.Kind, extensionRef.Name)
	}
	if filterFunc == nil {
		return "", nil, fmt.Errorf("undefined filterFunc for filter extension ref %s/%s/%s", extensionRef.Group, extensionRef.Kind, extensionRef.Name)
	}

	return filterFunc(string(extensionRef.Name), namespace)
}

func (l *httpRouteLoader) loadHTTPServers(namespace string, backendRef gatev1.HTTPBackendRef) (*dynamic.ServersLoadBalancer, *dynamic.ServersTransport, *metav1.Condition) {
	route := l.route
	backendAddresses, svcPort, err := l.p.getBackendAddresses(namespace, backendRef.BackendRef)
	if err != nil {
		return nil, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonBackendNotFound),
			Message:            fmt.Sprintf("Cannot load HTTPBackendRef %s/%s: %s", namespace, backendRef.Name, err),
		}
	}

	backendTLS, errCondition := l.p.loadBackendTLS(l.gatewayName, l.gatewayNamespace, l.listener, route.Generation, namespace, backendRef.Name, svcPort, l.statusReport)
	if errCondition != nil {
		return nil, nil, errCondition
	}
	serversTransport := newServersTransport(backendTLS)

	lb := &dynamic.ServersLoadBalancer{}
	lb.SetDefaults()

	// If a ServersTransport is set, it means a BackendTLSPolicy matched the service port, and we can safely assume the protocol is HTTPS.
	// When no ServersTransport is set, we need to determine the protocol based on the service port.
	protocol := "https"
	if serversTransport == nil {
		protocol, err = getHTTPServiceProtocol(svcPort)
		if err != nil {
			return nil, nil, &metav1.Condition{
				Type:               string(gatev1.RouteConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: route.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.RouteReasonUnsupportedProtocol),
				Message:            fmt.Sprintf("Cannot load HTTPBackendRef %s/%s: %s", namespace, backendRef.Name, err),
			}
		}
	}

	for _, ba := range backendAddresses {
		lb.Servers = append(lb.Servers, dynamic.Server{
			URL: fmt.Sprintf("%s://%s", protocol, net.JoinHostPort(ba.IP, strconv.Itoa(int(ba.Port)))),
		})
	}
	return lb, serversTransport, nil
}

// loadRequestMirror loads the backend of a RequestMirror filter as a rule backend, and returns the middleware mirroring to it.
// The loaded service is named after the middleware, so that each mirror of a rule has its own.
func (l *httpRouteLoader) loadRequestMirror(ctx context.Context, conf *dynamic.Configuration, name string, filter *gatev1.HTTPRequestMirrorFilter) (*dynamic.Middleware, error) {
	if filter == nil {
		return nil, errors.New("filter request mirror undefined")
	}

	backendRef := gatev1.HTTPBackendRef{
		BackendRef: gatev1.BackendRef{BackendObjectReference: filter.BackendRef},
	}

	// The rule backendRequest timeout bounds the round trip whose response is sent to the client,
	// it does not apply to the mirrored requests, whose responses are ignored.
	serviceName, errCondition := l.loadService(ctx, conf, name, 0, backendRef, 0, nil)
	if errCondition != nil {
		return nil, errors.New(errCondition.Message)
	}

	numerator, denominator := requestMirrorFraction(filter)

	return &dynamic.Middleware{
		RequestMirror: &dynamic.RequestMirror{
			Service:     serviceName,
			Numerator:   numerator,
			Denominator: denominator,
		},
	}, nil
}

// requestMirrorFraction returns the fraction of requests to mirror, out of range values being clamped,
// as the CRD validation is not enforced on every cluster.
func requestMirrorFraction(filter *gatev1.HTTPRequestMirrorFilter) (int64, int64) {
	switch {
	case filter.Percent != nil:
		return int64(min(max(*filter.Percent, 0), 100)), 100

	case filter.Fraction != nil:
		denominator := int64(ptr.Deref(filter.Fraction.Denominator, 100))
		if denominator <= 0 {
			return 0, 1
		}
		return min(max(int64(filter.Fraction.Numerator), 0), denominator), denominator

	default:
		return 1, 1
	}
}

func buildHostRule(hostnames []gatev1.Hostname) (string, int) {
	var rules []string
	var priority int

	for _, hostname := range hostnames {
		host := string(hostname)

		if priority < len(host) {
			priority = len(host)
		}

		rules = append(rules, fmt.Sprintf("Host(%q)", hostnameMatcherValue(host)))
	}

	switch len(rules) {
	case 0:
		return "", 0
	case 1:
		return rules[0], priority
	default:
		return fmt.Sprintf("(%s)", strings.Join(rules, " || ")), priority
	}
}

// buildMatchRule builds the route rule and computes its priority.
// The current priority computing is rather naive but aims to fulfill Conformance tests suite requirement.
// The priority is computed to match the following precedence order:
//
// * "Exact" path match (+100000).
// * "Prefix" path match with largest number of characters (+10000 + nb_characters*100).
// * Method match (+1000).
// * Largest number of header matches (+100 each).
// * Largest number of query param matches (+10 each).
//
// In case of multiple matches for a route, the maximum priority among all matches is retain.
func buildMatchRule(hostnames []gatev1.Hostname, match gatev1.HTTPRouteMatch) (string, int) {
	path := ptr.Deref(match.Path, gatev1.HTTPPathMatch{
		Type:  new(gatev1.PathMatchPathPrefix),
		Value: new("/"),
	})

	var priority int
	var matchRules []string

	pathRule, pathPriority := buildPathRule(path)
	matchRules = append(matchRules, pathRule)
	priority += pathPriority

	if match.Method != nil {
		matchRules = append(matchRules, fmt.Sprintf("Method(%q)", *match.Method))
		priority += 1000
	}

	headerRules, headersPriority := buildHeaderRules(match.Headers)
	matchRules = append(matchRules, headerRules...)
	priority += headersPriority

	queryParamRules, queryParamsPriority := buildQueryParamRules(match.QueryParams)
	matchRules = append(matchRules, queryParamRules...)
	priority += queryParamsPriority

	matchRulesStr := strings.Join(matchRules, " && ")

	hostRule, hostPriority := buildHostRule(hostnames)

	if hostRule == "" {
		return matchRulesStr, priority
	}

	// A route with a host should match over the same route with no host.
	priority += hostPriority
	return hostRule + " && " + matchRulesStr, priority
}

func buildPathRule(pathMatch gatev1.HTTPPathMatch) (string, int) {
	pathType := ptr.Deref(pathMatch.Type, gatev1.PathMatchPathPrefix)
	pathValue := ptr.Deref(pathMatch.Value, "/")

	switch pathType {
	case gatev1.PathMatchExact:
		return fmt.Sprintf("Path(%q)", pathValue), 100000

	case gatev1.PathMatchPathPrefix:
		// PathPrefix(`/`) rule is a catch-all,
		// here we ensure it would be evaluated last.
		if pathValue == "/" {
			return `PathPrefix("/")`, 1
		}

		pv := strings.TrimSuffix(pathValue, "/")
		return fmt.Sprintf("(Path(%q) || PathPrefix(%q))", pv, fmt.Sprintf("%s/", pv)), 10000 + len(pathValue)*100

	case gatev1.PathMatchRegularExpression:
		return fmt.Sprintf("PathRegexp(%q)", pathValue), 10000 + len(pathValue)*100

	default:
		return `PathPrefix("/")`, 1
	}
}

func buildHeaderRules(headers []gatev1.HTTPHeaderMatch) ([]string, int) {
	var (
		rules    []string
		priority int
	)
	for _, header := range headers {
		typ := ptr.Deref(header.Type, gatev1.HeaderMatchExact)
		switch typ {
		case gatev1.HeaderMatchExact:
			rules = append(rules, fmt.Sprintf("Header(%q,%q)", header.Name, header.Value))
		case gatev1.HeaderMatchRegularExpression:
			rules = append(rules, fmt.Sprintf("HeaderRegexp(%q,%q)", header.Name, header.Value))
		}
		priority += 100
	}

	return rules, priority
}

func buildQueryParamRules(queryParams []gatev1.HTTPQueryParamMatch) ([]string, int) {
	var (
		rules    []string
		priority int
	)
	for _, qp := range queryParams {
		typ := ptr.Deref(qp.Type, gatev1.QueryParamMatchExact)
		switch typ {
		case gatev1.QueryParamMatchExact:
			rules = append(rules, fmt.Sprintf("Query(%q,%q)", qp.Name, qp.Value))
		case gatev1.QueryParamMatchRegularExpression:
			rules = append(rules, fmt.Sprintf("QueryRegexp(%q,%q)", qp.Name, qp.Value))
		}
		priority += 10
	}

	return rules, priority
}

// createRequestHeaderModifier does not enforce/check the configuration,
// as the spec indicates that either the webhook or CEL (since v1.0 GA Release) should enforce that.
func createRequestHeaderModifier(filter *gatev1.HTTPHeaderFilter) *dynamic.Middleware {
	sets := map[string]string{}
	for _, header := range filter.Set {
		sets[string(header.Name)] = header.Value
	}

	adds := map[string]string{}
	for _, header := range filter.Add {
		adds[string(header.Name)] = header.Value
	}

	return &dynamic.Middleware{
		RequestHeaderModifier: &dynamic.HeaderModifier{
			Set:    sets,
			Add:    adds,
			Remove: filter.Remove,
		},
	}
}

// createResponseHeaderModifier does not enforce/check the configuration,
// as the spec indicates that either the webhook or CEL (since v1.0 GA Release) should enforce that.
func createResponseHeaderModifier(filter *gatev1.HTTPHeaderFilter) *dynamic.Middleware {
	sets := map[string]string{}
	for _, header := range filter.Set {
		sets[string(header.Name)] = header.Value
	}

	adds := map[string]string{}
	for _, header := range filter.Add {
		adds[string(header.Name)] = header.Value
	}

	return &dynamic.Middleware{
		ResponseHeaderModifier: &dynamic.HeaderModifier{
			Set:    sets,
			Add:    adds,
			Remove: filter.Remove,
		},
	}
}

func createRequestRedirect(filter *gatev1.HTTPRequestRedirectFilter, pathMatch gatev1.HTTPPathMatch) *dynamic.Middleware {
	var hostname *string
	if filter.Hostname != nil {
		hostname = new(string(*filter.Hostname))
	}

	var port *string
	filterScheme := ptr.Deref(filter.Scheme, "")
	if filterScheme == schemeHTTP || filterScheme == schemeHTTPS {
		port = new("")
	}
	if filter.Port != nil {
		port = new(strconv.Itoa(int(*filter.Port)))
	}

	var path *string
	var pathPrefix *string
	if filter.Path != nil {
		switch filter.Path.Type {
		case gatev1.FullPathHTTPPathModifier:
			path = filter.Path.ReplaceFullPath
		case gatev1.PrefixMatchHTTPPathModifier:
			path = filter.Path.ReplacePrefixMatch
			pathPrefix = pathMatch.Value
		}
	}

	return &dynamic.Middleware{
		RequestRedirect: &dynamic.RequestRedirect{
			Scheme:     filter.Scheme,
			Hostname:   hostname,
			Port:       port,
			Path:       path,
			PathPrefix: pathPrefix,
			StatusCode: ptr.Deref(filter.StatusCode, http.StatusFound),
		},
	}
}

func createURLRewrite(filter *gatev1.HTTPURLRewriteFilter, pathMatch gatev1.HTTPPathMatch) (*dynamic.Middleware, error) {
	if filter.Path == nil && filter.Hostname == nil {
		return nil, errors.New("empty configuration")
	}

	var host *string
	if filter.Hostname != nil {
		host = new(string(*filter.Hostname))
	}

	var path *string
	var pathPrefix *string
	if filter.Path != nil {
		switch filter.Path.Type {
		case gatev1.FullPathHTTPPathModifier:
			path = filter.Path.ReplaceFullPath
		case gatev1.PrefixMatchHTTPPathModifier:
			path = filter.Path.ReplacePrefixMatch
			pathPrefix = pathMatch.Value
		}
	}

	return &dynamic.Middleware{
		URLRewrite: &dynamic.URLRewrite{
			Hostname:   host,
			Path:       path,
			PathPrefix: pathPrefix,
		},
	}, nil
}

func createCORS(filter *gatev1.HTTPCORSFilter) *dynamic.Middleware {
	var allowOrigins, allowOriginsRegex []string
	for _, origin := range filter.AllowOrigins {
		if prefix, suffix, found := strings.Cut(string(origin), "*"); found {
			switch {
			case prefix != "" || suffix != "":
				allowOriginsRegex = append(allowOriginsRegex, "^"+regexp.QuoteMeta(prefix)+`.*`+regexp.QuoteMeta(suffix)+"$")
			default:
				// Convert to regex so the middleware echoes the specific request origin rather than the literal "*".
				// The Gateway API conformance tests require the specific origin, even when AllowCredentials is false/unset.
				allowOriginsRegex = append(allowOriginsRegex, ".*")
			}
			continue
		}

		allowOrigins = append(allowOrigins, string(origin))
	}

	var allowMethods []string
	for _, m := range filter.AllowMethods {
		allowMethods = append(allowMethods, string(m))
	}

	var allowHeaders []string
	for _, h := range filter.AllowHeaders {
		allowHeaders = append(allowHeaders, string(h))
	}

	var exposeHeaders []string
	for _, h := range filter.ExposeHeaders {
		exposeHeaders = append(exposeHeaders, string(h))
	}

	return &dynamic.Middleware{
		Headers: &dynamic.Headers{
			AccessControlAllowCredentials:     ptr.Deref(filter.AllowCredentials, false),
			AccessControlAllowOriginList:      allowOrigins,
			AccessControlAllowOriginListRegex: allowOriginsRegex,
			AccessControlAllowMethods:         allowMethods,
			AccessControlAllowHeaders:         allowHeaders,
			AccessControlExposeHeaders:        exposeHeaders,
			AccessControlMaxAge:               new(int64(filter.MaxAge)),
		},
	}
}

func getHTTPServiceProtocol(portSpec corev1.ServicePort) (string, error) {
	if portSpec.Protocol != corev1.ProtocolTCP {
		return "", errors.New("only TCP protocol is supported")
	}

	if portSpec.AppProtocol == nil {
		// Without an appProtocol the backend is assumed to serve plain HTTP.
		return schemeHTTP, nil
	}

	switch ap := strings.ToLower(*portSpec.AppProtocol); ap {
	case appProtocolH2C:
		return schemeH2C, nil
	case appProtocolHTTP, appProtocolWS:
		return schemeHTTP, nil
	case appProtocolHTTPS, appProtocolWSS:
		return schemeHTTPS, nil
	default:
		return "", fmt.Errorf("unsupported application protocol %s", ap)
	}
}

func mergeHTTPConfiguration(from, to *dynamic.Configuration) {
	if from == nil || from.HTTP == nil || to == nil {
		return
	}

	if to.HTTP == nil {
		to.HTTP = from.HTTP
		return
	}

	if to.HTTP.Routers == nil {
		to.HTTP.Routers = map[string]*dynamic.Router{}
	}
	maps.Copy(to.HTTP.Routers, from.HTTP.Routers)

	if to.HTTP.Middlewares == nil {
		to.HTTP.Middlewares = map[string]*dynamic.Middleware{}
	}
	maps.Copy(to.HTTP.Middlewares, from.HTTP.Middlewares)

	if to.HTTP.Services == nil {
		to.HTTP.Services = map[string]*dynamic.Service{}
	}
	maps.Copy(to.HTTP.Services, from.HTTP.Services)

	if to.HTTP.ServersTransports == nil {
		to.HTTP.ServersTransports = map[string]*dynamic.ServersTransport{}
	}
	maps.Copy(to.HTTP.ServersTransports, from.HTTP.ServersTransports)
}
