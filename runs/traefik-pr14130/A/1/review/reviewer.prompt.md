This directory holds one Go package, `pkg/`. Some of its files were just added
or changed; they are marked in the map below, and their full text follows it.
You may read any other file under pkg/.

Review the package as a code reviewer would. List the refactoring that should
be done now. Focus on the new and changed code and how it fits the package.

### pkg/

| file | lines | mark | declarations |
| --- | --- | --- | --- |
| annotations.go | 54 |  | annotationsPrefix (const), ServiceConfig (type), Service (type), parseServiceAnnotations (func), convertAnnotations (func) |
| annotations_test.go | 89 |  | Test_parseServiceConfig (func), Test_convertAnnotations (func) |
| client.go | 982 | CHANGED | resyncPeriod (const), clientWrapper (type), createClientFromConfig (func), newClientImpl (func), newInClusterClient (func), newExternalClusterClientFromFile (func), newExternalClusterClient (func), clientWrapper.WatchAll (method), clientWrapper.ListNamespaces (method), clientWrapper.ListHTTPRoutes (method), clientWrapper.ListGRPCRoutes (method), clientWrapper.ListTCPRoutes (method), clientWrapper.ListTLSRoutes (method), clientWrapper.ListReferenceGrants (method), clientWrapper.ListGateways (method), clientWrapper.ListListenerSets (method), clientWrapper.ListGatewayClasses (method), clientWrapper.ListEndpointSlicesForService (method), clientWrapper.ListBackendTLSPoliciesForService (method), clientWrapper.GetService (method), clientWrapper.GetSecret (method), clientWrapper.GetConfigMap (method), clientWrapper.UpdateGatewayClassStatus (method), clientWrapper.UpdateGatewayStatus (method), clientWrapper.UpdateHTTPRouteStatus (method), clientWrapper.UpdateGRPCRouteStatus (method), clientWrapper.UpdateTCPRouteStatus (method), clientWrapper.UpdateTLSRouteStatus (method), clientWrapper.UpdateBackendTLSPolicyStatus (method), clientWrapper.UpdateListenerSetStatus (method), clientWrapper.lookupNamespace (method), clientWrapper.isWatchedNamespace (method), mergeRouteParentStatuses (func), gatewayConditionTypes (var), listenerConditionTypes (var), withoutForeignConditions (func), mergeGatewayStatus (func), gatewayStatusEqual (func), policyAncestorStatusesEqual (func), policyAncestorStatusEqual (func), routeParentStatusesEqual (func), routeParentStatusEqual (func), listenersStatusEqual (func), conditionsEqual (func), listenerSetStatusEqual (func), listenerEntryStatusesEqual (func) |
| client_test.go | 553 | CHANGED | Test_gatewayStatusEquals (func), Test_listenerSetStatusEqual (func), Test_mergeGatewayStatus (func), Test_withoutForeignConditions (func) |
| features.go | 96 | CHANGED | SupportedFeatures (var), extendedGatewayFeatures (func), extendedTLSRouteFeatures (func), extendedHTTPRouteFeatures (func), extendedGRPCRouteFeatures (func), extendedBackendTLSPolicyFeatures (func) |
| grpcroute.go | 575 | CHANGED | Provider.loadGRPCRoute (method), grpcRouteRouter (type), Provider.loadGRPCRouteConfiguration (method), Provider.loadGRPCService (method), Provider.loadGRPCBackendRef (method), Provider.loadGRPCMiddlewares (method), Provider.loadGRPCServers (method), buildGRPCMatchRule (func), buildGRPCMethodRule (func), buildGRPCHeaderRules (func), getGRPCServiceProtocol (func) |
| grpcroute_test.go | 251 |  | Test_buildGRPCMatchRule (func), Test_buildGRPCMethodRule (func), Test_buildGRPCHeaderRules (func) |
| httproute.go | 1242 | CHANGED | Provider.loadHTTPRoute (method), httpRouteRouter (type), Provider.loadHTTPRouteConfiguration (method), Provider.loadWRRService (method), Provider.loadService (method), Provider.loadHTTPBackendRef (method), mirrorServiceLoader (type), Provider.newMirrorServiceLoader (method), Provider.loadMiddlewares (method), Provider.loadHTTPRouteFilterExtensionRef (method), Provider.loadHTTPServers (method), Provider.loadServersTransport (method), buildHostRule (func), buildMatchRule (func), buildPathRule (func), buildHeaderRules (func), buildQueryParamRules (func), createRequestHeaderModifier (func), createResponseHeaderModifier (func), createRequestRedirect (func), createURLRewrite (func), Provider.createRetry (method), createRequestMirror (func), mirrorRatio (func), createCORS (func), getHTTPServiceProtocol (func), mergeHTTPConfiguration (func) |
| httproute_test.go | 412 | CHANGED | Test_buildHostRule (func), Test_buildMatchRule (func), Test_createCORS (func), Test_createRequestMirror (func) |
| kubernetes.go | 2357 | CHANGED | ProviderName (const), controllerName (const), groupCore (const), groupGateway (const), kindGateway (const), kindTraefikService (const), kindHTTPRoute (const), kindGRPCRoute (const), kindTCPRoute (const), kindTLSRoute (const), kindService (const), kindListenerSet (const), kindConfigMap (const), kindSecret (const), appProtocolHTTP (const), appProtocolHTTPS (const), appProtocolH2C (const), appProtocolWS (const), appProtocolWSS (const), schemeHTTP (const), schemeHTTPS (const), schemeH2C (const), routeReasonHostnameConflict (const), listenerReasonFrontendValidationConflict (const), conditionNoErrorMessage (const), NamespacedName (type), Provider (type), Provider.SetDefaults (method), Entrypoint (type), StatusAddress (type), ServiceRef (type), BuildFilterFunc (type), BuildBackendFunc (type), ExtensionBuilderRegistry (type), listenerOwner (type), gatewayListener (type), gatewayWithListeners (type), Provider.RegisterFilterFuncs (method), Provider.RegisterBackendFuncs (method), Provider.SetRouterTransform (method), Provider.Init (method), Provider.Provide (method), Provider.applyRouterTransform (method), Provider.managesGateway (method), Provider.newK8sClient (method), Provider.loadConfigurationFromGateways (method), Provider.loadHTTPAndGRPCRoutes (method), Provider.loadGatewayListeners (method), uniqListener (type), Provider.buildListenerRouters (method), dropChildlessListenerRouters (func), mostSpecificListener (func), hostnameCovers (func), buildListenerRule (func), hostnameMatcherValue (func), Provider.loadListenerSetListeners (method), frontendValidation (type), Provider.resolveFrontendValidation (method), Provider.resolveCACertificateRefs (method), Provider.resolveBackendClientCertificate (method), makeGatewayResolvedRefsCondition (func), gatewayResolvedRefsCondition (func), Provider.makeGatewayStatus (method), Provider.gatewayAddresses (method), Provider.entryPointName (method), Provider.isReferenceGranted (method), Provider.getTLSCert (method), Provider.allowedNamespaces (method), backendAddress (type), Provider.getBackendAddresses (method), hasInsecureFrontendValidationMode (func), supportedRouteKinds (func), allowedRouteKinds (func), findMatchingHostnames (func), findMatchingHostname (func), allowRoute (func), compareRoutes (func), servedRuleKey (type), servedRules (type), servedRules.register (method), listenerRef (type), attachedRoutes (type), attachedRoutes.Record (method), attachedRoutes.Conflicts (method), gatewayListenersForParentRef (type), matchingGatewayListenersForParentRef (func), gatewayForOwner (func), matchListener (func), makeRouterName (func), makeListenerRouterName (func), getTLSConfig (func), getCertificateBlocks (func), throttleEvents (func), isTraefikService (func), isInternalService (func), isCrossProviderNamespaceAllowed (func), makeListenerKey (func), listenerClaims (type), newListenerClaims (func), makeListenerConflictConditions (func), listenerSetRefsGateway (func), Provider.isListenerSetAllowed (method), makeListenerSetStatus (func), makeListenerSetCondition (func), filterReferenceGrantsFrom (func), referenceGrantMatchesFrom (func), filterReferenceGrantsTo (func), referenceGrantMatchesTo (func), groupToString (func), kindToString (func), upsertGatewayClassConditionAccepted (func) |
| kubernetes_test.go | 16087 | CHANGED | defaultListenerTLSOptions (func), init (func), listenerCert (const), listenerKey (const), TestGatewayClassLabelSelector (func), TestGatewayClassFieldSelector (func), TestGatewayClassAndLabelSelectorExclusive (func), TestGatewayScoping (func), TestLoadHTTPRoutes (func), TestLoadHTTPRoutes_backendExtensionRef (func), TestLoadHTTPRoutes_filterExtensionRef (func), TestLoadGRPCRoutes (func), TestLoadGRPCRoutes_filterExtensionRef (func), TestLoadTCPRoutes (func), TestLoadTLSRoutes (func), TestLoadMixedRoutes (func), TestLoadRoutesWithReferenceGrants (func), Test_matchingGatewayListener (func), Test_loadRoutes_multipleGatewaysParentRefs (func), Test_loadHTTPAndGRPCRoutes_hostnameConflict (func), Test_attachedRoutes_conflicts (func), Test_mergeRouteParentStatuses (func), Test_managesGateway (func), Test_matchListener (func), Test_allowRoute (func), Test_findMatchingHostnames (func), Test_allowedRouteKinds (func), Test_makeListenerKey (func), Test_referenceGrantMatchesFrom (func), Test_referenceGrantMatchesTo (func), Test_gatewayAddresses (func), Test_makeGatewayStatus (func), Test_resolveFrontendValidation (func), Test_resolveBackendClientCertificate (func), Test_loadConfigurationFromGateways_backendClientCertificate (func), Test_loadGatewayListeners (func), Test_loadGatewayListeners_frontendValidationConflict (func), Test_isCrossProviderNamespaceAllowed (func), Test_makeListenerRouterName (func), TestCrossProviderNamespaces_HTTPRoute (func), TestCrossProviderNamespaces_TCPRoute (func), TestCrossProviderNamespaces_TLSRoute (func), findCondition (func), newGatewaySimpleClientSet (func), Test_loadConfigurationFromGateways_TLSCertificatesOrder (func), Test_loadConfigurationFromGateways_BackendTLSPolicyAncestorNamespace (func), certificateCommonName (func), readResources (func), Test_listenerSetRefsGateway (func), Test_isListenerSetAllowed (func), Test_makeListenerSetStatus (func), Test_loadListenerSetListeners (func) |
| status.go | 149 |  | statusReport (type), newStatusReport (func), statusReport.Flush (method), statusReport.RecordGatewayClassStatus (method), statusReport.RecordGatewayStatus (method), statusReport.RecordListenerSetStatus (method), statusReport.RecordHTTPRouteStatus (method), statusReport.RecordGRPCRouteStatus (method), statusReport.RecordTCPRouteStatus (method), statusReport.RecordTLSRouteStatus (method), statusReport.RecordBackendTLSPolicyStatus (method) |
| status_test.go | 152 |  | TestStatusReport_RecordGatewayClassStatus (func), TestStatusReport_RecordGatewayStatus (func), TestStatusReport_RecordHTTPRouteStatus (func), TestStatusReport_RecordGRPCRouteStatus (func), TestStatusReport_RecordTCPRouteStatus (func), TestStatusReport_RecordTLSRouteStatus (func), TestStatusReport_RecordBackendTLSPolicyStatus (func) |
| tcproute.go | 392 |  | Provider.loadTCPRoutes (method), tcpRouteRouter (type), Provider.loadTCPRoute (method), Provider.loadTCPWRRService (method), Provider.loadTCPService (method), Provider.loadTCPServers (method), Provider.loadTCPBackendRef (method), mergeTCPConfiguration (func) |
| tlsroute.go | 637 | CHANGED | Provider.loadTLSRoutes (method), tlsRouteRouter (type), Provider.loadTLSRoute (method), Provider.loadTLSWRRService (method), Provider.loadTLSService (method), Provider.loadTLSServers (method), Provider.loadTCPServersTransport (method), hostSNIRule (func) |
| tlsroute_test.go | 76 |  | Test_hostSNIRule (func) |

#### pkg/client.go

```go
package gateway

import (
	"context"
	"errors"
	"fmt"
	"os"
	"reflect"
	"slices"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/traefik/traefik/v3/pkg/provider/kubernetes/k8s"
	"github.com/traefik/traefik/v3/pkg/types"
	corev1 "k8s.io/api/core/v1"
	discoveryv1 "k8s.io/api/discovery/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	ktypes "k8s.io/apimachinery/pkg/types"
	kinformers "k8s.io/client-go/informers"
	kclientset "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/retry"
	"k8s.io/utils/ptr"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
	gateclientset "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned"
	gateinformers "sigs.k8s.io/gateway-api/pkg/client/informers/externalversions"
)

const resyncPeriod = 10 * time.Minute

type clientWrapper struct {
	csGateway gateclientset.Interface
	csKube    kclientset.Interface

	factoryNamespace    kinformers.SharedInformerFactory
	factoryGatewayClass gateinformers.SharedInformerFactory
	factoriesGateway    map[string]gateinformers.SharedInformerFactory
	factoriesKube       map[string]kinformers.SharedInformerFactory
	factoriesSecret     map[string]kinformers.SharedInformerFactory

	isNamespaceAll    bool
	watchedNamespaces []string

	gatewayClass  string
	labelSelector string
}

func createClientFromConfig(c *rest.Config, qps, burst int) (*clientWrapper, error) {
	c.QPS = float32(qps)
	c.Burst = burst

	csGateway, err := gateclientset.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	csKube, err := kclientset.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	return newClientImpl(csKube, csGateway), nil
}

func newClientImpl(csKube kclientset.Interface, csGateway gateclientset.Interface) *clientWrapper {
	return &clientWrapper{
		csGateway:        csGateway,
		csKube:           csKube,
		factoriesGateway: make(map[string]gateinformers.SharedInformerFactory),
		factoriesKube:    make(map[string]kinformers.SharedInformerFactory),
		factoriesSecret:  make(map[string]kinformers.SharedInformerFactory),
	}
}

// newInClusterClient returns a new Provider client that is expected to run
// inside the cluster.
func newInClusterClient(endpoint string, qps, burst int) (*clientWrapper, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to create in-cluster configuration: %w", err)
	}

	if endpoint != "" {
		config.Host = endpoint
	}

	return createClientFromConfig(config, qps, burst)
}

func newExternalClusterClientFromFile(file string, qps, burst int) (*clientWrapper, error) {
	configFromFlags, err := clientcmd.BuildConfigFromFlags("", file)
	if err != nil {
		return nil, err
	}
	return createClientFromConfig(configFromFlags, qps, burst)
}

// newExternalClusterClient returns a new Provider client that may run outside of the cluster.
// The endpoint parameter must not be empty.
func newExternalClusterClient(endpoint, caFilePath string, token types.FileOrContent, qps, burst int) (*clientWrapper, error) {
	if endpoint == "" {
		return nil, errors.New("endpoint missing for external cluster client")
	}

	tokenData, err := token.Read()
	if err != nil {
		return nil, fmt.Errorf("read token: %w", err)
	}

	config := &rest.Config{
		Host:        endpoint,
		BearerToken: string(tokenData),
	}

	if caFilePath != "" {
		caData, err := os.ReadFile(caFilePath)
		if err != nil {
			return nil, fmt.Errorf("failed to read CA file %s: %w", caFilePath, err)
		}

		config.TLSClientConfig = rest.TLSClientConfig{CAData: caData}
	}

	return createClientFromConfig(config, qps, burst)
}

// WatchAll starts namespace-specific controllers for all relevant kinds.
func (c *clientWrapper) WatchAll(namespaces []string, stopCh <-chan struct{}) (<-chan any, error) {
	eventCh := make(chan any, 1)
	eventHandler := &k8s.ResourceEventHandler{Ev: eventCh}

	if len(namespaces) == 0 {
		namespaces = []string{metav1.NamespaceAll}
		c.isNamespaceAll = true
	}

	c.watchedNamespaces = namespaces

	notOwnedByHelm := func(opts *metav1.ListOptions) {
		opts.LabelSelector = "owner!=helm"
	}

	gatewayClassOptions := func(options *metav1.ListOptions) {
		switch {
		case c.labelSelector != "":
			options.LabelSelector = c.labelSelector
		case c.gatewayClass != "":
			options.FieldSelector = fields.OneTermEqualSelector("metadata.name", c.gatewayClass).String()
		}
	}

	c.factoryNamespace = kinformers.NewSharedInformerFactoryWithOptions(c.csKube, resyncPeriod, kinformers.WithTransform(k8s.StripManagedFields))
	_, err := c.factoryNamespace.Core().V1().Namespaces().Informer().AddEventHandler(eventHandler)
	if err != nil {
		return nil, err
	}

	c.factoryGatewayClass = gateinformers.NewSharedInformerFactoryWithOptions(c.csGateway, resyncPeriod, gateinformers.WithTweakListOptions(gatewayClassOptions), gateinformers.WithTransform(k8s.StripManagedFields))
	_, err = c.factoryGatewayClass.Gateway().V1().GatewayClasses().Informer().AddEventHandler(eventHandler)
	if err != nil {
		return nil, err
	}

	for _, ns := range namespaces {
		factoryKube := kinformers.NewSharedInformerFactoryWithOptions(c.csKube, resyncPeriod, kinformers.WithNamespace(ns), kinformers.WithTransform(k8s.StripManagedFields))
		_, err = factoryKube.Core().V1().Services().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}
		endpointSliceInformer := factoryKube.Discovery().V1().EndpointSlices().Informer()
		if err = endpointSliceInformer.AddIndexers(k8s.EndpointSliceByServiceNameIndexers); err != nil {
			return nil, err
		}
		_, err = endpointSliceInformer.AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}

		factoryGateway := gateinformers.NewSharedInformerFactoryWithOptions(c.csGateway, resyncPeriod, gateinformers.WithNamespace(ns), gateinformers.WithTransform(k8s.StripManagedFields))
		_, err = factoryGateway.Gateway().V1().Gateways().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}
		_, err = factoryGateway.Gateway().V1().HTTPRoutes().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}
		_, err = factoryGateway.Gateway().V1().GRPCRoutes().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}
		_, err = factoryGateway.Gateway().V1().ReferenceGrants().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}
		_, err = factoryGateway.Gateway().V1().BackendTLSPolicies().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}
		_, err = factoryKube.Core().V1().ConfigMaps().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}

		_, err = factoryGateway.Gateway().V1().TLSRoutes().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}

		_, err = factoryGateway.Gateway().V1().ListenerSets().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}

		_, err = factoryGateway.Gateway().V1().TCPRoutes().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}

		factorySecret := kinformers.NewSharedInformerFactoryWithOptions(c.csKube, resyncPeriod, kinformers.WithNamespace(ns), kinformers.WithTweakListOptions(notOwnedByHelm), kinformers.WithTransform(k8s.StripManagedFields))
		_, err = factorySecret.Core().V1().Secrets().Informer().AddEventHandler(eventHandler)
		if err != nil {
			return nil, err
		}

		c.factoriesGateway[ns] = factoryGateway
		c.factoriesKube[ns] = factoryKube
		c.factoriesSecret[ns] = factorySecret
	}

	c.factoryNamespace.Start(stopCh)
	c.factoryGatewayClass.Start(stopCh)

	for _, ns := range namespaces {
		c.factoriesGateway[ns].Start(stopCh)
		c.factoriesKube[ns].Start(stopCh)
		c.factoriesSecret[ns].Start(stopCh)
	}

	for t, ok := range c.factoryNamespace.WaitForCacheSync(stopCh) {
		if !ok {
			return nil, fmt.Errorf("timed out waiting for controller caches to sync %s", t.String())
		}
	}

	for t, ok := range c.factoryGatewayClass.WaitForCacheSync(stopCh) {
		if !ok {
			return nil, fmt.Errorf("timed out waiting for controller caches to sync %s", t.String())
		}
	}

	for _, ns := range namespaces {
		for t, ok := range c.factoriesGateway[ns].WaitForCacheSync(stopCh) {
			if !ok {
				return nil, fmt.Errorf("timed out waiting for controller caches to sync %s in namespace %q", t.String(), ns)
			}
		}

		for t, ok := range c.factoriesKube[ns].WaitForCacheSync(stopCh) {
			if !ok {
				return nil, fmt.Errorf("timed out waiting for controller caches to sync %s in namespace %q", t.String(), ns)
			}
		}

		for t, ok := range c.factoriesSecret[ns].WaitForCacheSync(stopCh) {
			if !ok {
				return nil, fmt.Errorf("timed out waiting for controller caches to sync %s in namespace %q", t.String(), ns)
			}
		}
	}

	return eventCh, nil
}

func (c *clientWrapper) ListNamespaces(selector labels.Selector) ([]string, error) {
	ns, err := c.factoryNamespace.Core().V1().Namespaces().Lister().List(selector)
	if err != nil {
		return nil, err
	}

	var namespaces []string
	for _, namespace := range ns {
		if !c.isWatchedNamespace(namespace.Name) {
			log.Warn().Msgf("Namespace %q is not within %q watched namespaces", selector, namespace)
			continue
		}
		namespaces = append(namespaces, namespace.Name)
	}
	return namespaces, nil
}

func (c *clientWrapper) ListHTTPRoutes() ([]*gatev1.HTTPRoute, error) {
	var httpRoutes []*gatev1.HTTPRoute
	for _, namespace := range c.watchedNamespaces {
		routes, err := c.factoriesGateway[c.lookupNamespace(namespace)].Gateway().V1().HTTPRoutes().Lister().HTTPRoutes(namespace).List(labels.Everything())
		if err != nil {
			return nil, fmt.Errorf("listing HTTP routes in namespace %s", namespace)
		}

		httpRoutes = append(httpRoutes, routes...)
	}

	return httpRoutes, nil
}

func (c *clientWrapper) ListGRPCRoutes() ([]*gatev1.GRPCRoute, error) {
	var grpcRoutes []*gatev1.GRPCRoute
	for _, namespace := range c.watchedNamespaces {
		routes, err := c.factoriesGateway[c.lookupNamespace(namespace)].Gateway().V1().GRPCRoutes().Lister().GRPCRoutes(namespace).List(labels.Everything())
		if err != nil {
			return nil, fmt.Errorf("listing GRPC routes in namespace %s", namespace)
		}

		grpcRoutes = append(grpcRoutes, routes...)
	}

	return grpcRoutes, nil
}

func (c *clientWrapper) ListTCPRoutes() ([]*gatev1.TCPRoute, error) {
	var tcpRoutes []*gatev1.TCPRoute
	for _, namespace := range c.watchedNamespaces {
		routes, err := c.factoriesGateway[c.lookupNamespace(namespace)].Gateway().V1().TCPRoutes().Lister().TCPRoutes(namespace).List(labels.Everything())
		if err != nil {
			return nil, fmt.Errorf("listing TCP routes in namespace %s", namespace)
		}

		tcpRoutes = append(tcpRoutes, routes...)
	}

	return tcpRoutes, nil
}

func (c *clientWrapper) ListTLSRoutes() ([]*gatev1.TLSRoute, error) {
	var tlsRoutes []*gatev1.TLSRoute
	for _, namespace := range c.watchedNamespaces {
		routes, err := c.factoriesGateway[c.lookupNamespace(namespace)].Gateway().V1().TLSRoutes().Lister().TLSRoutes(namespace).List(labels.Everything())
		if err != nil {
			return nil, fmt.Errorf("listing TLS routes in namespace %s", namespace)
		}

		tlsRoutes = append(tlsRoutes, routes...)
	}

	return tlsRoutes, nil
}

func (c *clientWrapper) ListReferenceGrants(namespace string) ([]*gatev1.ReferenceGrant, error) {
	if !c.isWatchedNamespace(namespace) {
		return nil, fmt.Errorf("failed to get ReferenceGrants: namespace %s is not within watched namespaces", namespace)
	}

	referenceGrants, err := c.factoriesGateway[c.lookupNamespace(namespace)].Gateway().V1().ReferenceGrants().Lister().ReferenceGrants(namespace).List(labels.Everything())
	if err != nil {
		return nil, err
	}

	return referenceGrants, nil
}

func (c *clientWrapper) ListGateways() []*gatev1.Gateway {
	var result []*gatev1.Gateway

	for ns, factory := range c.factoriesGateway {
		gateways, err := factory.Gateway().V1().Gateways().Lister().Gateways(ns).List(labels.Everything())
		if err != nil {
			log.Error().Err(err).Msgf("Failed to list Gateways in namespace %s", ns)
			continue
		}
		result = append(result, gateways...)
	}

	return result
}

func (c *clientWrapper) ListListenerSets() []*gatev1.ListenerSet {
	var result []*gatev1.ListenerSet

	for ns, factory := range c.factoriesGateway {
		listenerSets, err := factory.Gateway().V1().ListenerSets().Lister().ListenerSets(ns).List(labels.Everything())
		if err != nil {
			log.Error().Err(err).Msgf("Failed to list ListenerSets in namespace %s", ns)
			continue
		}
		result = append(result, listenerSets...)
	}

	return result
}

func (c *clientWrapper) ListGatewayClasses() ([]*gatev1.GatewayClass, error) {
	return c.factoryGatewayClass.Gateway().V1().GatewayClasses().Lister().List(labels.Everything())
}

// ListEndpointSlicesForService returns the EndpointSlices for the given service name in the given namespace.
func (c *clientWrapper) ListEndpointSlicesForService(namespace, serviceName string) ([]*discoveryv1.EndpointSlice, error) {
	if !c.isWatchedNamespace(namespace) {
		return nil, fmt.Errorf("failed to get endpointslices for service %s/%s: namespace is not within watched namespaces", namespace, serviceName)
	}

	return k8s.EndpointSlicesByServiceName(
		c.factoriesKube[c.lookupNamespace(namespace)].Discovery().V1().EndpointSlices().Informer().GetIndexer(),
		namespace,
		serviceName,
	)
}

// ListBackendTLSPoliciesForService returns the BackendTLSPolicy for the given service name in the given namespace.
func (c *clientWrapper) ListBackendTLSPoliciesForService(namespace, serviceName string) ([]*gatev1.BackendTLSPolicy, error) {
	if !c.isWatchedNamespace(namespace) {
		return nil, fmt.Errorf("failed to get BackendTLSPolicies for service %s/%s: namespace is not within watched namespaces", namespace, serviceName)
	}

	policies, err := c.factoriesGateway[c.lookupNamespace(namespace)].Gateway().V1().BackendTLSPolicies().Lister().BackendTLSPolicies(namespace).List(labels.Everything())
	if err != nil {
		return nil, fmt.Errorf("failed to list BackendTLSPolicies in namespace %s", namespace)
	}

	var servicePolicies []*gatev1.BackendTLSPolicy
	for _, policy := range policies {
		for _, ref := range policy.Spec.TargetRefs {
			// The policy does not target the service.
			if (ref.Group != "" && ref.Group != groupCore) || ref.Kind != kindService || string(ref.Name) != serviceName {
				continue
			}

			servicePolicies = append(servicePolicies, policy)
		}
	}

	return servicePolicies, nil
}

// GetService returns the named service from the given namespace.
func (c *clientWrapper) GetService(namespace, name string) (*corev1.Service, error) {
	if !c.isWatchedNamespace(namespace) {
		return nil, fmt.Errorf("failed to get service %s/%s: namespace is not within watched namespaces", namespace, name)
	}
	return c.factoriesKube[c.lookupNamespace(namespace)].Core().V1().Services().Lister().Services(namespace).Get(name)
}

// GetSecret returns the named secret from the given namespace.
func (c *clientWrapper) GetSecret(namespace, name string) (*corev1.Secret, error) {
	if !c.isWatchedNamespace(namespace) {
		return nil, fmt.Errorf("failed to get secret %s/%s: namespace is not within watched namespaces", namespace, name)
	}
	return c.factoriesSecret[c.lookupNamespace(namespace)].Core().V1().Secrets().Lister().Secrets(namespace).Get(name)
}

// GetConfigMap returns the named configMap from the given namespace.
func (c *clientWrapper) GetConfigMap(namespace, name string) (*corev1.ConfigMap, error) {
	if !c.isWatchedNamespace(namespace) {
		return nil, fmt.Errorf("failed to get configMap %s/%s: namespace is not within watched namespaces", namespace, name)
	}
	return c.factoriesKube[c.lookupNamespace(namespace)].Core().V1().ConfigMaps().Lister().ConfigMaps(namespace).Get(name)
}

func (c *clientWrapper) UpdateGatewayClassStatus(ctx context.Context, name string, status gatev1.GatewayClassStatus) error {
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		currentGatewayClass, err := c.factoryGatewayClass.Gateway().V1().GatewayClasses().Lister().Get(name)
		if err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		if conditionsEqual(currentGatewayClass.Status.Conditions, status.Conditions) {
			return nil
		}

		currentGatewayClass = currentGatewayClass.DeepCopy()
		currentGatewayClass.Status = status

		if _, err = c.csGateway.GatewayV1().GatewayClasses().UpdateStatus(ctx, currentGatewayClass, metav1.UpdateOptions{}); err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update GatewayClass %s status: %w", name, err)
	}

	return nil
}

func (c *clientWrapper) UpdateGatewayStatus(ctx context.Context, gateway ktypes.NamespacedName, status gatev1.GatewayStatus) error {
	if !c.isWatchedNamespace(gateway.Namespace) {
		return fmt.Errorf("cannot update Gateway status %s/%s: namespace is not within watched namespaces", gateway.Namespace, gateway.Name)
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		currentGateway, err := c.factoriesGateway[c.lookupNamespace(gateway.Namespace)].Gateway().V1().Gateways().Lister().Gateways(gateway.Namespace).Get(gateway.Name)
		if err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		if gatewayStatusEqual(withoutForeignConditions(currentGateway.Status), status) {
			return nil
		}

		currentGateway = currentGateway.DeepCopy()
		currentGateway.Status = mergeGatewayStatus(currentGateway.Status, status)

		if _, err = c.csGateway.GatewayV1().Gateways(gateway.Namespace).UpdateStatus(ctx, currentGateway, metav1.UpdateOptions{}); err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update Gateway %s/%s status: %w", gateway.Namespace, gateway.Name, err)
	}

	return nil
}

func (c *clientWrapper) UpdateHTTPRouteStatus(ctx context.Context, route ktypes.NamespacedName, gateways []gatewayWithListeners, status gatev1.HTTPRouteStatus) error {
	if !c.isWatchedNamespace(route.Namespace) {
		return fmt.Errorf("updating HTTPRoute status %s/%s: namespace is not within watched namespaces", route.Namespace, route.Name)
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		currentRoute, err := c.factoriesGateway[c.lookupNamespace(route.Namespace)].Gateway().V1().HTTPRoutes().Lister().HTTPRoutes(route.Namespace).Get(route.Name)
		if err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		parentStatuses := mergeRouteParentStatuses(route.Namespace, currentRoute.Status.Parents, status.Parents, gateways)

		// do not update status when nothing has changed.
		if routeParentStatusesEqual(currentRoute.Status.Parents, parentStatuses) {
			return nil
		}

		currentRoute = currentRoute.DeepCopy()
		currentRoute.Status = gatev1.HTTPRouteStatus{
			Parents: parentStatuses,
		}

		if _, err = c.csGateway.GatewayV1().HTTPRoutes(route.Namespace).UpdateStatus(ctx, currentRoute, metav1.UpdateOptions{}); err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update HTTPRoute %s/%s status: %w", route.Namespace, route.Name, err)
	}

	return nil
}

func (c *clientWrapper) UpdateGRPCRouteStatus(ctx context.Context, route ktypes.NamespacedName, gateways []gatewayWithListeners, status gatev1.GRPCRouteStatus) error {
	if !c.isWatchedNamespace(route.Namespace) {
		return fmt.Errorf("updating GRPCRoute status %s/%s: namespace is not within watched namespaces", route.Namespace, route.Name)
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		currentRoute, err := c.factoriesGateway[c.lookupNamespace(route.Namespace)].Gateway().V1().GRPCRoutes().Lister().GRPCRoutes(route.Namespace).Get(route.Name)
		if err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		parentStatuses := mergeRouteParentStatuses(route.Namespace, currentRoute.Status.Parents, status.Parents, gateways)

		// do not update status when nothing has changed.
		if routeParentStatusesEqual(currentRoute.Status.Parents, parentStatuses) {
			return nil
		}

		currentRoute = currentRoute.DeepCopy()
		currentRoute.Status = gatev1.GRPCRouteStatus{
			Parents: parentStatuses,
		}

		if _, err = c.csGateway.GatewayV1().GRPCRoutes(route.Namespace).UpdateStatus(ctx, currentRoute, metav1.UpdateOptions{}); err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update GRPCRoute %q status: %w", route.Name, err)
	}

	return nil
}

func (c *clientWrapper) UpdateTCPRouteStatus(ctx context.Context, route ktypes.NamespacedName, gateways []gatewayWithListeners, status gatev1.TCPRouteStatus) error {
	if !c.isWatchedNamespace(route.Namespace) {
		return fmt.Errorf("updating TCPRoute status %s/%s: namespace is not within watched namespaces", route.Namespace, route.Name)
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		currentRoute, err := c.factoriesGateway[c.lookupNamespace(route.Namespace)].Gateway().V1().TCPRoutes().Lister().TCPRoutes(route.Namespace).Get(route.Name)
		if err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		parentStatuses := mergeRouteParentStatuses(route.Namespace, currentRoute.Status.Parents, status.Parents, gateways)

		// do not update status when nothing has changed.
		if routeParentStatusesEqual(currentRoute.Status.Parents, parentStatuses) {
			return nil
		}

		currentRoute = currentRoute.DeepCopy()
		currentRoute.Status = gatev1.TCPRouteStatus{
			Parents: parentStatuses,
		}

		if _, err = c.csGateway.GatewayV1().TCPRoutes(route.Namespace).UpdateStatus(ctx, currentRoute, metav1.UpdateOptions{}); err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update TCPRoute %s/%s status: %w", route.Namespace, route.Name, err)
	}

	return nil
}

func (c *clientWrapper) UpdateTLSRouteStatus(ctx context.Context, route ktypes.NamespacedName, gateways []gatewayWithListeners, status gatev1.TLSRouteStatus) error {
	if !c.isWatchedNamespace(route.Namespace) {
		return fmt.Errorf("updating TLSRoute status %s/%s: namespace is not within watched namespaces", route.Namespace, route.Name)
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		currentRoute, err := c.factoriesGateway[c.lookupNamespace(route.Namespace)].Gateway().V1().TLSRoutes().Lister().TLSRoutes(route.Namespace).Get(route.Name)
		if err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		parentStatuses := mergeRouteParentStatuses(route.Namespace, currentRoute.Status.Parents, status.Parents, gateways)

		// do not update status when nothing has changed.
		if routeParentStatusesEqual(currentRoute.Status.Parents, parentStatuses) {
			return nil
		}

		currentRoute = currentRoute.DeepCopy()
		currentRoute.Status = gatev1.TLSRouteStatus{
			Parents: parentStatuses,
		}

		if _, err = c.csGateway.GatewayV1().TLSRoutes(route.Namespace).UpdateStatus(ctx, currentRoute, metav1.UpdateOptions{}); err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update TLSRoute %s/%s status: %w", route.Namespace, route.Name, err)
	}

	return nil
}

func (c *clientWrapper) UpdateBackendTLSPolicyStatus(ctx context.Context, policy ktypes.NamespacedName, status gatev1.PolicyStatus) error {
	if !c.isWatchedNamespace(policy.Namespace) {
		return fmt.Errorf("updating BackendTLSPolicy status %s/%s: namespace is not within watched namespaces", policy.Namespace, policy.Name)
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		currentPolicy, err := c.factoriesGateway[c.lookupNamespace(policy.Namespace)].Gateway().V1().BackendTLSPolicies().Lister().BackendTLSPolicies(policy.Namespace).Get(policy.Name)
		if err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		ancestorStatuses := make([]gatev1.PolicyAncestorStatus, len(status.Ancestors))
		copy(ancestorStatuses, status.Ancestors)

		for _, ancestorStatus := range currentPolicy.Status.Ancestors {
			// Keep statuses added by other gateway controllers.
			if ancestorStatus.ControllerName != controllerName {
				ancestorStatuses = append(ancestorStatuses, ancestorStatus)
				continue
			}
		}

		if len(ancestorStatuses) > 16 {
			return fmt.Errorf("failed to update BackendTLSPolicy %s/%s status: PolicyAncestor statuses count exceeds 16", policy.Namespace, policy.Name)
		}

		// Do not update status when nothing has changed.
		if policyAncestorStatusesEqual(currentPolicy.Status.Ancestors, ancestorStatuses) {
			return nil
		}

		currentPolicy = currentPolicy.DeepCopy()
		currentPolicy.Status = gatev1.PolicyStatus{
			Ancestors: ancestorStatuses,
		}

		if _, err = c.csGateway.GatewayV1().BackendTLSPolicies(policy.Namespace).UpdateStatus(ctx, currentPolicy, metav1.UpdateOptions{}); err != nil {
			// We have to return err itself here (not wrapped inside another error)
			// so that RetryOnConflict can identify it correctly.
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update BackendTLSPolicy %s/%s status: %w", policy.Namespace, policy.Name, err)
	}

	return nil
}

func (c *clientWrapper) UpdateListenerSetStatus(ctx context.Context, listenerSet ktypes.NamespacedName, status gatev1.ListenerSetStatus) error {
	if !c.isWatchedNamespace(listenerSet.Namespace) {
		return fmt.Errorf("cannot update ListenerSet status %s/%s: namespace is not within watched namespaces", listenerSet.Namespace, listenerSet.Name)
	}

	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		currentListenerSet, err := c.factoriesGateway[c.lookupNamespace(listenerSet.Namespace)].Gateway().V1().ListenerSets().Lister().ListenerSets(listenerSet.Namespace).Get(listenerSet.Name)
		if err != nil {
			// We have to return err itself here (not wrapped inside another error) so that RetryOnConflict can identify it correctly.
			return err
		}

		if listenerSetStatusEqual(currentListenerSet.Status, status) {
			return nil
		}

		currentListenerSet = currentListenerSet.DeepCopy()
		currentListenerSet.Status = status

		if _, err = c.csGateway.GatewayV1().ListenerSets(listenerSet.Namespace).UpdateStatus(ctx, currentListenerSet, metav1.UpdateOptions{}); err != nil {
			// We have to return err itself here (not wrapped inside another error) so that RetryOnConflict can identify it correctly.
			return err
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("failed to update ListenerSet %s/%s status: %w", listenerSet.Namespace, listenerSet.Name, err)
	}

	return nil
}

// lookupNamespace returns the lookup namespace listenerKey for the given namespace.
// When listening on all namespaces, it returns the client-go identifier ("")
// for all-namespaces. Otherwise, it returns the given namespace.
// The distinction is necessary because we index all informers on the special
// identifier iff all-namespaces are requested but receive specific namespace
// identifiers from the Kubernetes API, so we have to bridge this gap.
func (c *clientWrapper) lookupNamespace(namespace string) string {
	if c.isNamespaceAll {
		return metav1.NamespaceAll
	}
	return namespace
}

// isWatchedNamespace checks to ensure that the namespace is being watched before we request
// it to ensure we don't panic by requesting an out-of-watch object.
func (c *clientWrapper) isWatchedNamespace(namespace string) bool {
	if c.isNamespaceAll {
		return true
	}

	return slices.Contains(c.watchedNamespaces, namespace)
}

// mergeRouteParentStatuses returns the desired parentStatuses augmented with the current parentStatuses Traefik does not own.
func mergeRouteParentStatuses(routeNamespace string, currentParents, desiredParents []gatev1.RouteParentStatus, gateways []gatewayWithListeners) []gatev1.RouteParentStatus {
	parentStatuses := make([]gatev1.RouteParentStatus, len(desiredParents))
	copy(parentStatuses, desiredParents)

	for _, currentParent := range currentParents {
		// Keep statuses added by other gateway controllers.
		if currentParent.ControllerName != controllerName {
			parentStatuses = append(parentStatuses, currentParent)
			continue
		}

		// TODO: Implement a mechanism to clean up old parentStatus for gateways that no instance manages.
		// Also, keep statuses from gateways managed by other Traefik instances.
		// We consider a status managed by the current instance when the parentRef targets one of the managed gateways.
		// A ListenerSet parentRef is managed when one of the managed gateways loads the listeners of that ListenerSet.
		// SectionName or Port is not used.
		owner := listenerOwner{
			Kind:      string(ptr.Deref(currentParent.ParentRef.Kind, kindGateway)),
			Namespace: string(ptr.Deref(currentParent.ParentRef.Namespace, gatev1.Namespace(routeNamespace))),
			Name:      string(currentParent.ParentRef.Name),
		}
		if gatewayForOwner(gateways, owner) == nil {
			parentStatuses = append(parentStatuses, currentParent)
		}
	}

	return parentStatuses
}

// gatewayConditionTypes and listenerConditionTypes are the condition types Traefik writes,
// conditions of any other type are owned by other controllers and must be preserved.
var (
	gatewayConditionTypes = []string{
		string(gatev1.GatewayConditionAccepted),
		string(gatev1.GatewayConditionProgrammed),
		string(gatev1.GatewayConditionResolvedRefs),
	}
	listenerConditionTypes = []string{
		string(gatev1.ListenerConditionAccepted),
		string(gatev1.ListenerConditionConflicted),
		string(gatev1.ListenerConditionProgrammed),
		string(gatev1.ListenerConditionResolvedRefs),
	}
)

// withoutForeignConditions returns the status without the Gateway and listener conditions not owned by Traefik.
func withoutForeignConditions(status gatev1.GatewayStatus) gatev1.GatewayStatus {
	filtered := status.DeepCopy()
	filtered.Conditions = slices.DeleteFunc(filtered.Conditions, func(condition metav1.Condition) bool {
		return !slices.Contains(gatewayConditionTypes, condition.Type)
	})

	for i := range filtered.Listeners {
		filtered.Listeners[i].Conditions = slices.DeleteFunc(filtered.Listeners[i].Conditions, func(condition metav1.Condition) bool {
			return !slices.Contains(listenerConditionTypes, condition.Type)
		})
	}

	return *filtered
}

// mergeGatewayStatus returns the desired status, with the Gateway and listener conditions
// of the current status that are not owned by Traefik.
func mergeGatewayStatus(current, desired gatev1.GatewayStatus) gatev1.GatewayStatus {
	status := desired.DeepCopy()
	for _, condition := range current.Conditions {
		if !slices.Contains(gatewayConditionTypes, condition.Type) {
			status.Conditions = append(status.Conditions, condition)
		}
	}

	for i, listener := range status.Listeners {
		idx := slices.IndexFunc(current.Listeners, func(l gatev1.ListenerStatus) bool {
			return l.Name == listener.Name
		})
		if idx < 0 {
			continue
		}

		for _, condition := range current.Listeners[idx].Conditions {
			if !slices.Contains(listenerConditionTypes, condition.Type) {
				status.Listeners[i].Conditions = append(status.Listeners[i].Conditions, condition)
			}
		}
	}

	return *status
}

func gatewayStatusEqual(statusA, statusB gatev1.GatewayStatus) bool {
	return reflect.DeepEqual(statusA.Addresses, statusB.Addresses) &&
		reflect.DeepEqual(statusA.AttachedListenerSets, statusB.AttachedListenerSets) &&
		listenersStatusEqual(statusA.Listeners, statusB.Listeners) &&
		conditionsEqual(statusA.Conditions, statusB.Conditions)
}

func policyAncestorStatusesEqual(policyAncestorStatusesA, policyAncestorStatusesB []gatev1.PolicyAncestorStatus) bool {
	if len(policyAncestorStatusesA) != len(policyAncestorStatusesB) {
		return false
	}

	for _, sA := range policyAncestorStatusesA {
		if !slices.ContainsFunc(policyAncestorStatusesB, func(sB gatev1.PolicyAncestorStatus) bool {
			return policyAncestorStatusEqual(sB, sA)
		}) {
			return false
		}
	}

	for _, sB := range policyAncestorStatusesB {
		if !slices.ContainsFunc(policyAncestorStatusesA, func(sA gatev1.PolicyAncestorStatus) bool {
			return policyAncestorStatusEqual(sA, sB)
		}) {
			return false
		}
	}

	return true
}

func policyAncestorStatusEqual(sA, sB gatev1.PolicyAncestorStatus) bool {
	return sA.ControllerName == sB.ControllerName &&
		reflect.DeepEqual(sA.AncestorRef, sB.AncestorRef) &&
		conditionsEqual(sA.Conditions, sB.Conditions)
}

func routeParentStatusesEqual(routeParentStatusesA, routeParentStatusesB []gatev1.RouteParentStatus) bool {
	if len(routeParentStatusesA) != len(routeParentStatusesB) {
		return false
	}

	for _, sA := range routeParentStatusesA {
		if !slices.ContainsFunc(routeParentStatusesB, func(sB gatev1.RouteParentStatus) bool {
			return routeParentStatusEqual(sB, sA)
		}) {
			return false
		}
	}

	for _, sB := range routeParentStatusesB {
		if !slices.ContainsFunc(routeParentStatusesA, func(sA gatev1.RouteParentStatus) bool {
			return routeParentStatusEqual(sA, sB)
		}) {
			return false
		}
	}

	return true
}

func routeParentStatusEqual(sA, sB gatev1.RouteParentStatus) bool {
	return sA.ControllerName == sB.ControllerName &&
		reflect.DeepEqual(sA.ParentRef, sB.ParentRef) &&
		conditionsEqual(sA.Conditions, sB.Conditions)
}

func listenersStatusEqual(listenerA, listenerB []gatev1.ListenerStatus) bool {
	return slices.EqualFunc(listenerA, listenerB, func(lA gatev1.ListenerStatus, lB gatev1.ListenerStatus) bool {
		return lA.Name == lB.Name &&
			lA.AttachedRoutes == lB.AttachedRoutes &&
			conditionsEqual(lA.Conditions, lB.Conditions)
	})
}

func conditionsEqual(conditionsA, conditionsB []metav1.Condition) bool {
	return slices.EqualFunc(conditionsA, conditionsB, func(cA metav1.Condition, cB metav1.Condition) bool {
		return cA.Type == cB.Type &&
			cA.Reason == cB.Reason &&
			cA.Status == cB.Status &&
			cA.Message == cB.Message &&
			cA.ObservedGeneration == cB.ObservedGeneration
	})
}

func listenerSetStatusEqual(statusA, statusB gatev1.ListenerSetStatus) bool {
	return listenerEntryStatusesEqual(statusA.Listeners, statusB.Listeners) &&
		conditionsEqual(statusA.Conditions, statusB.Conditions)
}

func listenerEntryStatusesEqual(listenersA, listenersB []gatev1.ListenerEntryStatus) bool {
	return slices.EqualFunc(listenersA, listenersB, func(lA gatev1.ListenerEntryStatus, lB gatev1.ListenerEntryStatus) bool {
		return lA.Name == lB.Name &&
			lA.AttachedRoutes == lB.AttachedRoutes &&
			conditionsEqual(lA.Conditions, lB.Conditions)
	})
}

```

#### pkg/client_test.go

```go
package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func Test_gatewayStatusEquals(t *testing.T) {
	testCases := []struct {
		desc     string
		statusA  gatev1.GatewayStatus
		statusB  gatev1.GatewayStatus
		expected bool
	}{
		{
			desc:     "Empty",
			statusA:  gatev1.GatewayStatus{},
			statusB:  gatev1.GatewayStatus{},
			expected: true,
		},
		{
			desc: "Same status",
			statusA: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{
					{
						Type:   "foobar",
						Reason: "foobar",
					},
				},
				Listeners: []gatev1.ListenerStatus{
					{
						Name: "foo",
						Conditions: []metav1.Condition{
							{
								Type:   "foobar",
								Reason: "foobar",
							},
						},
					},
				},
			},
			statusB: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{
					{
						Type:   "foobar",
						Reason: "foobar",
					},
				},
				Listeners: []gatev1.ListenerStatus{
					{
						Name: "foo",
						Conditions: []metav1.Condition{
							{
								Type:   "foobar",
								Reason: "foobar",
							},
						},
					},
				},
			},
			expected: true,
		},
		{
			desc: "Listeners length not equal",
			statusA: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{},
			},
			statusB: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{},
				},
			},
			expected: false,
		},
		{
			desc: "Gateway conditions length not equal",
			statusA: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{},
			},
			statusB: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{
					{},
				},
			},
			expected: false,
		},
		{
			desc: "Gateway conditions different types",
			statusA: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{
					{
						Type: "foobar",
					},
				},
			},
			statusB: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{
					{
						Type: "foobir",
					},
				},
			},
			expected: false,
		},
		{
			desc: "Gateway conditions same types but different reason",
			statusA: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{
					{
						Type: "foobar",
					},
				},
			},
			statusB: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{
					{
						Type:   "foobar",
						Reason: "Another reason",
					},
				},
			},
			expected: false,
		},
		{
			desc: "Gateway listeners conditions length",
			statusA: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Name:       "foo",
						Conditions: []metav1.Condition{},
					},
				},
			},
			statusB: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Name: "foo",
						Conditions: []metav1.Condition{
							{},
						},
					},
				},
			},
			expected: false,
		},
		{
			desc: "Gateway listeners conditions same types but different status",
			statusA: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Conditions: []metav1.Condition{
							{
								Type: "foobar",
							},
						},
					},
				},
			},
			statusB: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Conditions: []metav1.Condition{
							{
								Type:   "foobar",
								Status: "Another status",
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			desc: "Gateway listeners conditions same types but different message",
			statusA: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Conditions: []metav1.Condition{
							{
								Type: "foobar",
							},
						},
					},
				},
			},
			statusB: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Conditions: []metav1.Condition{
							{
								Type:    "foobar",
								Message: "Another status",
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			desc: "Gateway listeners conditions same types/reason but different names",
			statusA: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Name: "foo",
						Conditions: []metav1.Condition{
							{
								Type:   "foobar",
								Reason: "foobar",
							},
						},
					},
				},
			},
			statusB: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Name: "bar",
						Conditions: []metav1.Condition{
							{
								Type:   "foobar",
								Reason: "foobar",
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			desc: "Gateway listeners with same conditions but different number of attached routes",
			statusA: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Name:           "foo",
						AttachedRoutes: 1,
						Conditions: []metav1.Condition{
							{
								Type:   "foobar",
								Reason: "foobar",
							},
						},
					},
				},
			},
			statusB: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{
					{
						Name:           "foo",
						AttachedRoutes: 2,
						Conditions: []metav1.Condition{
							{
								Type:   "foobar",
								Reason: "foobar",
							},
						},
					},
				},
			},
			expected: false,
		},
		{
			desc: "Gateway with different attached listener sets count",
			statusA: gatev1.GatewayStatus{
				AttachedListenerSets: new(int32(1)),
			},
			statusB: gatev1.GatewayStatus{
				AttachedListenerSets: new(int32(2)),
			},
			expected: false,
		},
		{
			desc: "Gateway with same attached listener sets count",
			statusA: gatev1.GatewayStatus{
				AttachedListenerSets: new(int32(1)),
			},
			statusB: gatev1.GatewayStatus{
				AttachedListenerSets: new(int32(1)),
			},
			expected: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			result := gatewayStatusEqual(test.statusA, test.statusB)

			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_listenerSetStatusEqual(t *testing.T) {
	testCases := []struct {
		desc     string
		statusA  gatev1.ListenerSetStatus
		statusB  gatev1.ListenerSetStatus
		expected bool
	}{
		{
			desc:     "Empty",
			statusA:  gatev1.ListenerSetStatus{},
			statusB:  gatev1.ListenerSetStatus{},
			expected: true,
		},
		{
			desc: "Same status",
			statusA: gatev1.ListenerSetStatus{
				Conditions: []metav1.Condition{{Type: "Accepted", Reason: "Accepted"}},
				Listeners: []gatev1.ListenerEntryStatus{{
					Name:           "foo",
					AttachedRoutes: 1,
					Conditions:     []metav1.Condition{{Type: "Programmed", Reason: "Programmed"}},
				}},
			},
			statusB: gatev1.ListenerSetStatus{
				Conditions: []metav1.Condition{{Type: "Accepted", Reason: "Accepted"}},
				Listeners: []gatev1.ListenerEntryStatus{{
					Name:           "foo",
					AttachedRoutes: 1,
					Conditions:     []metav1.Condition{{Type: "Programmed", Reason: "Programmed"}},
				}},
			},
			expected: true,
		},
		{
			desc: "Top-level conditions differ",
			statusA: gatev1.ListenerSetStatus{
				Conditions: []metav1.Condition{{Type: "Accepted", Reason: "Accepted"}},
			},
			statusB: gatev1.ListenerSetStatus{
				Conditions: []metav1.Condition{{Type: "Accepted", Reason: "NotAllowed"}},
			},
			expected: false,
		},
		{
			desc: "Listener entries length differ",
			statusA: gatev1.ListenerSetStatus{
				Listeners: []gatev1.ListenerEntryStatus{{Name: "foo"}},
			},
			statusB:  gatev1.ListenerSetStatus{},
			expected: false,
		},
		{
			desc: "Listener entry names differ",
			statusA: gatev1.ListenerSetStatus{
				Listeners: []gatev1.ListenerEntryStatus{{Name: "foo"}},
			},
			statusB: gatev1.ListenerSetStatus{
				Listeners: []gatev1.ListenerEntryStatus{{Name: "bar"}},
			},
			expected: false,
		},
		{
			desc: "Listener entry attached routes differ",
			statusA: gatev1.ListenerSetStatus{
				Listeners: []gatev1.ListenerEntryStatus{{Name: "foo", AttachedRoutes: 1}},
			},
			statusB: gatev1.ListenerSetStatus{
				Listeners: []gatev1.ListenerEntryStatus{{Name: "foo", AttachedRoutes: 2}},
			},
			expected: false,
		},
		{
			desc: "Listener entry conditions differ",
			statusA: gatev1.ListenerSetStatus{
				Listeners: []gatev1.ListenerEntryStatus{{
					Name:       "foo",
					Conditions: []metav1.Condition{{Type: "Programmed", Status: metav1.ConditionTrue}},
				}},
			},
			statusB: gatev1.ListenerSetStatus{
				Listeners: []gatev1.ListenerEntryStatus{{
					Name:       "foo",
					Conditions: []metav1.Condition{{Type: "Programmed", Status: metav1.ConditionFalse}},
				}},
			},
			expected: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			result := listenerSetStatusEqual(test.statusA, test.statusB)

			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_mergeGatewayStatus(t *testing.T) {
	accepted := metav1.Condition{Type: string(gatev1.GatewayConditionAccepted), Status: metav1.ConditionTrue}
	staleAccepted := metav1.Condition{Type: string(gatev1.GatewayConditionAccepted), Status: metav1.ConditionFalse}
	resolvedRefs := metav1.Condition{Type: string(gatev1.GatewayConditionResolvedRefs), Status: metav1.ConditionTrue}
	staleResolvedRefs := metav1.Condition{Type: string(gatev1.GatewayConditionResolvedRefs), Status: metav1.ConditionFalse}
	listenerAccepted := metav1.Condition{Type: string(gatev1.ListenerConditionAccepted), Status: metav1.ConditionTrue}
	staleConflicted := metav1.Condition{Type: string(gatev1.ListenerConditionConflicted), Status: metav1.ConditionTrue}
	foreign := metav1.Condition{Type: "example.com/Foreign", Status: metav1.ConditionTrue}

	desired := gatev1.GatewayStatus{
		Conditions: []metav1.Condition{accepted, resolvedRefs},
		Listeners: []gatev1.ListenerStatus{{
			Name:       "http",
			Conditions: []metav1.Condition{listenerAccepted},
		}},
	}

	testCases := []struct {
		desc     string
		current  gatev1.GatewayStatus
		expected gatev1.GatewayStatus
	}{
		{
			desc:     "no current status",
			expected: desired,
		},
		{
			desc: "keeps foreign Gateway condition",
			current: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{staleAccepted, foreign},
			},
			expected: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{accepted, resolvedRefs, foreign},
				Listeners:  desired.Listeners,
			},
		},
		{
			desc: "replaces stale ResolvedRefs Gateway condition",
			current: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{staleResolvedRefs},
			},
			expected: desired,
		},
		{
			desc: "keeps foreign listener condition",
			current: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{{
					Name:       "http",
					Conditions: []metav1.Condition{foreign},
				}},
			},
			expected: gatev1.GatewayStatus{
				Conditions: desired.Conditions,
				Listeners: []gatev1.ListenerStatus{{
					Name:       "http",
					Conditions: []metav1.Condition{listenerAccepted, foreign},
				}},
			},
		},
		{
			desc: "drops owned listener condition no longer desired",
			current: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{{
					Name:       "http",
					Conditions: []metav1.Condition{listenerAccepted, staleConflicted},
				}},
			},
			expected: desired,
		},
		{
			desc: "drops status of listener no longer desired",
			current: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{{
					Name:       "removed",
					Conditions: []metav1.Condition{foreign},
				}},
			},
			expected: desired,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			result := mergeGatewayStatus(test.current, desired)

			assert.Equal(t, test.expected, result)
		})
	}
}

func Test_withoutForeignConditions(t *testing.T) {
	accepted := metav1.Condition{Type: string(gatev1.GatewayConditionAccepted), Status: metav1.ConditionTrue}
	resolvedRefs := metav1.Condition{Type: string(gatev1.GatewayConditionResolvedRefs), Status: metav1.ConditionTrue}
	listenerAccepted := metav1.Condition{Type: string(gatev1.ListenerConditionAccepted), Status: metav1.ConditionTrue}
	foreign := metav1.Condition{Type: "example.com/Foreign", Status: metav1.ConditionTrue}

	testCases := []struct {
		desc     string
		status   gatev1.GatewayStatus
		expected gatev1.GatewayStatus
	}{
		{
			desc: "keeps owned conditions",
			status: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{accepted, resolvedRefs},
				Listeners: []gatev1.ListenerStatus{{
					Name:       "http",
					Conditions: []metav1.Condition{listenerAccepted},
				}},
			},
			expected: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{accepted, resolvedRefs},
				Listeners: []gatev1.ListenerStatus{{
					Name:       "http",
					Conditions: []metav1.Condition{listenerAccepted},
				}},
			},
		},
		{
			desc: "drops foreign Gateway condition",
			status: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{foreign, accepted},
			},
			expected: gatev1.GatewayStatus{
				Conditions: []metav1.Condition{accepted},
			},
		},
		{
			desc: "drops foreign listener condition",
			status: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{{
					Name:       "http",
					Conditions: []metav1.Condition{foreign, listenerAccepted},
				}},
			},
			expected: gatev1.GatewayStatus{
				Listeners: []gatev1.ListenerStatus{{
					Name:       "http",
					Conditions: []metav1.Condition{listenerAccepted},
				}},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			result := withoutForeignConditions(test.status)

			assert.Equal(t, test.expected, result)
		})
	}
}

```

#### pkg/features.go

```go
package gateway

import (
	"sync"

	"k8s.io/apimachinery/pkg/util/sets"
	"sigs.k8s.io/gateway-api/pkg/features"
)

var SupportedFeatures = sync.OnceValue(func() []features.FeatureName {
	featureSet := sets.New[features.Feature]().
		Insert(features.GatewayCoreFeatures.UnsortedList()...).
		Insert(features.GatewayExtendedFeatures.Intersection(extendedGatewayFeatures()).UnsortedList()...).
		Insert(features.HTTPRouteCoreFeatures.UnsortedList()...).
		Insert(features.HTTPRouteExtendedFeatures.Intersection(extendedHTTPRouteFeatures()).UnsortedList()...).
		Insert(features.ReferenceGrantCoreFeatures.UnsortedList()...).
		Insert(features.BackendTLSPolicyCoreFeatures.UnsortedList()...).
		Insert(features.BackendTLSPolicyExtendedFeatures.Intersection(extendedBackendTLSPolicyFeatures()).UnsortedList()...).
		Insert(features.GRPCRouteCoreFeatures.UnsortedList()...).
		Insert(features.GRPCRouteExtendedFeatures.Intersection(extendedGRPCRouteFeatures()).UnsortedList()...).
		Insert(features.TLSRouteCoreFeatures.UnsortedList()...).
		Insert(features.TLSRouteExtendedFeatures.Intersection(extendedTLSRouteFeatures()).UnsortedList()...).
		Insert(features.TCPRouteFeature)

	featureNames := make([]features.FeatureName, 0, featureSet.Len())
	for f := range featureSet {
		featureNames = append(featureNames, f.Name)
	}
	return featureNames
})

// extendedGatewayFeatures returns the supported extended Gateway features.
func extendedGatewayFeatures() sets.Set[features.Feature] {
	return sets.New(
		features.GatewayPort8080Feature,
		features.GatewayHTTPListenerIsolationFeature,
		features.GatewayHTTPSListenerDetectMisdirectedRequestsFeature,
		features.ListenerSetFeature,
		features.GatewayFrontendClientCertificateValidationFeature,
		features.GatewayFrontendClientCertificateValidationInsecureFallbackFeature,
		features.GatewayBackendClientCertificateFeature,
	)
}

// extendedTLSRouteFeatures returns the supported extended TLS Route features.
func extendedTLSRouteFeatures() sets.Set[features.Feature] {
	return sets.New(
		features.TLSRouteModeTerminateFeature,
		features.TLSRouteModeMixedFeature,
	)
}

// extendedHTTPRouteFeatures returns the supported extended HTTP Route features.
func extendedHTTPRouteFeatures() sets.Set[features.Feature] {
	return sets.New(
		features.HTTPRouteQueryParamMatchingFeature,
		features.HTTPRouteMethodMatchingFeature,
		features.HTTPRoutePathRedirectFeature,
		features.HTTPRoutePortRedirectFeature,
		features.HTTPRouteSchemeRedirectFeature,
		features.HTTPRoute303RedirectStatusCodeFeature,
		features.HTTPRoute307RedirectStatusCodeFeature,
		features.HTTPRoute308RedirectStatusCodeFeature,
		features.HTTPRouteHostRewriteFeature,
		features.HTTPRoutePathRewriteFeature,
		features.HTTPRouteResponseHeaderModificationFeature,
		features.HTTPRouteBackendProtocolH2CFeature,
		features.HTTPRouteBackendProtocolWebSocketFeature,
		features.HTTPRouteDestinationPortMatchingFeature,
		features.HTTPRouteBackendRequestHeaderModificationFeature,
		features.HTTPRouteNamedRouteRule,
		features.HTTPRouteParentRefPortFeature,
		features.HTTPRouteCORS,
		features.HTTPRouteRequestTimeoutFeature,
		features.HTTPRouteBackendTimeoutFeature,
		features.HTTPRouteRetryCodesFeature,
		features.HTTPRouteRetryConnectionErrorFeature,
		features.HTTPRouteRequestMirrorFeature,
		features.HTTPRouteRequestMultipleMirrorsFeature,
		features.HTTPRouteRequestPercentageMirrorFeature,
	)
}

// extendedGRPCRouteFeatures returns the supported extended GRPC Route features.
func extendedGRPCRouteFeatures() sets.Set[features.Feature] {
	return sets.New(
		features.GRPCRouteNamedRouteRule,
	)
}

// extendedBackendTLSPolicyFeatures returns the supported extended BackendTLSPolicy features.
func extendedBackendTLSPolicyFeatures() sets.Set[features.Feature] {
	return sets.New(
		features.BackendTLSPolicySanValidationFeature,
	)
}

```

#### pkg/grpcroute.go

```go
package gateway

import (
	"context"
	"errors"
	"fmt"
	"net"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"google.golang.org/grpc/codes"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (p *Provider) loadGRPCRoute(ctx context.Context, gateways []gatewayWithListeners, route *gatev1.GRPCRoute, conf *dynamic.Configuration, attachedRoutes attachedRoutes, servedRules servedRules, statusReport *statusReport) {
	logger := log.Ctx(ctx).With().
		Str("grpc_route", route.Name).
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

			if accepted && !allowRoute(listener, route.Namespace, kindGRPCRoute) {
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
			if accepted && attachedRoutes.Conflicts(ref, kindGRPCRoute, hostnames) {
				if acceptedCondition.Status == metav1.ConditionFalse {
					acceptedCondition.Reason = string(routeReasonHostnameConflict)
				}
				accepted = false
			}

			if accepted {
				// Gateway listener should have AttachedRoutes set even when Gateway has unresolved refs.
				listener.Status.AttachedRoutes++

				attachedRoutes.Record(ref, kindGRPCRoute, hostnames)
			}

			// The ResolvedRefs condition must be reported for every parentRef,
			// even when the route does not attach to the listener.
			routerConfs, condition := p.loadGRPCRouteConfiguration(logger.WithContext(ctx), match.GatewayName, match.GatewayNamespace, listener, route, hostnames, statusReport)
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

		statusReport.RecordGRPCRouteStatus(ktypes.NamespacedName{Namespace: route.Namespace, Name: route.Name}, gatev1.RouteParentStatus{
			ParentRef:      match.ParentRef,
			ControllerName: controllerName,
			Conditions:     parentStatusConditions,
		})
	}
}

// grpcRouteRouter is a router of a GRPCRoute, with the configuration of its resources.
type grpcRouteRouter struct {
	Name string
	Conf *dynamic.Configuration
}

func (p *Provider) loadGRPCRouteConfiguration(ctx context.Context, gatewayName, gatewayNamespace string, listener gatewayListener, route *gatev1.GRPCRoute, hostnames []gatev1.Hostname, statusReport *statusReport) ([]grpcRouteRouter, metav1.Condition) {
	var routers []grpcRouteRouter

	condition := metav1.Condition{
		Type:               string(gatev1.RouteConditionResolvedRefs),
		Status:             metav1.ConditionTrue,
		ObservedGeneration: route.Generation,
		LastTransitionTime: metav1.Now(),
		Reason:             string(gatev1.RouteConditionResolvedRefs),
	}

	for ri, routeRule := range route.Spec.Rules {
		matches := routeRule.Matches
		if len(matches) == 0 {
			matches = []gatev1.GRPCRouteMatch{{}}
		}

		for _, match := range matches {
			rule, priority := buildGRPCMatchRule(hostnames, match)
			router := dynamic.Router{
				// "default" stands for the default rule syntax in Traefik v3, i.e. the v3 syntax.
				RuleSyntax: "default",
				Rule:       rule,
				Priority:   priority,
				ParentRefs: listener.RouterNames,
			}

			routerName := makeRouterName(strings.ToLower(kindGRPCRoute), rule, route.Namespace, route.Name, gatewayNamespace, gatewayName, listener.EPName, ri)

			routerConf := &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           make(map[string]*dynamic.Router),
					Middlewares:       make(map[string]*dynamic.Middleware),
					Services:          make(map[string]*dynamic.Service),
					ServersTransports: make(map[string]*dynamic.ServersTransport),
				},
			}

			var err error
			router.Middlewares, err = p.loadGRPCMiddlewares(routerConf, route.Namespace, routerName, routeRule.Filters)
			switch {
			case err != nil:
				log.Ctx(ctx).Error().Err(err).Msg("Unable to load GRPC route filters")

				errWrrName := routerName + "-err-wrr"
				routerConf.HTTP.Services[errWrrName] = &dynamic.Service{
					Weighted: &dynamic.WeightedRoundRobin{
						Services: []dynamic.WRRService{
							{
								Name: "invalid-grpcroute-filter",
								GRPCStatus: &dynamic.GRPCStatus{
									Code: codes.Unavailable,
									Msg:  "Service Unavailable",
								},
								Weight: new(1),
							},
						},
					},
				}
				router.Service = errWrrName

			default:
				var serviceCondition *metav1.Condition
				router.Service, serviceCondition = p.loadGRPCService(gatewayName, gatewayNamespace, listener, routerConf, routerName, routeRule, route, statusReport)
				if serviceCondition != nil {
					condition = *serviceCondition
				}
			}

			routerConf.HTTP.Routers[routerName] = &router
			routers = append(routers, grpcRouteRouter{Name: routerName, Conf: routerConf})
		}
	}

	return routers, condition
}

func (p *Provider) loadGRPCService(gatewayName, gatewayNamespace string, listener gatewayListener, conf *dynamic.Configuration, routerName string, routeRule gatev1.GRPCRouteRule, route *gatev1.GRPCRoute, statusReport *statusReport) (string, *metav1.Condition) {
	name := routerName + "-wrr"
	if _, ok := conf.HTTP.Services[name]; ok {
		return name, nil
	}

	var wrr dynamic.WeightedRoundRobin
	var condition *metav1.Condition
	for bi, backendRef := range routeRule.BackendRefs {
		svcName, svc, errCondition := p.loadGRPCBackendRef(gatewayName, gatewayNamespace, listener, conf, routerName, route, bi, backendRef, statusReport)
		weight := new(int(ptr.Deref(backendRef.Weight, 1)))
		if errCondition != nil {
			condition = errCondition
			wrr.Services = append(wrr.Services, dynamic.WRRService{
				Name: svcName,
				GRPCStatus: &dynamic.GRPCStatus{
					Code: codes.Unavailable,
					Msg:  "Service Unavailable",
				},
				Weight: weight,
			})
			continue
		}

		if svc != nil {
			conf.HTTP.Services[svcName] = svc
		}

		wrr.Services = append(wrr.Services, dynamic.WRRService{
			Name:   svcName,
			Weight: weight,
		})
	}

	conf.HTTP.Services[name] = &dynamic.Service{Weighted: &wrr}
	return name, condition
}

func (p *Provider) loadGRPCBackendRef(gatewayName, gatewayNamespace string, listener gatewayListener, conf *dynamic.Configuration, routerName string, route *gatev1.GRPCRoute, backendIndex int, backendRef gatev1.GRPCBackendRef, statusReport *statusReport) (string, *dynamic.Service, *metav1.Condition) {
	kind := ptr.Deref(backendRef.Kind, kindService)

	group := groupCore
	if backendRef.Group != nil && *backendRef.Group != "" {
		group = string(*backendRef.Group)
	}

	namespace := route.Namespace
	if backendRef.Namespace != nil && *backendRef.Namespace != "" {
		namespace = string(*backendRef.Namespace)
	}

	serviceName := fmt.Sprintf("%s-svc-%s-%s-%d", routerName, namespace, backendRef.Name, backendIndex)

	if group != groupCore || kind != kindService {
		return serviceName, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonInvalidKind),
			Message:            fmt.Sprintf("Cannot load GRPCBackendRef %s/%s/%s/%s: only Kubernetes services are supported", group, kind, namespace, backendRef.Name),
		}
	}

	if err := p.isReferenceGranted(kindGRPCRoute, route.Namespace, group, string(kind), string(backendRef.Name), namespace); err != nil {
		return serviceName, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot load GRPCBackendRef %s/%s/%s/%s: %s", group, kind, namespace, backendRef.Name, err),
		}
	}

	port := ptr.Deref(backendRef.Port, gatev1.PortNumber(0))
	if port == 0 {
		return serviceName, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonUnsupportedProtocol),
			Message:            fmt.Sprintf("Cannot load GRPCBackendRef %s/%s/%s/%s: port is required", group, kind, namespace, backendRef.Name),
		}
	}

	lb, st, errCondition := p.loadGRPCServers(gatewayName, gatewayNamespace, namespace, route, backendRef, listener, statusReport)
	if errCondition != nil {
		return serviceName, nil, errCondition
	}

	if st != nil {
		lb.ServersTransport = serviceName
		conf.HTTP.ServersTransports[serviceName] = st
	}

	return serviceName, &dynamic.Service{LoadBalancer: lb}, nil
}

func (p *Provider) loadGRPCMiddlewares(conf *dynamic.Configuration, namespace, routerName string, filters []gatev1.GRPCRouteFilter) ([]string, error) {
	type namedMiddleware struct {
		Name   string
		Config *dynamic.Middleware
	}

	var middlewares []namedMiddleware
	for i, filter := range filters {
		name := fmt.Sprintf("%s-%s-%d", routerName, strings.ToLower(string(filter.Type)), i)
		switch filter.Type {
		case gatev1.GRPCRouteFilterRequestHeaderModifier:
			middlewares = append(middlewares, namedMiddleware{
				name,
				createRequestHeaderModifier(filter.RequestHeaderModifier),
			})

		case gatev1.GRPCRouteFilterExtensionRef:
			name, middleware, err := p.loadHTTPRouteFilterExtensionRef(namespace, filter.ExtensionRef)
			if err != nil {
				return nil, fmt.Errorf("loading ExtensionRef filter %s: %w", filter.Type, err)
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

func (p *Provider) loadGRPCServers(gatewayName, gatewayNamespace, namespace string, route *gatev1.GRPCRoute, backendRef gatev1.GRPCBackendRef, listener gatewayListener, statusReport *statusReport) (*dynamic.ServersLoadBalancer, *dynamic.ServersTransport, *metav1.Condition) {
	backendAddresses, svcPort, err := p.getBackendAddresses(namespace, backendRef.BackendRef)
	if err != nil {
		return nil, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonBackendNotFound),
			Message:            fmt.Sprintf("Cannot load GRPCBackendRef %s/%s: %s", namespace, backendRef.Name, err),
		}
	}

	backendTLSPolicies, err := p.client.ListBackendTLSPoliciesForService(namespace, string(backendRef.Name))
	if err != nil {
		return nil, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot list BackendTLSPolicies for Service %s/%s: %s", namespace, string(backendRef.Name), err),
		}
	}

	// Sort BackendTLSPolicies by creation timestamp, then by name to match the BackendTLSPolicy requirements.
	slices.SortStableFunc(backendTLSPolicies, func(a, b *gatev1.BackendTLSPolicy) int {
		cmpTime := a.CreationTimestamp.Time.Compare(b.CreationTimestamp.Time)
		if cmpTime == 0 {
			return strings.Compare(a.Name, b.Name)
		}
		return cmpTime
	})

	var serversTransport *dynamic.ServersTransport
	for _, policy := range backendTLSPolicies {
		for _, targetRef := range policy.Spec.TargetRefs {
			// Skip targetRefs that doesn't match the backendRef,
			// since a BackendTLSPolicy can select multiple services.
			if targetRef.Name != backendRef.Name {
				continue
			}
			// Skip the targetRef if the sectionName doesn't match the backendRef port.
			if targetRef.SectionName != nil && svcPort.Name != string(*targetRef.SectionName) {
				continue
			}

			policyAncestorStatus := gatev1.PolicyAncestorStatus{
				AncestorRef: gatev1.ParentReference{
					Group:       new(gatev1.Group(groupGateway)),
					Kind:        new(gatev1.Kind(kindGateway)),
					Namespace:   new(gatev1.Namespace(gatewayNamespace)),
					Name:        gatev1.ObjectName(gatewayName),
					SectionName: new(gatev1.SectionName(listener.SectionName)),
				},
				ControllerName: controllerName,
			}

			// Multiple BackendTLSPolicies can match the same service port, meaning that there is a conflict.
			if serversTransport != nil {
				policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions,
					metav1.Condition{
						Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
						Status:             metav1.ConditionFalse,
						ObservedGeneration: policy.Generation,
						LastTransitionTime: metav1.Now(),
						Reason:             string(gatev1.BackendTLSPolicyReasonResolvedRefs),
					},
					metav1.Condition{
						Type:               string(gatev1.PolicyConditionAccepted),
						Status:             metav1.ConditionFalse,
						ObservedGeneration: policy.Generation,
						LastTransitionTime: metav1.Now(),
						Reason:             string(gatev1.PolicyReasonConflicted),
					},
				)

				statusReport.RecordBackendTLSPolicyStatus(ktypes.NamespacedName{Namespace: policy.Namespace, Name: policy.Name}, policyAncestorStatus)

				continue
			}

			var resolvedRefCondition metav1.Condition
			serversTransport, resolvedRefCondition = p.loadServersTransport(namespace, policy, listener.BackendClientCertificate)

			policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions, resolvedRefCondition)
			if resolvedRefCondition.Status == metav1.ConditionFalse {
				policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions, metav1.Condition{
					Type:               string(gatev1.PolicyConditionAccepted),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.BackendTLSPolicyReasonNoValidCACertificate),
				})
			} else {
				policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions, metav1.Condition{
					Type:               string(gatev1.PolicyConditionAccepted),
					Status:             metav1.ConditionTrue,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.PolicyReasonAccepted),
				})
			}

			statusReport.RecordBackendTLSPolicyStatus(ktypes.NamespacedName{Namespace: policy.Namespace, Name: policy.Name}, policyAncestorStatus)

			// When something went wrong during the loading of a ServersTransport,
			// we stop here and return a route condition error.
			if resolvedRefCondition.Status == metav1.ConditionFalse {
				return nil, nil, &metav1.Condition{
					Type:               string(gatev1.RouteConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: route.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.RouteReasonRefNotPermitted),
					Message:            fmt.Sprintf("Cannot apply BackendTLSPolicy for Service %s/%s: %s", namespace, string(backendRef.Name), resolvedRefCondition.Message),
				}
			}
		}
	}

	// If a ServersTransport is set, it means a BackendTLSPolicy matched the service port, and we can safely assume the protocol is HTTPS.
	// When no ServersTransport is set, we need to determine the protocol based on the service port.
	protocol := "https"
	if serversTransport == nil {
		protocol, err = getGRPCServiceProtocol(svcPort)
		if err != nil {
			return nil, nil, &metav1.Condition{
				Type:               string(gatev1.RouteConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: route.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.RouteReasonUnsupportedProtocol),
				Message:            fmt.Sprintf("Cannot load GRPCBackendRef %s/%s: %s", namespace, backendRef.Name, err),
			}
		}
	}

	lb := &dynamic.ServersLoadBalancer{}
	lb.SetDefaults()

	for _, ba := range backendAddresses {
		lb.Servers = append(lb.Servers, dynamic.Server{
			URL: fmt.Sprintf("%s://%s", protocol, net.JoinHostPort(ba.IP, strconv.Itoa(int(ba.Port)))),
		})
	}

	return lb, serversTransport, nil
}

func buildGRPCMatchRule(hostnames []gatev1.Hostname, match gatev1.GRPCRouteMatch) (string, int) {
	var matchRules []string

	methodRule := buildGRPCMethodRule(match.Method)
	matchRules = append(matchRules, methodRule)

	headerRules := buildGRPCHeaderRules(match.Headers)
	matchRules = append(matchRules, headerRules...)

	matchRulesStr := strings.Join(matchRules, " && ")

	hostRule, priority := buildHostRule(hostnames)
	if hostRule == "" {
		return matchRulesStr, len(matchRulesStr)
	}
	return hostRule + " && " + matchRulesStr, priority + len(matchRulesStr)
}

func buildGRPCMethodRule(method *gatev1.GRPCMethodMatch) string {
	if method == nil {
		return `PathPrefix("/")`
	}

	isExact := method.Type == nil || *method.Type == gatev1.GRPCMethodMatchExact

	sExpr := "[^/]+"
	if s := ptr.Deref(method.Service, ""); s != "" {
		if isExact {
			sExpr = regexp.QuoteMeta(s)
		} else {
			sExpr = s
		}
	}

	mExpr := "[^/]+"
	if m := ptr.Deref(method.Method, ""); m != "" {
		if isExact {
			mExpr = regexp.QuoteMeta(m)
		} else {
			mExpr = m
		}
	}

	return fmt.Sprintf("PathRegexp(%q)", fmt.Sprintf("/%s/%s", sExpr, mExpr))
}

func buildGRPCHeaderRules(headers []gatev1.GRPCHeaderMatch) []string {
	var rules []string
	for _, header := range headers {
		switch ptr.Deref(header.Type, gatev1.GRPCHeaderMatchExact) {
		case gatev1.GRPCHeaderMatchExact:
			rules = append(rules, fmt.Sprintf("Header(%q,%q)", header.Name, header.Value))
		case gatev1.GRPCHeaderMatchRegularExpression:
			rules = append(rules, fmt.Sprintf("HeaderRegexp(%q,%q)", header.Name, header.Value))
		}
	}

	return rules
}

func getGRPCServiceProtocol(portSpec corev1.ServicePort) (string, error) {
	if portSpec.Protocol != corev1.ProtocolTCP {
		return "", errors.New("only TCP protocol is supported")
	}

	if portSpec.AppProtocol == nil {
		return schemeH2C, nil
	}

	switch ap := strings.ToLower(*portSpec.AppProtocol); ap {
	case appProtocolH2C:
		return schemeH2C, nil
	case appProtocolHTTPS:
		return schemeHTTPS, nil
	default:
		return "", fmt.Errorf("unsupported application protocol %s", ap)
	}
}

```

#### pkg/httproute.go

```go
package gateway

import (
	"context"
	"errors"
	"fmt"
	"maps"
	"net"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/provider"
	"github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
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
			routerConfs, condition := p.loadHTTPRouteConfiguration(logger.WithContext(ctx), match.GatewayName, match.GatewayNamespace, listener, route, hostnames, statusReport)
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

func (p *Provider) loadHTTPRouteConfiguration(ctx context.Context, gatewayName, gatewayNamespace string, listener gatewayListener, route *gatev1.HTTPRoute, hostnames []gatev1.Hostname, statusReport *statusReport) ([]httpRouteRouter, metav1.Condition) {
	var routers []httpRouteRouter

	condition := metav1.Condition{
		Type:               string(gatev1.RouteConditionResolvedRefs),
		Status:             metav1.ConditionTrue,
		ObservedGeneration: route.Generation,
		LastTransitionTime: metav1.Now(),
		Reason:             string(gatev1.RouteConditionResolvedRefs),
	}

	loadMirrorService := p.newMirrorServiceLoader(gatewayName, gatewayNamespace, listener, route, statusReport)

	for ri, routeRule := range route.Spec.Rules {
		var respondingTimeouts *dynamic.RouterRespondingTimeouts
		var backendRequestTimeout time.Duration
		if routeRule.Timeouts != nil {
			if routeRule.Timeouts.Request != nil {
				// The CRD CEL validation guarantees a valid GEP-2257 duration string.
				// Both timeouts count from the same instant, so they share a single deadline bounding the whole request.
				// A zero duration disables the timeout, lifting the entrypoint read and write deadlines, as GEP-1742 requires.
				if d, err := time.ParseDuration(string(*routeRule.Timeouts.Request)); err != nil {
					// TODO: Maybe it would be better to handle the error and put the routers in error.
					log.Ctx(ctx).Error().Err(err).Msg("Ignoring HTTPRoute timeouts.request: invalid duration")
				} else {
					respondingTimeouts = &dynamic.RouterRespondingTimeouts{
						ReadTimeout:  new(ptypes.Duration(d)),
						WriteTimeout: new(ptypes.Duration(d)),
					}
				}
			}

			if routeRule.Timeouts.BackendRequest != nil {
				// The CRD CEL validation guarantees a valid GEP-2257 duration string.
				if d, err := time.ParseDuration(string(*routeRule.Timeouts.BackendRequest)); err != nil {
					log.Ctx(ctx).Error().Err(err).Msg("Ignoring HTTPRoute timeouts.backendRequest: invalid duration")
				} else {
					backendRequestTimeout = d
				}
			}
		}

		var retry *dynamic.Middleware
		if routeRule.Retry != nil {
			retry = p.createRetry(ctx, routeRule.Retry)
		}

		for _, match := range routeRule.Matches {
			rule, priority := buildMatchRule(hostnames, match)
			router := dynamic.Router{
				// "default" stands for the default rule syntax in Traefik v3, i.e. the v3 syntax.
				RuleSyntax:         "default",
				Rule:               rule,
				Priority:           priority + len(route.Spec.Rules) - ri,
				ParentRefs:         listener.RouterNames,
				RespondingTimeouts: respondingTimeouts,
			}

			routerName := makeRouterName(strings.ToLower(kindHTTPRoute), rule, route.Namespace, route.Name, gatewayNamespace, gatewayName, listener.EPName, ri)

			routerConf := &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           make(map[string]*dynamic.Router),
					Middlewares:       make(map[string]*dynamic.Middleware),
					Services:          make(map[string]*dynamic.Service),
					ServersTransports: make(map[string]*dynamic.ServersTransport),
				},
			}

			var err error
			// TODO loadMiddlewares errors could change the condition.
			router.Middlewares, err = p.loadMiddlewares(ctx, routerConf, route.Namespace, routerName, routeRule.Filters, match.Path, loadMirrorService)
			// The retry middleware comes after the filters so that only the call to the backend is retried.
			if err == nil && retry != nil {
				retryName := routerName + "-retry"
				routerConf.HTTP.Middlewares[retryName] = retry
				router.Middlewares = append(router.Middlewares, retryName)
			}

			switch {
			case err != nil:
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

			case len(routeRule.BackendRefs) == 1 && isInternalService(routeRule.BackendRefs[0].BackendRef):
				if !isCrossProviderNamespaceAllowed(p.CrossProviderNamespaces, route.Namespace) {
					condition = metav1.Condition{
						Type:               string(gatev1.RouteConditionResolvedRefs),
						Status:             metav1.ConditionFalse,
						ObservedGeneration: route.Generation,
						LastTransitionTime: metav1.Now(),
						Reason:             string(gatev1.RouteReasonRefNotPermitted),
						Message:            fmt.Sprintf("Cannot load HTTPRoute BackendRef %s: internal service reference is not allowed: HTTPRoute namespace %q is not in crossProviderNamespaces", routeRule.BackendRefs[0].Name, route.Namespace),
					}
				} else {
					router.Service = string(routeRule.BackendRefs[0].Name)
				}

			default:
				var serviceCondition *metav1.Condition
				router.Service, serviceCondition = p.loadWRRService(ctx, gatewayName, gatewayNamespace, listener, routerConf, routerName, routeRule, route, match.Path, backendRequestTimeout, statusReport)
				if serviceCondition != nil {
					condition = *serviceCondition
				}
			}

			p.applyRouterTransform(ctx, &router, route)

			routerConf.HTTP.Routers[routerName] = &router
			routers = append(routers, httpRouteRouter{Name: routerName, Conf: routerConf})
		}
	}

	return routers, condition
}

func (p *Provider) loadWRRService(ctx context.Context, gatewayName, gatewayNamespace string, listener gatewayListener, conf *dynamic.Configuration, routerName string, routeRule gatev1.HTTPRouteRule, route *gatev1.HTTPRoute, pathMatch *gatev1.HTTPPathMatch, backendRequestTimeout time.Duration, statusReport *statusReport) (string, *metav1.Condition) {
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
		svcName, errCondition := p.loadService(ctx, gatewayName, gatewayNamespace, listener, conf, routerName, route, bi, backendRef, pathMatch, backendRequestTimeout, statusReport)
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
func (p *Provider) loadService(ctx context.Context, gatewayName, gatewayNamespace string, listener gatewayListener, conf *dynamic.Configuration, routerName string, route *gatev1.HTTPRoute, backendIndex int, backendRef gatev1.HTTPBackendRef, pathMatch *gatev1.HTTPPathMatch, backendRequestTimeout time.Duration, statusReport *statusReport) (string, *metav1.Condition) {
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

	if err := p.isReferenceGranted(kindHTTPRoute, route.Namespace, group, string(kind), string(backendRef.Name), namespace); err != nil {
		return serviceName, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot load HTTPBackendRef %s/%s/%s/%s: %s", group, kind, namespace, backendRef.Name, err),
		}
	}

	loadMirrorService := p.newMirrorServiceLoader(gatewayName, gatewayNamespace, listener, route, statusReport)
	middlewares, err := p.loadMiddlewares(ctx, conf, route.Namespace, serviceName, backendRef.Filters, pathMatch, loadMirrorService)
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
		name, service, err := p.loadHTTPBackendRef(namespace, backendRef)
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

	lb, st, errCondition := p.loadHTTPServers(gatewayName, gatewayNamespace, namespace, route, backendRef, listener, statusReport)
	if errCondition != nil {
		return serviceName, errCondition
	}

	// A zero duration means no timeout, as GEP-1742 requires.
	if backendRequestTimeout != 0 {
		if st == nil {
			st = &dynamic.ServersTransport{}
		}

		st.ForwardingTimeouts = &dynamic.ForwardingTimeouts{}
		st.ForwardingTimeouts.SetDefaults()
		// The response header timeout stops once the response headers are received, so it does not bound reading the response body.
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

// mirrorServiceLoader loads the backend of a RequestMirror filter into conf, as a service named after the mirror middleware.
type mirrorServiceLoader func(ctx context.Context, conf *dynamic.Configuration, middlewareName string, backendRef gatev1.BackendObjectReference) (string, *metav1.Condition)

func (p *Provider) newMirrorServiceLoader(gatewayName, gatewayNamespace string, listener gatewayListener, route *gatev1.HTTPRoute, statusReport *statusReport) mirrorServiceLoader {
	return func(ctx context.Context, conf *dynamic.Configuration, middlewareName string, backendRef gatev1.BackendObjectReference) (string, *metav1.Condition) {
		// The rule timeouts.backendRequest bounds the requests to the rule backends, not the mirrored copies, whose responses are discarded.
		httpBackendRef := gatev1.HTTPBackendRef{BackendRef: gatev1.BackendRef{BackendObjectReference: backendRef}}
		return p.loadService(ctx, gatewayName, gatewayNamespace, listener, conf, middlewareName, route, 0, httpBackendRef, nil, 0, statusReport)
	}
}

func (p *Provider) loadMiddlewares(ctx context.Context, conf *dynamic.Configuration, namespace, parentName string, filters []gatev1.HTTPRouteFilter, pathMatch *gatev1.HTTPPathMatch, loadMirrorService mirrorServiceLoader) ([]string, error) {
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
			name, middleware, err := p.loadHTTPRouteFilterExtensionRef(namespace, filter.ExtensionRef)
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
			// As per the spec, an invalid mirror must not fail the request, so it is dropped.
			if filter.RequestMirror == nil {
				log.Ctx(ctx).Debug().Str("middleware", name).Msg("Ignoring HTTPRoute RequestMirror filter: requestMirror undefined")
				continue
			}

			serviceName, errCondition := loadMirrorService(ctx, conf, name, filter.RequestMirror.BackendRef)
			if errCondition != nil {
				log.Ctx(ctx).Debug().Str("middleware", name).Msgf("Ignoring HTTPRoute RequestMirror filter: %s", errCondition.Message)
				continue
			}

			middlewares = append(middlewares, namedMiddleware{
				name,
				createRequestMirror(filter.RequestMirror, serviceName),
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

func (p *Provider) loadHTTPServers(gatewayName, gatewayNamespace, namespace string, route *gatev1.HTTPRoute, backendRef gatev1.HTTPBackendRef, listener gatewayListener, statusReport *statusReport) (*dynamic.ServersLoadBalancer, *dynamic.ServersTransport, *metav1.Condition) {
	backendAddresses, svcPort, err := p.getBackendAddresses(namespace, backendRef.BackendRef)
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

	backendTLSPolicies, err := p.client.ListBackendTLSPoliciesForService(namespace, string(backendRef.Name))
	if err != nil {
		return nil, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot list BackendTLSPolicies for Service %s/%s: %s", namespace, string(backendRef.Name), err),
		}
	}

	// Sort BackendTLSPolicies by creation timestamp, then by name to match the BackendTLSPolicy requirements.
	slices.SortStableFunc(backendTLSPolicies, func(a, b *gatev1.BackendTLSPolicy) int {
		cmpTime := a.CreationTimestamp.Time.Compare(b.CreationTimestamp.Time)
		if cmpTime == 0 {
			return strings.Compare(a.Name, b.Name)
		}
		return cmpTime
	})

	var serversTransport *dynamic.ServersTransport
	for _, policy := range backendTLSPolicies {
		for _, targetRef := range policy.Spec.TargetRefs {
			// Skip targetRefs that doesn't match the backendRef,
			// since a BackendTLSPolicy can select multiple services.
			if targetRef.Name != backendRef.Name {
				continue
			}
			// Skip the targetRef if the sectionName doesn't match the backendRef port.
			if targetRef.SectionName != nil && svcPort.Name != string(*targetRef.SectionName) {
				continue
			}

			policyAncestorStatus := gatev1.PolicyAncestorStatus{
				AncestorRef: gatev1.ParentReference{
					Group:       new(gatev1.Group(groupGateway)),
					Kind:        new(gatev1.Kind(kindGateway)),
					Namespace:   new(gatev1.Namespace(gatewayNamespace)),
					Name:        gatev1.ObjectName(gatewayName),
					SectionName: new(gatev1.SectionName(listener.SectionName)),
				},
				ControllerName: controllerName,
			}

			// Multiple BackendTLSPolicies can match the same service port, meaning that there is a conflict.
			if serversTransport != nil {
				policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions,
					metav1.Condition{
						Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
						Status:             metav1.ConditionFalse,
						ObservedGeneration: policy.Generation,
						LastTransitionTime: metav1.Now(),
						Reason:             string(gatev1.BackendTLSPolicyReasonResolvedRefs),
					},
					metav1.Condition{
						Type:               string(gatev1.PolicyConditionAccepted),
						Status:             metav1.ConditionFalse,
						ObservedGeneration: policy.Generation,
						LastTransitionTime: metav1.Now(),
						Reason:             string(gatev1.PolicyReasonConflicted),
					},
				)

				statusReport.RecordBackendTLSPolicyStatus(ktypes.NamespacedName{Namespace: policy.Namespace, Name: policy.Name}, policyAncestorStatus)

				continue
			}

			var resolvedRefCondition metav1.Condition
			serversTransport, resolvedRefCondition = p.loadServersTransport(namespace, policy, listener.BackendClientCertificate)

			policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions, resolvedRefCondition)
			if resolvedRefCondition.Status == metav1.ConditionFalse {
				policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions, metav1.Condition{
					Type:               string(gatev1.PolicyConditionAccepted),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.BackendTLSPolicyReasonNoValidCACertificate),
				})
			} else {
				policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions, metav1.Condition{
					Type:               string(gatev1.PolicyConditionAccepted),
					Status:             metav1.ConditionTrue,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.PolicyReasonAccepted),
				})
			}

			statusReport.RecordBackendTLSPolicyStatus(ktypes.NamespacedName{Namespace: policy.Namespace, Name: policy.Name}, policyAncestorStatus)

			// When something wen wrong during the loading of a ServersTransport,
			// we stop here and return a route condition error.
			if resolvedRefCondition.Status == metav1.ConditionFalse {
				return nil, nil, &metav1.Condition{
					Type:               string(gatev1.RouteConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: route.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.RouteReasonRefNotPermitted),
					Message:            fmt.Sprintf("Cannot apply BackendTLSPolicy for Service %s/%s: %s", namespace, string(backendRef.Name), resolvedRefCondition.Message),
				}
			}
		}
	}

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

func (p *Provider) loadServersTransport(namespace string, policy *gatev1.BackendTLSPolicy, clientCertificate *tls.Certificate) (*dynamic.ServersTransport, metav1.Condition) {
	st := &dynamic.ServersTransport{
		ServerName: string(policy.Spec.Validation.Hostname),
	}
	if clientCertificate != nil {
		st.Certificates = tls.Certificates{*clientCertificate}
	}

	if len(policy.Spec.Validation.SubjectAltNames) > 0 {
		// Per the Gateway API specification the Hostname should only be used for authentication
		// and not for certificate validation. Thus, if SubjectAltNames is specified, we ignore
		// the Hostname validation by setting the InsecureSkipVerify option to true.
		st.InsecureSkipVerify = true

		for _, san := range policy.Spec.Validation.SubjectAltNames {
			switch san.Type {
			case gatev1.URISubjectAltNameType:
				st.PeerCertSANs = append(st.PeerCertSANs, tls.SAN{
					Type:  tls.SANURIType,
					Value: string(san.URI),
				})
			case gatev1.HostnameSubjectAltNameType:
				st.PeerCertSANs = append(st.PeerCertSANs, tls.SAN{
					Type:  tls.SANDNSNameType,
					Value: string(san.Hostname),
				})
			default:
				return nil, metav1.Condition{
					Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.BackendTLSPolicyReasonInvalidKind),
					Message:            fmt.Sprintf("Unsupported SubjectAltName type %q; only URI and Hostname types are supported", san.Type),
				}
			}
		}
	}

	if policy.Spec.Validation.WellKnownCACertificates != nil {
		return st, metav1.Condition{
			Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
			Status:             metav1.ConditionTrue,
			ObservedGeneration: policy.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.BackendTLSPolicyReasonResolvedRefs),
		}
	}

	for _, caCertRef := range policy.Spec.Validation.CACertificateRefs {
		if (caCertRef.Group != "" && caCertRef.Group != groupCore) || (caCertRef.Kind != kindConfigMap && caCertRef.Kind != kindSecret) {
			return nil, metav1.Condition{
				Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: policy.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.BackendTLSPolicyReasonInvalidKind),
				Message:            "Only ConfigMaps and Secrets are supported",
			}
		}

		var caCRT string
		switch caCertRef.Kind {
		case kindConfigMap:
			configmap, err := p.client.GetConfigMap(namespace, string(caCertRef.Name))
			if err != nil {
				return nil, metav1.Condition{
					Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.BackendTLSPolicyReasonInvalidCACertificateRef),
					Message:            fmt.Sprintf("getting configmap %s/%s: %s", namespace, string(caCertRef.Name), err),
				}
			}
			caCRT = configmap.Data["ca.crt"]
		case kindSecret:
			secret, err := p.client.GetSecret(namespace, string(caCertRef.Name))
			if err != nil {
				return nil, metav1.Condition{
					Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.BackendTLSPolicyReasonInvalidCACertificateRef),
					Message:            fmt.Sprintf("getting secret %s/%s: %s", namespace, string(caCertRef.Name), err),
				}
			}
			caCRT = string(secret.Data["ca.crt"])
		}

		if caCRT == "" {
			return nil, metav1.Condition{
				Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: policy.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.BackendTLSPolicyReasonInvalidCACertificateRef),
				Message:            fmt.Sprintf("%s %s/%s does not have a ca.crt", caCertRef.Kind, namespace, string(caCertRef.Name)),
			}
		}

		st.RootCAs = append(st.RootCAs, types.FileOrContent(caCRT))
	}

	return st, metav1.Condition{
		Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
		Status:             metav1.ConditionTrue,
		ObservedGeneration: policy.Generation,
		LastTransitionTime: metav1.Now(),
		Reason:             string(gatev1.BackendTLSPolicyReasonResolvedRefs),
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

func (p *Provider) createRetry(ctx context.Context, retry *gatev1.HTTPRouteRetry) *dynamic.Middleware {
	// Traefik counts the initial request in the attempts, whereas the Gateway API only counts the retries.
	// The Gateway API defaults attempts to 1.
	middleware := &dynamic.Retry{
		Attempts:            ptr.Deref(retry.Attempts, 1) + 1,
		MaxRequestBodyBytes: new(p.RetryMaxRequestBodyBytes),
	}

	for _, code := range retry.Codes {
		middleware.Status = append(middleware.Status, strconv.Itoa(int(code)))
	}

	if retry.Backoff != nil {
		// The CRD CEL validation guarantees a valid GEP-2257 duration string.
		if d, err := time.ParseDuration(string(*retry.Backoff)); err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("Ignoring HTTPRoute retry.backoff: invalid duration")
		} else {
			middleware.InitialInterval = ptypes.Duration(d)
		}
	}

	return &dynamic.Middleware{Retry: middleware}
}

func createRequestMirror(filter *gatev1.HTTPRequestMirrorFilter, serviceName string) *dynamic.Middleware {
	numerator, denominator := mirrorRatio(filter)

	return &dynamic.Middleware{
		RequestMirror: &dynamic.RequestMirror{
			Service:     serviceName,
			Numerator:   numerator,
			Denominator: denominator,
		},
	}
}

// mirrorRatio returns the share of requests to mirror.
// The CRD validation bounds percent and fraction, but out of range values are clamped rather than trusted.
func mirrorRatio(filter *gatev1.HTTPRequestMirrorFilter) (numerator, denominator int64) {
	switch {
	case filter.Percent != nil:
		return int64(min(max(*filter.Percent, 0), 100)), 100

	case filter.Fraction != nil:
		d := ptr.Deref(filter.Fraction.Denominator, 100)
		if d <= 0 {
			return 0, 1
		}
		return int64(min(max(filter.Fraction.Numerator, 0), d)), int64(d)

	default:
		return 1, 1
	}
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

```

#### pkg/httproute_test.go

```go
package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func Test_buildHostRule(t *testing.T) {
	testCases := []struct {
		desc             string
		hostnames        []gatev1.Hostname
		expectedRule     string
		expectedPriority int
		expectErr        bool
	}{
		{
			desc:         "Empty (should not happen)",
			expectedRule: "",
		},
		{
			desc: "One Host",
			hostnames: []gatev1.Hostname{
				"Foo",
			},
			expectedRule:     `Host("Foo")`,
			expectedPriority: 3,
		},
		{
			desc: "Multiple Hosts",
			hostnames: []gatev1.Hostname{
				"Foo",
				"Bar",
				"Bir",
			},
			expectedRule:     `(Host("Foo") || Host("Bar") || Host("Bir"))`,
			expectedPriority: 3,
		},
		{
			desc: "Several Host and wildcard",
			hostnames: []gatev1.Hostname{
				"*.bar.foo",
				"bar.foo",
				"foo.foo",
			},
			expectedRule:     `(Host("**.bar.foo") || Host("bar.foo") || Host("foo.foo"))`,
			expectedPriority: 9,
		},
		{
			desc: "Host with wildcard",
			hostnames: []gatev1.Hostname{
				"*.bar.foo",
			},
			expectedRule:     `Host("**.bar.foo")`,
			expectedPriority: 9,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			rule, priority := buildHostRule(test.hostnames)
			assert.Equal(t, test.expectedRule, rule)
			assert.Equal(t, test.expectedPriority, priority)
		})
	}
}

func Test_buildMatchRule(t *testing.T) {
	testCases := []struct {
		desc             string
		match            gatev1.HTTPRouteMatch
		hostnames        []gatev1.Hostname
		expectedRule     string
		expectedPriority int
		expectedError    bool
	}{
		{
			desc:             "Empty rule and matches",
			expectedRule:     `PathPrefix("/")`,
			expectedPriority: 1,
		},
		{
			desc:             "One Host rule without match",
			hostnames:        []gatev1.Hostname{"foo.com"},
			expectedRule:     `Host("foo.com") && PathPrefix("/")`,
			expectedPriority: 8,
		},
		{
			desc: "One HTTPRouteMatch with nil HTTPHeaderMatch",
			match: gatev1.HTTPRouteMatch{
				Path: new(gatev1.HTTPPathMatch{
					Type:  new(gatev1.PathMatchPathPrefix),
					Value: new("/"),
				}),
				Headers: nil,
			},
			expectedRule:     `PathPrefix("/")`,
			expectedPriority: 1,
		},
		{
			desc: "One HTTPRouteMatch with nil HTTPHeaderMatch Type",
			match: gatev1.HTTPRouteMatch{
				Path: new(gatev1.HTTPPathMatch{
					Type:  new(gatev1.PathMatchPathPrefix),
					Value: new("/"),
				}),
				Headers: []gatev1.HTTPHeaderMatch{
					{Name: "foo", Value: "bar"},
				},
			},
			expectedRule:     `PathPrefix("/") && Header("foo","bar")`,
			expectedPriority: 101,
		},
		{
			desc:             "One HTTPRouteMatch with nil HTTPPathMatch",
			match:            gatev1.HTTPRouteMatch{Path: nil},
			expectedRule:     `PathPrefix("/")`,
			expectedPriority: 1,
		},
		{
			desc: "One HTTPRouteMatch with nil HTTPPathMatch Type",
			match: gatev1.HTTPRouteMatch{
				Path: &gatev1.HTTPPathMatch{
					Type:  nil,
					Value: new("/foo/"),
				},
			},
			expectedRule:     `(Path("/foo") || PathPrefix("/foo/"))`,
			expectedPriority: 10500,
		},
		{
			desc: "One HTTPRouteMatch with nil HTTPPathMatch Values",
			match: gatev1.HTTPRouteMatch{
				Path: &gatev1.HTTPPathMatch{
					Type:  new(gatev1.PathMatchExact),
					Value: nil,
				},
			},
			expectedRule:     `Path("/")`,
			expectedPriority: 100000,
		},
		{
			desc: "One Path",
			match: gatev1.HTTPRouteMatch{
				Path: &gatev1.HTTPPathMatch{
					Type:  new(gatev1.PathMatchExact),
					Value: new("/foo/"),
				},
			},
			expectedRule:     `Path("/foo/")`,
			expectedPriority: 100000,
		},
		{
			desc: "Path && Header",
			match: gatev1.HTTPRouteMatch{
				Path: &gatev1.HTTPPathMatch{
					Type:  new(gatev1.PathMatchExact),
					Value: new("/foo/"),
				},
				Headers: []gatev1.HTTPHeaderMatch{
					{
						Type:  new(gatev1.HeaderMatchExact),
						Name:  "my-header",
						Value: "foo",
					},
				},
			},
			expectedRule:     `Path("/foo/") && Header("my-header","foo")`,
			expectedPriority: 100100,
		},
		{
			desc:      "Host && Path && Header",
			hostnames: []gatev1.Hostname{"foo.com"},
			match: gatev1.HTTPRouteMatch{
				Path: &gatev1.HTTPPathMatch{
					Type:  new(gatev1.PathMatchExact),
					Value: new("/foo/"),
				},
				Headers: []gatev1.HTTPHeaderMatch{
					{
						Type:  new(gatev1.HeaderMatchExact),
						Name:  "my-header",
						Value: "foo",
					},
				},
			},
			expectedRule:     `Host("foo.com") && Path("/foo/") && Header("my-header","foo")`,
			expectedPriority: 100107,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			rule, priority := buildMatchRule(test.hostnames, test.match)
			assert.Equal(t, test.expectedRule, rule)
			assert.Equal(t, test.expectedPriority, priority)
		})
	}
}

func Test_createCORS(t *testing.T) {
	testCases := []struct {
		desc     string
		filter   *gatev1.HTTPCORSFilter
		expected *dynamic.Middleware
	}{
		{
			desc:   "Empty filter",
			filter: &gatev1.HTTPCORSFilter{},
			expected: &dynamic.Middleware{
				Headers: &dynamic.Headers{
					AccessControlMaxAge: new(int64(0)),
				},
			},
		},
		{
			desc: "Plain origins",
			filter: &gatev1.HTTPCORSFilter{
				AllowOrigins: []gatev1.CORSOrigin{"https://foo.example.com", "https://bar.example.com"},
			},
			expected: &dynamic.Middleware{
				Headers: &dynamic.Headers{
					AccessControlAllowOriginList: []string{"https://foo.example.com", "https://bar.example.com"},
					AccessControlMaxAge:          new(int64(0)),
				},
			},
		},
		{
			desc: "Wildcard origin becomes catch-all regex",
			filter: &gatev1.HTTPCORSFilter{
				AllowOrigins: []gatev1.CORSOrigin{"*"},
			},
			expected: &dynamic.Middleware{
				Headers: &dynamic.Headers{
					AccessControlAllowOriginListRegex: []string{`.*`},
					AccessControlMaxAge:               new(int64(0)),
				},
			},
		},
		{
			desc: "Origin with bare wildcard host",
			filter: &gatev1.HTTPCORSFilter{
				AllowOrigins: []gatev1.CORSOrigin{"https://*"},
			},
			expected: &dynamic.Middleware{
				Headers: &dynamic.Headers{
					AccessControlAllowOriginListRegex: []string{`^https://.*$`},
					AccessControlMaxAge:               new(int64(0)),
				},
			},
		},
		{
			desc: "Origin with wildcard subdomain becomes regex",
			filter: &gatev1.HTTPCORSFilter{
				AllowOrigins: []gatev1.CORSOrigin{"https://*.example.com"},
			},
			expected: &dynamic.Middleware{
				Headers: &dynamic.Headers{
					AccessControlAllowOriginListRegex: []string{`^https://.*\.example\.com$`},
					AccessControlMaxAge:               new(int64(0)),
				},
			},
		},
		{
			desc: "Mixed plain and wildcard origins",
			filter: &gatev1.HTTPCORSFilter{
				AllowOrigins: []gatev1.CORSOrigin{"https://foo.example.com", "https://*.example.com"},
			},
			expected: &dynamic.Middleware{
				Headers: &dynamic.Headers{
					AccessControlAllowOriginList:      []string{"https://foo.example.com"},
					AccessControlAllowOriginListRegex: []string{`^https://.*\.example\.com$`},
					AccessControlMaxAge:               new(int64(0)),
				},
			},
		},
		{
			desc: "All fields set",
			filter: &gatev1.HTTPCORSFilter{
				AllowOrigins:     []gatev1.CORSOrigin{"https://foo.example.com"},
				AllowCredentials: new(true),
				AllowMethods:     []gatev1.HTTPMethodWithWildcard{"GET", "POST"},
				AllowHeaders:     []gatev1.HTTPHeaderName{"X-Foo", "X-Bar"},
				ExposeHeaders:    []gatev1.HTTPHeaderName{"X-Baz"},
				MaxAge:           600,
			},
			expected: &dynamic.Middleware{
				Headers: &dynamic.Headers{
					AccessControlAllowCredentials: true,
					AccessControlAllowOriginList:  []string{"https://foo.example.com"},
					AccessControlAllowMethods:     []string{"GET", "POST"},
					AccessControlAllowHeaders:     []string{"X-Foo", "X-Bar"},
					AccessControlExposeHeaders:    []string{"X-Baz"},
					AccessControlMaxAge:           new(int64(600)),
				},
			},
		},
		{
			desc: "AllowCredentials explicitly false",
			filter: &gatev1.HTTPCORSFilter{
				AllowCredentials: new(false),
			},
			expected: &dynamic.Middleware{
				Headers: &dynamic.Headers{
					AccessControlAllowCredentials: false,
					AccessControlMaxAge:           new(int64(0)),
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.expected, createCORS(test.filter))
		})
	}
}

func Test_createRequestMirror(t *testing.T) {
	testCases := []struct {
		desc                string
		percent             *int32
		fraction            *gatev1.Fraction
		expectedNumerator   int64
		expectedDenominator int64
	}{
		{
			desc:                "Neither percent nor fraction",
			expectedNumerator:   1,
			expectedDenominator: 1,
		},
		{
			desc:                "Percent",
			percent:             new(int32(42)),
			expectedNumerator:   42,
			expectedDenominator: 100,
		},
		{
			desc:                "Percent below 0",
			percent:             new(int32(-1)),
			expectedNumerator:   0,
			expectedDenominator: 100,
		},
		{
			desc:                "Percent above 100",
			percent:             new(int32(101)),
			expectedNumerator:   100,
			expectedDenominator: 100,
		},
		{
			desc:                "Fraction",
			fraction:            &gatev1.Fraction{Numerator: 1, Denominator: new(int32(4))},
			expectedNumerator:   1,
			expectedDenominator: 4,
		},
		{
			desc:                "Fraction with default denominator",
			fraction:            &gatev1.Fraction{Numerator: 5},
			expectedNumerator:   5,
			expectedDenominator: 100,
		},
		{
			desc:                "Fraction with zero denominator",
			fraction:            &gatev1.Fraction{Numerator: 1, Denominator: new(int32(0))},
			expectedNumerator:   0,
			expectedDenominator: 1,
		},
		{
			desc:                "Fraction with negative denominator",
			fraction:            &gatev1.Fraction{Numerator: 1, Denominator: new(int32(-2))},
			expectedNumerator:   0,
			expectedDenominator: 1,
		},
		{
			desc:                "Fraction with numerator above denominator",
			fraction:            &gatev1.Fraction{Numerator: 3, Denominator: new(int32(2))},
			expectedNumerator:   2,
			expectedDenominator: 2,
		},
		{
			desc:                "Fraction with negative numerator",
			fraction:            &gatev1.Fraction{Numerator: -1, Denominator: new(int32(2))},
			expectedNumerator:   0,
			expectedDenominator: 2,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			filter := &gatev1.HTTPRequestMirrorFilter{Percent: test.percent, Fraction: test.fraction}

			expected := &dynamic.Middleware{
				RequestMirror: &dynamic.RequestMirror{
					Service:     "mirror",
					Numerator:   test.expectedNumerator,
					Denominator: test.expectedDenominator,
				},
			}
			assert.Equal(t, expected, createRequestMirror(filter, "mirror"))
		})
	}
}

```

#### pkg/kubernetes.go

```go
package gateway

import (
	"cmp"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"maps"
	"os"
	"reflect"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/rs/zerolog/log"
	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/job"
	"github.com/traefik/traefik/v3/pkg/observability/logs"
	"github.com/traefik/traefik/v3/pkg/provider"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	"github.com/traefik/traefik/v3/pkg/provider/kubernetes/k8s"
	"github.com/traefik/traefik/v3/pkg/safe"
	"github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	ktypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

const (
	// ProviderName is the Kubernetes Gateway API provider name.
	ProviderName = "kubernetesgateway"

	controllerName = "traefik.io/gateway-controller"

	groupCore    = "core"
	groupGateway = "gateway.networking.k8s.io"

	kindGateway        = "Gateway"
	kindTraefikService = "TraefikService"
	kindHTTPRoute      = "HTTPRoute"
	kindGRPCRoute      = "GRPCRoute"
	kindTCPRoute       = "TCPRoute"
	kindTLSRoute       = "TLSRoute"
	kindService        = "Service"
	kindListenerSet    = "ListenerSet"
	kindConfigMap      = "ConfigMap"
	kindSecret         = "Secret"

	appProtocolHTTP  = "http"
	appProtocolHTTPS = "https"
	appProtocolH2C   = "kubernetes.io/h2c"
	appProtocolWS    = "kubernetes.io/ws"
	appProtocolWSS   = "kubernetes.io/wss"

	schemeHTTP  = "http"
	schemeHTTPS = "https"
	schemeH2C   = "h2c"

	// routeReasonHostnameConflict is raised when an HTTP and a GRPC route are
	// attached to the same listener with intersecting hostnames.
	routeReasonHostnameConflict gatev1.RouteConditionReason = "HostnameConflict"

	// listenerReasonFrontendValidationConflict reports a frontend validation that differs from the one an older Gateway applies on the port.
	listenerReasonFrontendValidationConflict gatev1.ListenerConditionReason = "FrontendValidationConflict"

	conditionNoErrorMessage = "No error found"
)

// NamespacedName holds a Kubernetes resource reference with namespace and name.
type NamespacedName struct {
	Namespace string `description:"Defines the resource namespace." json:"namespace,omitempty" toml:"namespace,omitempty" yaml:"namespace,omitempty"`
	Name      string `description:"Defines the resource name." json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty"`
}

// Provider holds configurations of the provider.
type Provider struct {
	Endpoint         string              `description:"Kubernetes server endpoint (required for external cluster client)." json:"endpoint,omitempty" toml:"endpoint,omitempty" yaml:"endpoint,omitempty"`
	Token            types.FileOrContent `description:"Kubernetes bearer token (not needed for in-cluster client). It accepts either a token value or a file path to the token." json:"token,omitempty" toml:"token,omitempty" yaml:"token,omitempty" loggable:"false"`
	QPS              int                 `description:"Defines the maximum QPS to the Kubernetes API server. Setting this to a negative value will disable client-side ratelimiting." json:"qps,omitempty" toml:"qps,omitempty" yaml:"qps,omitempty" export:"true"`
	Burst            int                 `description:"Defines the maximum burst of requests to the Kubernetes API server." json:"burst,omitempty" toml:"burst,omitempty" yaml:"burst,omitempty" export:"true"`
	CertAuthFilePath string              `description:"Kubernetes certificate authority file path (not needed for in-cluster client)." json:"certAuthFilePath,omitempty" toml:"certAuthFilePath,omitempty" yaml:"certAuthFilePath,omitempty"`
	Namespaces       []string            `description:"Kubernetes namespaces." json:"namespaces,omitempty" toml:"namespaces,omitempty" yaml:"namespaces,omitempty" export:"true"`
	GatewayClass     string              `description:"Name of the GatewayClass to watch for. If empty, all GatewayClasses are processed." json:"gatewayClass,omitempty" toml:"gatewayClass,omitempty" yaml:"gatewayClass,omitempty" export:"true"`
	// Deprecated: please use the GatewayClass option instead.
	LabelSelector            string           `description:"Kubernetes label selector to select specific GatewayClasses (Deprecated, please use the GatewayClass option instead)." json:"labelSelector,omitempty" toml:"labelSelector,omitempty" yaml:"labelSelector,omitempty" export:"true"`
	Gateways                 []NamespacedName `description:"Scopes the provider to specific Gateways." json:"gateways,omitempty" toml:"gateways,omitempty" yaml:"gateways,omitempty" export:"true"`
	ThrottleDuration         ptypes.Duration  `description:"Kubernetes refresh throttle duration" json:"throttleDuration,omitempty" toml:"throttleDuration,omitempty" yaml:"throttleDuration,omitempty" export:"true"`
	ExperimentalChannel      bool             `description:"Toggles Experimental Channel resources support. Requires the Experimental Channel CRDs." json:"experimentalChannel,omitempty" toml:"experimentalChannel,omitempty" yaml:"experimentalChannel,omitempty" export:"true"`
	StatusAddress            *StatusAddress   `description:"Defines the Kubernetes Gateway status address." json:"statusAddress,omitempty" toml:"statusAddress,omitempty" yaml:"statusAddress,omitempty" export:"true"`
	NativeLBByDefault        bool             `description:"Defines whether to use Native Kubernetes load-balancing by default." json:"nativeLBByDefault,omitempty" toml:"nativeLBByDefault,omitempty" yaml:"nativeLBByDefault,omitempty" export:"true"`
	CrossProviderNamespaces  []string         `description:"List of namespaces from which Gateway API routes are allowed to declare TraefikService backendRef references." json:"crossProviderNamespaces,omitempty" toml:"crossProviderNamespaces,omitempty" yaml:"crossProviderNamespaces,omitempty" export:"true"`
	RetryMaxRequestBodyBytes int64            `description:"Defines the maximum request body size (in bytes) buffered for HTTPRoute retries. If negative, the limit is disabled." json:"retryMaxRequestBodyBytes,omitempty" toml:"retryMaxRequestBodyBytes,omitempty" yaml:"retryMaxRequestBodyBytes,omitempty" export:"true"`

	EntryPoints map[string]Entrypoint `json:"-" toml:"-" yaml:"-" label:"-" file:"-"`

	// groupKindFilterFuncs is the list of allowed Group and Kinds for the Filter ExtensionRef objects.
	groupKindFilterFuncs map[string]map[string]BuildFilterFunc
	// groupKindBackendFuncs is the list of allowed Group and Kinds for the Backend ExtensionRef objects.
	groupKindBackendFuncs map[string]map[string]BuildBackendFunc

	routerTransform k8s.RouterTransform
	client          *clientWrapper
}

func (p *Provider) SetDefaults() {
	p.QPS = 50    // the default value for the QPS is 10x the default Kubernetes client QPS value.
	p.Burst = 100 // the default value for the Burst is 10x the default Kubernetes client Burst value.
	p.RetryMaxRequestBodyBytes = dynamic.RetryDefaultMaxRequestBodyBytes
}

// Entrypoint defines the available entry points.
type Entrypoint struct {
	Address        string
	HasHTTPTLSConf bool
}

// StatusAddress holds the Gateway Status address configuration.
type StatusAddress struct {
	IP       string     `description:"IP used to set Kubernetes Gateway status address." json:"ip,omitempty" toml:"ip,omitempty" yaml:"ip,omitempty"`
	Hostname string     `description:"Hostname used for Kubernetes Gateway status address." json:"hostname,omitempty" toml:"hostname,omitempty" yaml:"hostname,omitempty"`
	Service  ServiceRef `description:"Published Kubernetes Service to copy status addresses from." json:"service" toml:"service,omitempty" yaml:"service,omitempty"`
}

// ServiceRef holds a Kubernetes service reference.
type ServiceRef struct {
	Name      string `description:"Name of the Kubernetes service." json:"name,omitempty" toml:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `description:"Namespace of the Kubernetes service." json:"namespace,omitempty" toml:"namespace,omitempty" yaml:"namespace,omitempty"`
}

// BuildFilterFunc returns the name of the filter and the related dynamic.Middleware if needed.
type BuildFilterFunc func(name, namespace string) (string, *dynamic.Middleware, error)

// BuildBackendFunc returns the name of the backend and the related dynamic.Service if needed.
type BuildBackendFunc func(name, namespace string) (string, *dynamic.Service, error)

type ExtensionBuilderRegistry interface {
	RegisterFilterFuncs(group, kind string, builderFunc BuildFilterFunc)
	RegisterBackendFuncs(group, kind string, builderFunc BuildBackendFunc)
}

// listenerOwner identifies the resource declaring listeners a route parentRef targets, a Gateway or a ListenerSet.
type listenerOwner struct {
	Kind      string
	Namespace string
	Name      string
}

type gatewayListener struct {
	SectionName string

	Port              gatev1.PortNumber
	Protocol          gatev1.ProtocolType
	TLS               *gatev1.ListenerTLSConfig
	Hostname          *gatev1.Hostname
	Status            *gatev1.ListenerStatus
	AllowedNamespaces []string
	AllowedRouteKinds []string

	Attached bool

	EPName string

	// RouterNames holds one parent router per entry point hostname
	// the listener is the most specific match for.
	RouterNames []string

	// ListenerSet is the ListenerSet declaring this listener, nil when the Gateway declares it itself.
	ListenerSet *gatev1.ListenerSet

	// BackendClientCertificate is the client certificate of the parent Gateway, nil when it has none,
	// presented to the backends targeted by a BackendTLSPolicy.
	BackendClientCertificate *tls.Certificate
}

type gatewayWithListeners struct {
	gateway   *gatev1.Gateway
	listeners []gatewayListener

	// backendClientCertificateCondition explains why the Gateway backend client certificate cannot be resolved, nil when it can.
	backendClientCertificateCondition *metav1.Condition
}

// RegisterFilterFuncs registers an allowed Group, Kind, and builder for the Filter ExtensionRef objects.
func (p *Provider) RegisterFilterFuncs(group, kind string, builderFunc BuildFilterFunc) {
	if p.groupKindFilterFuncs == nil {
		p.groupKindFilterFuncs = map[string]map[string]BuildFilterFunc{}
	}

	if p.groupKindFilterFuncs[group] == nil {
		p.groupKindFilterFuncs[group] = map[string]BuildFilterFunc{}
	}

	p.groupKindFilterFuncs[group][kind] = builderFunc
}

// RegisterBackendFuncs registers an allowed Group, Kind, and builder for the Backend ExtensionRef objects.
func (p *Provider) RegisterBackendFuncs(group, kind string, builderFunc BuildBackendFunc) {
	if p.groupKindBackendFuncs == nil {
		p.groupKindBackendFuncs = map[string]map[string]BuildBackendFunc{}
	}

	if p.groupKindBackendFuncs[group] == nil {
		p.groupKindBackendFuncs[group] = map[string]BuildBackendFunc{}
	}

	p.groupKindBackendFuncs[group][kind] = builderFunc
}

func (p *Provider) SetRouterTransform(routerTransform k8s.RouterTransform) {
	p.routerTransform = routerTransform
}

// Init the provider.
func (p *Provider) Init() error {
	logger := log.With().Str(logs.ProviderName, ProviderName).Logger()

	var err error
	p.client, err = p.newK8sClient(logger.WithContext(context.Background()))
	if err != nil {
		return fmt.Errorf("creating k8s client: %w", err)
	}

	return nil
}

// Provide allows the k8s provider to provide configurations to traefik using the given configuration channel.
func (p *Provider) Provide(configurationChan chan<- dynamic.Message, pool *safe.Pool) error {
	logger := log.With().Str(logs.ProviderName, ProviderName).Logger()
	ctxLog := logger.WithContext(context.Background())

	if p.CrossProviderNamespaces != nil {
		logger.Warn().Msgf("Cross-provider references are restricted to namespaces %v (see CrossProviderNamespaces option)", p.CrossProviderNamespaces)
	}

	pool.GoCtx(func(ctxPool context.Context) {
		operation := func() error {
			eventsChan, err := p.client.WatchAll(p.Namespaces, ctxPool.Done())
			if err != nil {
				logger.Error().Err(err).Msg("Error watching kubernetes events")
				timer := time.NewTimer(1 * time.Second)
				select {
				case <-timer.C:
					return err
				case <-ctxPool.Done():
					return nil
				}
			}

			throttleDuration := time.Duration(p.ThrottleDuration)
			throttledChan := throttleEvents(ctxLog, throttleDuration, pool, eventsChan)
			if throttledChan != nil {
				eventsChan = throttledChan
			}

			for {
				select {
				case <-ctxPool.Done():
					return nil
				case <-eventsChan:
					// Note that event is the *first* event that came in during this throttling interval -- if we're hitting our throttle, we may have dropped events.
					// This is fine, because we don't treat different event types differently.
					// But if we do in the future, we'll need to track more information about the dropped events.
					conf, statusReport, err := p.loadConfigurationFromGateways(ctxLog)
					if err != nil {
						logger.Error().Err(err).Msg("Unable to load configuration from Gateways")
					} else {
						configurationChan <- dynamic.Message{
							ProviderName:  ProviderName,
							Configuration: conf,
						}

						// Flush regardless of whether the dynamic configuration changed: the
						// statusReport is independent of confHash and may carry writes even
						// when the data plane has nothing new to consume (e.g. a GatewayClass
						// that's now Accepted but has no Gateway pointing at it yet).
						statusReport.Flush(ctxLog, p.client)
					}

					// If we're throttling,
					// we sleep here for the throttle duration to enforce that we don't refresh faster than our throttle.
					// time.Sleep returns immediately if p.ThrottleDuration is 0 (no throttle).
					time.Sleep(throttleDuration)
				}
			}
		}

		notify := func(err error, time time.Duration) {
			logger.Error().Err(err).Msgf("Provider error, retrying in %s", time)
		}
		err := backoff.RetryNotify(safe.OperationWithRecover(operation), backoff.WithContext(job.NewBackOff(backoff.NewExponentialBackOff()), ctxPool), notify)
		if err != nil {
			logger.Error().Err(err).Msg("Cannot retrieve data")
		}
	})

	return nil
}

func (p *Provider) applyRouterTransform(ctx context.Context, rt *dynamic.Router, route *gatev1.HTTPRoute) {
	if p.routerTransform == nil {
		return
	}

	if err := p.routerTransform.Apply(ctx, rt, route); err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Apply router transform")
	}
}

// managesGateway reports whether the provider should reconcile the given Gateways.
// When the Gateways scoping option is set, only the referenced namespace/name entries are managed.
func (p *Provider) managesGateway(gateway *gatev1.Gateway) bool {
	if len(p.Gateways) == 0 {
		return true
	}

	for _, g := range p.Gateways {
		if gateway.Namespace == g.Namespace && gateway.Name == g.Name {
			return true
		}
	}

	return false
}

func (p *Provider) newK8sClient(ctx context.Context) (*clientWrapper, error) {
	if p.GatewayClass != "" && p.LabelSelector != "" {
		return nil, errors.New("gatewayClass and labelSelector options are mutually exclusive")
	}

	logger := log.Ctx(ctx)

	if p.LabelSelector != "" {
		// Label selector validation
		_, err := labels.Parse(p.LabelSelector)
		if err != nil {
			return nil, fmt.Errorf("invalid label selector: %q", p.LabelSelector)
		}

		logger.Info().Msgf("Label selector is: %q", p.LabelSelector)
		logger.Warn().Msg("LabelSelector option is deprecated, please use the GatewayClass option instead.")
	}

	var (
		err    error
		client *clientWrapper
	)
	switch {
	case os.Getenv("KUBERNETES_SERVICE_HOST") != "" && os.Getenv("KUBERNETES_SERVICE_PORT") != "":
		logger.Info().Str("endpoint", p.Endpoint).Msg("Creating in-cluster Provider client")
		client, err = newInClusterClient(p.Endpoint, p.QPS, p.Burst)
	case os.Getenv("KUBECONFIG") != "":
		logger.Info().Msgf("Creating cluster-external Provider client from KUBECONFIG %s", os.Getenv("KUBECONFIG"))
		client, err = newExternalClusterClientFromFile(os.Getenv("KUBECONFIG"), p.QPS, p.Burst)
	default:
		logger.Info().Str("endpoint", p.Endpoint).Msg("Creating cluster-external Provider client")
		client, err = newExternalClusterClient(p.Endpoint, p.CertAuthFilePath, p.Token, p.QPS, p.Burst)
	}

	if err != nil {
		return nil, err
	}

	client.gatewayClass = p.GatewayClass
	client.labelSelector = p.LabelSelector

	return client, nil
}

// TODO Handle errors and update resources statuses (gatewayClass, gateway).
func (p *Provider) loadConfigurationFromGateways(ctx context.Context) (*dynamic.Configuration, *statusReport, error) {
	statusReport := newStatusReport()
	conf := &dynamic.Configuration{
		HTTP: &dynamic.HTTPConfiguration{
			Routers:           map[string]*dynamic.Router{},
			Middlewares:       map[string]*dynamic.Middleware{},
			Services:          map[string]*dynamic.Service{},
			ServersTransports: map[string]*dynamic.ServersTransport{},
		},
		TCP: &dynamic.TCPConfiguration{
			Routers:           map[string]*dynamic.TCPRouter{},
			Middlewares:       map[string]*dynamic.TCPMiddleware{},
			Services:          map[string]*dynamic.TCPService{},
			ServersTransports: map[string]*dynamic.TCPServersTransport{},
		},
		UDP: &dynamic.UDPConfiguration{
			Routers:  map[string]*dynamic.UDPRouter{},
			Services: map[string]*dynamic.UDPService{},
		},
		TLS: &dynamic.TLSConfiguration{
			Options: map[string]tls.Options{},
		},
	}

	addresses, err := p.gatewayAddresses()
	if err != nil {
		return nil, nil, fmt.Errorf("getting gateway addresses: %w", err)
	}

	gatewayClasses, err := p.client.ListGatewayClasses()
	if err != nil {
		return nil, nil, fmt.Errorf("listing gateway classes: %w", err)
	}

	var supportedFeatures []gatev1.SupportedFeature
	for _, feature := range SupportedFeatures() {
		supportedFeatures = append(supportedFeatures, gatev1.SupportedFeature{Name: gatev1.FeatureName(feature)})
	}
	slices.SortFunc(supportedFeatures, func(a, b gatev1.SupportedFeature) int {
		return strings.Compare(string(a.Name), string(b.Name))
	})

	gatewayClassNames := map[string]struct{}{}
	for _, gatewayClass := range gatewayClasses {
		if gatewayClass.Spec.ControllerName != controllerName {
			continue
		}

		gatewayClassNames[gatewayClass.Name] = struct{}{}

		status := gatev1.GatewayClassStatus{
			Conditions: upsertGatewayClassConditionAccepted(gatewayClass.Status.Conditions, metav1.Condition{
				Type:               string(gatev1.GatewayClassConditionStatusAccepted),
				Status:             metav1.ConditionTrue,
				ObservedGeneration: gatewayClass.Generation,
				Reason:             "Handled",
				Message:            "Handled by Traefik controller",
				LastTransitionTime: metav1.Now(),
			}),
			SupportedFeatures: supportedFeatures,
		}

		statusReport.RecordGatewayClassStatus(gatewayClass.Name, status)
	}

	var gateways []*gatev1.Gateway
	for _, gateway := range p.client.ListGateways() {
		if _, ok := gatewayClassNames[string(gateway.Spec.GatewayClassName)]; !ok {
			continue
		}
		if !p.managesGateway(gateway) {
			continue
		}
		gateways = append(gateways, gateway)
	}

	slices.SortStableFunc(gateways, func(a, b *gatev1.Gateway) int {
		return cmp.Or(a.GetCreationTimestamp().Time.Compare(b.GetCreationTimestamp().Time),
			strings.Compare(a.GetNamespace(), b.GetNamespace()),
			strings.Compare(a.GetName(), b.GetName()))
	})

	// ListenerSets are listed once and dispatched to their parent Gateway below,
	// oldest first, as the oldest ListenerSet wins a listener conflict with its siblings (GEP-1713).
	listenerSets := p.client.ListListenerSets()
	slices.SortStableFunc(listenerSets, func(a, b *gatev1.ListenerSet) int {
		return cmp.Or(a.GetCreationTimestamp().Time.Compare(b.GetCreationTimestamp().Time),
			strings.Compare(a.GetNamespace(), b.GetNamespace()),
			strings.Compare(a.GetName(), b.GetName()))
	})

	var gatewaysWithListeners []gatewayWithListeners
	frontendValidations := make(map[gatev1.PortNumber]*tls.ClientAuth)
	for _, gateway := range gateways {
		logger := log.Ctx(ctx).With().
			Str("gateway", gateway.Name).
			Str("namespace", gateway.Namespace).
			Logger()

		claims := newListenerClaims(frontendValidations)

		listeners := p.loadGatewayListeners(logger.WithContext(ctx), gateway, nil, claims, conf)

		listeners = append(listeners, p.loadListenerSetListeners(logger.WithContext(ctx), gateway, listenerSets, claims, conf, statusReport)...)

		clientCertificate, clientCertificateCondition := p.resolveBackendClientCertificate(gateway)
		if clientCertificateCondition != nil {
			logger.Error().Str("reason", clientCertificateCondition.Reason).Msg(clientCertificateCondition.Message)
		}
		// The client certificate is set at the Gateway level, and applies to the listeners of its ListenerSets too.
		for i := range listeners {
			listeners[i].BackendClientCertificate = clientCertificate
		}

		gatewaysWithListeners = append(gatewaysWithListeners, gatewayWithListeners{
			gateway:                           gateway,
			listeners:                         listeners,
			backendClientCertificateCondition: clientCertificateCondition,
		})
	}

	statusReport.gatewayListeners = gatewaysWithListeners

	// The isolation of a listener depends on the other listeners of its entry point.
	listenerRouters := p.buildListenerRouters(gatewaysWithListeners, frontendValidations, conf)

	p.loadHTTPAndGRPCRoutes(ctx, gatewaysWithListeners, conf, statusReport)

	p.loadTLSRoutes(ctx, gatewaysWithListeners, conf, statusReport)

	p.loadTCPRoutes(ctx, gatewaysWithListeners, conf, statusReport)

	// A listener with no route attached gives a parent router with no child,
	// which the router manager reports in error as it has no service either.
	dropChildlessListenerRouters(conf, listenerRouters)

	for _, gwl := range gatewaysWithListeners {
		logger := log.Ctx(ctx).With().
			Str("gateway", gwl.gateway.Name).
			Str("namespace", gwl.gateway.Namespace).
			Logger()

		// A Gateway is accepted as soon as one of the listeners serving it is valid, whoever declares it.
		// GEP-1713 forbids programming the listeners of a ListenerSet whose parent Gateway is not accepted,
		// which holds on its own here:
		// only a valid listener is programmed, and a valid one makes the Gateway accepted.
		// It is decided before makeGatewayStatus reports the valid listeners with their conditions.
		accepted := len(gwl.listeners) == 0 || slices.ContainsFunc(gwl.listeners, func(listener gatewayListener) bool {
			return len(listener.Status.Conditions) == 0
		})

		gatewayStatus, errConditions := p.makeGatewayStatus(gwl.gateway, gwl.listeners, gwl.backendClientCertificateCondition, addresses, accepted)
		if len(errConditions) > 0 {
			messages := map[string]struct{}{}
			for _, condition := range errConditions {
				messages[condition.Message] = struct{}{}
			}
			var conditionsErr error
			for message := range messages {
				conditionsErr = errors.Join(conditionsErr, errors.New(message))
			}
			logger.Debug().
				Err(conditionsErr).
				Msg("Gateway Not Accepted")
		}

		// The refused ListenerSets have no listener, and got their status when refused.
		listenerSetListeners := make(map[*gatev1.ListenerSet][]gatewayListener)
		for _, listener := range gwl.listeners {
			if listener.ListenerSet != nil {
				listenerSetListeners[listener.ListenerSet] = append(listenerSetListeners[listener.ListenerSet], listener)
			}
		}

		var attachedListenerSets int32
		for listenerSet, listeners := range listenerSetListeners {
			listenerSetStatus, listenerSetAccepted := makeListenerSetStatus(listenerSet.Generation, listeners, accepted)
			statusReport.RecordListenerSetStatus(ktypes.NamespacedName{Namespace: listenerSet.Namespace, Name: listenerSet.Name}, listenerSetStatus)
			if listenerSetAccepted {
				attachedListenerSets++
			}
		}
		gatewayStatus.AttachedListenerSets = &attachedListenerSets

		statusReport.RecordGatewayStatus(ktypes.NamespacedName{Name: gwl.gateway.Name, Namespace: gwl.gateway.Namespace}, gatewayStatus)
	}

	return conf, statusReport, nil
}

// loadHTTPAndGRPCRoutes loads the HTTPRoutes and the GRPCRoutes together, ordered by
// creation timestamp then by "{namespace}/{name}".
// As stated in the specification, when an HTTPRoute and a GRPCRoute attached to the same
// listener have intersecting hostnames, only the first one in that order is accepted.
func (p *Provider) loadHTTPAndGRPCRoutes(ctx context.Context, gateways []gatewayWithListeners, conf *dynamic.Configuration, statusReport *statusReport) {
	routes := make([]metav1.Object, 0)

	httpRoutes, err := p.client.ListHTTPRoutes()
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Unable to list HTTPRoutes")
	}
	for _, route := range httpRoutes {
		routes = append(routes, route)
	}

	grpcRoutes, err := p.client.ListGRPCRoutes()
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Msg("Unable to list GRPCRoutes")
	}
	for _, route := range grpcRoutes {
		routes = append(routes, route)
	}

	slices.SortStableFunc(routes, compareRoutes)

	attached := make(attachedRoutes)
	served := make(servedRules)
	for _, route := range routes {
		switch route := route.(type) {
		case *gatev1.HTTPRoute:
			p.loadHTTPRoute(ctx, gateways, route, conf, attached, served, statusReport)
		case *gatev1.GRPCRoute:
			p.loadGRPCRoute(ctx, gateways, route, conf, attached, served, statusReport)
		}
	}
}

// loadGatewayListeners loads the listeners of the given ListenerSet, or of the given Gateway when nil,
// and records what the valid ones claim in claims.
func (p *Provider) loadGatewayListeners(ctx context.Context, gateway *gatev1.Gateway, listenerSet *gatev1.ListenerSet, claims *listenerClaims, conf *dynamic.Configuration) []gatewayListener {
	listeners, generation := gateway.Spec.Listeners, gateway.Generation
	// The resource declaring the listeners drives the ReferenceGrant checks,
	// and the namespace "Same" resolves to in allowedRoutes.
	ownerKind, ownerNamespace := kindGateway, gateway.Namespace

	if listenerSet != nil {
		// A ListenerEntry mirrors a Gateway Listener field for field.
		listeners = make([]gatev1.Listener, 0, len(listenerSet.Spec.Listeners))
		for _, entry := range listenerSet.Spec.Listeners {
			listeners = append(listeners, gatev1.Listener(entry))
		}
		generation = listenerSet.Generation
		ownerKind, ownerNamespace = kindListenerSet, listenerSet.Namespace
	}

	tlsCerts := make(map[string]*tls.CertAndStores)
	gatewayListeners := make([]gatewayListener, len(listeners))

	var defaultFrontendValidation frontendValidation
	frontendValidationPerPort := make(map[gatev1.PortNumber]frontendValidation)

	if gateway.Spec.TLS != nil && gateway.Spec.TLS.Frontend != nil {
		defaultFrontendValidation = p.resolveFrontendValidation(gateway, gateway.Spec.TLS.Frontend.Default.Validation)
		for _, entry := range gateway.Spec.TLS.Frontend.PerPort {
			frontendValidationPerPort[entry.Port] = p.resolveFrontendValidation(gateway, entry.TLS.Validation)
		}
	}

	for i, listener := range listeners {
		gatewayListeners[i] = gatewayListener{
			SectionName: string(listener.Name),
			Port:        listener.Port,
			Protocol:    listener.Protocol,
			TLS:         listener.TLS,
			Hostname:    listener.Hostname,
			ListenerSet: listenerSet,
			Status: &gatev1.ListenerStatus{
				Name:           listener.Name,
				SupportedKinds: []gatev1.RouteGroupKind{},
				Conditions:     []metav1.Condition{},
			},
		}

		// The listener protocol is validated first, so that an unsupported protocol
		// is reported as such instead of being masked by the entryPoint lookup,
		// which cannot succeed for a protocol Traefik does not know about.
		supportedKinds, conditions := supportedRouteKinds(generation, listener.Protocol)
		if len(conditions) > 0 {
			gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, conditions...)
			continue
		}

		ep, err := p.entryPointName(listener.Port, listener.Protocol)
		if err != nil {
			// update "Detached" status with "PortUnavailable" reason
			gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, metav1.Condition{
				Type:               string(gatev1.ListenerConditionAccepted),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.ListenerReasonPortUnavailable),
				Message:            fmt.Sprintf("Cannot find entryPoint for %s: %v", ownerKind, err),
			})

			continue
		}
		gatewayListeners[i].EPName = ep

		allowedRoutes := ptr.Deref(listener.AllowedRoutes, gatev1.AllowedRoutes{Namespaces: &gatev1.RouteNamespaces{From: new(gatev1.NamespacesFromSame)}})
		gatewayListeners[i].AllowedNamespaces, err = p.allowedNamespaces(ownerNamespace, allowedRoutes.Namespaces)
		if err != nil {
			// update "ResolvedRefs" status true with "InvalidRoutesRef" reason
			gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, metav1.Condition{
				Type:               string(gatev1.ListenerConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: generation,
				LastTransitionTime: metav1.Now(),
				Reason:             "InvalidRouteNamespacesSelector", // Should never happen as the selector is validated by kubernetes
				Message:            fmt.Sprintf("Invalid route namespaces selector: %v", err),
			})

			continue
		}

		routeKinds, conditions := allowedRouteKinds(generation, listener, supportedKinds)
		for _, kind := range routeKinds {
			gatewayListeners[i].AllowedRouteKinds = append(gatewayListeners[i].AllowedRouteKinds, string(kind.Kind))
		}
		gatewayListeners[i].Status.SupportedKinds = routeKinds
		if len(conditions) > 0 {
			gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, conditions...)
			continue
		}

		listenerKey := makeListenerKey(listener)

		if _, ok := claims.listeners[listenerKey]; ok {
			gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions,
				makeListenerConflictConditions(generation, gatev1.ListenerReasonHostnameConflict, "A listener with the same protocol, port and hostname already exists")...)

			continue
		}

		// A TCP listener cannot share its port with an HTTP, HTTPS or TLS listener,
		// as the connections could not be matched to a single listener (Gateway API listener distinctness).
		isTCP := listener.Protocol == gatev1.TCPProtocolType
		if tcp, claimed := claims.tcpPorts[listener.Port]; claimed && tcp != isTCP {
			gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions,
				makeListenerConflictConditions(generation, gatev1.ListenerReasonProtocolConflict, "A listener with an incompatible protocol already exists on the same port")...)

			continue
		}

		claims.listeners[listenerKey] = struct{}{}
		claims.tcpPorts[listener.Port] = isTCP

		if (listener.Protocol == gatev1.HTTPProtocolType || listener.Protocol == gatev1.TCPProtocolType) && listener.TLS != nil {
			gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, metav1.Condition{
				Type:               string(gatev1.ListenerConditionAccepted),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: generation,
				LastTransitionTime: metav1.Now(),
				Reason:             "InvalidTLSConfiguration", // TODO check the spec if a proper reason is introduced at some point
				Message:            "TLS configuration must no be defined when using HTTP or TCP protocol",
			})

			continue
		}

		// TLS
		if listener.Protocol == gatev1.HTTPSProtocolType || listener.Protocol == gatev1.TLSProtocolType {
			if listener.TLS == nil {
				gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, metav1.Condition{
					Type:               string(gatev1.ListenerConditionAccepted),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: generation,
					LastTransitionTime: metav1.Now(),
					Reason:             "InvalidTLSConfiguration", // TODO check the spec if a proper reason is introduced at some point
					Message:            fmt.Sprintf("No TLS configuration for %s Listener %s:%d and protocol %q", ownerKind, listener.Name, listener.Port, listener.Protocol),
				})
				continue
			}

			tlsMode := ptr.Deref(listener.TLS.Mode, gatev1.TLSModeTerminate)
			isTLSPassthrough := tlsMode == gatev1.TLSModePassthrough

			if isTLSPassthrough && len(listener.TLS.CertificateRefs) > 0 {
				log.Ctx(ctx).Warn().Msg("In case of Passthrough TLS mode, no TLS settings take effect as the TLS session from the client is NOT terminated at the Gateway")
			}

			// Allowed configurations:
			// Protocol TLS -> Passthrough -> TLSRoute
			// Protocol TLS -> Terminate -> TLSRoute
			// Protocol HTTPS -> Terminate -> HTTPRoute
			if isTLSPassthrough && listener.Protocol == gatev1.HTTPSProtocolType {
				gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, metav1.Condition{
					Type:               string(gatev1.ListenerConditionAccepted),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.ListenerReasonUnsupportedProtocol),
					Message:            "HTTPS protocol is not supported with TLS mode Passthrough",
				})
				continue
			}

			if !isTLSPassthrough {
				if len(listener.TLS.CertificateRefs) == 0 {
					gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, metav1.Condition{
						Type:               string(gatev1.ListenerConditionResolvedRefs),
						Status:             metav1.ConditionFalse,
						ObservedGeneration: generation,
						LastTransitionTime: metav1.Now(),
						Reason:             string(gatev1.ListenerReasonInvalidCertificateRef),
						Message:            "One TLS CertificateRef is required in Terminate mode",
					})
					continue
				}

				var errCertConditions []metav1.Condition
				listenerTLSCerts := make(map[string]*tls.CertAndStores)
				for _, certificateRef := range listener.TLS.CertificateRefs {
					if certificateRef.Kind == nil || *certificateRef.Kind != kindSecret || certificateRef.Group == nil || (*certificateRef.Group != "" && *certificateRef.Group != groupCore) {
						errCertConditions = append(errCertConditions, metav1.Condition{
							Type:               string(gatev1.ListenerConditionResolvedRefs),
							Status:             metav1.ConditionFalse,
							ObservedGeneration: generation,
							LastTransitionTime: metav1.Now(),
							Reason:             string(gatev1.ListenerReasonInvalidCertificateRef),
							Message:            fmt.Sprintf("Unsupported TLS CertificateRef group/kind: %s/%s", groupToString(certificateRef.Group), kindToString(certificateRef.Kind)),
						})
						continue
					}

					certificateNamespace := string(ptr.Deref(certificateRef.Namespace, gatev1.Namespace(ownerNamespace)))
					if err := p.isReferenceGranted(ownerKind, ownerNamespace, groupCore, kindSecret, string(certificateRef.Name), certificateNamespace); err != nil {
						errCertConditions = append(errCertConditions, metav1.Condition{
							Type:               string(gatev1.ListenerConditionResolvedRefs),
							Status:             metav1.ConditionFalse,
							ObservedGeneration: generation,
							LastTransitionTime: metav1.Now(),
							Reason:             string(gatev1.ListenerReasonRefNotPermitted),
							Message:            fmt.Sprintf("Cannot reference CertificateRef %s/%s: %s", certificateNamespace, certificateRef.Name, err),
						})
						continue
					}

					configKey := certificateNamespace + "/" + string(certificateRef.Name)
					if _, tlsExists := listenerTLSCerts[configKey]; !tlsExists {
						tlsCert, err := p.getTLSCert(certificateRef.Name, certificateNamespace)
						if err != nil {
							errCertConditions = append(errCertConditions, metav1.Condition{
								Type:               string(gatev1.ListenerConditionResolvedRefs),
								Status:             metav1.ConditionFalse,
								ObservedGeneration: generation,
								LastTransitionTime: metav1.Now(),
								Reason:             string(gatev1.ListenerReasonInvalidCertificateRef),
								Message:            fmt.Sprintf("Cannot load CertificateRef %s/%s: %s", certificateNamespace, certificateRef.Name, err),
							})
							continue
						}
						listenerTLSCerts[configKey] = tlsCert
					}
				}

				if len(errCertConditions) > 0 {
					gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, errCertConditions...)
					gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions, metav1.Condition{
						Type:               string(gatev1.ListenerConditionProgrammed),
						Status:             metav1.ConditionFalse,
						ObservedGeneration: generation,
						LastTransitionTime: metav1.Now(),
						Reason:             string(gatev1.ListenerReasonInvalid),
						Message:            "Invalid CertificateRefs",
					})
					continue
				}

				// Frontend validation only applies to HTTPS listeners.
				if listener.Protocol == gatev1.HTTPSProtocolType {
					frontendValidation := defaultFrontendValidation
					if perPortFV, ok := frontendValidationPerPort[listener.Port]; ok {
						frontendValidation = perPortFV
					}

					// An invalid CA certificate reference leaves the listener unattached, like an invalid certificate reference.
					if frontendValidation.errCondition != nil {
						gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions,
							*frontendValidation.errCondition,
							metav1.Condition{
								Type:               string(gatev1.ListenerConditionAccepted),
								Status:             metav1.ConditionFalse,
								ObservedGeneration: generation,
								LastTransitionTime: metav1.Now(),
								Reason:             string(gatev1.ListenerReasonNoValidCACertificate),
								Message:            "No valid CA certificate found in CACertificateRefs",
							},
							metav1.Condition{
								Type:               string(gatev1.ListenerConditionProgrammed),
								Status:             metav1.ConditionFalse,
								ObservedGeneration: generation,
								LastTransitionTime: metav1.Now(),
								Reason:             string(gatev1.ListenerReasonInvalid),
								Message:            "Invalid CA certificate configuration",
							})
						continue
					}

					// Gateways sharing a port must declare the same frontend validation, the oldest Gateway wins.
					if applied, ok := claims.frontendValidations[listener.Port]; ok && !reflect.DeepEqual(applied, frontendValidation.clientAuth) {
						gatewayListeners[i].Status.Conditions = append(gatewayListeners[i].Status.Conditions,
							makeListenerConflictConditions(generation, listenerReasonFrontendValidationConflict, "An older Gateway already applies a different frontend validation on this port")...)
						continue
					}

					claims.frontendValidations[listener.Port] = frontendValidation.clientAuth
				}

				// Only copy if the certificate TLS config is not already known.
				for key, listenerTLSCert := range listenerTLSCerts {
					if _, ok := tlsCerts[key]; !ok {
						tlsCerts[key] = listenerTLSCert
					}
				}
			}
		}

		gatewayListeners[i].Attached = true
	}

	if len(tlsCerts) > 0 {
		conf.TLS.Certificates = append(conf.TLS.Certificates, getTLSConfig(tlsCerts)...)
	}

	return gatewayListeners
}

// uniqListener identifies a unique listener configuration.
// The protocol is part of it because a TLS parent router and a plain one are built on the two
// distinct handlers of an entry point, and because merging them would put the routes of an HTTP
// listener behind the TLS configuration of an HTTPS one.
//
// TODO: The listener TLS options (listener.TLS.Options) have to be part of it once supported.
type uniqListener struct {
	epName   string
	protocol gatev1.ProtocolType
	port     gatev1.PortNumber
}

// buildListenerRouters builds a parent router per entry point hostname, scoped to the requests it
// is the most specific match for. Electing the listener per Gateway isolates the listeners of a
// Gateway without hiding the routes of the other Gateways sharing the entry point.
func (p *Provider) buildListenerRouters(gateways []gatewayWithListeners, frontendValidations map[gatev1.PortNumber]*tls.ClientAuth, conf *dynamic.Configuration) []string {
	var listenerRouterNames []string

	hostnamesByListener := map[uniqListener][]string{}
	for _, gateway := range gateways {
		for _, listener := range gateway.listeners {
			if !listener.Attached ||
				listener.Protocol != gatev1.HTTPProtocolType && listener.Protocol != gatev1.HTTPSProtocolType {
				continue
			}

			uniq := uniqListener{epName: listener.EPName, protocol: listener.Protocol, port: listener.Port}

			hostname := string(ptr.Deref(listener.Hostname, ""))
			if !slices.Contains(hostnamesByListener[uniq], hostname) {
				hostnamesByListener[uniq] = append(hostnamesByListener[uniq], hostname)
			}
		}
	}

	uniqListeners := slices.SortedFunc(maps.Keys(hostnamesByListener), func(a, b uniqListener) int {
		return cmp.Or(cmp.Compare(a.epName, b.epName), cmp.Compare(a.protocol, b.protocol))
	})

	for _, uniq := range uniqListeners {
		hostnames := hostnamesByListener[uniq]
		slices.Sort(hostnames)

		for _, hostname := range hostnames {
			listenerRouterName := makeListenerRouterName(uniq, hostname)

			listenerRouter := &dynamic.Router{
				Rule:        buildListenerRule(hostname, hostnames),
				EntryPoints: []string{uniq.epName},
			}

			for _, gateway := range gateways {
				listener := mostSpecificListener(gateway.listeners, uniq, hostname)
				if listener == nil {
					continue
				}

				listener.RouterNames = append(listener.RouterNames, listenerRouterName)
			}

			if uniq.protocol == gatev1.HTTPSProtocolType {
				listenerTLSOptions := tls.Options{}
				listenerTLSOptions.SetDefaults()

				if validation := frontendValidations[uniq.port]; validation != nil {
					listenerTLSOptions.ClientAuth = *validation
				}

				conf.TLS.Options[listenerRouterName] = listenerTLSOptions

				listenerRouter.TLS = &dynamic.RouterTLSConfig{
					Options: listenerRouterName,
				}
			}

			conf.HTTP.Routers[listenerRouterName] = listenerRouter
			listenerRouterNames = append(listenerRouterNames, listenerRouterName)
		}
	}

	return listenerRouterNames
}

// dropChildlessListenerRouters removes the parent routers no route is attached to.
func dropChildlessListenerRouters(conf *dynamic.Configuration, listenerRouterNames []string) {
	parents := map[string]struct{}{}
	for _, router := range conf.HTTP.Routers {
		for _, parent := range router.ParentRefs {
			parents[parent] = struct{}{}
		}
	}

	for _, name := range listenerRouterNames {
		if _, ok := parents[name]; ok {
			continue
		}

		delete(conf.HTTP.Routers, name)
		delete(conf.TLS.Options, name)
	}
}

func mostSpecificListener(listeners []gatewayListener, uniq uniqListener, hostname string) *gatewayListener {
	var elected *gatewayListener
	for i, listener := range listeners {
		if listener.EPName != uniq.epName || listener.Protocol != uniq.protocol || !listener.Attached {
			continue
		}

		listenerHostname := string(ptr.Deref(listener.Hostname, ""))
		if !hostnameCovers(listenerHostname, hostname) {
			continue
		}

		// A hostname covered by the elected one is a more specific match.
		if elected == nil || hostnameCovers(string(ptr.Deref(elected.Hostname, "")), listenerHostname) {
			elected = &listeners[i]
		}
	}

	return elected
}

// hostnameCovers reports whether every request matching hostname also matches listenerHostname.
func hostnameCovers(listenerHostname, hostname string) bool {
	if listenerHostname == "" {
		return true
	}

	return findMatchingHostname(gatev1.Hostname(listenerHostname), gatev1.Hostname(hostname)) != ""
}

func buildListenerRule(hostname string, entryPointHostnames []string) string {
	rule := `Host("*")`
	if hostname != "" {
		rule = fmt.Sprintf("Host(%q)", hostnameMatcherValue(hostname))
	}

	var exclusions []string
	for _, entryPointHostname := range entryPointHostnames {
		if entryPointHostname == hostname || !hostnameCovers(hostname, entryPointHostname) {
			continue
		}

		exclusions = append(exclusions, fmt.Sprintf("Host(%q)", hostnameMatcherValue(entryPointHostname)))
	}

	if len(exclusions) == 0 {
		return rule
	}

	return fmt.Sprintf("(%s) && !(%s)", rule, strings.Join(exclusions, " || "))
}

// hostnameMatcherValue returns the Host matcher value for a Gateway API hostname,
// whose wildcard spans one or more labels, unlike the Traefik single one.
func hostnameMatcherValue(hostname string) string {
	if suffix, ok := strings.CutPrefix(hostname, "*."); ok {
		return "**." + suffix
	}

	return hostname
}

// loadListenerSetListeners loads the listeners of the ListenerSets referencing the given Gateway,
// and reports the status of the ones its AllowedListeners policy refuses.
func (p *Provider) loadListenerSetListeners(ctx context.Context, gateway *gatev1.Gateway, listenerSets []*gatev1.ListenerSet, claims *listenerClaims, conf *dynamic.Configuration, statusReport *statusReport) []gatewayListener {
	var listeners []gatewayListener
	for _, listenerSet := range listenerSets {
		if !listenerSetRefsGateway(listenerSet, gateway) {
			continue
		}

		if !p.isListenerSetAllowed(ctx, gateway, listenerSet) {
			log.Ctx(ctx).Debug().
				Str("listenerset", listenerSet.Name).
				Str("namespace", listenerSet.Namespace).
				Msg("ListenerSet not allowed by Gateway's AllowedListeners")

			const message = "ListenerSet is not allowed by the Gateway's AllowedListeners policy"
			statusReport.RecordListenerSetStatus(ktypes.NamespacedName{Namespace: listenerSet.Namespace, Name: listenerSet.Name}, gatev1.ListenerSetStatus{
				Conditions: []metav1.Condition{
					makeListenerSetCondition(gatev1.ListenerSetConditionAccepted, metav1.ConditionFalse, listenerSet.Generation, string(gatev1.ListenerSetReasonNotAllowed), message),
					makeListenerSetCondition(gatev1.ListenerSetConditionProgrammed, metav1.ConditionFalse, listenerSet.Generation, string(gatev1.ListenerSetReasonNotAllowed), message),
				},
			})

			continue
		}

		listeners = append(listeners, p.loadGatewayListeners(ctx, gateway, listenerSet, claims, conf)...)
	}

	return listeners
}

type frontendValidation struct {
	clientAuth   *tls.ClientAuth
	errCondition *metav1.Condition
}

// resolveFrontendValidation resolves the client authentication of a frontend validation.
func (p *Provider) resolveFrontendValidation(gateway *gatev1.Gateway, validation *gatev1.FrontendTLSValidation) frontendValidation {
	if validation == nil {
		return frontendValidation{}
	}

	caCert, errCondition := p.resolveCACertificateRefs(gateway, validation.CACertificateRefs)
	if errCondition != nil {
		return frontendValidation{errCondition: errCondition}
	}

	clientAuthType := tls.RequireAndVerifyClientCert
	if validation.Mode == gatev1.AllowInsecureFallback {
		clientAuthType = tls.RequestClientCert
	}

	return frontendValidation{
		clientAuth: &tls.ClientAuth{
			CAFiles:        []types.FileOrContent{types.FileOrContent(caCert)},
			ClientAuthType: clientAuthType,
		},
	}
}

// resolveCACertificateRefs returns the referenced CA certificate, or the ResolvedRefs condition explaining why it cannot.
// Only the Core support is implemented: a single ConfigMap reference with a ca.crt key.
func (p *Provider) resolveCACertificateRefs(gateway *gatev1.Gateway, refs []gatev1.ObjectReference) (string, *metav1.Condition) {
	if len(refs) != 1 {
		return "", &metav1.Condition{
			Type:               string(gatev1.ListenerConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.ListenerReasonInvalidCACertificateRef),
			Message:            fmt.Sprintf("Unsupported number of CACertificateRefs: %d, only one reference is supported", len(refs)),
		}
	}

	ref := refs[0]
	if (ref.Group != "" && ref.Group != groupCore) || ref.Kind != kindConfigMap {
		return "", &metav1.Condition{
			Type:               string(gatev1.ListenerConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.ListenerReasonInvalidCACertificateKind),
			Message:            fmt.Sprintf("Unsupported CACertificateRef group/kind: %s/%s", ref.Group, ref.Kind),
		}
	}

	refNamespace := string(ptr.Deref(ref.Namespace, gatev1.Namespace(gateway.Namespace)))
	if err := p.isReferenceGranted(kindGateway, gateway.Namespace, groupCore, string(ref.Kind), string(ref.Name), refNamespace); err != nil {
		return "", &metav1.Condition{
			Type:               string(gatev1.ListenerConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.ListenerReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot reference CACertificateRef: %s/%s: %s", ref.Group, ref.Kind, err),
		}
	}

	cm, err := p.client.GetConfigMap(refNamespace, string(ref.Name))
	if err != nil {
		return "", &metav1.Condition{
			Type:               string(gatev1.ListenerConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.ListenerReasonInvalidCACertificateRef),
			Message:            fmt.Sprintf("Cannot resolve configmap: %s/%s: %s", refNamespace, ref.Name, err),
		}
	}

	caCert := cm.Data["ca.crt"]
	if caCert == "" {
		return "", &metav1.Condition{
			Type:               string(gatev1.ListenerConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.ListenerReasonInvalidCACertificateRef),
			Message:            fmt.Sprintf("Cannot find ca.crt: %s %s/%s", ref.Kind, refNamespace, ref.Name),
		}
	}

	return caCert, nil
}

// resolveBackendClientCertificate returns the client certificate the Gateway presents to its backends,
// nil when it has none, or the ResolvedRefs condition explaining why it cannot be resolved.
func (p *Provider) resolveBackendClientCertificate(gateway *gatev1.Gateway) (*tls.Certificate, *metav1.Condition) {
	if gateway.Spec.TLS == nil || gateway.Spec.TLS.Backend == nil || gateway.Spec.TLS.Backend.ClientCertificateRef == nil {
		return nil, nil
	}

	ref := gateway.Spec.TLS.Backend.ClientCertificateRef
	refGroup := ptr.Deref(ref.Group, "")
	refKind := ptr.Deref(ref.Kind, kindSecret)
	if (refGroup != "" && refGroup != groupCore) || refKind != kindSecret {
		return nil, makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionFalse, gatev1.GatewayReasonInvalidClientCertificateRef,
			fmt.Sprintf("Unsupported ClientCertificateRef group/kind: %s/%s", refGroup, refKind))
	}

	refNamespace := string(ptr.Deref(ref.Namespace, gatev1.Namespace(gateway.Namespace)))
	if err := p.isReferenceGranted(kindGateway, gateway.Namespace, groupCore, kindSecret, string(ref.Name), refNamespace); err != nil {
		return nil, makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionFalse, gatev1.GatewayReasonRefNotPermitted,
			fmt.Sprintf("Cannot reference ClientCertificateRef %s/%s: %s", refNamespace, ref.Name, err))
	}

	cert, err := p.getTLSCert(ref.Name, refNamespace)
	if err != nil {
		return nil, makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionFalse, gatev1.GatewayReasonInvalidClientCertificateRef,
			fmt.Sprintf("Cannot load ClientCertificateRef %s/%s: %s", refNamespace, ref.Name, err))
	}

	return &cert.Certificate, nil
}

func makeGatewayResolvedRefsCondition(generation int64, status metav1.ConditionStatus, reason gatev1.GatewayConditionReason, message string) *metav1.Condition {
	return &metav1.Condition{
		Type:               string(gatev1.GatewayConditionResolvedRefs),
		Status:             status,
		ObservedGeneration: generation,
		LastTransitionTime: metav1.Now(),
		Reason:             string(reason),
		Message:            message,
	}
}

// gatewayResolvedRefsCondition returns the Gateway ResolvedRefs condition: the backend client certificate condition when
// it cannot be resolved, otherwise whether the listeners declared by the Gateway itself have resolved their references.
func gatewayResolvedRefsCondition(gateway *gatev1.Gateway, listeners []gatewayListener, clientCertificateCondition *metav1.Condition) metav1.Condition {
	if clientCertificateCondition != nil {
		return *clientCertificateCondition
	}

	unresolved := slices.ContainsFunc(listeners, func(listener gatewayListener) bool {
		return listener.ListenerSet == nil && meta.IsStatusConditionFalse(listener.Status.Conditions, string(gatev1.ListenerConditionResolvedRefs))
	})
	if unresolved {
		return *makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionFalse, gatev1.GatewayReasonListenersNotResolved,
			"At least one Listener has unresolved references")
	}

	return *makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionTrue, gatev1.GatewayReasonResolvedRefs, conditionNoErrorMessage)
}

func (p *Provider) makeGatewayStatus(gateway *gatev1.Gateway, listeners []gatewayListener, clientCertificateCondition *metav1.Condition, addresses []gatev1.GatewayStatusAddress, accepted bool) (gatev1.GatewayStatus, []metav1.Condition) {
	gatewayStatus := gatev1.GatewayStatus{
		Addresses: addresses,
		// ResolvedRefs is reported whatever the outcome of the Accepted and Programmed conditions.
		Conditions: []metav1.Condition{gatewayResolvedRefsCondition(gateway, listeners, clientCertificateCondition)},
	}

	// The spec requires the InsecureFrontendValidationMode condition when any validation allows an insecure fallback.
	if hasInsecureFrontendValidationMode(gateway) {
		gatewayStatus.Conditions = append(gatewayStatus.Conditions, metav1.Condition{
			Type:               string(gatev1.GatewayConditionInsecureFrontendValidationMode),
			Status:             metav1.ConditionTrue,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.GatewayReasonConfigurationChanged),
			Message:            "FrontendValidationMode is set to AllowInsecureFallback",
		})
	}

	var errorConditions []metav1.Condition
	for _, listener := range listeners {
		// The ListenerSet listeners are reported in their own ListenerSet status.
		if listener.ListenerSet != nil {
			continue
		}

		if len(listener.Status.Conditions) == 0 {
			listener.Status.Conditions = append(listener.Status.Conditions,
				metav1.Condition{
					Type:               string(gatev1.ListenerConditionAccepted),
					Status:             metav1.ConditionTrue,
					ObservedGeneration: gateway.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.ListenerReasonAccepted),
					Message:            conditionNoErrorMessage,
				},
				metav1.Condition{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionTrue,
					ObservedGeneration: gateway.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.ListenerReasonResolvedRefs),
					Message:            conditionNoErrorMessage,
				},
				metav1.Condition{
					Type:               string(gatev1.ListenerConditionProgrammed),
					Status:             metav1.ConditionTrue,
					ObservedGeneration: gateway.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.ListenerReasonProgrammed),
					Message:            conditionNoErrorMessage,
				},
			)

			// TODO: refactor
			gatewayStatus.Listeners = append(gatewayStatus.Listeners, *listener.Status)
			continue
		}

		errorConditions = append(errorConditions, listener.Status.Conditions...)
		gatewayStatus.Listeners = append(gatewayStatus.Listeners, *listener.Status)
	}

	// Traefik supports no infrastructure parameters, and the specification requires
	// a parametersRef that cannot be resolved to be reported instead of ignored.
	if gateway.Spec.Infrastructure != nil && gateway.Spec.Infrastructure.ParametersRef != nil {
		condition := metav1.Condition{
			Type:               string(gatev1.GatewayConditionAccepted),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.GatewayReasonInvalidParameters),
			Message:            "Gateway infrastructure parametersRef is not supported",
		}
		gatewayStatus.Conditions = append(gatewayStatus.Conditions, condition)

		return gatewayStatus, append(errorConditions, condition)
	}

	if !accepted {
		gatewayStatus.Conditions = append(gatewayStatus.Conditions,
			// update "Accepted" status with "Accepted" reason
			metav1.Condition{
				Type:               string(gatev1.GatewayConditionAccepted),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: gateway.Generation,
				Reason:             string(gatev1.GatewayReasonListenersNotValid),
				Message:            "At least one Listener must be valid",
				LastTransitionTime: metav1.Now(),
			},
			// update "Programmed" status with "Programmed" reason
			metav1.Condition{
				Type:               string(gatev1.GatewayConditionProgrammed),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: gateway.Generation,
				Reason:             string(gatev1.GatewayReasonInvalid),
				Message:            "No Listener is valid",
				LastTransitionTime: metav1.Now(),
			},
		)

		return gatewayStatus, errorConditions
	}

	acceptedConditionReason := gatev1.GatewayReasonAccepted
	acceptedConditionMessage := "Gateway successfully scheduled"
	programmedConditionMessage := "Gateway successfully programmed"
	if len(errorConditions) > 0 {
		acceptedConditionReason = gatev1.GatewayReasonListenersNotValid
		acceptedConditionMessage = "Gateway successfully scheduled, but some Listeners are not valid"
		programmedConditionMessage = "Gateway successfully programmed, but some Listeners are not valid"
	}

	gatewayStatus.Conditions = append(gatewayStatus.Conditions,
		// update "Accepted" status with "Accepted" reason
		metav1.Condition{
			Type:               string(gatev1.GatewayConditionAccepted),
			Status:             metav1.ConditionTrue,
			ObservedGeneration: gateway.Generation,
			Reason:             string(acceptedConditionReason),
			Message:            acceptedConditionMessage,
			LastTransitionTime: metav1.Now(),
		},
		// update "Programmed" status with "Programmed" reason
		metav1.Condition{
			Type:               string(gatev1.GatewayConditionProgrammed),
			Status:             metav1.ConditionTrue,
			ObservedGeneration: gateway.Generation,
			Reason:             string(gatev1.GatewayReasonProgrammed),
			Message:            programmedConditionMessage,
			LastTransitionTime: metav1.Now(),
		},
	)

	return gatewayStatus, nil
}

func (p *Provider) gatewayAddresses() ([]gatev1.GatewayStatusAddress, error) {
	if p.StatusAddress == nil {
		return nil, nil
	}

	if p.StatusAddress.IP != "" {
		return []gatev1.GatewayStatusAddress{{
			Type:  new(gatev1.IPAddressType),
			Value: p.StatusAddress.IP,
		}}, nil
	}

	if p.StatusAddress.Hostname != "" {
		return []gatev1.GatewayStatusAddress{{
			Type:  new(gatev1.HostnameAddressType),
			Value: p.StatusAddress.Hostname,
		}}, nil
	}

	svcRef := p.StatusAddress.Service
	if svcRef.Name != "" && svcRef.Namespace != "" {
		svc, err := p.client.GetService(svcRef.Namespace, svcRef.Name)
		if err != nil {
			return nil, fmt.Errorf("getting service: %w", err)
		}

		var addresses []gatev1.GatewayStatusAddress
		for _, addr := range svc.Status.LoadBalancer.Ingress {
			switch {
			case addr.IP != "":
				addresses = append(addresses, gatev1.GatewayStatusAddress{
					Type:  new(gatev1.IPAddressType),
					Value: addr.IP,
				})

			case addr.Hostname != "":
				addresses = append(addresses, gatev1.GatewayStatusAddress{
					Type:  new(gatev1.HostnameAddressType),
					Value: addr.Hostname,
				})
			}
		}
		return addresses, nil
	}

	return nil, errors.New("empty Gateway status address configuration")
}

func (p *Provider) entryPointName(port gatev1.PortNumber, protocol gatev1.ProtocolType) (string, error) {
	portStr := strconv.FormatInt(int64(port), 10)

	for name, entryPoint := range p.EntryPoints {
		if strings.HasSuffix(entryPoint.Address, ":"+portStr) {
			// If the protocol is HTTP the entryPoint must have no TLS conf
			// Not relevant for gatev1.TLSProtocolType && gatev1.TCPProtocolType
			if protocol == gatev1.HTTPProtocolType && entryPoint.HasHTTPTLSConf {
				continue
			}

			return name, nil
		}
	}

	return "", fmt.Errorf("no matching entryPoint for port %d and protocol %q", port, protocol)
}

func (p *Provider) isReferenceGranted(fromKind, fromNamespace, toGroup, toKind, toName, toNamespace string) error {
	if toNamespace == fromNamespace {
		return nil
	}

	refGrants, err := p.client.ListReferenceGrants(toNamespace)
	if err != nil {
		return fmt.Errorf("listing ReferenceGrant: %w", err)
	}

	refGrants = filterReferenceGrantsFrom(refGrants, groupGateway, fromKind, fromNamespace)
	refGrants = filterReferenceGrantsTo(refGrants, toGroup, toKind, toName)
	if len(refGrants) == 0 {
		return errors.New("missing ReferenceGrant")
	}

	return nil
}

func (p *Provider) getTLSCert(secretName gatev1.ObjectName, namespace string) (*tls.CertAndStores, error) {
	secret, err := p.client.GetSecret(namespace, string(secretName))
	if err != nil {
		return nil, fmt.Errorf("getting secret: %w", err)
	}

	cert, key, err := getCertificateBlocks(secret, namespace, string(secretName))
	if err != nil {
		return nil, fmt.Errorf("getting certificate blocks: %w", err)
	}

	certAndStore := &tls.CertAndStores{
		CertFile: types.FileOrContent(cert),
		KeyFile:  types.FileOrContent(key),
	}
	if _, err := certAndStore.GetCertificate(); err != nil {
		return nil, fmt.Errorf("validating certificate: %w", err)
	}

	return certAndStore, nil
}

func (p *Provider) allowedNamespaces(gatewayNamespace string, routeNamespaces *gatev1.RouteNamespaces) ([]string, error) {
	if routeNamespaces == nil || routeNamespaces.From == nil {
		return []string{gatewayNamespace}, nil
	}

	switch *routeNamespaces.From {
	case gatev1.NamespacesFromAll:
		return []string{metav1.NamespaceAll}, nil

	case gatev1.NamespacesFromSame:
		return []string{gatewayNamespace}, nil

	case gatev1.NamespacesFromSelector:
		selector, err := metav1.LabelSelectorAsSelector(routeNamespaces.Selector)
		if err != nil {
			return nil, fmt.Errorf("malformed selector: %w", err)
		}

		return p.client.ListNamespaces(selector)
	}

	return nil, fmt.Errorf("unsupported RouteSelectType: %q", *routeNamespaces.From)
}

type backendAddress struct {
	IP   string
	Port int32
}

func (p *Provider) getBackendAddresses(namespace string, ref gatev1.BackendRef) ([]backendAddress, corev1.ServicePort, error) {
	if ref.Port == nil {
		return nil, corev1.ServicePort{}, errors.New("port is required for Kubernetes Service reference")
	}

	service, err := p.client.GetService(namespace, string(ref.Name))
	if err != nil {
		return nil, corev1.ServicePort{}, fmt.Errorf("getting service: %w", err)
	}
	if service.Spec.Type == corev1.ServiceTypeExternalName {
		return nil, corev1.ServicePort{}, errors.New("type ExternalName is not supported for Kubernetes Service reference")
	}

	var svcPort *corev1.ServicePort
	for _, p := range service.Spec.Ports {
		if p.Port == *ref.Port {
			svcPort = &p
			break
		}
	}
	if svcPort == nil {
		return nil, corev1.ServicePort{}, fmt.Errorf("service port %d not found", *ref.Port)
	}

	annotationsConfig, err := parseServiceAnnotations(service.Annotations)
	if err != nil {
		return nil, corev1.ServicePort{}, fmt.Errorf("parsing service annotations config: %w", err)
	}

	nativeLB := p.NativeLBByDefault
	if annotationsConfig.Service.NativeLB != nil {
		nativeLB = *annotationsConfig.Service.NativeLB
	}

	if nativeLB {
		if service.Spec.ClusterIP == "" || service.Spec.ClusterIP == "None" {
			return nil, corev1.ServicePort{}, fmt.Errorf("no clusterIP found for service: %s/%s", service.Namespace, service.Name)
		}

		return []backendAddress{{
			IP:   service.Spec.ClusterIP,
			Port: svcPort.Port,
		}}, *svcPort, nil
	}

	endpointSlices, err := p.client.ListEndpointSlicesForService(namespace, string(ref.Name))
	if err != nil {
		return nil, corev1.ServicePort{}, fmt.Errorf("getting endpointslices: %w", err)
	}
	if len(endpointSlices) == 0 {
		return nil, corev1.ServicePort{}, errors.New("endpointslices not found")
	}

	uniqAddresses := map[string]struct{}{}
	backendServers := make([]backendAddress, 0)
	for _, endpointSlice := range endpointSlices {
		var port int32
		for _, p := range endpointSlice.Ports {
			if p.Name != nil && svcPort.Name == *p.Name {
				port = ptr.Deref(p.Port, 0)
				break
			}
		}
		if port == 0 {
			continue
		}

		for _, endpoint := range endpointSlice.Endpoints {
			if endpoint.Conditions.Ready == nil || !*endpoint.Conditions.Ready {
				continue
			}

			for _, address := range endpoint.Addresses {
				if _, ok := uniqAddresses[address]; ok {
					continue
				}

				uniqAddresses[address] = struct{}{}
				backendServers = append(backendServers, backendAddress{
					IP:   address,
					Port: port,
				})
			}
		}
	}

	return backendServers, *svcPort, nil
}

// hasInsecureFrontendValidationMode reports whether a validation of the Gateway allows an insecure fallback.
func hasInsecureFrontendValidationMode(gateway *gatev1.Gateway) bool {
	if gateway.Spec.TLS == nil {
		return false
	}

	frontend := gateway.Spec.TLS.Frontend
	if frontend == nil {
		return false
	}

	if frontend.Default.Validation != nil && frontend.Default.Validation.Mode == gatev1.AllowInsecureFallback {
		return true
	}

	for _, perPort := range frontend.PerPort {
		if perPort.TLS.Validation != nil && perPort.TLS.Validation.Mode == gatev1.AllowInsecureFallback {
			return true
		}
	}

	return false
}

func supportedRouteKinds(gatewayGeneration int64, protocol gatev1.ProtocolType) ([]gatev1.RouteGroupKind, []metav1.Condition) {
	group := gatev1.Group(gatev1.GroupName)

	switch protocol {
	case gatev1.TCPProtocolType:
		return []gatev1.RouteGroupKind{{Kind: kindTCPRoute, Group: &group}}, nil

	case gatev1.HTTPProtocolType, gatev1.HTTPSProtocolType:
		return []gatev1.RouteGroupKind{
			{Kind: kindHTTPRoute, Group: &group},
			{Kind: kindGRPCRoute, Group: &group},
		}, nil

	case gatev1.TLSProtocolType:
		return []gatev1.RouteGroupKind{
			{Kind: kindTLSRoute, Group: &group},
		}, nil
	}

	return nil, []metav1.Condition{{
		Type:               string(gatev1.ListenerConditionAccepted),
		Status:             metav1.ConditionFalse,
		ObservedGeneration: gatewayGeneration,
		LastTransitionTime: metav1.Now(),
		Reason:             string(gatev1.ListenerReasonUnsupportedProtocol),
		Message:            fmt.Sprintf("Unsupported listener protocol %q", protocol),
	}}
}

func allowedRouteKinds(generation int64, listener gatev1.Listener, supportedKinds []gatev1.RouteGroupKind) ([]gatev1.RouteGroupKind, []metav1.Condition) {
	if listener.AllowedRoutes == nil || len(listener.AllowedRoutes.Kinds) == 0 {
		return supportedKinds, nil
	}

	var conditions []metav1.Condition
	routeKinds := []gatev1.RouteGroupKind{}
	uniqRouteKinds := map[gatev1.Kind]struct{}{}
	for _, routeKind := range listener.AllowedRoutes.Kinds {
		var isSupported bool
		for _, kind := range supportedKinds {
			if routeKind.Kind == kind.Kind && routeKind.Group != nil && *routeKind.Group == *kind.Group {
				isSupported = true
				break
			}
		}

		if !isSupported {
			conditions = append(conditions, metav1.Condition{
				Type:               string(gatev1.ListenerConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.ListenerReasonInvalidRouteKinds),
				Message:            fmt.Sprintf("Listener protocol %q does not support RouteGroupKind %s/%s", listener.Protocol, groupToString(routeKind.Group), routeKind.Kind),
			})
			continue
		}

		if _, exists := uniqRouteKinds[routeKind.Kind]; !exists {
			routeKinds = append(routeKinds, routeKind)
			uniqRouteKinds[routeKind.Kind] = struct{}{}
		}
	}

	return routeKinds, conditions
}

func findMatchingHostnames(listenerHostname *gatev1.Hostname, routeHostnames []gatev1.Hostname) ([]gatev1.Hostname, bool) {
	if listenerHostname == nil {
		return routeHostnames, true
	}

	if len(routeHostnames) == 0 {
		return []gatev1.Hostname{*listenerHostname}, true
	}

	matches := map[gatev1.Hostname]struct{}{}
	for _, routeHostname := range routeHostnames {
		if match := findMatchingHostname(*listenerHostname, routeHostname); match != "" {
			matches[match] = struct{}{}
			continue
		}

		if match := findMatchingHostname(routeHostname, *listenerHostname); match != "" {
			matches[match] = struct{}{}
			continue
		}
	}

	return slices.Sorted(maps.Keys(matches)), len(matches) > 0
}

func findMatchingHostname(h1, h2 gatev1.Hostname) gatev1.Hostname {
	if h1 == h2 {
		return h1
	}

	if !strings.HasPrefix(string(h1), "*.") {
		return ""
	}

	trimmedH1 := strings.TrimPrefix(string(h1), "*")

	if !strings.HasSuffix(string(h2), trimmedH1) {
		return ""
	}

	// h1 is a wildcard that encompasses h2, so h2 is always
	// the more specific hostname (the correct intersection).
	return h2
}

func allowRoute(listener gatewayListener, routeNamespace, routeKind string) bool {
	if !slices.Contains(listener.AllowedRouteKinds, routeKind) {
		return false
	}

	return slices.ContainsFunc(listener.AllowedNamespaces, func(allowedNamespace string) bool {
		return allowedNamespace == corev1.NamespaceAll || allowedNamespace == routeNamespace
	})
}

// compareRoutes orders the routes as the specification asks for the ones matching
// a request equally well to be discriminated.
func compareRoutes(a, b metav1.Object) int {
	return cmp.Or(a.GetCreationTimestamp().Time.Compare(b.GetCreationTimestamp().Time),
		strings.Compare(a.GetNamespace(), b.GetNamespace()),
		strings.Compare(a.GetName(), b.GetName()))
}

// servedRuleKey identifies a rule within the scope its router competes in.
// A TCP router competes in the muxer of its entry points,
// and an HTTP router in the muxer of its listener parent routers, as it has no entry points of its own.
type servedRuleKey struct {
	Scope string
	Rule  string
}

// servedRules keeps the router that serves each rule.
// Two routers that compete and have the same rule match the same requests.
// Thus only the first router that a muxer evaluates can serve a request.
type servedRules map[servedRuleKey]string

// register keeps the given router as the router that serves its rule within the given scope.
// If a different router serves that rule already, register gives its name, and true.
func (sr servedRules) register(name string, scope []string, rule string) (string, bool) {
	key := servedRuleKey{Scope: strings.Join(scope, ","), Rule: rule}

	if servedBy, served := sr[key]; served {
		return servedBy, true
	}

	sr[key] = name

	return "", false
}

// listenerRef identifies a listener of a Gateway.
type listenerRef struct {
	SectionName      string
	GatewayNamespace string
	GatewayName      string
}

// attachedRoutes holds, for each listener, the kind of the route already attached to a hostname.
// It allows detecting the hostname Conflicts between the HTTPRoutes and the GRPCRoutes attached to the same listener.
type attachedRoutes map[listenerRef]map[gatev1.Hostname]string

func (ar attachedRoutes) Record(ref listenerRef, kind string, hostnames []gatev1.Hostname) {
	if ar[ref] == nil {
		ar[ref] = make(map[gatev1.Hostname]string)
	}

	// A route without hostname matches all of them.
	if len(hostnames) == 0 {
		ar[ref][""] = kind
		return
	}

	for _, hostname := range hostnames {
		ar[ref][hostname] = kind
	}
}

// Conflicts returns whether one of the given hostnames is already attached to the listener by a route of another kind.
func (ar attachedRoutes) Conflicts(ref listenerRef, kind string, hostnames []gatev1.Hostname) bool {
	for attachedHostname, attachedKind := range ar[ref] {
		if attachedKind == kind {
			continue
		}

		// The empty hostname stands for all the hostnames,
		// so an empty hostname on either side means they intersect.
		if attachedHostname == "" || len(hostnames) == 0 {
			return true
		}

		for _, hostname := range hostnames {
			if findMatchingHostname(attachedHostname, hostname) != "" || findMatchingHostname(hostname, attachedHostname) != "" {
				return true
			}
		}
	}

	return false
}

// gatewayListenersForParentRef associates a route ParentRef with the Listeners of
// the Gateway it refers to, among the Gateways managed by this controller.
type gatewayListenersForParentRef struct {
	ParentRef gatev1.ParentReference

	GatewayName      string
	GatewayNamespace string

	Listeners []gatewayListener
}

// matchingGatewayListenersForParentRef returns, for each parentRef referring to a Gateway or a ListenerSet managed by this controller,
// the listeners this parent declares.
// parentRefs that do not refer to one of our Gateways or ListenerSets are omitted.
func matchingGatewayListenersForParentRef(gateways []gatewayWithListeners, routeNamespace string, parentRefs []gatev1.ParentReference) []gatewayListenersForParentRef {
	var matches []gatewayListenersForParentRef

	for _, parentRef := range parentRefs {
		if ptr.Deref(parentRef.Group, gatev1.GroupName) != gatev1.GroupName {
			continue
		}

		owner := listenerOwner{
			Kind:      string(ptr.Deref(parentRef.Kind, kindGateway)),
			Namespace: string(ptr.Deref(parentRef.Namespace, gatev1.Namespace(routeNamespace))),
			Name:      string(parentRef.Name),
		}
		if owner.Kind != kindGateway && owner.Kind != kindListenerSet {
			continue
		}

		gateway := gatewayForOwner(gateways, owner)
		if gateway == nil {
			continue
		}

		// All the parent listeners are kept:
		// whether each listener is actually targeted (SectionName, Port) is decided when loading the route,
		// so that ResolvedRefs is reported even for parentRefs that match no listener.
		var listeners []gatewayListener
		for _, listener := range gateway.listeners {
			// A Gateway parent targets the listeners the Gateway declares itself,
			// and a ListenerSet parent only the listeners of that ListenerSet.
			if owner.Kind == kindGateway && listener.ListenerSet == nil ||
				owner.Kind == kindListenerSet && listener.ListenerSet != nil && listener.ListenerSet.Namespace == owner.Namespace && listener.ListenerSet.Name == owner.Name {
				listeners = append(listeners, listener)
			}
		}

		matches = append(matches, gatewayListenersForParentRef{
			ParentRef:        parentRef,
			GatewayName:      gateway.gateway.Name,
			GatewayNamespace: gateway.gateway.Namespace,
			Listeners:        listeners,
		})
	}

	return matches
}

// gatewayForOwner returns the managed Gateway serving the listeners of the given owner, a Gateway or a ListenerSet,
// or nil when this controller does not manage that owner.
func gatewayForOwner(gateways []gatewayWithListeners, owner listenerOwner) *gatewayWithListeners {
	for _, gateway := range gateways {
		if owner.Kind == kindListenerSet {
			for _, listener := range gateway.listeners {
				if listener.ListenerSet != nil && listener.ListenerSet.Namespace == owner.Namespace && listener.ListenerSet.Name == owner.Name {
					return &gateway
				}
			}
			continue
		}

		if gateway.gateway.Namespace == owner.Namespace && gateway.gateway.Name == owner.Name {
			return &gateway
		}
	}

	return nil
}

func matchListener(listener gatewayListener, parentRef gatev1.ParentReference) bool {
	sectionName := string(ptr.Deref(parentRef.SectionName, ""))
	if sectionName != "" && sectionName != listener.SectionName {
		return false
	}

	if parentRef.Port != nil && *parentRef.Port != listener.Port {
		return false
	}

	return true
}

func makeRouterName(kind, rule, namespace, name, gatewayNamespace, gatewayName, epName string, ruleIndex int) string {
	label := provider.Normalize(fmt.Sprintf("%s-%s-%s-gw-%s-%s-ep-%s-%d", kind, namespace, name, gatewayNamespace, gatewayName, epName, ruleIndex))

	h := sha256.New()

	for _, c := range []string{namespace, name, gatewayNamespace, gatewayName, epName, strconv.Itoa(ruleIndex)} {
		// Length-prefixing to avoid ambiguity between distinct components with embedded delimiter.
		fmt.Fprintf(h, "%d:%s", len(c), c)
	}

	// As explained in https://pkg.go.dev/hash#Hash,
	// Write never returns an error.
	h.Write([]byte(rule))

	return fmt.Sprintf("%s-%.10x", label, h.Sum(nil))
}

// makeListenerRouterName hashes the hostname, as provider.Normalize drops the characters
// telling two of them apart: the "*.example.com" and "example.com" hostnames of an entry
// point both normalize to the "listener-web-http-example-com" label.
func makeListenerRouterName(uniq uniqListener, hostname string) string {
	protocol := strings.ToLower(string(uniq.protocol))

	label := provider.Normalize(fmt.Sprintf("listener-%s-%s-%s", uniq.epName, protocol, hostname))

	h := sha256.New()

	for _, c := range []string{uniq.epName, protocol, hostname} {
		// Length-prefixing to avoid ambiguity between distinct components with embedded delimiter.
		fmt.Fprintf(h, "%d:%s", len(c), c)
	}

	return fmt.Sprintf("%s-%.10x", label, h.Sum(nil))
}

func getTLSConfig(tlsConfigs map[string]*tls.CertAndStores) []*tls.CertAndStores {
	var secretNames []string
	for secretName := range tlsConfigs {
		secretNames = append(secretNames, secretName)
	}
	sort.Strings(secretNames)

	var configs []*tls.CertAndStores
	for _, secretName := range secretNames {
		configs = append(configs, tlsConfigs[secretName])
	}

	return configs
}

func getCertificateBlocks(secret *corev1.Secret, namespace, secretName string) (string, string, error) {
	var missingEntries []string

	tlsCrtData, tlsCrtExists := secret.Data["tls.crt"]
	if !tlsCrtExists {
		missingEntries = append(missingEntries, "tls.crt")
	}

	tlsKeyData, tlsKeyExists := secret.Data["tls.key"]
	if !tlsKeyExists {
		missingEntries = append(missingEntries, "tls.key")
	}

	if len(missingEntries) > 0 {
		return "", "", fmt.Errorf("secret %s/%s is missing the following TLS data entries: %s",
			namespace, secretName, strings.Join(missingEntries, ", "))
	}

	cert := string(tlsCrtData)
	if cert == "" {
		missingEntries = append(missingEntries, "tls.crt")
	}

	key := string(tlsKeyData)
	if key == "" {
		missingEntries = append(missingEntries, "tls.key")
	}

	if len(missingEntries) > 0 {
		return "", "", fmt.Errorf("secret %s/%s contains the following empty TLS data entries: %s",
			namespace, secretName, strings.Join(missingEntries, ", "))
	}

	return cert, key, nil
}

func throttleEvents(ctx context.Context, throttleDuration time.Duration, pool *safe.Pool, eventsChan <-chan any) chan any {
	if throttleDuration == 0 {
		return nil
	}
	// Create a buffered channel to hold the pending event (if we're delaying processing the event due to throttling)
	eventsChanBuffered := make(chan any, 1)

	// Run a goroutine that reads events from eventChan and does a non-blocking write to pendingEvent.
	// This guarantees that writing to eventChan will never block,
	// and that pendingEvent will have something in it if there's been an event since we read from that channel.
	pool.GoCtx(func(ctxPool context.Context) {
		for {
			select {
			case <-ctxPool.Done():
				return
			case nextEvent := <-eventsChan:
				select {
				case eventsChanBuffered <- nextEvent:
				default:
					// We already have an event in eventsChanBuffered, so we'll do a refresh as soon as our throttle allows us to.
					// It's fine to drop the event and keep whatever's in the buffer -- we don't do different things for different events
					log.Ctx(ctx).Debug().Msgf("Dropping event kind %T due to throttling", nextEvent)
				}
			}
		}
	})

	return eventsChanBuffered
}

func isTraefikService(ref gatev1.BackendRef) bool {
	if ref.Kind == nil || ref.Group == nil {
		return false
	}

	return *ref.Group == traefikv1alpha1.GroupName && *ref.Kind == kindTraefikService
}

func isInternalService(ref gatev1.BackendRef) bool {
	return isTraefikService(ref) && strings.HasSuffix(string(ref.Name), "@internal")
}

// isCrossProviderNamespaceAllowed reports whether the given namespace is allowed to use cross-provider references.
func isCrossProviderNamespaceAllowed(allowList []string, namespace string) bool {
	if allowList == nil {
		return true
	}

	return slices.Contains(allowList, namespace)
}

// makeListenerKey joins protocol, hostname, and port of a listener into a string key.
func makeListenerKey(l gatev1.Listener) string {
	var hostname gatev1.Hostname
	if l.Hostname != nil {
		hostname = *l.Hostname
	}

	return fmt.Sprintf("%s|%s|%d", l.Protocol, hostname, l.Port)
}

// listenerClaims holds what the listeners already loaded for a Gateway claim,
// so that the listeners loaded after them are rejected when they conflict.
type listenerClaims struct {
	// listeners holds the keys of the listeners, made of their protocol, hostname and port.
	listeners map[string]struct{}
	// tcpPorts reports, for each claimed port, whether TCP listeners or hostname based ones claim it.
	tcpPorts map[gatev1.PortNumber]bool
	// frontendValidations holds the frontend validation applied on each port, shared by every Gateway of the load.
	frontendValidations map[gatev1.PortNumber]*tls.ClientAuth
}

func newListenerClaims(frontendValidations map[gatev1.PortNumber]*tls.ClientAuth) *listenerClaims {
	return &listenerClaims{
		listeners:           make(map[string]struct{}),
		tcpPorts:            make(map[gatev1.PortNumber]bool),
		frontendValidations: frontendValidations,
	}
}

func makeListenerConflictConditions(generation int64, reason gatev1.ListenerConditionReason, message string) []metav1.Condition {
	return []metav1.Condition{
		{
			Type:               string(gatev1.ListenerConditionAccepted),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(reason),
			Message:            message,
		},
		{
			Type:               string(gatev1.ListenerConditionProgrammed),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(reason),
			Message:            message,
		},
		{
			Type:               string(gatev1.ListenerConditionConflicted),
			Status:             metav1.ConditionTrue,
			ObservedGeneration: generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(reason),
			Message:            message,
		},
	}
}

// listenerSetRefsGateway returns true if the ListenerSet's ParentRef references the given gateway.
func listenerSetRefsGateway(ls *gatev1.ListenerSet, gw *gatev1.Gateway) bool {
	ref := ls.Spec.ParentRef

	if ref.Group != nil && string(*ref.Group) != gatev1.GroupName {
		return false
	}
	if ref.Kind != nil && string(*ref.Kind) != kindGateway {
		return false
	}
	if string(ref.Name) != gw.Name {
		return false
	}

	refNS := ls.Namespace
	if ref.Namespace != nil {
		refNS = string(*ref.Namespace)
	}
	return refNS == gw.Namespace
}

// isListenerSetAllowed checks whether the gateway's AllowedListeners permits the given ListenerSet.
func (p *Provider) isListenerSetAllowed(ctx context.Context, gw *gatev1.Gateway, ls *gatev1.ListenerSet) bool {
	if gw.Spec.AllowedListeners == nil {
		return false
	}

	ns := gw.Spec.AllowedListeners.Namespaces
	if ns == nil || ns.From == nil {
		return false
	}

	switch *ns.From {
	case gatev1.NamespacesFromNone:
		return false
	case gatev1.NamespacesFromSame:
		return ls.Namespace == gw.Namespace
	case gatev1.NamespacesFromAll:
		return true
	case gatev1.NamespacesFromSelector:
		if ns.Selector == nil {
			return false
		}
		selector, err := metav1.LabelSelectorAsSelector(ns.Selector)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("Invalid AllowedListeners namespace selector")
			return false
		}
		namespaces, err := p.client.ListNamespaces(selector)
		if err != nil {
			log.Ctx(ctx).Error().Err(err).Msg("Unable to list namespaces for AllowedListeners selector")
			return false
		}
		return slices.Contains(namespaces, ls.Namespace)
	}

	return false
}

func makeListenerSetStatus(listenerSetGeneration int64, listeners []gatewayListener, parentAccepted bool) (gatev1.ListenerSetStatus, bool) {
	var status gatev1.ListenerSetStatus
	var validListeners int
	for _, listener := range listeners {
		// A ListenerEntryStatus mirrors a ListenerStatus field for field.
		entryStatus := gatev1.ListenerEntryStatus(*listener.Status)
		if len(entryStatus.Conditions) == 0 {
			validListeners++

			entryStatus.Conditions = []metav1.Condition{
				{
					Type:               string(gatev1.ListenerEntryConditionAccepted),
					Status:             metav1.ConditionTrue,
					ObservedGeneration: listenerSetGeneration,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.ListenerEntryReasonAccepted),
					Message:            conditionNoErrorMessage,
				},
				{
					Type:               string(gatev1.ListenerEntryConditionResolvedRefs),
					Status:             metav1.ConditionTrue,
					ObservedGeneration: listenerSetGeneration,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.ListenerEntryReasonResolvedRefs),
					Message:            conditionNoErrorMessage,
				},
				{
					Type:               string(gatev1.ListenerEntryConditionProgrammed),
					Status:             metav1.ConditionTrue,
					ObservedGeneration: listenerSetGeneration,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.ListenerEntryReasonProgrammed),
					Message:            conditionNoErrorMessage,
				},
			}
		}

		status.Listeners = append(status.Listeners, entryStatus)
	}

	switch {
	case !parentAccepted:
		const message = "Parent Gateway is not accepted"
		status.Conditions = []metav1.Condition{
			makeListenerSetCondition(gatev1.ListenerSetConditionAccepted, metav1.ConditionFalse, listenerSetGeneration, string(gatev1.ListenerSetReasonParentNotAccepted), message),
			// The Programmed condition documents this reason, but v1.6.1 defines no constant for it.
			makeListenerSetCondition(gatev1.ListenerSetConditionProgrammed, metav1.ConditionFalse, listenerSetGeneration, "ParentNotProgrammed", message),
		}

		return status, false

	case len(status.Listeners) > 0 && validListeners == 0:
		// A ListenerSet with no valid listener at all is neither accepted nor programmed, per the spec.
		const message = "No valid listener"
		status.Conditions = []metav1.Condition{
			makeListenerSetCondition(gatev1.ListenerSetConditionAccepted, metav1.ConditionFalse, listenerSetGeneration, string(gatev1.ListenerSetReasonListenersNotValid), message),
			makeListenerSetCondition(gatev1.ListenerSetConditionProgrammed, metav1.ConditionFalse, listenerSetGeneration, string(gatev1.ListenerSetReasonListenersNotValid), message),
		}

		return status, false

	case validListeners < len(status.Listeners):
		// The valid listeners are programmed, so the ListenerSet is accepted and programmed even though some listeners have errors.
		status.Conditions = []metav1.Condition{
			makeListenerSetCondition(gatev1.ListenerSetConditionAccepted, metav1.ConditionTrue, listenerSetGeneration, string(gatev1.ListenerSetReasonListenersNotValid), "Some listeners have errors"),
			makeListenerSetCondition(gatev1.ListenerSetConditionProgrammed, metav1.ConditionTrue, listenerSetGeneration, string(gatev1.ListenerSetReasonProgrammed), "Valid listeners programmed"),
		}

	default:
		status.Conditions = []metav1.Condition{
			makeListenerSetCondition(gatev1.ListenerSetConditionAccepted, metav1.ConditionTrue, listenerSetGeneration, string(gatev1.ListenerSetReasonAccepted), "ListenerSet accepted"),
			makeListenerSetCondition(gatev1.ListenerSetConditionProgrammed, metav1.ConditionTrue, listenerSetGeneration, string(gatev1.ListenerSetReasonProgrammed), "ListenerSet programmed"),
		}
	}

	return status, true
}

func makeListenerSetCondition(conditionType gatev1.ListenerSetConditionType, status metav1.ConditionStatus, generation int64, reason, message string) metav1.Condition {
	return metav1.Condition{
		Type:               string(conditionType),
		Status:             status,
		ObservedGeneration: generation,
		LastTransitionTime: metav1.Now(),
		Reason:             reason,
		Message:            message,
	}
}

func filterReferenceGrantsFrom(referenceGrants []*gatev1.ReferenceGrant, group, kind, namespace string) []*gatev1.ReferenceGrant {
	var matchingReferenceGrants []*gatev1.ReferenceGrant
	for _, referenceGrant := range referenceGrants {
		if referenceGrantMatchesFrom(referenceGrant, group, kind, namespace) {
			matchingReferenceGrants = append(matchingReferenceGrants, referenceGrant)
		}
	}
	return matchingReferenceGrants
}

func referenceGrantMatchesFrom(referenceGrant *gatev1.ReferenceGrant, group, kind, namespace string) bool {
	for _, from := range referenceGrant.Spec.From {
		sanitizedGroup := string(from.Group)
		if sanitizedGroup == "" {
			sanitizedGroup = groupCore
		}
		if string(from.Namespace) != namespace || string(from.Kind) != kind || sanitizedGroup != group {
			continue
		}
		return true
	}
	return false
}

func filterReferenceGrantsTo(referenceGrants []*gatev1.ReferenceGrant, group, kind, name string) []*gatev1.ReferenceGrant {
	var matchingReferenceGrants []*gatev1.ReferenceGrant
	for _, referenceGrant := range referenceGrants {
		if referenceGrantMatchesTo(referenceGrant, group, kind, name) {
			matchingReferenceGrants = append(matchingReferenceGrants, referenceGrant)
		}
	}
	return matchingReferenceGrants
}

func referenceGrantMatchesTo(referenceGrant *gatev1.ReferenceGrant, group, kind, name string) bool {
	for _, to := range referenceGrant.Spec.To {
		sanitizedGroup := string(to.Group)
		if sanitizedGroup == "" {
			sanitizedGroup = groupCore
		}
		if string(to.Kind) != kind || sanitizedGroup != group || (to.Name != nil && string(*to.Name) != name) {
			continue
		}
		return true
	}
	return false
}

func groupToString(p *gatev1.Group) string {
	if p == nil {
		return "<nil>"
	}
	return string(*p)
}

func kindToString(p *gatev1.Kind) string {
	if p == nil {
		return "<nil>"
	}
	return string(*p)
}

func upsertGatewayClassConditionAccepted(conditions []metav1.Condition, condition metav1.Condition) []metav1.Condition {
	var conds []metav1.Condition
	for _, c := range conditions {
		if c.Type == string(gatev1.GatewayClassConditionStatusAccepted) {
			continue
		}
		conds = append(conds, c)
	}
	return append(conds, condition)
}

```

#### pkg/kubernetes_test.go

```go
package gateway

import (
	"cmp"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	ptypes "github.com/traefik/paerser/types"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/provider"
	traefikv1alpha1 "github.com/traefik/traefik/v3/pkg/provider/kubernetes/crd/traefikio/v1alpha1"
	"github.com/traefik/traefik/v3/pkg/provider/kubernetes/k8s"
	"github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
	"google.golang.org/grpc/codes"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ktypes "k8s.io/apimachinery/pkg/types"
	kubefake "k8s.io/client-go/kubernetes/fake"
	kscheme "k8s.io/client-go/kubernetes/scheme"
	ktesting "k8s.io/client-go/testing"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
	gatefake "sigs.k8s.io/gateway-api/pkg/client/clientset/versioned/fake"
)

var _ provider.Provider = (*Provider)(nil)

func defaultListenerTLSOptions() tls.Options {
	var opts tls.Options
	opts.SetDefaults()
	return opts
}

func init() {
	// required by k8s.MustParseYaml
	if err := gatev1.Install(kscheme.Scheme); err != nil {
		panic(err)
	}
}

const (
	listenerCert string = `-----BEGIN CERTIFICATE-----
MIIBqDCCAU6gAwIBAgIUYOsr0QgHOBq4kYRCL5+TDdVt6bQwCgYIKoZIzj0EAwIw
FjEUMBIGA1UEAwwLZXhhbXBsZS5jb20wHhcNMjUxMDEwMDcxNzMwWhcNMzUxMDA4
MDcxNzMwWjAWMRQwEgYDVQQDDAtleGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqG
SM49AwEHA0IABDOriw3ZQ7wIhWrbPS6JFQT3bToNCF00vSV5fab6TbXyL8tlsGre
TRIF2EwgsteMOkxGKKSlDvuaDwq8p/qV+0ujejB4MB0GA1UdDgQWBBRMEkuexXQh
UtDgRg1KBv72CDq+EzAfBgNVHSMEGDAWgBRMEkuexXQhUtDgRg1KBv72CDq+EzAP
BgNVHRMBAf8EBTADAQH/MCUGA1UdEQQeMByCC2V4YW1wbGUuY29tgg0qLmV4YW1w
bGUuY29tMAoGCCqGSM49BAMCA0gAMEUCIQDs87Vk0swA6HgOJjROye1mxD83qcGy
peFgoxV93Dy+cwIgV0MMEJJbVsTyZK3EQ++Xc5rEL78nrJ+YIEV+rCUWj5U=
-----END CERTIFICATE-----`
	listenerKey string = `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgnwgL5DY4UB14sM6f
DikQdtqh2QW1ArfF4fc1UFzifdGhRANCAAQzq4sN2UO8CIVq2z0uiRUE9206DQhd
NL0leX2m+k218i/LZbBq3k0SBdhMILLXjDpMRiikpQ77mg8KvKf6lftL
-----END PRIVATE KEY-----`
)

func TestGatewayClassLabelSelector(t *testing.T) {
	k8sObjects, gwObjects := readResources(t, []string{"gatewayclass_labelselector.yaml"})

	kubeClient := kubefake.NewClientset(k8sObjects...)
	gwClient := newGatewaySimpleClientSet(t, gwObjects...)

	client := newClientImpl(kubeClient, gwClient)

	// This is initialized by the Provider init method but this cannot be called in a unit test.
	client.labelSelector = "name=traefik-internal"

	eventCh, err := client.WatchAll(nil, make(chan struct{}))
	require.NoError(t, err)

	if len(k8sObjects) > 0 || len(gwObjects) > 0 {
		// just wait for the first event
		<-eventCh
	}

	p := Provider{
		EntryPoints:   map[string]Entrypoint{"http": {Address: ":9080"}},
		StatusAddress: &StatusAddress{IP: "1.2.3.4"},
		client:        client,
	}

	_, statusReport, err := p.loadConfigurationFromGateways(t.Context())
	require.NoError(t, err)

	statusReport.Flush(t.Context(), p.client)

	gw, err := gwClient.GatewayV1().Gateways("default").Get(t.Context(), "traefik-external", metav1.GetOptions{})
	require.NoError(t, err)

	assert.Empty(t, gw.Status.Addresses)

	gw, err = gwClient.GatewayV1().Gateways("default").Get(t.Context(), "traefik-internal", metav1.GetOptions{})
	require.NoError(t, err)
	require.Len(t, gw.Status.Addresses, 1)
	require.NotNil(t, gw.Status.Addresses[0].Type)

	assert.Equal(t, gatev1.IPAddressType, *gw.Status.Addresses[0].Type)
	assert.Equal(t, "1.2.3.4", gw.Status.Addresses[0].Value)
}

func TestGatewayClassFieldSelector(t *testing.T) {
	// The resources are only there to make the informer emit a first event to wait for.
	_, gwObjects := readResources(t, []string{"gatewayclass_labelselector.yaml"})

	gwClient := newGatewaySimpleClientSet(t, gwObjects...)
	client := newClientImpl(kubefake.NewClientset(), gwClient)
	client.gatewayClass = "traefik-internal"

	eventCh, err := client.WatchAll(nil, make(chan struct{}))
	require.NoError(t, err)

	<-eventCh

	// The fake clientset ignores field selectors, so the lister cannot show the filtering.
	// The API server does the filtering. This test only checks the selector is sent to it.
	var (
		listed  bool
		watched bool
	)
	for _, action := range gwClient.Actions() {
		if action.GetResource().Resource != "gatewayclasses" {
			continue
		}

		switch a := action.(type) {
		case ktesting.ListAction:
			listed = true
			assert.Equal(t, "metadata.name=traefik-internal", a.GetListRestrictions().Fields.String())
		case ktesting.WatchAction:
			watched = true
			assert.Equal(t, "metadata.name=traefik-internal", a.GetWatchRestrictions().Fields.String())
		}
	}

	assert.True(t, listed)
	assert.True(t, watched)
}

func TestGatewayClassAndLabelSelectorExclusive(t *testing.T) {
	p := Provider{
		GatewayClass:  "traefik-internal",
		LabelSelector: "name=traefik-internal",
	}

	_, err := p.newK8sClient(t.Context())
	require.EqualError(t, err, "gatewayClass and labelSelector options are mutually exclusive")
}

func TestGatewayScoping(t *testing.T) {
	k8sObjects, gwObjects := readResources(t, []string{"gateway_scoping.yaml"})

	kubeClient := kubefake.NewClientset(k8sObjects...)
	gwClient := newGatewaySimpleClientSet(t, gwObjects...)

	client := newClientImpl(kubeClient, gwClient)

	eventCh, err := client.WatchAll(nil, make(chan struct{}))
	require.NoError(t, err)

	if len(k8sObjects) > 0 || len(gwObjects) > 0 {
		<-eventCh
	}

	p := Provider{
		EntryPoints:   map[string]Entrypoint{"http": {Address: ":9080"}},
		StatusAddress: &StatusAddress{IP: "1.2.3.4"},
		// Scope the provider to a single Gateway.
		Gateways: []NamespacedName{{Namespace: "default", Name: "traefik-internal"}},
		client:   client,
	}

	_, statusReport, err := p.loadConfigurationFromGateways(t.Context())
	require.NoError(t, err)

	statusReport.Flush(t.Context(), p.client)

	// The scoped Gateway is managed and gets a status address.
	gw, err := gwClient.GatewayV1().Gateways("default").Get(t.Context(), "traefik-internal", metav1.GetOptions{})
	require.NoError(t, err)
	require.Len(t, gw.Status.Addresses, 1)
	assert.Equal(t, "1.2.3.4", gw.Status.Addresses[0].Value)

	// The other Gateway is ignored.
	gw, err = gwClient.GatewayV1().Gateways("default").Get(t.Context(), "traefik-external", metav1.GetOptions{})
	require.NoError(t, err)
	assert.Empty(t, gw.Status.Addresses)
}

func TestLoadHTTPRoutes(t *testing.T) {
	testCases := []struct {
		desc         string
		ingressClass string
		paths        []string
		expected     *dynamic.Configuration
		entryPoints  map[string]Entrypoint
		nativeLB     bool
		// retryMaxRequestBodyBytes defaults to dynamic.RetryDefaultMaxRequestBodyBytes when zero.
		retryMaxRequestBodyBytes int64
	}{
		{
			desc: "Empty",
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because missing entry point",
			paths: []string{"services.yml", "httproute/simple.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":443",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because no http route defined",
			paths: []string{"services.yml", "httproute/without_httproute.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by missing GatewayClass",
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			paths: []string{"services.yml", "httproute/without_gatewayclass.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by unknown GatewayClass controller desc",
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			paths: []string{"services.yml", "httproute/gatewayclass_with_unknown_controller.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Router with service in error caused by wrong TargetPort",
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			paths: []string{"services.yml", "httproute/with_wrong_service_port.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
										Status: new(500),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by HTTPS without TLS",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":443",
			}},
			paths: []string{"services.yml", "httproute/with_protocol_https_without_tls.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by HTTPS with TLS passthrough",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":443",
			}},
			paths: []string{"services.yml", "httproute/with_protocol_https_with_tls_passthrough.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by HTTPRoute with protocol TLS",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":443",
			}},
			paths: []string{"services.yml", "httproute/with_protocol_tls.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by HTTPRoute with protocol TCP",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":443",
			}},
			paths: []string{"services.yml", "httproute/with_protocol_tcp.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by TCPRoute with protocol HTTP",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":9000",
			}},
			paths: []string{"services.yml", "tcproute/with_protocol_http.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by TCPRoute with protocol HTTPS",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":9000",
			}},
			paths: []string{"services.yml", "tcproute/with_protocol_https.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by TLSRoute with protocol TCP",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":9001",
			}},
			paths: []string{"services.yml", "tlsroute/with_protocol_tcp.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by TLSRoute with protocol HTTP",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":9001",
			}},
			paths: []string{"services.yml", "tlsroute/with_protocol_http.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by TLSRoute with protocol HTTPS",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":9001",
			}},
			paths: []string{"services.yml", "tlsroute/with_protocol_https.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused use http entrypoint with tls activated with HTTPRoute",
			entryPoints: map[string]Entrypoint{"websecure": {
				Address:        ":443",
				HasHTTPTLSConf: true,
			}},
			paths: []string{"services.yml", "httproute/simple_with_tls_entrypoint.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because no tcp route defined tls protocol",
			paths: []string{"services.yml", "tcproute/without_tcproute_tls_protocol.yml"},
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8080",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc: "Empty caused by HTTPRoute with TLS configuration",
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			paths: []string{"services.yml", "httproute/with_tls_configuration.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute",
			paths: []string{"services.yml", "httproute/simple.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Listeners isolation with a specificity chain spread over three Gateways",
			paths: []string{"services.yml", "httproute/listener_isolation_three_gateways.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `(Host("*")) && !(Host("**.example.com") || Host("abc.example.com"))`,
						},
						"listener-web-http-example-com-e7b9da953ca48d642d6b": {
							EntryPoints: []string{"web"},
							Rule:        `(Host("**.example.com")) && !(Host("abc.example.com"))`,
						},
						"listener-web-http-abc-example-com-05dd4f497bcb27f2cdca": {
							EntryPoints: []string{"web"},
							Rule:        `Host("abc.example.com")`,
						},
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb": {
							ParentRefs: []string{
								"listener-web-http-5b3226d4ebd42b7200f1",
								"listener-web-http-example-com-e7b9da953ca48d642d6b",
							},
							Service:    "httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb-wrr",
							Rule:       `Path("/catchall")`,
							Priority:   100001,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b": {
							ParentRefs: []string{"listener-web-http-abc-example-com-05dd4f497bcb27f2cdca"},
							Service:    "httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b-wrr",
							Rule:       `Host("abc.example.com") && Path("/precise")`,
							Priority:   100016,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-wildcard-gw-default-my-wildcard-gateway-ep-web-0-63284dea2e7614ab00a2": {
							ParentRefs: []string{
								"listener-web-http-example-com-e7b9da953ca48d642d6b",
								"listener-web-http-abc-example-com-05dd4f497bcb27f2cdca",
							},
							Service:    "httproute-default-http-app-wildcard-gw-default-my-wildcard-gateway-ep-web-0-63284dea2e7614ab00a2-wrr",
							Rule:       `Host("**.example.com") && Path("/wildcard")`,
							Priority:   100014,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce": {
							ParentRefs: []string{
								"listener-web-http-5b3226d4ebd42b7200f1",
								"listener-web-http-example-com-e7b9da953ca48d642d6b",
								"listener-web-http-abc-example-com-05dd4f497bcb27f2cdca",
							},
							Service:    "httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce-wrr",
							Rule:       `Path("/other")`,
							Priority:   100001,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-wildcard-gw-default-my-wildcard-gateway-ep-web-0-63284dea2e7614ab00a2-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-wildcard-gw-default-my-wildcard-gateway-ep-web-0-63284dea2e7614ab00a2-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-wildcard-gw-default-my-wildcard-gateway-ep-web-0-63284dea2e7614ab00a2-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Listeners isolation with a hostname listener and a catch-all listener on two Gateways",
			paths: []string{"services.yml", "httproute/listener_isolation.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `(Host("*")) && !(Host("abc.example.com"))`,
						},
						"listener-web-http-abc-example-com-05dd4f497bcb27f2cdca": {
							EntryPoints: []string{"web"},
							Rule:        `Host("abc.example.com")`,
						},
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb-wrr",
							Rule:       `Path("/catchall")`,
							Priority:   100001,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b": {
							ParentRefs: []string{"listener-web-http-abc-example-com-05dd4f497bcb27f2cdca"},
							Service:    "httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b-wrr",
							Rule:       `Host("abc.example.com") && Path("/precise")`,
							Priority:   100016,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce": {
							ParentRefs: []string{
								"listener-web-http-5b3226d4ebd42b7200f1",
								"listener-web-http-abc-example-com-05dd4f497bcb27f2cdca",
							},
							Service:    "httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce-wrr",
							Rule:       `Path("/other")`,
							Priority:   100001,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-web-0-6bf410c47a90e258b1eb-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-precise-gw-default-my-gateway-ep-web-0-5b6dacf4c7da9367767b-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-other-gw-default-my-other-gateway-ep-web-0-d7590587db124f0d35ce-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute with a wildcard listener hostname",
			paths: []string{"services.yml", "httproute/simple_with_wildcard_listener_hostname.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-example-com-e7b9da953ca48d642d6b": {
							EntryPoints: []string{"web"},
							Rule:        `Host("**.example.com")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-2944dbefe4493e703281": {
							ParentRefs: []string{"listener-web-http-example-com-e7b9da953ca48d642d6b"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-2944dbefe4493e703281-wrr",
							Rule:       `Host("**.example.com") && Path("/bar")`,
							Priority:   100014,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-2944dbefe4493e703281-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-2944dbefe4493e703281-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-2944dbefe4493e703281-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, without backendRefs",
			paths: []string{"services.yml", "httproute/without_backendrefs.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-b0d3512b71b7695aa511": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-b0d3512b71b7695aa511-wrr",
							Rule:       `Host("foo.com") && Path("/omitted")`,
							Priority:   100009,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-28c8ec797c4b0f521cf7": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-28c8ec797c4b0f521cf7-wrr",
							Rule:       `Host("foo.com") && Path("/empty")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-b0d3512b71b7695aa511-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "no-backend-refs",
										Status: new(500),
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-28c8ec797c4b0f521cf7-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "no-backend-refs",
										Status: new(500),
										Weight: new(1),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute with request and backendRequest timeouts",
			paths: []string{"services.yml", "httproute/with_request_timeout.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs:         []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:            "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:               `Host("foo.com") && Path("/bar")`,
							Priority:           100010,
							RuleSyntax:         "default",
							RespondingTimeouts: &dynamic.RouterRespondingTimeouts{ReadTimeout: new(ptypes.Duration(10 * time.Second)), WriteTimeout: new(ptypes.Duration(10 * time.Second))},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74": {
							ParentRefs:         []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:            "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-wrr",
							Rule:               `Host("foo.com") && Path("/baz")`,
							Priority:           100009,
							RuleSyntax:         "default",
							RespondingTimeouts: &dynamic.RouterRespondingTimeouts{ReadTimeout: new(ptypes.Duration(0)), WriteTimeout: new(ptypes.Duration(0))},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-c465583028340322911a": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-c465583028340322911a-wrr",
							Rule:       `Host("foo.com") && Path("/buz")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-c465583028340322911a-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-c465583028340322911a-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-c465583028340322911a-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-c465583028340322911a-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							ForwardingTimeouts: &dynamic.ForwardingTimeouts{
								DialTimeout:           ptypes.Duration(30 * time.Second),
								ResponseHeaderTimeout: ptypes.Duration(5 * time.Second),
								IdleConnTimeout:       ptypes.Duration(90 * time.Second),
								PingTimeout:           ptypes.Duration(15 * time.Second),
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-c465583028340322911a-svc-default-whoami-0": {
							ForwardingTimeouts: &dynamic.ForwardingTimeouts{
								DialTimeout:           ptypes.Duration(30 * time.Second),
								ResponseHeaderTimeout: ptypes.Duration(5 * time.Second),
								IdleConnTimeout:       ptypes.Duration(90 * time.Second),
								PingTimeout:           ptypes.Duration(15 * time.Second),
							},
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute with retry",
			paths: []string{"services.yml", "httproute/with_retry.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100009,
							RuleSyntax: "default",
							Middlewares: []string{
								"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-requestheadermodifier-0",
								"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-retry",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-err-wrr",
							Rule:       `Host("foo.com") && Path("/baz")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-requestheadermodifier-0": {
							RequestHeaderModifier: &dynamic.HeaderModifier{
								Set: map[string]string{"X-Foo": "bar"},
								Add: map[string]string{},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-retry": {
							Retry: &dynamic.Retry{
								Attempts:            4,
								Status:              []string{"500", "503"},
								InitialInterval:     ptypes.Duration(100 * time.Millisecond),
								MaxRequestBodyBytes: new(dynamic.RetryDefaultMaxRequestBodyBytes),
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-err-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "invalid-httproute-filter",
										Status: new(500),
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute with empty retry",
			paths: []string{"services.yml", "httproute/with_empty_retry.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							Middlewares: []string{
								"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-retry",
							},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-retry": {
							Retry: &dynamic.Retry{
								Attempts:            2,
								MaxRequestBodyBytes: new(dynamic.RetryDefaultMaxRequestBodyBytes),
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:                     "HTTPRoute with retry and custom retryMaxRequestBodyBytes",
			paths:                    []string{"services.yml", "httproute/with_empty_retry.yml"},
			retryMaxRequestBodyBytes: -1,
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							Middlewares: []string{
								"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-retry",
							},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-retry": {
							Retry: &dynamic.Retry{
								Attempts:            2,
								MaxRequestBodyBytes: new(int64(-1)),
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, with api@internal service",
			paths: []string{"services.yml", "httproute/simple_to_api_internal.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "api@internal",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute with protocol HTTPS",
			paths: []string{"services.yml", "httproute/with_protocol_https.yml"},
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":443",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "HTTPRoute with multiple TLS certificateRefs",
			paths: []string{"services.yml", "httproute/with_multiple_certificaterefs.yml"},
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":443",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "HTTPRoute with multiple certificateRefs, one missing",
			paths: []string{"services.yml", "httproute/with_multiple_certificaterefs_one_missing.yml"},
			entryPoints: map[string]Entrypoint{"websecure": {
				Address: ":443",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, with multiple hosts",
			paths: []string{"services.yml", "httproute/with_multiple_host.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-fe758dd8380cc8acabbc": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-fe758dd8380cc8acabbc-wrr",
							Rule:       `(Host("foo.com") || Host("bar.com")) && PathPrefix("/")`,
							Priority:   9,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-fe758dd8380cc8acabbc-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-fe758dd8380cc8acabbc-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-fe758dd8380cc8acabbc-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, with two hosts one wildcard",
			paths: []string{"services.yml", "with_two_hosts_one_wildcard.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-5af0a8dc466bcb103301": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-5af0a8dc466bcb103301-wrr",
							Rule:       `(Host("foo.com") || Host("**.bar.com")) && PathPrefix("/")`,
							Priority:   11,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-5af0a8dc466bcb103301-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-5af0a8dc466bcb103301-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-5af0a8dc466bcb103301-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, with one host and a wildcard",
			paths: []string{"services.yml", "with_two_hosts_wildcard.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-764e519b9557509caa7d": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-764e519b9557509caa7d-wrr",
							Rule:       `(Host("foo.com") || Host("**.foo.com")) && PathPrefix("/")`,
							Priority:   11,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-764e519b9557509caa7d-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-764e519b9557509caa7d-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-764e519b9557509caa7d-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "One HTTPRoute with two different rules",
			paths: []string{"services.yml", "httproute/two_rules.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100009,
							RuleSyntax: "default",
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-71ae3ac5f5f8583c5cfa": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Rule:       `Host("foo.com") && Path("/bir")`,
							Priority:   100008,
							RuleSyntax: "default",
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-71ae3ac5f5f8583c5cfa-wrr",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-71ae3ac5f5f8583c5cfa-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-71ae3ac5f5f8583c5cfa-svc-default-whoami2-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-71ae3ac5f5f8583c5cfa-svc-default-whoami2-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.3:8080",
									},
									{
										URL: "http://10.10.0.4:8080",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "One HTTPRoute with one rule two targets",
			paths: []string{"services.yml", "httproute/one_rule_two_targets.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami2-1",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami2-1": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.3:8080",
									},
									{
										URL: "http://10.10.0.4:8080",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Two Gateways and one HTTPRoute",
			paths: []string{"services.yml", "httproute/with_two_gateways_one_httproute.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {
					Address: ":80",
				},
				"websecure": {
					Address: ":443",
				},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-http-ep-web-0-779a0bc558422ee1b154": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-http-ep-web-0-779a0bc558422ee1b154-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-1-gw-default-my-gateway-https-ep-websecure-0-036eecf3c21d00eac117": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-https-ep-websecure-0-036eecf3c21d00eac117-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-http-ep-web-0-779a0bc558422ee1b154-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-http-ep-web-0-779a0bc558422ee1b154-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-https-ep-websecure-0-036eecf3c21d00eac117-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-https-ep-websecure-0-036eecf3c21d00eac117-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-http-ep-web-0-779a0bc558422ee1b154-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-https-ep-websecure-0-036eecf3c21d00eac117-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "Gateway with two listeners and one HTTPRoute",
			paths: []string{"services.yml", "httproute/with_two_listeners_one_httproute.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {
					Address: ":80",
				},
				"websecure": {
					Address: ":443",
				},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, with several rules",
			paths: []string{"services.yml", "httproute/with_several_rules.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-cbebdedef34a6261eced": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Rule:       `Host("foo.com") && (Path("/bar") || PathPrefix("/bar/")) && Header("my-header","foo") && Header("my-header2","bar")`,
							Priority:   10610,
							RuleSyntax: "default",
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-cbebdedef34a6261eced-wrr",
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-74d0c698c3be747a738f": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Rule:       `Host("foo.com") && PathRegexp("^/buzz/[0-9]+$")`,
							Priority:   11408,
							RuleSyntax: "default",
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-74d0c698c3be747a738f-wrr",
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-d4501fe0597184abd40e": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Rule:       `Host("foo.com") && Path("/bar") && Header("my-header","bar")`,
							Priority:   100109,
							RuleSyntax: "default",
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-d4501fe0597184abd40e-wrr",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-cbebdedef34a6261eced-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-cbebdedef34a6261eced-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-74d0c698c3be747a738f-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-74d0c698c3be747a738f-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-d4501fe0597184abd40e-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-d4501fe0597184abd40e-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-cbebdedef34a6261eced-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-2-74d0c698c3be747a738f-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-d4501fe0597184abd40e-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, with method matching",
			paths: []string{"services.yml", "httproute/with_method_matching.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-eaf4f488cc5127097136": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Rule:       `Host("foo.com") && (Path("/foo") || PathPrefix("/foo/")) && Method("GET")`,
							Priority:   11408,
							RuleSyntax: "default",
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-eaf4f488cc5127097136-wrr",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-eaf4f488cc5127097136-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-eaf4f488cc5127097136-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-eaf4f488cc5127097136-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, with query param matching",
			paths: []string{"services.yml", "httproute/with_query_param_matching.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-8dc41dbb469b11d7ce9c": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Rule:       `Host("foo.com") && (Path("/foo") || PathPrefix("/foo/")) && Query("foo","bar") && QueryRegexp("baz","buz")`,
							Priority:   10428,
							RuleSyntax: "default",
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-8dc41dbb469b11d7ce9c-wrr",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-8dc41dbb469b11d7ce9c-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-8dc41dbb469b11d7ce9c-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-8dc41dbb469b11d7ce9c-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute with Same namespace selector",
			paths: []string{"services.yml", "httproute/with_namespace_same.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":80"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe-wrr",
							Rule:       `Host("foo.com") && Path("/foo")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute with All namespace selector",
			paths: []string{"services.yml", "httproute/with_namespace_all.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":80"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe-wrr",
							Rule:       `Host("foo.com") && Path("/foo")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d-wrr",
							Rule:       `Host("bar.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d-svc-bar-whoami-bar-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-71c892303a6d70cb4efe-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d-svc-bar-whoami-bar-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.11:80",
									},
									{
										URL: "http://10.10.0.12:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute with namespace selector",
			paths: []string{"services.yml", "httproute/with_namespace_selector.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":80"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d-wrr",
							Rule:       `Host("bar.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d-svc-bar-whoami-bar-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-eb47eeb38d713d8e351d-svc-bar-whoami-bar-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.11:80",
									},
									{
										URL: "http://10.10.0.12:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, request header modifier",
			paths: []string{"services.yml", "httproute/filter_request_header_modifier.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664": {
							ParentRefs:  []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr",
							Rule:        `Host("example.org") && PathPrefix("/")`,
							Priority:    13,
							RuleSyntax:  "default",
							Middlewares: []string{"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestheadermodifier-0"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestheadermodifier-0": {
							RequestHeaderModifier: &dynamic.HeaderModifier{
								Set:    map[string]string{"X-Foo": "Bar"},
								Add:    map[string]string{"X-Bar": "Foo"},
								Remove: []string{"X-Baz"},
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, response header modifier",
			paths: []string{"services.yml", "httproute/filter_response_header_modifier.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664": {
							ParentRefs:  []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr",
							Rule:        `Host("example.org") && PathPrefix("/")`,
							Priority:    13,
							RuleSyntax:  "default",
							Middlewares: []string{"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-responseheadermodifier-0"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-responseheadermodifier-0": {
							ResponseHeaderModifier: &dynamic.HeaderModifier{
								Set:    map[string]string{"X-Foo": "Bar"},
								Add:    map[string]string{"X-Bar": "Foo"},
								Remove: []string{"X-Baz"},
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, request mirrors mixed with other filters",
			paths: []string{"services.yml", "httproute/filter_request_mirror.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr",
							Rule:       `Host("example.org") && PathPrefix("/")`,
							Priority:   13,
							RuleSyntax: "default",
							Middlewares: []string{
								"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestheadermodifier-0",
								"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-1",
								"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-responseheadermodifier-2",
								"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-3",
							},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestheadermodifier-0": {
							RequestHeaderModifier: &dynamic.HeaderModifier{
								Set: map[string]string{"X-Foo": "Bar"},
								Add: map[string]string{},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-1": {
							RequestMirror: &dynamic.RequestMirror{
								Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-1-svc-default-whoami2-0",
								Numerator:   1,
								Denominator: 1,
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-responseheadermodifier-2": {
							ResponseHeaderModifier: &dynamic.HeaderModifier{
								Set: map[string]string{"X-Bar": "Foo"},
								Add: map[string]string{},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-3": {
							RequestMirror: &dynamic.RequestMirror{
								Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-3-svc-default-whoami-0",
								Numerator:   50,
								Denominator: 100,
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0-requestmirror-0": {
							RequestMirror: &dynamic.RequestMirror{
								Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0-requestmirror-0-svc-default-whoami2-0",
								Numerator:   1,
								Denominator: 4,
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0": {
							Middlewares: []string{"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0-requestmirror-0"},
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-1-svc-default-whoami2-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.3:8080",
									},
									{
										URL: "http://10.10.0.4:8080",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-3-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0-requestmirror-0-svc-default-whoami2-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.3:8080",
									},
									{
										URL: "http://10.10.0.4:8080",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0": {
							ForwardingTimeouts: &dynamic.ForwardingTimeouts{
								DialTimeout:           ptypes.Duration(30 * time.Second),
								ResponseHeaderTimeout: ptypes.Duration(5 * time.Second),
								IdleConnTimeout:       ptypes.Duration(90 * time.Second),
								PingTimeout:           ptypes.Duration(15 * time.Second),
							},
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, invalid request mirrors are dropped",
			paths: []string{"services.yml", "httproute/filter_request_mirror_missing_service.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr",
							Rule:       `Host("example.org") && PathPrefix("/")`,
							Priority:   13,
							RuleSyntax: "default",
							Middlewares: []string{
								"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-1",
							},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-1": {
							RequestMirror: &dynamic.RequestMirror{
								Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-1-svc-default-whoami2-0",
								Numerator:   1,
								Denominator: 1,
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestmirror-1-svc-default-whoami2-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.3:8080",
									},
									{
										URL: "http://10.10.0.4:8080",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, backend filter request header modifier",
			paths: []string{"services.yml", "httproute/backend_filter_request_header_modifier.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-78ef4770613eae69b5da": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-78ef4770613eae69b5da-wrr",
							Rule:       `Host("foo.com") && (Path("/bar") || PathPrefix("/bar/"))`,
							Priority:   10408,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-78ef4770613eae69b5da-svc-default-whoami-0-requestheadermodifier-0": {
							RequestHeaderModifier: &dynamic.HeaderModifier{
								Set:    map[string]string{"X-Foo": "Bar"},
								Add:    map[string]string{"X-Bar": "Foo"},
								Remove: []string{"X-Baz"},
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-78ef4770613eae69b5da-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-78ef4770613eae69b5da-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-78ef4770613eae69b5da-svc-default-whoami-0": {
							Middlewares: []string{"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-78ef4770613eae69b5da-svc-default-whoami-0-requestheadermodifier-0"},
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, redirect HTTP to HTTPS",
			paths: []string{"services.yml", "httproute/filter_http_to_https.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664": {
							ParentRefs:  []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr",
							Rule:        `Host("example.org") && PathPrefix("/")`,
							Priority:    13,
							RuleSyntax:  "default",
							Middlewares: []string{"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestredirect-0"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestredirect-0": {
							RequestRedirect: &dynamic.RequestRedirect{
								Scheme:     new("https"),
								Port:       new(""),
								StatusCode: http.StatusMovedPermanently,
							},
						},
					},

					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, redirect HTTP to HTTPS with hostname",
			paths: []string{"services.yml", "httproute/filter_http_to_https_with_hostname_and_port.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664": {
							ParentRefs:  []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr",
							Rule:        `Host("example.org") && PathPrefix("/")`,
							Priority:    13,
							RuleSyntax:  "default",
							Middlewares: []string{"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestredirect-0"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-requestredirect-0": {
							RequestRedirect: &dynamic.RequestRedirect{
								Hostname:   new("example.com"),
								Port:       new("443"),
								StatusCode: http.StatusFound,
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-c683caccb22c43e5e664-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute URL rewrite FullPath",
			paths: []string{"services.yml", "httproute/filter_url_rewrite_fullpath.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d": {
							ParentRefs:  []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-wrr",
							Rule:        `Host("example.com") && (Path("/foo") || PathPrefix("/foo/"))`,
							RuleSyntax:  "default",
							Priority:    10412,
							Middlewares: []string{"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-urlrewrite-0"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-urlrewrite-0": {
							URLRewrite: &dynamic.URLRewrite{
								Path: new("/bar"),
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-svc-default-whoami-0",
										Weight: func(i int) *int { return &i }(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute URL rewrite Hostname",
			paths: []string{"services.yml", "httproute/filter_url_rewrite_hostname.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d": {
							ParentRefs:  []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-wrr",
							Rule:        `Host("example.com") && (Path("/foo") || PathPrefix("/foo/"))`,
							RuleSyntax:  "default",
							Priority:    10412,
							Middlewares: []string{"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-urlrewrite-0"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-urlrewrite-0": {
							URLRewrite: &dynamic.URLRewrite{
								Hostname: new("www.foo.bar"),
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-svc-default-whoami-0",
										Weight: func(i int) *int { return &i }(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute URL rewrite Combined",
			paths: []string{"services.yml", "httproute/filter_url_rewrite_combined.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d": {
							ParentRefs:  []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:     "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-wrr",
							Rule:        `Host("example.com") && (Path("/foo") || PathPrefix("/foo/"))`,
							RuleSyntax:  "default",
							Priority:    10412,
							Middlewares: []string{"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-urlrewrite-0"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-urlrewrite-0": {
							URLRewrite: &dynamic.URLRewrite{
								Hostname:   new("www.foo.bar"),
								Path:       new("/xyz"),
								PathPrefix: new("/foo"),
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-svc-default-whoami-0",
										Weight: func(i int) *int { return &i }(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-d1e350a646631250e99d-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute and BackendTLSPolicy with CA certificate",
			paths: []string{"services.yml", "httproute/with_backend_tls_policy.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.1:80",
									},
									{
										URL: "https://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							ServerName: "whoami",
							RootCAs: []types.FileOrContent{
								"CA1",
								"CA2",
								"CA1-secret",
								"CA2-secret",
							},
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute and BackendTLSPolicy with a Gateway backend client certificate",
			paths: []string{"services.yml", "httproute/with_backend_tls_policy_and_client_certificate.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100009,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-wrr",
							Rule:       `Host("foo.com") && Path("/baz")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.1:80",
									},
									{
										URL: "https://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-svc-default-whoami2-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-1-6bebfd4e0e692b59df74-svc-default-whoami2-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.3:8080",
									},
									{
										URL: "http://10.10.0.4:8080",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							ServerName: "whoami",
							RootCAs:    []types.FileOrContent{"CA1"},
							Certificates: tls.Certificates{
								{
									CertFile: types.FileOrContent(listenerCert),
									KeyFile:  types.FileOrContent(listenerKey),
								},
							},
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute and BackendTLSPolicy with backendRequest timeout",
			paths: []string{"services.yml", "httproute/with_backend_tls_policy_and_backend_request_timeout.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.1:80",
									},
									{
										URL: "https://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							ServerName: "whoami",
							RootCAs: []types.FileOrContent{
								"CA1",
								"CA2",
								"CA1-secret",
								"CA2-secret",
							},
							ForwardingTimeouts: &dynamic.ForwardingTimeouts{
								DialTimeout:           ptypes.Duration(30 * time.Second),
								ResponseHeaderTimeout: ptypes.Duration(5 * time.Second),
								IdleConnTimeout:       ptypes.Duration(90 * time.Second),
								PingTimeout:           ptypes.Duration(15 * time.Second),
							},
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute and BackendTLSPolicy with subjectAltNames",
			paths: []string{"services.yml", "httproute/with_backend_tls_policy_san.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.1:80",
									},
									{
										URL: "https://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							ServerName:         "whoami",
							InsecureSkipVerify: true,
							PeerCertSANs: []tls.SAN{
								{Type: tls.SANDNSNameType, Value: "whoami.default.svc.cluster.local"},
								{Type: tls.SANURIType, Value: "spiffe://cluster.local/ns/default/sa/whoami"},
							},
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute and BackendTLSPolicy with System CA",
			paths: []string{"services.yml", "httproute/with_backend_tls_policy_system.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.1:80",
									},
									{
										URL: "https://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							ServerName: "whoami",
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:     "Simple HTTPRoute with NativeLBByDefault enabled",
			paths:    []string{"services.yml", "httproute/simple.yml"},
			nativeLB: true,
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.10.1:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute with NativeLB annotation",
			paths: []string{"services.yml", "httproute/simple_nativelb.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-native-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-native-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.10.1:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:     "Simple HTTPRoute with NativeLBByDefault enabled but service has disabled nativelb",
			paths:    []string{"services.yml", "httproute/simple_nativelb_disabled.yml"},
			nativeLB: true,
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-native-disabled-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-native-disabled-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.20:80",
									},
									{
										URL: "http://10.10.0.21:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute with nil endpointslice port name",
			paths: []string{"services.yml", "httproute/with_nil_endpointslice_port_name.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-nil-port-name-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-nil-port-name-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy:       dynamic.BalancerStrategyWRR,
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute with nil endpointslice port value",
			paths: []string{"services.yml", "httproute/with_nil_endpointslice_port_value.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-nil-port-value-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-nil-port-value-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy:       dynamic.BalancerStrategyWRR,
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoutes with identical rules",
			paths: []string{"services.yml", "httproute/with_identical_rules.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-foo-com-5f20d578e2e3eccac466": {
							EntryPoints: []string{"web"},
							Rule:        `Host("foo.com")`,
						},
						"httproute-default-http-app-catch-all-gw-default-my-gateway-ep-web-0-402d8b0319543c8e637b": {
							Service:    "httproute-default-http-app-catch-all-gw-default-my-gateway-ep-web-0-402d8b0319543c8e637b-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   9,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-foo-com-5f20d578e2e3eccac466"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-catch-all-gw-default-my-gateway-ep-web-0-402d8b0319543c8e637b-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-catch-all-gw-default-my-gateway-ep-web-0-402d8b0319543c8e637b-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-catch-all-gw-default-my-gateway-ep-web-0-402d8b0319543c8e637b-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute attached to a ListenerSet HTTPS listener",
			paths: []string{"services.yml", "httproute/with_listenerset.yml"},
			entryPoints: map[string]Entrypoint{
				"web":       {Address: ":80"},
				"websecure": {Address: ":443"},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282": {
							Service: "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-wrr",
							Rule:    `Host("foo.example.com") && Path("/bar")`,
							ParentRefs: []string{
								"listener-websecure-https-439feda8e3f1dbdca529",
							},
							RuleSyntax: "default",
							Priority:   100016,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{
								"websecure",
							},
							Rule: `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								Strategy:       dynamic.BalancerStrategy("wrr"),
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100000000),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					Middlewares:       map[string]*dynamic.Middleware{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Services:          map[string]*dynamic.TCPService{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							CipherSuites: []string{
								"TLS_AES_128_GCM_SHA256",
								"TLS_AES_256_GCM_SHA384",
								"TLS_CHACHA20_POLY1305_SHA256",
								"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
								"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
								"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
								"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
								"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
								"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
								"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
								"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
								"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
								"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
							},
							ALPNProtocols: []string{
								"h2",
								"http/1.1",
								"acme-tls/1",
							},
						},
					},
				},
			},
		},
		{
			desc:  "Gateway tls.frontend Default and PerPort validation register TLS options",
			paths: []string{"services.yml", "httproute/with_frontend_validation.yml"},
			entryPoints: map[string]Entrypoint{
				"websecure":  {Address: ":443"},
				"websecure2": {Address: ":8443"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"listener-websecure2-https-dd662c6921e9998d0cca": {
							EntryPoints: []string{"websecure2"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure2-https-dd662c6921e9998d0cca",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure2-0-e22d24f42cd0b2e1fc86": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure2-0-e22d24f42cd0b2e1fc86-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure2-https-dd662c6921e9998d0cca"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-398c6f76284d21e6e3b3-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure2-0-e22d24f42cd0b2e1fc86-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure2-0-e22d24f42cd0b2e1fc86-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure2-0-e22d24f42cd0b2e1fc86-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": func() tls.Options {
							opts := defaultListenerTLSOptions()
							opts.ClientAuth = tls.ClientAuth{
								CAFiles:        []types.FileOrContent{"CA1"},
								ClientAuthType: tls.RequireAndVerifyClientCert,
							}
							return opts
						}(),
						"listener-websecure2-https-dd662c6921e9998d0cca": func() tls.Options {
							opts := defaultListenerTLSOptions()
							opts.ClientAuth = tls.ClientAuth{
								CAFiles:        []types.FileOrContent{"CA1"},
								ClientAuthType: tls.RequestClientCert,
							}
							return opts
						}(),
					},
				},
			},
		},
		{
			desc:  "Gateways sharing an entry point with the same frontend validation keep their listeners isolated",
			paths: []string{"services.yml", "httproute/with_frontend_validation_listener_isolation.yml"},
			entryPoints: map[string]Entrypoint{
				"websecure": {Address: ":443"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `(Host("*")) && !(Host("foo.com"))`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"listener-websecure-https-foo-com-c4461032f627a1b9033b": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("foo.com")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-foo-com-c4461032f627a1b9033b",
							},
						},
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269": {
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529", "listener-websecure-https-foo-com-c4461032f627a1b9033b"},
							Service:    "httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269-wrr",
							Rule:       `Path("/catchall")`,
							Priority:   100001,
							RuleSyntax: "default",
						},
						"httproute-default-http-app-secure-gw-default-my-secure-gateway-ep-websecure-0-b7eaf4c1f751a09c9658": {
							ParentRefs: []string{"listener-websecure-https-foo-com-c4461032f627a1b9033b"},
							Service:    "httproute-default-http-app-secure-gw-default-my-secure-gateway-ep-websecure-0-b7eaf4c1f751a09c9658-wrr",
							Rule:       `Host("foo.com") && Path("/secure")`,
							Priority:   100008,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-secure-gw-default-my-secure-gateway-ep-websecure-0-b7eaf4c1f751a09c9658-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-secure-gw-default-my-secure-gateway-ep-websecure-0-b7eaf4c1f751a09c9658-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-secure-gw-default-my-secure-gateway-ep-websecure-0-b7eaf4c1f751a09c9658-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": func() tls.Options {
							opts := defaultListenerTLSOptions()
							opts.ClientAuth = tls.ClientAuth{
								CAFiles:        []types.FileOrContent{"CA1"},
								ClientAuthType: tls.RequireAndVerifyClientCert,
							}
							return opts
						}(),
						"listener-websecure-https-foo-com-c4461032f627a1b9033b": func() tls.Options {
							opts := defaultListenerTLSOptions()
							opts.ClientAuth = tls.ClientAuth{
								CAFiles:        []types.FileOrContent{"CA1"},
								ClientAuthType: tls.RequireAndVerifyClientCert,
							}
							return opts
						}(),
					},
				},
			},
		},
		{
			desc:  "Catch-all listeners sharing an entry point with different frontend validations are conflicted, the oldest Gateway wins",
			paths: []string{"services.yml", "httproute/with_frontend_validation_shared_entrypoint.yml"},
			entryPoints: map[string]Entrypoint{
				"websecure": {Address: ":443"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269": {
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
							Service:    "httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269-wrr",
							Rule:       `Path("/catchall")`,
							Priority:   100001,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-catchall-gw-default-my-gateway-ep-websecure-0-d0a1acd3dc8dc5905269-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.expected == nil {
				return
			}

			k8sObjects, gwObjects := readResources(t, test.paths)

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints:              test.entryPoints,
				NativeLBByDefault:        test.nativeLB,
				RetryMaxRequestBodyBytes: cmp.Or(test.retryMaxRequestBodyBytes, dynamic.RetryDefaultMaxRequestBodyBytes),
				client:                   client,
			}

			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			assert.Equal(t, test.expected, conf)
		})
	}
}

func TestLoadHTTPRoutes_backendExtensionRef(t *testing.T) {
	testCases := []struct {
		desc                  string
		paths                 []string
		groupKindBackendFuncs map[string]map[string]BuildBackendFunc
		expected              *dynamic.Configuration
		entryPoints           map[string]Entrypoint
	}{
		{
			desc:  "Simple HTTPRoute with TraefikService",
			paths: []string{"services.yml", "httproute/simple_with_TraefikService.yml"},
			groupKindBackendFuncs: map[string]map[string]BuildBackendFunc{
				traefikv1alpha1.GroupName: {"TraefikService": func(name, namespace string) (string, *dynamic.Service, error) {
					return name, nil, nil
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "whoami",
										Weight: new(1),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute with TraefikService with service configuration",
			paths: []string{"services.yml", "httproute/simple_with_TraefikService.yml"},
			groupKindBackendFuncs: map[string]map[string]BuildBackendFunc{
				traefikv1alpha1.GroupName: {"TraefikService": func(name, namespace string) (string, *dynamic.Service, error) {
					return name, &dynamic.Service{LoadBalancer: &dynamic.ServersLoadBalancer{Strategy: dynamic.BalancerStrategyWRR, Servers: []dynamic.Server{{URL: "foobar"}}}}, nil
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "whoami",
										Weight: new(1),
									},
								},
							},
						},
						"whoami": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "foobar"},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute with invalid TraefikService kind",
			paths: []string{"services.yml", "httproute/simple_with_TraefikService.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
										Status: new(500),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute with backendFunc error",
			paths: []string{"services.yml", "httproute/simple_with_TraefikService.yml"},
			groupKindBackendFuncs: map[string]map[string]BuildBackendFunc{
				traefikv1alpha1.GroupName: {"TraefikService": func(name, namespace string) (string, *dynamic.Service, error) {
					return "", nil, errors.New("BOOM")
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
										Status: new(500),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, with myservice@file service",
			paths: []string{"services.yml", "httproute/simple_cross_provider.yml"},
			groupKindBackendFuncs: map[string]map[string]BuildBackendFunc{
				traefikv1alpha1.GroupName: {"TraefikService": func(name, namespace string) (string, *dynamic.Service, error) {
					// func should never be executed in case of cross-provider reference.
					return "", nil, errors.New("BOOM")
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "service@file",
										Weight: new(1),
									},
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-1",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-1": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple HTTPRoute, with appProtocol service",
			paths: []string{"services.yml", "httproute/with_app_protocol_service.yml"},
			groupKindBackendFuncs: map[string]map[string]BuildBackendFunc{
				traefikv1alpha1.GroupName: {"TraefikService": func(name, namespace string) (string, *dynamic.Service, error) {
					// func should never be executed in case of cross-provider reference.
					return "", nil, errors.New("BOOM")
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd": {
							Service:    "httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-h2c-0",
										Weight: new(1),
									},
									{
										Name:   "httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-ws-1",
										Weight: new(1),
									},
									{
										Name:   "httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-wss-2",
										Weight: new(1),
									},
									{
										Name:   "httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-HTTP-3",
										Weight: new(1),
									},
									{
										Name:   "httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-HTTPS-4",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-h2c-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "h2c://10.10.0.13:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-ws-1": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.14:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-wss-2": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.15:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-HTTPS-4": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.16:8443",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-multi-protocols-gw-default-my-gateway-ep-web-0-bbdaa7d1e726d90930fd-svc-default-whoami-HTTP-3": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.17:8080",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.expected == nil {
				return
			}

			k8sObjects, gwObjects := readResources(t, test.paths)

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints: test.entryPoints,
				client:      client,
			}

			for group, kindFuncs := range test.groupKindBackendFuncs {
				for kind, backendFunc := range kindFuncs {
					p.RegisterBackendFuncs(group, kind, backendFunc)
				}
			}
			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			assert.Equal(t, test.expected, conf)
		})
	}
}

func TestLoadHTTPRoutes_filterExtensionRef(t *testing.T) {
	testCases := []struct {
		desc                 string
		groupKindFilterFuncs map[string]map[string]BuildFilterFunc
		expected             *dynamic.Configuration
		entryPoints          map[string]Entrypoint
	}{
		{
			desc: "ExtensionRef filter",
			groupKindFilterFuncs: map[string]map[string]BuildFilterFunc{
				traefikv1alpha1.GroupName: {"Middleware": func(name, namespace string) (string, *dynamic.Middleware, error) {
					return namespace + "-" + name, nil, nil
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Middlewares: []string{
								"default-my-first-middleware",
								"default-my-second-middleware",
							},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-bar-whoami-bar-1",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						// Cross-namespace backendRef: its ExtensionRef filter must resolve to the route's own namespace,
						// not the backendRef's namespace.
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-bar-whoami-bar-1": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.11:80",
									},
									{
										URL: "http://10.10.0.12:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
							Middlewares: []string{"default-my-backend-middleware"},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "ExtensionRef filter with middleware creation",
			groupKindFilterFuncs: map[string]map[string]BuildFilterFunc{
				traefikv1alpha1.GroupName: {"Middleware": func(name, namespace string) (string, *dynamic.Middleware, error) {
					return namespace + "-" + name, &dynamic.Middleware{Headers: &dynamic.Headers{CustomRequestHeaders: map[string]string{"Test-Header": "Test"}}}, nil
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Middlewares: []string{
								"default-my-first-middleware",
								"default-my-second-middleware",
							},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"default-my-first-middleware":   {Headers: &dynamic.Headers{CustomRequestHeaders: map[string]string{"Test-Header": "Test"}}},
						"default-my-second-middleware":  {Headers: &dynamic.Headers{CustomRequestHeaders: map[string]string{"Test-Header": "Test"}}},
						"default-my-backend-middleware": {Headers: &dynamic.Headers{CustomRequestHeaders: map[string]string{"Test-Header": "Test"}}},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0",
										Weight: new(1),
									},
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-bar-whoami-bar-1",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						// Cross-namespace backendRef: its ExtensionRef filter must resolve to the route's own namespace,
						// not the backendRef's namespace.
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-bar-whoami-bar-1": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.11:80",
									},
									{
										URL: "http://10.10.0.12:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
							Middlewares: []string{"default-my-backend-middleware"},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Unknown ExtensionRef filter",
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-err-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-err-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "invalid-httproute-filter",
										Weight: new(1),
										Status: new(500),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "ExtensionRef filter with filterFunc error",
			groupKindFilterFuncs: map[string]map[string]BuildFilterFunc{
				traefikv1alpha1.GroupName: {"Middleware": func(name, namespace string) (string, *dynamic.Middleware, error) {
					return "", nil, errors.New("BOOM")
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-err-wrr",
							Rule:       `Host("foo.com") && Path("/bar")`,
							Priority:   100008,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-err-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "invalid-httproute-filter",
										Weight: new(1),
										Status: new(500),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.expected == nil {
				return
			}

			k8sObjects, gwObjects := readResources(t, []string{"services.yml", "httproute/filter_extension_ref.yml"})

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints: test.entryPoints,
				client:      client,
			}

			for group, kindFuncs := range test.groupKindFilterFuncs {
				for kind, filterFunc := range kindFuncs {
					p.RegisterFilterFuncs(group, kind, filterFunc)
				}
			}
			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			assert.NoError(t, err)

			assert.Equal(t, test.expected, conf)
		})
	}
}

func TestLoadGRPCRoutes(t *testing.T) {
	testCases := []struct {
		desc        string
		paths       []string
		expected    *dynamic.Configuration
		entryPoints map[string]Entrypoint
	}{
		{
			desc:  "Simple GRPCRoute and BackendTLSPolicy with CA certificate",
			paths: []string{"services.yml", "grpcroute/with_backend_tls_policy.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.1:80",
									},
									{
										URL: "https://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							ServerName: "whoami",
							RootCAs: []types.FileOrContent{
								"CA1",
								"CA2",
								"CA1-secret",
								"CA2-secret",
							},
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "GRPCRoute and BackendTLSPolicy with a Gateway backend client certificate",
			paths: []string{"services.yml", "grpcroute/with_backend_tls_policy_and_client_certificate.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.1:80",
									},
									{
										URL: "https://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							ServerName: "whoami",
							RootCAs:    []types.FileOrContent{"CA1"},
							Certificates: tls.Certificates{
								{
									CertFile: types.FileOrContent(listenerCert),
									KeyFile:  types.FileOrContent(listenerKey),
								},
							},
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple GRPCRoute and BackendTLSPolicy with subjectAltNames",
			paths: []string{"services.yml", "grpcroute/with_backend_tls_policy_san.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.1:80",
									},
									{
										URL: "https://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							ServerName:         "whoami",
							InsecureSkipVerify: true,
							PeerCertSANs: []tls.SAN{
								{Type: tls.SANDNSNameType, Value: "whoami.default.svc.cluster.local"},
								{Type: tls.SANURIType, Value: "spiffe://cluster.local/ns/default/sa/whoami"},
							},
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple GRPCRoute and BackendTLSPolicy with System CA",
			paths: []string{"services.yml", "grpcroute/with_backend_tls_policy_system.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "https://10.10.0.1:80",
									},
									{
										URL: "https://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
								ServersTransport: "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							ServerName: "whoami",
						},
					},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "GRPCRoute Gateway tls.frontend Default and PerPort validation register TLS options",
			paths: []string{"services.yml", "grpcroute/with_frontend_validation.yml"},
			entryPoints: map[string]Entrypoint{
				"websecure":  {Address: ":443"},
				"websecure2": {Address: ":8443"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"listener-websecure2-https-dd662c6921e9998d0cca": {
							EntryPoints: []string{"websecure2"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure2-https-dd662c6921e9998d0cca",
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure-0-5015e40a36a5159a5ca2": {
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure-0-5015e40a36a5159a5ca2-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure2-0-c1153de85fa10a42dc33": {
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure2-0-c1153de85fa10a42dc33-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure2-https-dd662c6921e9998d0cca"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure-0-5015e40a36a5159a5ca2-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure-0-5015e40a36a5159a5ca2-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure-0-5015e40a36a5159a5ca2-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "h2c://10.10.0.1:80",
									},
									{
										URL: "h2c://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure2-0-c1153de85fa10a42dc33-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure2-0-c1153de85fa10a42dc33-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-websecure2-0-c1153de85fa10a42dc33-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "h2c://10.10.0.1:80",
									},
									{
										URL: "h2c://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": func() tls.Options {
							opts := defaultListenerTLSOptions()
							opts.ClientAuth = tls.ClientAuth{
								CAFiles:        []types.FileOrContent{"CA1"},
								ClientAuthType: tls.RequireAndVerifyClientCert,
							}
							return opts
						}(),
						"listener-websecure2-https-dd662c6921e9998d0cca": func() tls.Options {
							opts := defaultListenerTLSOptions()
							opts.ClientAuth = tls.ClientAuth{
								CAFiles:        []types.FileOrContent{"CA1"},
								ClientAuthType: tls.RequestClientCert,
							}
							return opts
						}(),
					},
				},
			},
		},
		{
			desc:  "GRPCRoute Gateways sharing an entry point with the same frontend validation keep their listeners isolated",
			paths: []string{"services.yml", "grpcroute/with_frontend_validation_listener_isolation.yml"},
			entryPoints: map[string]Entrypoint{
				"websecure": {Address: ":443"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `(Host("*")) && !(Host("foo.com"))`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"listener-websecure-https-foo-com-c4461032f627a1b9033b": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("foo.com")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-foo-com-c4461032f627a1b9033b",
							},
						},
						"grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57": {
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529", "listener-websecure-https-foo-com-c4461032f627a1b9033b"},
							Service:    "grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   15,
							RuleSyntax: "default",
						},
						"grpcroute-default-grpc-app-secure-gw-default-my-secure-gateway-ep-websecure-0-3aae484bce849b96c90c": {
							ParentRefs: []string{"listener-websecure-https-foo-com-c4461032f627a1b9033b"},
							Service:    "grpcroute-default-grpc-app-secure-gw-default-my-secure-gateway-ep-websecure-0-3aae484bce849b96c90c-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "h2c://10.10.0.1:80"},
									{URL: "h2c://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"grpcroute-default-grpc-app-secure-gw-default-my-secure-gateway-ep-websecure-0-3aae484bce849b96c90c-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-secure-gw-default-my-secure-gateway-ep-websecure-0-3aae484bce849b96c90c-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-secure-gw-default-my-secure-gateway-ep-websecure-0-3aae484bce849b96c90c-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "h2c://10.10.0.1:80"},
									{URL: "h2c://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": func() tls.Options {
							opts := defaultListenerTLSOptions()
							opts.ClientAuth = tls.ClientAuth{
								CAFiles:        []types.FileOrContent{"CA1"},
								ClientAuthType: tls.RequireAndVerifyClientCert,
							}
							return opts
						}(),
						"listener-websecure-https-foo-com-c4461032f627a1b9033b": func() tls.Options {
							opts := defaultListenerTLSOptions()
							opts.ClientAuth = tls.ClientAuth{
								CAFiles:        []types.FileOrContent{"CA1"},
								ClientAuthType: tls.RequireAndVerifyClientCert,
							}
							return opts
						}(),
					},
				},
			},
		},
		{
			desc:  "GRPCRoute catch-all listeners sharing an entry point with different frontend validations are conflicted, the oldest Gateway wins",
			paths: []string{"services.yml", "grpcroute/with_frontend_validation_shared_entrypoint.yml"},
			entryPoints: map[string]Entrypoint{
				"websecure": {Address: ":443"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57": {
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
							Service:    "grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   15,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-catchall-gw-default-my-gateway-ep-websecure-0-4b61347f65fc7f1a8a57-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "h2c://10.10.0.1:80"},
									{URL: "h2c://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
				},
			},
		},
		{
			desc:  "GRPCRoute with nil endpointslice port name",
			paths: []string{"services.yml", "grpcroute/with_nil_endpointslice_port_name.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":80"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-nil-port-name-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-nil-port-name-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy:       dynamic.BalancerStrategyWRR,
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "GRPCRoute with nil endpointslice port value",
			paths: []string{"services.yml", "grpcroute/with_nil_endpointslice_port_value.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":80"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-nil-port-value-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-nil-port-value-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy:       dynamic.BalancerStrategyWRR,
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "GRPCRoutes with identical rules",
			paths: []string{"services.yml", "grpcroute/with_identical_rules.yml"},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-foo-com-5f20d578e2e3eccac466": {
							EntryPoints: []string{"web"},
							Rule:        `Host("foo.com")`,
						},
						"grpcroute-default-grpc-app-catch-all-gw-default-my-gateway-ep-web-0-24ad420c9f1560b62d23": {
							Service:    "grpcroute-default-grpc-app-catch-all-gw-default-my-gateway-ep-web-0-24ad420c9f1560b62d23-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-foo-com-5f20d578e2e3eccac466"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-catch-all-gw-default-my-gateway-ep-web-0-24ad420c9f1560b62d23-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-catch-all-gw-default-my-gateway-ep-web-0-24ad420c9f1560b62d23-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-catch-all-gw-default-my-gateway-ep-web-0-24ad420c9f1560b62d23-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "h2c://10.10.0.1:80",
									},
									{
										URL: "h2c://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.expected == nil {
				return
			}

			k8sObjects, gwObjects := readResources(t, test.paths)

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				<-eventCh
			}

			p := Provider{
				EntryPoints: test.entryPoints,
				client:      client,
			}

			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			assert.Equal(t, test.expected, conf)
		})
	}
}

func TestLoadGRPCRoutes_filterExtensionRef(t *testing.T) {
	testCases := []struct {
		desc                 string
		groupKindFilterFuncs map[string]map[string]BuildFilterFunc
		expected             *dynamic.Configuration
		entryPoints          map[string]Entrypoint
	}{
		{
			desc: "ExtensionRef filter",
			groupKindFilterFuncs: map[string]map[string]BuildFilterFunc{
				traefikv1alpha1.GroupName: {"Middleware": func(name, namespace string) (string, *dynamic.Middleware, error) {
					return namespace + "-" + name, nil, nil
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Middlewares: []string{
								"default-my-first-middleware",
								"default-my-second-middleware",
							},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "h2c://10.10.0.1:80",
									},
									{
										URL: "h2c://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "ExtensionRef filter with middleware creation",
			groupKindFilterFuncs: map[string]map[string]BuildFilterFunc{
				traefikv1alpha1.GroupName: {"Middleware": func(name, namespace string) (string, *dynamic.Middleware, error) {
					return namespace + "-" + name, &dynamic.Middleware{Headers: &dynamic.Headers{CustomRequestHeaders: map[string]string{"Test-Header": "Test"}}}, nil
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Middlewares: []string{
								"default-my-first-middleware",
								"default-my-second-middleware",
							},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{
						"default-my-first-middleware":  {Headers: &dynamic.Headers{CustomRequestHeaders: map[string]string{"Test-Header": "Test"}}},
						"default-my-second-middleware": {Headers: &dynamic.Headers{CustomRequestHeaders: map[string]string{"Test-Header": "Test"}}},
					},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "h2c://10.10.0.1:80",
									},
									{
										URL: "h2c://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Unknown ExtensionRef filter",
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-err-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-err-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "invalid-grpcroute-filter",
										Weight: new(1),
										GRPCStatus: &dynamic.GRPCStatus{
											Code: codes.Unavailable,
											Msg:  "Service Unavailable",
										},
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "ExtensionRef filter with filterFunc error",
			groupKindFilterFuncs: map[string]map[string]BuildFilterFunc{
				traefikv1alpha1.GroupName: {"Middleware": func(name, namespace string) (string, *dynamic.Middleware, error) {
					return "", nil, errors.New("BOOM")
				}},
			},
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":80",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296": {
							Service:    "grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-err-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-err-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "invalid-grpcroute-filter",
										Weight: new(1),
										GRPCStatus: &dynamic.GRPCStatus{
											Code: codes.Unavailable,
											Msg:  "Service Unavailable",
										},
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.expected == nil {
				return
			}

			k8sObjects, gwObjects := readResources(t, []string{"services.yml", "grpcroute/filter_extension_ref.yml"})

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints: test.entryPoints,
				client:      client,
			}

			for group, kindFuncs := range test.groupKindFilterFuncs {
				for kind, filterFunc := range kindFuncs {
					p.RegisterFilterFuncs(group, kind, filterFunc)
				}
			}
			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			assert.NoError(t, err)

			assert.Equal(t, test.expected, conf)
		})
	}
}

func TestLoadTCPRoutes(t *testing.T) {
	testCases := []struct {
		desc         string
		ingressClass string
		paths        []string
		expected     *dynamic.Configuration
		entryPoints  map[string]Entrypoint
		nativeLB     bool
	}{
		{
			desc: "Empty",
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because missing entry point",
			paths: []string{"services.yml", "tcproute/simple.yml"},
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8000",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because no tcp route defined",
			paths: []string{"services.yml", "tcproute/without_tcproute.yml"},
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8080",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by missing GatewayClass",
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8080",
			}},
			paths: []string{"services.yml", "tcproute/without_gatewayclass.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by unknown GatewayClass controller desc",
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8080",
			}},
			paths: []string{"services.yml", "tcproute/gatewayclass_with_unknown_controller.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by HTTPRoute with TLS configuration",
			entryPoints: map[string]Entrypoint{"web": {
				Address: ":8080",
			}},
			paths: []string{"services.yml", "tcproute/with_tls_configuration.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Router with service in error caused by wrong TargetPort",
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8080",
			}},
			paths: []string{"services.yml", "tcproute/with_wrong_service_port.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-TCP-app-1-gw-default-my-gateway-ep-TCP-0-b06faef5537f5b1e4c70": {
							EntryPoints: []string{"TCP"},
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
							Service:     "tcproute-default-TCP-app-1-gw-default-my-gateway-ep-TCP-0-b06faef5537f5b1e4c70-wrr",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-TCP-app-1-gw-default-my-gateway-ep-TCP-0-b06faef5537f5b1e4c70-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{{
									Name:   "tcproute-default-TCP-app-1-gw-default-my-gateway-ep-TCP-0-b06faef5537f5b1e4c70-err-lb",
									Weight: new(1),
								}},
							},
						},
						"tcproute-default-TCP-app-1-gw-default-my-gateway-ep-TCP-0-b06faef5537f5b1e4c70-err-lb": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TCPRoute",
			paths: []string{"services.yml", "tcproute/simple.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Multiple TCPRoute",
			paths: []string{"services.yml", "tcproute/with_multiple_routes.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp-1":   {Address: ":9000"},
				"tcp-2":   {Address: ":10000"},
				"not-tcp": {Address: ":11000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-1-0-538fb5c01603fd936dbf": {
							EntryPoints: []string{"tcp-1"},
							Service:     "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-1-0-538fb5c01603fd936dbf-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
						"tcproute-default-tcp-app-2-gw-default-my-tcp-gateway-ep-tcp-2-0-8eec98026c1dabcb62e0": {
							EntryPoints: []string{"tcp-2"},
							Service:     "tcproute-default-tcp-app-2-gw-default-my-tcp-gateway-ep-tcp-2-0-8eec98026c1dabcb62e0-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-1-0-538fb5c01603fd936dbf-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-1-0-538fb5c01603fd936dbf-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-2-gw-default-my-tcp-gateway-ep-tcp-2-0-8eec98026c1dabcb62e0-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-2-gw-default-my-tcp-gateway-ep-tcp-2-0-8eec98026c1dabcb62e0-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-1-0-538fb5c01603fd936dbf-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
						"tcproute-default-tcp-app-2-gw-default-my-tcp-gateway-ep-tcp-2-0-8eec98026c1dabcb62e0-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:10000",
									},
									{
										Address: "10.10.0.10:10000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TCPRoute with multiple rules",
			paths: []string{"services.yml", "tcproute/with_multiple_rules.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp-1":   {Address: ":9000"},
				"tcp-2":   {Address: ":10000"},
				"not-tcp": {Address: ":11000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-1-0-1767e543195002a45b0b": {
							EntryPoints: []string{"tcp-1"},
							Service:     "tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-1-0-1767e543195002a45b0b-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-1-0-1767e543195002a45b0b-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-1-0-1767e543195002a45b0b-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-1-0-1767e543195002a45b0b-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TCPRoute, with backendRef",
			paths: []string{"services.yml", "tcproute/simple_cross_provider.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "service@file",
										Weight: new(1),
									},
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da-svc-default-whoamitcp-1",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da-svc-default-whoamitcp-1": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty caused by TCPRoute attached to a TLS listener",
			paths: []string{"services.yml", "tcproute/with_protocol_tls.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Conflicting TCPRoutes on the same listener, only the oldest is bound",
			paths: []string{"services.yml", "tcproute/with_conflicting_routes.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-older-gw-default-my-tcp-gateway-ep-tcp-0-6aa3057a8ee9e5d61398": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-older-gw-default-my-tcp-gateway-ep-tcp-0-6aa3057a8ee9e5d61398-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-older-gw-default-my-tcp-gateway-ep-tcp-0-6aa3057a8ee9e5d61398-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-older-gw-default-my-tcp-gateway-ep-tcp-0-6aa3057a8ee9e5d61398-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-older-gw-default-my-tcp-gateway-ep-tcp-0-6aa3057a8ee9e5d61398-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TCPRoute with Same namespace selector",
			paths: []string{"services.yml", "tcproute/with_namespace_same.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-default-gw-default-my-tcp-gateway-ep-tcp-0-9c02509721599badb6c4": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-default-gw-default-my-tcp-gateway-ep-tcp-0-9c02509721599badb6c4-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-default-gw-default-my-tcp-gateway-ep-tcp-0-9c02509721599badb6c4-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-default-gw-default-my-tcp-gateway-ep-tcp-0-9c02509721599badb6c4-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-default-gw-default-my-tcp-gateway-ep-tcp-0-9c02509721599badb6c4-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TCPRoute with All namespace selector",
			paths: []string{"services.yml", "tcproute/with_namespace_all.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171-svc-bar-whoamitcp-bar-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171-svc-bar-whoamitcp-bar-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.13:9000",
									},
									{
										Address: "10.10.0.14:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TCPRoute with namespace selector",
			paths: []string{"services.yml", "tcproute/with_namespace_selector.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171-svc-bar-whoamitcp-bar-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-bar-tcp-app-bar-gw-default-my-tcp-gateway-ep-tcp-0-18313c5ba079a1d9e171-svc-bar-whoamitcp-bar-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.13:9000",
									},
									{
										Address: "10.10.0.14:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:     "Simple TCPRoute with NativeLBByDefault",
			paths:    []string{"services.yml", "tcproute/simple.yml"},
			nativeLB: true,
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.10.1:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TCPRoute with NativeLB annotation",
			paths: []string{"services.yml", "tcproute/simple_nativelb.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-native-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-native-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.10.1:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:     "Simple TCPRoute with NativeLBByDefault enabled but service has disabled nativelb",
			paths:    []string{"services.yml", "tcproute/simple_nativelb_disabled.yml"},
			nativeLB: true,
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-native-disabled-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-native-disabled-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.30:9000",
									},
									{
										Address: "10.10.0.31:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TCPRoute with nil endpointslice port name",
			paths: []string{"services.yml", "tcproute/with_nil_endpointslice_port_name.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-nil-port-name-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-nil-port-name-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TCPRoute with nil endpointslice port value",
			paths: []string{"services.yml", "tcproute/with_nil_endpointslice_port_value.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-nil-port-value-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-nil-port-value-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TCPRoutes with identical rules",
			paths: []string{"services.yml", "tcproute/with_identical_rules.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-first-gw-default-my-tcp-gateway-ep-tcp-0-79a3b41b9e16bdaa2841": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-first-gw-default-my-tcp-gateway-ep-tcp-0-79a3b41b9e16bdaa2841-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-first-gw-default-my-tcp-gateway-ep-tcp-0-79a3b41b9e16bdaa2841-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-first-gw-default-my-tcp-gateway-ep-tcp-0-79a3b41b9e16bdaa2841-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-first-gw-default-my-tcp-gateway-ep-tcp-0-79a3b41b9e16bdaa2841-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TCPRoute attached to a ListenerSet listener",
			paths: []string{"services.yml", "tcproute/with_listenerset.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":80"},
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Services:          map[string]*dynamic.Service{},
					Middlewares:       map[string]*dynamic.Middleware{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf": {
							EntryPoints: []string{
								"tcp",
							},
							Service:    "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr",
							Rule:       `HostSNI("*")`,
							RuleSyntax: "default",
						},
					},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-tcp-gateway-ep-tcp-0-b0953a7f472d14e31cbf-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.expected == nil {
				return
			}

			k8sObjects, gwObjects := readResources(t, test.paths)

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints:       test.entryPoints,
				NativeLBByDefault: test.nativeLB,
				client:            client,
			}

			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			assert.Equal(t, test.expected, conf)
		})
	}
}

func TestLoadTLSRoutes(t *testing.T) {
	testCases := []struct {
		desc         string
		ingressClass string
		paths        []string
		entryPoints  map[string]Entrypoint
		nativeLB     bool
		expected     *dynamic.Configuration
	}{
		{
			desc: "Empty",
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because no matching entry point",
			paths: []string{"services.yml", "tlsroute/simple_TLS_to_TCPRoute.yml"},
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8000",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because no tls route defined",
			paths: []string{"services.yml", "tlsroute/without_tlsroute.yml"},
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8080",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by missing GatewayClass",
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8080",
			}},
			paths: []string{"services.yml", "tlsroute/without_gatewayclass.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by unknown GatewayClass controller desc",
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8080",
			}},
			paths: []string{"services.yml", "tlsroute/gatewayclass_with_unknown_controller.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Router with service in error caused by wrong TargetPort",
			entryPoints: map[string]Entrypoint{"TCP": {
				Address: ":8080",
			}},
			paths: []string{"services.yml", "tlsroute/with_wrong_service_port.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-TCP-0-45c28fdcc89d443a020b": {
							EntryPoints: []string{"TCP"},
							Priority:    0,
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-TCP-0-45c28fdcc89d443a020b-wrr",
							TLS:         &dynamic.RouterTCPTLSConfig{},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-TCP-0-45c28fdcc89d443a020b-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{{
									Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-TCP-0-45c28fdcc89d443a020b-err-lb",
									Weight: new(1),
								}},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-TCP-0-45c28fdcc89d443a020b-err-lb": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc: "Empty caused by mixed routes with wrong parent ref",
			entryPoints: map[string]Entrypoint{
				"tcp": {
					Address: ":9000",
				},
				"tcp-tls": {
					Address: ":9443",
				},
				"http": {
					Address: ":80",
				},
			},
			paths: []string{"services.yml", "mixed/with_wrong_routes_selector.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty caused by TCPRoute attached to a TLS listener in Terminate mode",
			paths: []string{"services.yml", "tlsroute/simple_TLS_to_TCPRoute.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty caused by TCPRoute attached to a TLS listener in Passthrough mode",
			paths: []string{"services.yml", "tlsroute/simple_TLS_to_TCPRoute_passthrough.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TLS listener to TLSRoute",
			paths: []string{"services.yml", "tlsroute/simple_TLS_to_TLSRoute.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef": {
							EntryPoints: []string{"tcp"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr",
							Priority:    15,
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Multiple TLSRoute",
			paths: []string{"services.yml", "tlsroute/with_multiple_routes_kind.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
				"tcp": {Address: ":10000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-24d35d31a7578c5062a7": {
							EntryPoints: []string{"tcp"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-24d35d31a7578c5062a7-wrr",
							Priority:    0,
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-24d35d31a7578c5062a7-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-24d35d31a7578c5062a7-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-24d35d31a7578c5062a7-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:10000",
									},
									{
										Address: "10.10.0.10:10000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TLSRoute, with backendRef",
			paths: []string{"services.yml", "tlsroute/simple_cross_provider.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-efc0f8d903baa9fb6592": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-efc0f8d903baa9fb6592-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
							TLS:         &dynamic.RouterTCPTLSConfig{},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-efc0f8d903baa9fb6592-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "service@file",
										Weight: new(1),
									},
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-efc0f8d903baa9fb6592-svc-default-whoamitcp-1",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-efc0f8d903baa9fb6592-svc-default-whoamitcp-1": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "Simple TLSRoute, with Passthrough and TLS configuration should raise a warn",
			paths: []string{"services.yml", "tlsroute/with_passthrough_tls.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr",
							Priority:    15,
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TLSRoute, with Passthrough",
			paths: []string{"services.yml", "tlsroute/with_passthrough.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr",
							Priority:    15,
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TLSRoute, with single SNI matching",
			paths: []string{"services.yml", "tlsroute/with_SNI_matching.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr",
							Priority:    15,
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TLSRoute attached to TLS listeners with different hostnames on the same entry point",
			paths: []string{"services.yml", "tlsroute/with_multiple_listeners_same_port.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr",
							Priority:    15,
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-276f6c243a7761966c5b": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-276f6c243a7761966c5b-wrr",
							Priority:    15,
							Rule:        `HostSNI("bar.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-276f6c243a7761966c5b-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-276f6c243a7761966c5b-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-276f6c243a7761966c5b-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TLSRoute, with multiple SNI matching",
			paths: []string{"services.yml", "tlsroute/with_multiple_SNI_matching.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-035a9ab93ab8d0d12874": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-035a9ab93ab8d0d12874-wrr",
							Priority:    15,
							Rule:        `HostSNI("bar.example.com") || HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-035a9ab93ab8d0d12874-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-035a9ab93ab8d0d12874-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-035a9ab93ab8d0d12874-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TLSRoute with Same namespace selector",
			paths: []string{"services.yml", "tlsroute/with_namespace_same.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d-wrr",
							Priority:    11,
							Rule:        `HostSNI("foo.default")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TLSRoute with All namespace selector",
			paths: []string{"services.yml", "tlsroute/with_namespace_all.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d-wrr",
							Priority:    11,
							Rule:        `HostSNI("foo.default")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
						"tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b-wrr",
							Priority:    7,
							Rule:        `HostSNI("foo.bar")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b-svc-bar-whoamitcp-bar-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-0-76eff62990bc46a9697d-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
						"tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b-svc-bar-whoamitcp-bar-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.13:9000",
									},
									{
										Address: "10.10.0.14:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TLSRoute with namespace selector",
			paths: []string{"services.yml", "tlsroute/with_namespace_selector.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b-wrr",
							Priority:    7,
							Rule:        `HostSNI("foo.bar")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b-svc-bar-whoamitcp-bar-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-0-58a6fd0c8bc014d99d8b-svc-bar-whoamitcp-bar-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.13:9000",
									},
									{
										Address: "10.10.0.14:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TLSRoute with multiple rules",
			paths: []string{"services.yml", "tlsroute/with_multiple_rules.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp-1": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-gw-default-my-gateway-ep-tcp-1-0-2e99b4a7e81cc622b335": {
							EntryPoints: []string{"tcp-1"},
							Service:     "tlsroute-default-tls-app-gw-default-my-gateway-ep-tcp-1-0-2e99b4a7e81cc622b335-wrr",
							Priority:    0,
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-gw-default-my-gateway-ep-tcp-1-0-2e99b4a7e81cc622b335-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-gw-default-my-gateway-ep-tcp-1-0-2e99b4a7e81cc622b335-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-gw-default-my-gateway-ep-tcp-1-0-2e99b4a7e81cc622b335-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:     "Simple TLSRoute with NativeLBByDefault",
			paths:    []string{"services.yml", "tlsroute/simple_TLS_to_TLSRoute.yml"},
			nativeLB: true,
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef": {
							EntryPoints: []string{"tcp"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr",
							Priority:    15,
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.10.1:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TLSRoute with NativeLB annotation",
			paths: []string{"services.yml", "tlsroute/simple_nativelb.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef": {
							EntryPoints: []string{"tcp"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr",
							Priority:    15,
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-native-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-native-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.10.1:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Simple TLSRoute and BackendTLSPolicy with CA certificate",
			paths: []string{"services.yml", "tlsroute/with_backend_tls_policy.yml"},
			entryPoints: map[string]Entrypoint{"tls": {
				Address: ":9001",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-wrr",
							Rule:        `HostSNI("foo.com")`,
							Priority:    7,
							RuleSyntax:  "default",
							TLS:         &dynamic.RouterTCPTLSConfig{},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.1:80",
									},
									{
										Address: "10.10.0.2:80",
									},
								},
								ServersTransport: "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0": {
							TLS: &dynamic.TLSClientConfig{
								ServerName: "whoami",
								RootCAs: []types.FileOrContent{
									"CA1",
									"CA2",
									"CA1-secret",
									"CA2-secret",
								},
							},
						},
					},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "TLSRoute and BackendTLSPolicy with a Gateway backend client certificate",
			paths: []string{"services.yml", "tlsroute/with_backend_tls_policy_and_client_certificate.yml"},
			entryPoints: map[string]Entrypoint{"tls": {
				Address: ":9001",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-wrr",
							Rule:        `HostSNI("foo.com")`,
							Priority:    7,
							RuleSyntax:  "default",
							TLS:         &dynamic.RouterTCPTLSConfig{},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.1:80",
									},
									{
										Address: "10.10.0.2:80",
									},
								},
								ServersTransport: "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0": {
							TLS: &dynamic.TLSClientConfig{
								ServerName: "whoami",
								RootCAs:    []types.FileOrContent{"CA1"},
								Certificates: tls.Certificates{
									{
										CertFile: types.FileOrContent(listenerCert),
										KeyFile:  types.FileOrContent(listenerKey),
									},
								},
							},
						},
					},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "Simple TLSRoute and BackendTLSPolicy with subjectAltNames",
			paths: []string{"services.yml", "tlsroute/with_backend_tls_policy_san.yml"},
			entryPoints: map[string]Entrypoint{"tls": {
				Address: ":9001",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-wrr",
							Rule:        `HostSNI("foo.com")`,
							Priority:    7,
							RuleSyntax:  "default",
							TLS:         &dynamic.RouterTCPTLSConfig{},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.1:80",
									},
									{
										Address: "10.10.0.2:80",
									},
								},
								ServersTransport: "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0": {
							TLS: &dynamic.TLSClientConfig{
								ServerName:         "whoami",
								InsecureSkipVerify: true,
								PeerCertSANs: []tls.SAN{
									{Type: tls.SANDNSNameType, Value: "whoami.default.svc.cluster.local"},
									{Type: tls.SANURIType, Value: "spiffe://cluster.local/ns/default/sa/whoami"},
								},
							},
						},
					},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "Simple TLSRoute and BackendTLSPolicy with System CA",
			paths: []string{"services.yml", "tlsroute/with_backend_tls_policy_system.yml"},
			entryPoints: map[string]Entrypoint{"tls": {
				Address: ":9001",
			}},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-wrr",
							Rule:        `HostSNI("foo.com")`,
							Priority:    7,
							RuleSyntax:  "default",
							TLS:         &dynamic.RouterTCPTLSConfig{},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.1:80",
									},
									{
										Address: "10.10.0.2:80",
									},
								},
								ServersTransport: "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0",
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0": {
							TLS: &dynamic.TLSClientConfig{
								ServerName: "whoami",
							},
						},
					},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "TLSRoute with nil endpointslice port name",
			paths: []string{"services.yml", "tlsroute/with_nil_endpointslice_port_name.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef": {
							EntryPoints: []string{"tcp"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr",
							Priority:    15,
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-nil-port-name-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-nil-port-name-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TLSRoute with nil endpointslice port value",
			paths: []string{"services.yml", "tlsroute/with_nil_endpointslice_port_value.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef": {
							EntryPoints: []string{"tcp"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr",
							Priority:    15,
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-nil-port-value-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-tls-gateway-ep-tcp-0-d01849ae1c96d9ffe0ef-svc-default-whoamitcp-nil-port-value-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TLSRoutes with identical rules",
			paths: []string{"services.yml", "tlsroute/with_identical_rules.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tlsroute-default-tls-app-catch-all-gw-default-my-tls-gateway-ep-tcp-0-60f8cf9276253172d000": {
							EntryPoints: []string{"tcp"},
							Service:     "tlsroute-default-tls-app-catch-all-gw-default-my-tls-gateway-ep-tcp-0-60f8cf9276253172d000-wrr",
							Rule:        `HostSNI("foo.com")`,
							Priority:    7,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tlsroute-default-tls-app-catch-all-gw-default-my-tls-gateway-ep-tcp-0-60f8cf9276253172d000-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-catch-all-gw-default-my-tls-gateway-ep-tcp-0-60f8cf9276253172d000-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-catch-all-gw-default-my-tls-gateway-ep-tcp-0-60f8cf9276253172d000-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TLSRoute attached to a ListenerSet Passthrough listener",
			paths: []string{"services.yml", "tlsroute/with_listenerset.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":80"},
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Services:          map[string]*dynamic.Service{},
					Middlewares:       map[string]*dynamic.Middleware{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Service:  "deny-unknown-host",
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174": {
							EntryPoints: []string{
								"tls",
							},
							Service:    "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr",
							Rule:       `HostSNI("foo.example.com")`,
							RuleSyntax: "default",
							Priority:   15,
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.expected == nil {
				return
			}

			k8sObjects, gwObjects := readResources(t, test.paths)

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints:       test.entryPoints,
				NativeLBByDefault: test.nativeLB,
				client:            client,
			}

			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			assert.Equal(t, test.expected, conf)
		})
	}
}

func TestLoadMixedRoutes(t *testing.T) {
	testCases := []struct {
		desc         string
		ingressClass string
		paths        []string
		expected     *dynamic.Configuration
		entryPoints  map[string]Entrypoint
	}{
		{
			desc: "Empty",
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty caused by unsupported listener.Protocol",
			paths: []string{"services.yml", "mixed/with_bad_listener_protocol.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":9080"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc: "Empty caused by unsupported listener.Route.Kind",
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":9080"},
			},
			paths: []string{"services.yml", "mixed/with_bad_listener_route_kind.yml"},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty caused by listener.Protocol does not support listener.Route.Kind",
			paths: []string{"services.yml", "mixed/with_incompatible_protocol_and_route_kind.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":9080"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Mixed routes",
			paths: []string{"services.yml", "mixed/simple.yml"},
			entryPoints: map[string]Entrypoint{
				"web":       {Address: ":9080"},
				"websecure": {Address: ":9443"},
				"tcp":       {Address: ":9000"},
				"tls-1":     {Address: ":10000"},
				"tls-2":     {Address: ":11000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-2-0-34e311881843bc2963c9": {
							EntryPoints: []string{"tls-2"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-2-0-34e311881843bc2963c9-wrr",
							Priority:    24,
							Rule:        `HostSNI("pass.tls.foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-2-0-34e311881843bc2963c9-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-2-0-34e311881843bc2963c9-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-2-0-34e311881843bc2963c9-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-f20abd7604e9c4caa5f3": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-f20abd7604e9c4caa5f3-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-c319354d6e97e7ccf839": {
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-c319354d6e97e7ccf839-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-f20abd7604e9c4caa5f3-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-f20abd7604e9c4caa5f3-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-c319354d6e97e7ccf839-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-c319354d6e97e7ccf839-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-f20abd7604e9c4caa5f3-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-c319354d6e97e7ccf839-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "Empty caused by mixed routes with multiple listeners using same hostname, port and protocol",
			paths: []string{"services.yml", "mixed/with_multiple_listeners_using_same_hostname_port_protocol.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":9080"},
				"tcp": {Address: ":9000"},
				"tls": {Address: ":9001"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Mixed routes with Same namespace selector",
			paths: []string{"services.yml", "mixed/with_namespace_same.yml"},
			entryPoints: map[string]Entrypoint{
				"web":       {Address: ":9080"},
				"websecure": {Address: ":9443"},
				"tcp":       {Address: ":9000"},
				"tls-1":     {Address: ":10000"},
				"tls-2":     {Address: ":11000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8": {
							EntryPoints: []string{"tls-2"},
							Service:     "tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8-wrr",
							Priority:    24,
							Rule:        `HostSNI("pass.tls.foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432": {
							Service:    "httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9": {
							Service:    "httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "Mixed routes with All namespace selector",
			paths: []string{"services.yml", "mixed/with_namespace_all.yml"},
			entryPoints: map[string]Entrypoint{
				"web":       {Address: ":9080"},
				"websecure": {Address: ":9443"},
				"tcp":       {Address: ":9000"},
				"tls-1":     {Address: ":10000"},
				"tls-2":     {Address: ":11000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8": {
							EntryPoints: []string{"tls-2"},
							Service:     "tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8-wrr",
							Priority:    24,
							Rule:        `HostSNI("pass.tls.foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
						"tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-default-gw-default-my-gateway-ep-tls-2-0-30cc8c9fb5be933316f8-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
						"tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7-svc-bar-whoamitcp-bar-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.13:9000",
									},
									{
										Address: "10.10.0.14:9000",
									},
								},
							},
						},
						"tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7-svc-bar-whoamitcp-bar-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755": {
							Service:    "httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13": {
							Service:    "httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755-svc-bar-whoami-bar-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.11:80",
									},
									{
										URL: "http://10.10.0.12:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13-svc-bar-whoami-bar-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.11:80",
									},
									{
										URL: "http://10.10.0.12:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755-svc-bar-whoami-bar-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13-svc-bar-whoami-bar-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "Mixed routes with Selector Route Binding",
			paths: []string{"services.yml", "mixed/with_namespace_selector.yml"},
			entryPoints: map[string]Entrypoint{
				"web":       {Address: ":9080"},
				"websecure": {Address: ":9443"},
				"tcp":       {Address: ":9000"},
				"tls-1":     {Address: ":10000"},
				"tls-2":     {Address: ":11000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
						"tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-2-0-2a95ae92bba77504addc": {
							EntryPoints: []string{"tls-2"},
							Service:     "tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-2-0-2a95ae92bba77504addc-wrr",
							Priority:    24,
							Rule:        `HostSNI("pass.tls.foo.example.com")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7-svc-bar-whoamitcp-bar-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.13:9000",
									},
									{
										Address: "10.10.0.14:9000",
									},
								},
							},
						},
						"tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-2-0-2a95ae92bba77504addc-svc-bar-whoamitcp-bar-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.13:9000",
									},
									{
										Address: "10.10.0.14:9000",
									},
								},
							},
						},
						"tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-bar-tcp-app-bar-gw-default-my-gateway-ep-tcp-0-c5d9e24bf010118167e7-svc-bar-whoamitcp-bar-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-2-0-2a95ae92bba77504addc-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-bar-tls-app-bar-gw-default-my-gateway-ep-tls-2-0-2a95ae92bba77504addc-svc-bar-whoamitcp-bar-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755": {
							Service:    "httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13": {
							Service:    "httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755-svc-bar-whoami-bar-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.11:80",
									},
									{
										URL: "http://10.10.0.12:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13-svc-bar-whoami-bar-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.11:80",
									},
									{
										URL: "http://10.10.0.12:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-bar-http-app-bar-gw-default-my-gateway-ep-web-0-cf99cd1374e6056a6755-svc-bar-whoami-bar-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-bar-http-app-bar-gw-default-my-gateway-ep-websecure-0-135c6826c1baedf74c13-svc-bar-whoami-bar-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "Mixed routes with core group",
			paths: []string{"services.yml", "mixed/with_core_group.yml"},
			entryPoints: map[string]Entrypoint{
				"web":       {Address: ":9080"},
				"websecure": {Address: ":9443"},
				"tcp":       {Address: ":9000"},
				"tls":       {Address: ":10000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-default-gw-default-my-gateway-ep-tcp-0-df24b7f0f1d9b443e55b-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{"websecure"},
							Rule:        `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432": {
							Service:    "httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9": {
							Service:    "httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9-wrr",
							Rule:       `PathPrefix("/")`,
							Priority:   2,
							RuleSyntax: "default",
							ParentRefs: []string{"listener-websecure-https-439feda8e3f1dbdca529"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-web-0-3b55179f1d48675f1432-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
						"httproute-default-http-app-default-gw-default-my-gateway-ep-websecure-0-4a26538fbaaf60dc52c9-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": defaultListenerTLSOptions(),
					},
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
				},
			},
		},
		{
			desc:  "GRPCRoute created first wins the conflicting hostname over the HTTPRoute",
			paths: []string{"services.yml", "mixed/with_grpcroute_priority_over_httproute.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":9080"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"grpcroute-default-grpc-app-priority-gw-default-my-gateway-ep-web-0-46e5641ba0ccfa2eff8a": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "grpcroute-default-grpc-app-priority-gw-default-my-gateway-ep-web-0-46e5641ba0ccfa2eff8a-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   22,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"grpcroute-default-grpc-app-priority-gw-default-my-gateway-ep-web-0-46e5641ba0ccfa2eff8a-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "grpcroute-default-grpc-app-priority-gw-default-my-gateway-ep-web-0-46e5641ba0ccfa2eff8a-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"grpcroute-default-grpc-app-priority-gw-default-my-gateway-ep-web-0-46e5641ba0ccfa2eff8a-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "h2c://10.10.0.1:80"},
									{URL: "h2c://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "HTTPRoute created first wins the conflicting hostname over the GRPCRoute",
			paths: []string{"services.yml", "mixed/with_httproute_priority_over_grpcroute.yml"},
			entryPoints: map[string]Entrypoint{
				"web": {Address: ":9080"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-web-http-5b3226d4ebd42b7200f1": {
							EntryPoints: []string{"web"},
							Rule:        `Host("*")`,
						},
						"httproute-default-http-app-conflict-gw-default-my-gateway-ep-web-0-3fcf93756cef1b546fd2": {
							ParentRefs: []string{"listener-web-http-5b3226d4ebd42b7200f1"},
							Service:    "httproute-default-http-app-conflict-gw-default-my-gateway-ep-web-0-3fcf93756cef1b546fd2-wrr",
							Rule:       `Host("foo.com") && PathPrefix("/")`,
							Priority:   9,
							RuleSyntax: "default",
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-conflict-gw-default-my-gateway-ep-web-0-3fcf93756cef1b546fd2-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-conflict-gw-default-my-gateway-ep-web-0-3fcf93756cef1b546fd2-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
						"httproute-default-http-app-conflict-gw-default-my-gateway-ep-web-0-3fcf93756cef1b546fd2-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Strategy: dynamic.BalancerStrategyWRR,
								Servers: []dynamic.Server{
									{URL: "http://10.10.0.1:80"},
									{URL: "http://10.10.0.2:80"},
								},
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100 * time.Millisecond),
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "TLSRoute and TCPRoute with the same rule on the same entry point",
			paths: []string{"services.yml", "mixed/with_tlsroute_and_tcproute_on_same_entrypoint.yml"},
			entryPoints: map[string]Entrypoint{
				"tcp": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-gw-default-my-tls-gateway-ep-tcp-0-06fcfb4c19c97feaa646": {
							EntryPoints: []string{"tcp"},
							Service:     "tlsroute-default-tls-app-gw-default-my-tls-gateway-ep-tcp-0-06fcfb4c19c97feaa646-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
							TLS: &dynamic.RouterTCPTLSConfig{
								Passthrough: true,
							},
						},
						"tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-0-ada211f39fad8759e161": {
							EntryPoints: []string{"tcp"},
							Service:     "tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-0-ada211f39fad8759e161-wrr",
							Rule:        `HostSNI("*")`,
							RuleSyntax:  "default",
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-gw-default-my-tls-gateway-ep-tcp-0-06fcfb4c19c97feaa646-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tlsroute-default-tls-app-gw-default-my-tls-gateway-ep-tcp-0-06fcfb4c19c97feaa646-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tlsroute-default-tls-app-gw-default-my-tls-gateway-ep-tcp-0-06fcfb4c19c97feaa646-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{Address: "10.10.0.9:9000"},
									{Address: "10.10.0.10:9000"},
								},
							},
						},
						"tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-0-ada211f39fad8759e161-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{
									{
										Name:   "tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-0-ada211f39fad8759e161-svc-default-whoamitcp-0",
										Weight: new(1),
									},
								},
							},
						},
						"tcproute-default-tcp-app-gw-default-my-tcp-gateway-ep-tcp-0-ada211f39fad8759e161-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{Address: "10.10.0.9:9000"},
									{Address: "10.10.0.10:9000"},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "ListenerSet listener taking the hostname of a Gateway listener is not loaded",
			paths: []string{"services.yml", "mixed/with_listenerset_conflicts.yml"},
			entryPoints: map[string]Entrypoint{
				"web":       {Address: ":80"},
				"websecure": {Address: ":443"},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-ba3a418469df6d6c8b5a": {
							Service: "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-ba3a418469df6d6c8b5a-wrr",
							Rule:    `Host("baz.example.com") && Path("/bar")`,
							ParentRefs: []string{
								"listener-websecure-https-baz-example-com-dd68422a9ac9bfc6c47e",
							},
							RuleSyntax: "default",
							Priority:   100016,
						},
						"listener-websecure-https-baz-example-com-dd68422a9ac9bfc6c47e": {
							EntryPoints: []string{
								"websecure",
							},
							Rule: `Host("baz.example.com")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-baz-example-com-dd68422a9ac9bfc6c47e",
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-ba3a418469df6d6c8b5a-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								Strategy:       dynamic.BalancerStrategy("wrr"),
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100000000),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-ba3a418469df6d6c8b5a-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-ba3a418469df6d6c8b5a-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					Middlewares:       map[string]*dynamic.Middleware{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Services:          map[string]*dynamic.TCPService{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-baz-example-com-dd68422a9ac9bfc6c47e": {
							CipherSuites: []string{
								"TLS_AES_128_GCM_SHA256",
								"TLS_AES_256_GCM_SHA384",
								"TLS_CHACHA20_POLY1305_SHA256",
								"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
								"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
								"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
								"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
								"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
								"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
								"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
								"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
								"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
								"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
							},
							ALPNProtocols: []string{
								"h2",
								"http/1.1",
								"acme-tls/1",
							},
						},
					},
				},
			},
		},
		{
			desc:  "Older ListenerSet wins the listener conflict with its sibling",
			paths: []string{"services.yml", "mixed/with_listenerset_precedence.yml"},
			entryPoints: map[string]Entrypoint{
				"web":  {Address: ":80"},
				"web2": {Address: ":8080"},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"httproute-default-http-app-older-gw-default-my-gateway-ep-web2-0-4541b5743d681903cbd4": {
							Service: "httproute-default-http-app-older-gw-default-my-gateway-ep-web2-0-4541b5743d681903cbd4-wrr",
							Rule:    `Host("older.example.com") && Path("/bar")`,
							ParentRefs: []string{
								"listener-web2-http-4b299c63a1fd0bffdf3d",
							},
							RuleSyntax: "default",
							Priority:   100018,
						},
						"listener-web2-http-4b299c63a1fd0bffdf3d": {
							EntryPoints: []string{
								"web2",
							},
							Rule: `Host("*")`,
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-older-gw-default-my-gateway-ep-web2-0-4541b5743d681903cbd4-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								Strategy:       dynamic.BalancerStrategy("wrr"),
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100000000),
								},
							},
						},
						"httproute-default-http-app-older-gw-default-my-gateway-ep-web2-0-4541b5743d681903cbd4-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-older-gw-default-my-gateway-ep-web2-0-4541b5743d681903cbd4-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					Middlewares:       map[string]*dynamic.Middleware{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Services:          map[string]*dynamic.TCPService{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because the ListenerSet is not allowed by the Gateway",
			paths: []string{"services.yml", "mixed/with_listenerset_not_allowed.yml"},
			entryPoints: map[string]Entrypoint{
				"web":  {Address: ":80"},
				"web2": {Address: ":8080"},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Services:          map[string]*dynamic.Service{},
					Middlewares:       map[string]*dynamic.Middleware{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Services:          map[string]*dynamic.TCPService{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			// The Gateway declares no valid listener of its own,
			// so it is accepted through the listener its ListenerSet delegates, which serves the route.
			desc:  "ListenerSet serving a Gateway without a valid listener of its own",
			paths: []string{"services.yml", "mixed/with_listenerset_delegation.yml"},
			entryPoints: map[string]Entrypoint{
				"websecure": {Address: ":443"},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282": {
							Service: "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-wrr",
							Rule:    `Host("foo.example.com") && Path("/bar")`,
							ParentRefs: []string{
								"listener-websecure-https-439feda8e3f1dbdca529",
							},
							RuleSyntax: "default",
							Priority:   100016,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{
								"websecure",
							},
							Rule: `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								Strategy:       dynamic.BalancerStrategy("wrr"),
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100000000),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					Middlewares:       map[string]*dynamic.Middleware{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Services:          map[string]*dynamic.TCPService{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							CipherSuites: []string{
								"TLS_AES_128_GCM_SHA256",
								"TLS_AES_256_GCM_SHA384",
								"TLS_CHACHA20_POLY1305_SHA256",
								"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
								"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
								"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
								"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
								"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
								"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
								"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
								"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
								"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
								"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
							},
							ALPNProtocols: []string{
								"h2",
								"http/1.1",
								"acme-tls/1",
							},
						},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.expected == nil {
				return
			}

			k8sObjects, gwObjects := readResources(t, test.paths)

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints: test.entryPoints,
				client:      client,
			}

			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			assert.Equal(t, test.expected, conf)
		})
	}
}

func TestLoadRoutesWithReferenceGrants(t *testing.T) {
	testCases := []struct {
		desc         string
		ingressClass string
		paths        []string
		expected     *dynamic.Configuration
		entryPoints  map[string]Entrypoint
	}{
		{
			desc: "Empty",
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because ReferenceGrant for Secret is missing",
			paths: []string{"services.yml", "referencegrant/for_secret_missing.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because ReferenceGrant spec.from does not match secret",
			paths: []string{"services.yml", "referencegrant/for_secret_not_matching_from.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because ReferenceGrant spec.to does not match secret",
			paths: []string{"services.yml", "referencegrant/for_secret_not_matching_to.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "For Secret",
			paths: []string{"services.yml", "referencegrant/for_secret.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers: map[string]*dynamic.TCPRouter{
						"deny-unknown-host": {
							Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
							Priority: 1,
							Service:  "deny-unknown-host",
							TLS:      &dynamic.RouterTCPTLSConfig{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174": {
							EntryPoints: []string{"tls"},
							Service:     "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr",
							Rule:        `HostSNI("foo.example.com")`,
							RuleSyntax:  "default",
							Priority:    15,
							TLS:         &dynamic.RouterTCPTLSConfig{},
						},
					},
					Middlewares: map[string]*dynamic.TCPMiddleware{},
					Services: map[string]*dynamic.TCPService{
						"deny-unknown-host": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-wrr": {
							Weighted: &dynamic.TCPWeightedRoundRobin{
								Services: []dynamic.TCPWRRService{{
									Name:   "tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0",
									Weight: new(1),
								}},
							},
						},
						"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-85354ef9915a52028174-svc-default-whoamitcp-0": {
							LoadBalancer: &dynamic.TCPServersLoadBalancer{
								Servers: []dynamic.TCPServer{
									{
										Address: "10.10.0.9:9000",
									},
									{
										Address: "10.10.0.10:9000",
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because ReferenceGrant for Service is missing",
			paths: []string{"services.yml", "referencegrant/for_secret_missing.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because ReferenceGrant spec.from does not match service",
			paths: []string{"services.yml", "referencegrant/for_service_not_matching_from.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "Empty because ReferenceGrant spec.to does not match service",
			paths: []string{"services.yml", "referencegrant/for_service_not_matching_to.yml"},
			entryPoints: map[string]Entrypoint{
				"tls": {Address: ":9000"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Middlewares:       map[string]*dynamic.Middleware{},
					Services:          map[string]*dynamic.Service{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "For Service",
			paths: []string{"services.yml", "referencegrant/for_service.yml"},
			entryPoints: map[string]Entrypoint{
				"http": {Address: ":80"},
			},
			expected: &dynamic.Configuration{
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					Services:          map[string]*dynamic.TCPService{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"listener-http-http-foo-example-com-34297ca2d1326bfcdaad": {
							EntryPoints: []string{"http"},
							Rule:        `Host("foo.example.com")`,
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-http-0-f6b84730f282bab89081": {
							Rule:       `Host("foo.example.com") && PathPrefix("/")`,
							Service:    "httproute-default-http-app-1-gw-default-my-gateway-ep-http-0-f6b84730f282bab89081-wrr",
							RuleSyntax: "default",
							Priority:   17,
							ParentRefs: []string{"listener-http-http-foo-example-com-34297ca2d1326bfcdaad"},
						},
					},
					Middlewares: map[string]*dynamic.Middleware{},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-http-0-f6b84730f282bab89081-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-http-0-f6b84730f282bab89081-svc-default-whoami-bar-0",
										Weight: new(1),
										Status: new(500),
									},
								},
							},
						},
					},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
		{
			desc:  "For Secret from ListenerSet",
			paths: []string{"services.yml", "referencegrant/for_secret_from_listenerset.yml"},
			entryPoints: map[string]Entrypoint{
				"web":       {Address: ":80"},
				"websecure": {Address: ":443"},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers: map[string]*dynamic.Router{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282": {
							Service: "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-wrr",
							Rule:    `Host("foo.example.com") && Path("/bar")`,
							ParentRefs: []string{
								"listener-websecure-https-439feda8e3f1dbdca529",
							},
							RuleSyntax: "default",
							Priority:   100016,
						},
						"listener-websecure-https-439feda8e3f1dbdca529": {
							EntryPoints: []string{
								"websecure",
							},
							Rule: `Host("*")`,
							TLS: &dynamic.RouterTLSConfig{
								Options: "listener-websecure-https-439feda8e3f1dbdca529",
							},
						},
					},
					Services: map[string]*dynamic.Service{
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-svc-default-whoami-0": {
							LoadBalancer: &dynamic.ServersLoadBalancer{
								Servers: []dynamic.Server{
									{
										URL: "http://10.10.0.1:80",
									},
									{
										URL: "http://10.10.0.2:80",
									},
								},
								Strategy:       dynamic.BalancerStrategy("wrr"),
								PassHostHeader: new(true),
								ResponseForwarding: &dynamic.ResponseForwarding{
									FlushInterval: ptypes.Duration(100000000),
								},
							},
						},
						"httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-wrr": {
							Weighted: &dynamic.WeightedRoundRobin{
								Services: []dynamic.WRRService{
									{
										Name:   "httproute-default-http-app-1-gw-default-my-gateway-ep-websecure-0-3c8b64d308e92f9c9282-svc-default-whoami-0",
										Weight: new(1),
									},
								},
							},
						},
					},
					Middlewares:       map[string]*dynamic.Middleware{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Services:          map[string]*dynamic.TCPService{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TLS: &dynamic.TLSConfiguration{
					Certificates: []*tls.CertAndStores{
						{
							CertFile: types.FileOrContent(listenerCert),
							KeyFile:  types.FileOrContent(listenerKey),
						},
					},
					Options: map[string]tls.Options{
						"listener-websecure-https-439feda8e3f1dbdca529": {
							CipherSuites: []string{
								"TLS_AES_128_GCM_SHA256",
								"TLS_AES_256_GCM_SHA384",
								"TLS_CHACHA20_POLY1305_SHA256",
								"TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA",
								"TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA",
								"TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
								"TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
								"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
								"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
								"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
								"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
								"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256",
								"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256",
							},
							ALPNProtocols: []string{
								"h2",
								"http/1.1",
								"acme-tls/1",
							},
						},
					},
				},
			},
		},
		{
			desc:  "Empty because ReferenceGrant spec.from does not match ListenerSet",
			paths: []string{"services.yml", "referencegrant/for_secret_not_matching_from_listenerset.yml"},
			entryPoints: map[string]Entrypoint{
				"web":       {Address: ":80"},
				"websecure": {Address: ":443"},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{
					Routers:           map[string]*dynamic.Router{},
					Services:          map[string]*dynamic.Service{},
					Middlewares:       map[string]*dynamic.Middleware{},
					ServersTransports: map[string]*dynamic.ServersTransport{},
				},
				TCP: &dynamic.TCPConfiguration{
					Routers:           map[string]*dynamic.TCPRouter{},
					Services:          map[string]*dynamic.TCPService{},
					Middlewares:       map[string]*dynamic.TCPMiddleware{},
					ServersTransports: map[string]*dynamic.TCPServersTransport{},
				},
				UDP: &dynamic.UDPConfiguration{
					Routers:  map[string]*dynamic.UDPRouter{},
					Services: map[string]*dynamic.UDPService{},
				},
				TLS: &dynamic.TLSConfiguration{
					Options: map[string]tls.Options{},
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			if test.expected == nil {
				return
			}

			k8sObjects, gwObjects := readResources(t, test.paths)

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints: test.entryPoints,
				client:      client,
			}

			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			assert.Equal(t, test.expected, conf)
		})
	}
}

func Test_matchingGatewayListener(t *testing.T) {
	testCases := []struct {
		desc           string
		gateways       []gatewayWithListeners
		parentRefs     []gatev1.ParentReference
		routeNamespace string
		want           []gatewayListenersForParentRef
	}{
		{
			desc: "Unsupported group",
			gateways: []gatewayWithListeners{{
				gateway:   &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{{SectionName: "foo"}},
			}},
			parentRefs: []gatev1.ParentReference{{
				Group: new(gatev1.Group("foo")),
			}},
		},
		{
			desc: "Unsupported kind",
			gateways: []gatewayWithListeners{{
				gateway:   &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{{SectionName: "foo"}},
			}},
			parentRefs: []gatev1.ParentReference{{
				Group: new(gatev1.Group(gatev1.GroupName)),
				Kind:  new(gatev1.Kind("foo")),
			}},
		},
		{
			desc: "Namespace does not match the listener",
			gateways: []gatewayWithListeners{{
				gateway:   &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{{SectionName: "foo"}},
			}},
			parentRefs: []gatev1.ParentReference{{
				Namespace: new(gatev1.Namespace("foo")),
				Group:     new(gatev1.Group(gatev1.GroupName)),
				Kind:      new(gatev1.Kind("Gateway")),
			}},
		},
		{
			desc: "Route namespace defaulting does not match the listener",
			gateways: []gatewayWithListeners{{
				gateway:   &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{{SectionName: "foo"}},
			}},
			routeNamespace: "foo",
			parentRefs: []gatev1.ParentReference{{
				Group: new(gatev1.Group(gatev1.GroupName)),
				Kind:  new(gatev1.Kind("Gateway")),
			}},
		},
		{
			desc: "Name does not match the listener",
			gateways: []gatewayWithListeners{{
				gateway:   &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{{}},
			}},
			parentRefs: []gatev1.ParentReference{{
				Namespace: new(gatev1.Namespace("default")),
				Name:      "foo",
				Group:     new(gatev1.Group(gatev1.GroupName)),
				Kind:      new(gatev1.Kind("Gateway")),
			}},
		},
		{
			desc: "Match",
			gateways: []gatewayWithListeners{{
				gateway:   &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{{}},
			}},
			parentRefs: []gatev1.ParentReference{{
				Name:      "gateway",
				Namespace: new(gatev1.Namespace("default")),
				Group:     new(gatev1.Group(gatev1.GroupName)),
				Kind:      new(gatev1.Kind("Gateway")),
			}},
			want: []gatewayListenersForParentRef{{
				ParentRef: gatev1.ParentReference{
					Name:      "gateway",
					Namespace: new(gatev1.Namespace("default")),
					Group:     new(gatev1.Group(gatev1.GroupName)),
					Kind:      new(gatev1.Kind("Gateway")),
				},
				GatewayName:      "gateway",
				GatewayNamespace: "default",
				Listeners:        []gatewayListener{{}},
			}},
		},
		{
			desc: "Match with route namespace defaulting",
			gateways: []gatewayWithListeners{{
				gateway:   &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{{}},
			}},
			routeNamespace: "default",
			parentRefs: []gatev1.ParentReference{{
				Name:  "gateway",
				Group: new(gatev1.Group(gatev1.GroupName)),
				Kind:  new(gatev1.Kind("Gateway")),
			}},
			want: []gatewayListenersForParentRef{{
				ParentRef: gatev1.ParentReference{
					Name:  "gateway",
					Group: new(gatev1.Group(gatev1.GroupName)),
					Kind:  new(gatev1.Kind("Gateway")),
				},
				GatewayName:      "gateway",
				GatewayNamespace: "default",
				Listeners:        []gatewayListener{{}},
			}},
		},
		{
			// A Route may reference Gateways owned by other controllers; those
			// parentRefs must not be associated to (and thus get status written by)
			// this controller.
			desc: "Only parentRefs targeting our Gateways are returned",
			gateways: []gatewayWithListeners{{
				gateway:   &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{{}},
			}},
			parentRefs: []gatev1.ParentReference{
				{
					Name:      "gateway",
					Namespace: new(gatev1.Namespace("default")),
					Group:     new(gatev1.Group(gatev1.GroupName)),
					Kind:      new(gatev1.Kind("Gateway")),
				},
				{
					Name:      "other-gateway",
					Namespace: new(gatev1.Namespace("default")),
					Group:     new(gatev1.Group(gatev1.GroupName)),
					Kind:      new(gatev1.Kind("Gateway")),
				},
			},
			want: []gatewayListenersForParentRef{{
				ParentRef: gatev1.ParentReference{
					Name:      "gateway",
					Namespace: new(gatev1.Namespace("default")),
					Group:     new(gatev1.Group(gatev1.GroupName)),
					Kind:      new(gatev1.Kind("Gateway")),
				},
				GatewayName:      "gateway",
				GatewayNamespace: "default",
				Listeners:        []gatewayListener{{}},
			}},
		},
		{
			desc: "ParentRef is associated to all the listeners of its Gateway",
			gateways: []gatewayWithListeners{{
				gateway: &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{
					{SectionName: "web"},
					{SectionName: "websecure"},
				},
			}},
			parentRefs: []gatev1.ParentReference{{
				Name:      "gateway",
				Namespace: new(gatev1.Namespace("default")),
				Group:     new(gatev1.Group(gatev1.GroupName)),
				Kind:      new(gatev1.Kind("Gateway")),
			}},
			want: []gatewayListenersForParentRef{{
				ParentRef: gatev1.ParentReference{
					Name:      "gateway",
					Namespace: new(gatev1.Namespace("default")),
					Group:     new(gatev1.Group(gatev1.GroupName)),
					Kind:      new(gatev1.Kind("Gateway")),
				},
				GatewayName:      "gateway",
				GatewayNamespace: "default",
				Listeners: []gatewayListener{
					{SectionName: "web"},
					{SectionName: "websecure"},
				},
			}},
		},
		{
			// A parentRef must only be associated to the listeners of the Gateway it
			// targets, even when several of our Gateways are known.
			desc: "ParentRef is only associated to the referenced Gateway",
			gateways: []gatewayWithListeners{
				{
					gateway:   &gatev1.Gateway{Name: "gateway-a", Namespace: "default"},
					listeners: []gatewayListener{{SectionName: "web"}},
				},
				{
					gateway:   &gatev1.Gateway{Name: "gateway-b", Namespace: "default"},
					listeners: []gatewayListener{{SectionName: "web"}},
				},
			},
			parentRefs: []gatev1.ParentReference{{
				Name:      "gateway-a",
				Namespace: new(gatev1.Namespace("default")),
				Group:     new(gatev1.Group(gatev1.GroupName)),
				Kind:      new(gatev1.Kind("Gateway")),
			}},
			want: []gatewayListenersForParentRef{{
				ParentRef: gatev1.ParentReference{
					Name:      "gateway-a",
					Namespace: new(gatev1.Namespace("default")),
					Group:     new(gatev1.Group(gatev1.GroupName)),
					Kind:      new(gatev1.Kind("Gateway")),
				},
				GatewayName:      "gateway-a",
				GatewayNamespace: "default",
				Listeners:        []gatewayListener{{SectionName: "web"}},
			}},
		},
		{
			// A parentRef with a non-matching SectionName or Port still targets the
			// Gateway: all its listeners are returned so the route loader can report
			// ResolvedRefs even though the route attaches to none of them.
			desc: "ParentRef with a non-matching port still returns the Gateway listeners",
			gateways: []gatewayWithListeners{{
				gateway:   &gatev1.Gateway{Name: "gateway", Namespace: "default"},
				listeners: []gatewayListener{{SectionName: "web", Port: 80}},
			}},
			parentRefs: []gatev1.ParentReference{{
				Name:      "gateway",
				Namespace: new(gatev1.Namespace("default")),
				Group:     new(gatev1.Group(gatev1.GroupName)),
				Kind:      new(gatev1.Kind("Gateway")),
				Port:      new(gatev1.PortNumber(8080)),
			}},
			want: []gatewayListenersForParentRef{{
				ParentRef: gatev1.ParentReference{
					Name:      "gateway",
					Namespace: new(gatev1.Namespace("default")),
					Group:     new(gatev1.Group(gatev1.GroupName)),
					Kind:      new(gatev1.Kind("Gateway")),
					Port:      new(gatev1.PortNumber(8080)),
				},
				GatewayName:      "gateway",
				GatewayNamespace: "default",
				Listeners:        []gatewayListener{{SectionName: "web", Port: 80}},
			}},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			matches := matchingGatewayListenersForParentRef(test.gateways, test.routeNamespace, test.parentRefs)
			assert.Equal(t, test.want, matches)
		})
	}
}

// Test_loadRoutes_multipleGatewaysParentRefs ensures that, when a Route
// references several Gateways (including one owned by another controller), Traefik
// only manages the parents it is responsible for: it writes a single RouteParentStatus
// for its own parentRef, writes nothing for a Route that targets only a foreign
// Gateway, and counts each attached Route only once in the listener AttachedRoutes.
func Test_loadRoutes_multipleGatewaysParentRefs(t *testing.T) {
	testCases := []struct {
		desc           string
		path           string
		entryPoints    map[string]Entrypoint
		gatewayName    string
		route1, route2 string
		getParents     func(t *testing.T, gwClient *gatefake.Clientset, name string) []gatev1.RouteParentStatus
	}{
		{
			desc:        "HTTPRoute",
			path:        "httproute/with_multiple_gateways_parent_refs.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}},
			gatewayName: "my-gateway",
			route1:      "http-app-1",
			route2:      "http-app-2",
			getParents: func(t *testing.T, gwClient *gatefake.Clientset, name string) []gatev1.RouteParentStatus {
				t.Helper()
				route, err := gwClient.GatewayV1().HTTPRoutes("default").Get(t.Context(), name, metav1.GetOptions{})
				require.NoError(t, err)
				return route.Status.Parents
			},
		},
		{
			desc:        "GRPCRoute",
			path:        "grpcroute/with_multiple_gateways_parent_refs.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}},
			gatewayName: "my-gateway",
			route1:      "grpc-app-1",
			route2:      "grpc-app-2",
			getParents: func(t *testing.T, gwClient *gatefake.Clientset, name string) []gatev1.RouteParentStatus {
				t.Helper()
				route, err := gwClient.GatewayV1().GRPCRoutes("default").Get(t.Context(), name, metav1.GetOptions{})
				require.NoError(t, err)
				return route.Status.Parents
			},
		},
		{
			desc:        "TCPRoute",
			path:        "tcproute/with_multiple_gateways_parent_refs.yml",
			entryPoints: map[string]Entrypoint{"tcp": {Address: ":9000"}},
			gatewayName: "my-tcp-gateway",
			route1:      "tcp-app-1",
			route2:      "tcp-app-2",
			getParents: func(t *testing.T, gwClient *gatefake.Clientset, name string) []gatev1.RouteParentStatus {
				t.Helper()
				route, err := gwClient.GatewayV1().TCPRoutes("default").Get(t.Context(), name, metav1.GetOptions{})
				require.NoError(t, err)
				return route.Status.Parents
			},
		},
		{
			desc:        "TLSRoute",
			path:        "tlsroute/with_multiple_gateways_parent_refs.yml",
			entryPoints: map[string]Entrypoint{"tls": {Address: ":9000"}},
			gatewayName: "my-tls-gateway",
			route1:      "tls-app-1",
			route2:      "tls-app-2",
			getParents: func(t *testing.T, gwClient *gatefake.Clientset, name string) []gatev1.RouteParentStatus {
				t.Helper()
				route, err := gwClient.GatewayV1().TLSRoutes("default").Get(t.Context(), name, metav1.GetOptions{})
				require.NoError(t, err)
				return route.Status.Parents
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			k8sObjects, gwObjects := readResources(t, []string{"services.yml", test.path})

			gwClient := newGatewaySimpleClientSet(t, gwObjects...)
			client := newClientImpl(kubefake.NewClientset(k8sObjects...), gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)
			<-eventCh

			p := Provider{
				EntryPoints: test.entryPoints,
				client:      client,
			}

			_, statusReport, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			statusReport.Flush(t.Context(), client)

			// route1 references our Gateway and a foreign one: only our parentRef must
			// get a status entry, written under our controllerName and accepted.
			parents1 := test.getParents(t, gwClient, test.route1)
			require.Len(t, parents1, 1)
			assert.Equal(t, gatev1.GatewayController(controllerName), parents1[0].ControllerName)
			assert.Equal(t, gatev1.ObjectName(test.gatewayName), parents1[0].ParentRef.Name)

			acceptedCondition := findCondition(parents1[0].Conditions, gatev1.RouteConditionAccepted)
			require.NotNil(t, acceptedCondition)
			assert.Equal(t, metav1.ConditionTrue, acceptedCondition.Status)
			assert.Equal(t, string(gatev1.RouteReasonAccepted), acceptedCondition.Reason)

			// route2 references only the foreign Gateway: Traefik must not write any status.
			assert.Empty(t, test.getParents(t, gwClient, test.route2))

			// The listener must count route1 once, ignoring the foreign parentRef and
			// the foreign-only Route.
			gateway, err := gwClient.GatewayV1().Gateways("default").Get(t.Context(), test.gatewayName, metav1.GetOptions{})
			require.NoError(t, err)
			require.Len(t, gateway.Status.Listeners, 1)
			assert.Equal(t, int32(1), gateway.Status.Listeners[0].AttachedRoutes)
		})
	}
}

func Test_loadHTTPAndGRPCRoutes_hostnameConflict(t *testing.T) {
	testCases := []struct {
		desc            string
		path            string
		acceptedRoute   string
		acceptedIsHTTP  bool
		conflictedRoute string
	}{
		{
			desc:            "HTTPRoute created first",
			path:            "mixed/with_httproute_priority_over_grpcroute.yml",
			acceptedRoute:   "http-app-conflict",
			acceptedIsHTTP:  true,
			conflictedRoute: "grpc-app-conflict",
		},
		{
			desc:            "GRPCRoute created first",
			path:            "mixed/with_grpcroute_priority_over_httproute.yml",
			acceptedRoute:   "grpc-app-priority",
			conflictedRoute: "http-app-priority",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			k8sObjects, gwObjects := readResources(t, []string{"services.yml", test.path})

			gwClient := newGatewaySimpleClientSet(t, gwObjects...)
			client := newClientImpl(kubefake.NewClientset(k8sObjects...), gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)
			<-eventCh

			p := Provider{
				EntryPoints: map[string]Entrypoint{"web": {Address: ":9080"}},
				client:      client,
			}

			_, statusReport, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			statusReport.Flush(t.Context(), client)

			getParents := func(name string, isHTTP bool) []gatev1.RouteParentStatus {
				t.Helper()
				if isHTTP {
					route, err := gwClient.GatewayV1().HTTPRoutes("default").Get(t.Context(), name, metav1.GetOptions{})
					require.NoError(t, err)
					return route.Status.Parents
				}
				route, err := gwClient.GatewayV1().GRPCRoutes("default").Get(t.Context(), name, metav1.GetOptions{})
				require.NoError(t, err)
				return route.Status.Parents
			}

			acceptedParents := getParents(test.acceptedRoute, test.acceptedIsHTTP)
			require.Len(t, acceptedParents, 1)
			acceptedCondition := findCondition(acceptedParents[0].Conditions, gatev1.RouteConditionAccepted)
			require.NotNil(t, acceptedCondition)
			assert.Equal(t, metav1.ConditionTrue, acceptedCondition.Status)
			assert.Equal(t, string(gatev1.RouteReasonAccepted), acceptedCondition.Reason)

			conflictedParents := getParents(test.conflictedRoute, !test.acceptedIsHTTP)
			require.Len(t, conflictedParents, 1)
			conflictedCondition := findCondition(conflictedParents[0].Conditions, gatev1.RouteConditionAccepted)
			require.NotNil(t, conflictedCondition)
			assert.Equal(t, metav1.ConditionFalse, conflictedCondition.Status)
			assert.Equal(t, string(routeReasonHostnameConflict), conflictedCondition.Reason)

			gateway, err := gwClient.GatewayV1().Gateways("default").Get(t.Context(), "my-gateway", metav1.GetOptions{})
			require.NoError(t, err)
			require.Len(t, gateway.Status.Listeners, 1)
			assert.Equal(t, int32(1), gateway.Status.Listeners[0].AttachedRoutes)
		})
	}
}

func Test_attachedRoutes_conflicts(t *testing.T) {
	testCases := []struct {
		desc              string
		attachedKind      string
		attachedHostnames []gatev1.Hostname
		attachedListener  listenerRef
		kind              string
		hostnames         []gatev1.Hostname
		expected          bool
	}{
		{
			desc:              "same kind",
			attachedKind:      kindHTTPRoute,
			attachedHostnames: []gatev1.Hostname{"foo.com"},
			attachedListener: listenerRef{
				GatewayNamespace: "default",
				GatewayName:      "my-gateway",
				SectionName:      "web",
			},
			kind:      kindHTTPRoute,
			hostnames: []gatev1.Hostname{"foo.com"},
		},
		{
			desc:              "another listener",
			attachedKind:      kindHTTPRoute,
			attachedHostnames: []gatev1.Hostname{"foo.com"},
			attachedListener: listenerRef{
				GatewayNamespace: "default",
				GatewayName:      "my-gateway",
				SectionName:      "websecure",
			},
			kind:      kindGRPCRoute,
			hostnames: []gatev1.Hostname{"foo.com"},
		},
		{
			desc:              "disjoint hostnames",
			attachedKind:      kindHTTPRoute,
			attachedHostnames: []gatev1.Hostname{"foo.com"},
			attachedListener: listenerRef{
				GatewayNamespace: "default",
				GatewayName:      "my-gateway",
				SectionName:      "web",
			},
			kind:      kindGRPCRoute,
			hostnames: []gatev1.Hostname{"bar.com"},
		},
		{
			desc:              "same hostname",
			attachedKind:      kindHTTPRoute,
			attachedHostnames: []gatev1.Hostname{"foo.com"},
			attachedListener: listenerRef{
				GatewayNamespace: "default",
				GatewayName:      "my-gateway",
				SectionName:      "web",
			},
			kind:      kindGRPCRoute,
			hostnames: []gatev1.Hostname{"bar.com", "foo.com"},
			expected:  true,
		},
		{
			desc:              "attached wildcard hostname",
			attachedKind:      kindGRPCRoute,
			attachedHostnames: []gatev1.Hostname{"*.foo.com"},
			attachedListener: listenerRef{
				GatewayNamespace: "default",
				GatewayName:      "my-gateway",
				SectionName:      "web",
			},
			kind:      kindHTTPRoute,
			hostnames: []gatev1.Hostname{"bar.foo.com"},
			expected:  true,
		},
		{
			desc:              "wildcard hostname",
			attachedKind:      kindGRPCRoute,
			attachedHostnames: []gatev1.Hostname{"bar.foo.com"},
			attachedListener: listenerRef{
				GatewayNamespace: "default",
				GatewayName:      "my-gateway",
				SectionName:      "web",
			},
			kind:      kindHTTPRoute,
			hostnames: []gatev1.Hostname{"*.foo.com"},
			expected:  true,
		},
		{
			desc:              "attached without hostname",
			attachedKind:      kindHTTPRoute,
			attachedHostnames: nil,
			attachedListener: listenerRef{
				GatewayNamespace: "default",
				GatewayName:      "my-gateway",
				SectionName:      "web",
			},
			kind:      kindGRPCRoute,
			hostnames: []gatev1.Hostname{"foo.com"},
			expected:  true,
		},
		{
			desc:              "without hostname",
			attachedKind:      kindHTTPRoute,
			attachedHostnames: []gatev1.Hostname{"foo.com"},
			attachedListener: listenerRef{
				GatewayNamespace: "default",
				GatewayName:      "my-gateway",
				SectionName:      "web",
			},
			kind:      kindGRPCRoute,
			hostnames: nil,
			expected:  true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			attached := make(attachedRoutes)
			attached.Record(test.attachedListener, test.attachedKind, test.attachedHostnames)

			listener := listenerRef{GatewayNamespace: "default", GatewayName: "my-gateway", SectionName: "web"}
			assert.Equal(t, test.expected, attached.Conflicts(listener, test.kind, test.hostnames))
		})
	}
}

// Test_mergeRouteParentStatuses ensures the merge keeps the parentStatuses Traefik
// does not own: those written by another controller, and those written by another
// Traefik instance (our controllerName but a parentRef not targeting a managed
// Gateway), while replacing our stale statuses with the freshly computed ones.
func Test_mergeRouteParentStatuses(t *testing.T) {
	gateways := []gatewayWithListeners{{gateway: &gatev1.Gateway{Name: "my-gateway", Namespace: "default"}}}

	otherController := gatev1.RouteParentStatus{
		ControllerName: "example.com/other-controller",
		ParentRef:      gatev1.ParentReference{Namespace: new(gatev1.Namespace("default")), Name: "other-gateway"},
	}
	otherTraefik := gatev1.RouteParentStatus{
		ControllerName: controllerName,
		ParentRef:      gatev1.ParentReference{Namespace: new(gatev1.Namespace("default")), Name: "other-traefik-gateway"},
	}
	staleManaged := gatev1.RouteParentStatus{
		ControllerName: controllerName,
		ParentRef:      gatev1.ParentReference{Namespace: new(gatev1.Namespace("default")), Name: "my-gateway"},
	}
	// A parentRef without an explicit namespace defaults to the Route namespace,
	// and must still be recognized as targeting a managed Gateway.
	staleManagedDefaulted := gatev1.RouteParentStatus{
		ControllerName: controllerName,
		ParentRef:      gatev1.ParentReference{Name: "my-gateway"},
	}
	desired := gatev1.RouteParentStatus{
		ControllerName: controllerName,
		ParentRef:      gatev1.ParentReference{Namespace: new(gatev1.Namespace("default")), Name: "my-gateway"},
		Conditions: []metav1.Condition{{
			Type:   string(gatev1.RouteConditionAccepted),
			Status: metav1.ConditionTrue,
			Reason: string(gatev1.RouteReasonAccepted),
		}},
	}

	testCases := []struct {
		desc    string
		current []gatev1.RouteParentStatus
		want    []gatev1.RouteParentStatus
	}{
		{
			desc: "no current status",
			want: []gatev1.RouteParentStatus{desired},
		},
		{
			desc:    "keeps status owned by another controller",
			current: []gatev1.RouteParentStatus{otherController},
			want:    []gatev1.RouteParentStatus{desired, otherController},
		},
		{
			desc:    "keeps our status targeting a Gateway managed by another Traefik instance",
			current: []gatev1.RouteParentStatus{otherTraefik},
			want:    []gatev1.RouteParentStatus{desired, otherTraefik},
		},
		{
			desc:    "drops our stale status targeting a managed Gateway",
			current: []gatev1.RouteParentStatus{staleManaged},
			want:    []gatev1.RouteParentStatus{desired},
		},
		{
			desc:    "drops our stale status targeting a managed Gateway via namespace defaulting",
			current: []gatev1.RouteParentStatus{staleManagedDefaulted},
			want:    []gatev1.RouteParentStatus{desired},
		},
		{
			desc:    "mixed: keeps unmanaged, drops managed",
			current: []gatev1.RouteParentStatus{otherController, otherTraefik, staleManaged},
			want:    []gatev1.RouteParentStatus{desired, otherController, otherTraefik},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			got := mergeRouteParentStatuses("default", test.current, []gatev1.RouteParentStatus{desired}, gateways)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_managesGateway(t *testing.T) {
	testCases := []struct {
		desc     string
		gateways []NamespacedName
		gw       *gatev1.Gateway
		want     bool
	}{
		{
			desc: "no scoping manages everything",
			gw: &gatev1.Gateway{
				Namespace: "default",
				Name:      "foo",
			},
			want: true,
		},
		{
			desc:     "matching ref",
			gateways: []NamespacedName{{Namespace: "default", Name: "foo"}},
			gw: &gatev1.Gateway{
				Namespace: "default",
				Name:      "foo",
			},
			want: true,
		},
		{
			desc:     "different name",
			gateways: []NamespacedName{{Namespace: "default", Name: "foo"}},
			gw: &gatev1.Gateway{
				Namespace: "default",
				Name:      "bar",
			},
			want: false,
		},
		{
			desc:     "different namespace",
			gateways: []NamespacedName{{Namespace: "default", Name: "foo"}},
			gw: &gatev1.Gateway{
				Namespace: "other",
				Name:      "foo",
			},
			want: false,
		},
		{
			desc: "list matches first",
			gateways: []NamespacedName{
				{Namespace: "default", Name: "foo"},
				{Namespace: "other", Name: "bar"},
			},
			gw: &gatev1.Gateway{
				Namespace: "default",
				Name:      "foo",
			},
			want: true,
		},
		{
			desc: "list matches none",
			gateways: []NamespacedName{
				{Namespace: "default", Name: "foo"},
				{Namespace: "other", Name: "bar"},
			},
			gw: &gatev1.Gateway{
				Namespace: "default",
				Name:      "baz",
			},
			want: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			p := Provider{Gateways: test.gateways}
			assert.Equal(t, test.want, p.managesGateway(test.gw))
		})
	}
}

func Test_matchListener(t *testing.T) {
	testCases := []struct {
		desc       string
		gwListener gatewayListener
		parentRef  gatev1.ParentReference
		wantMatch  bool
	}{
		{
			desc: "Section do not match",
			gwListener: gatewayListener{
				SectionName: "foo",
				Port:        gatev1.PortNumber(80),
			},
			parentRef: gatev1.ParentReference{
				SectionName: new(gatev1.SectionName("bar")),
				Port:        new(gatev1.PortNumber(80)),
			},
		},
		{
			desc: "Section matches",
			gwListener: gatewayListener{
				SectionName: "foo",
				Port:        gatev1.PortNumber(80),
			},
			parentRef: gatev1.ParentReference{
				SectionName: new(gatev1.SectionName("foo")),
				Port:        new(gatev1.PortNumber(80)),
			},
			wantMatch: true,
		},
		{
			desc: "Port do not match",
			gwListener: gatewayListener{
				SectionName: "foo",
				Port:        gatev1.PortNumber(90),
			},
			parentRef: gatev1.ParentReference{
				SectionName: new(gatev1.SectionName("foo")),
				Port:        new(gatev1.PortNumber(80)),
			},
		},
		{
			desc: "Port matches",
			gwListener: gatewayListener{
				SectionName: "foo",
				Port:        gatev1.PortNumber(80),
			},
			parentRef: gatev1.ParentReference{
				SectionName: new(gatev1.SectionName("foo")),
				Port:        new(gatev1.PortNumber(80)),
			},
			wantMatch: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			gotMatch := matchListener(test.gwListener, test.parentRef)
			assert.Equal(t, test.wantMatch, gotMatch)
		})
	}
}

func Test_allowRoute(t *testing.T) {
	testCases := []struct {
		desc           string
		gwListener     gatewayListener
		routeNamespace string
		routeKind      string
		wantAllow      bool
	}{
		{
			desc: "Not allowed Kind",
			gwListener: gatewayListener{
				SectionName: "foo",
				AllowedRouteKinds: []string{
					"foo",
					"bar",
				},
			},
			routeKind: "baz",
			wantAllow: false,
		},
		{
			desc: "Allowed Kind",
			gwListener: gatewayListener{
				SectionName: "foo",
				AllowedRouteKinds: []string{
					"foo",
					"bar",
				},
				AllowedNamespaces: []string{
					corev1.NamespaceAll,
				},
			},
			routeKind: "bar",
			wantAllow: true,
		},
		{
			desc: "Not allowed namespace",
			gwListener: gatewayListener{
				SectionName: "foo",
				AllowedRouteKinds: []string{
					"foo",
				},
				AllowedNamespaces: []string{
					"foo",
					"bar",
				},
			},
			routeKind:      "foo",
			routeNamespace: "baz",
			wantAllow:      false,
		},
		{
			desc: "Allowed namespace",
			gwListener: gatewayListener{
				SectionName: "foo",
				AllowedRouteKinds: []string{
					"foo",
				},
				AllowedNamespaces: []string{
					"foo",
					"bar",
				},
			},
			routeKind:      "foo",
			routeNamespace: "foo",
			wantAllow:      true,
		},
		{
			desc: "Allowed namespace",
			gwListener: gatewayListener{
				SectionName: "foo",
				AllowedRouteKinds: []string{
					"foo",
				},
				AllowedNamespaces: []string{
					corev1.NamespaceAll,
					"bar",
				},
			},
			routeKind:      "foo",
			routeNamespace: "foo",
			wantAllow:      true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			gotAllow := allowRoute(test.gwListener, test.routeNamespace, test.routeKind)
			assert.Equal(t, test.wantAllow, gotAllow)
		})
	}
}

func Test_findMatchingHostnames(t *testing.T) {
	testCases := []struct {
		desc             string
		listenerHostname *gatev1.Hostname
		routeHostnames   []gatev1.Hostname
		want             []gatev1.Hostname
		wantOk           bool
	}{
		{
			desc:   "Empty",
			wantOk: true,
		},
		{
			desc:             "Only listener hostname",
			listenerHostname: new(gatev1.Hostname("foo.com")),
			want:             []gatev1.Hostname{"foo.com"},
			wantOk:           true,
		},
		{
			desc:           "Only Route hostname",
			routeHostnames: []gatev1.Hostname{"foo.com"},
			want:           []gatev1.Hostname{"foo.com"},
			wantOk:         true,
		},
		{
			desc:             "Matching hostname",
			listenerHostname: new(gatev1.Hostname("foo.com")),
			routeHostnames:   []gatev1.Hostname{"foo.com"},
			want:             []gatev1.Hostname{"foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Matching hostname with wildcard",
			listenerHostname: new(gatev1.Hostname("*.foo.com")),
			routeHostnames:   []gatev1.Hostname{"*.foo.com"},
			want:             []gatev1.Hostname{"*.foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Matching subdomain with listener wildcard",
			listenerHostname: new(gatev1.Hostname("*.foo.com")),
			routeHostnames:   []gatev1.Hostname{"bar.foo.com"},
			want:             []gatev1.Hostname{"bar.foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Matching subdomain wildcard with listener wildcard",
			listenerHostname: new(gatev1.Hostname("*.foo.com")),
			routeHostnames:   []gatev1.Hostname{"*.bar.foo.com"},
			want:             []gatev1.Hostname{"*.bar.foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Matching subsubdomain with listener wildcard",
			listenerHostname: new(gatev1.Hostname("*.foo.com")),
			routeHostnames:   []gatev1.Hostname{"baz.bar.foo.com"},
			want:             []gatev1.Hostname{"baz.bar.foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Matching subdomain with route hostname wildcard",
			listenerHostname: new(gatev1.Hostname("bar.foo.com")),
			routeHostnames:   []gatev1.Hostname{"*.foo.com"},
			want:             []gatev1.Hostname{"bar.foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Matching subsubdomain with route hostname wildcard",
			listenerHostname: new(gatev1.Hostname("baz.bar.foo.com")),
			routeHostnames:   []gatev1.Hostname{"*.foo.com"},
			want:             []gatev1.Hostname{"baz.bar.foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Non matching root domain with listener wildcard",
			listenerHostname: new(gatev1.Hostname("*.foo.com")),
			routeHostnames:   []gatev1.Hostname{"foo.com"},
		},
		{
			desc:             "Non matching root domain with route hostname wildcard",
			listenerHostname: new(gatev1.Hostname("foo.com")),
			routeHostnames:   []gatev1.Hostname{"*.foo.com"},
		},
		{
			desc:             "Multiple route hostnames with one matching route hostname",
			listenerHostname: new(gatev1.Hostname("*.foo.com")),
			routeHostnames:   []gatev1.Hostname{"bar.com", "test.foo.com", "test.buz.com"},
			want:             []gatev1.Hostname{"test.foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Multiple route hostnames with non matching route hostname",
			listenerHostname: new(gatev1.Hostname("*.fuz.com")),
			routeHostnames:   []gatev1.Hostname{"bar.com", "test.foo.com", "test.buz.com"},
		},
		{
			desc:             "Multiple route hostnames with multiple matching route hostnames",
			listenerHostname: new(gatev1.Hostname("*.foo.com")),
			routeHostnames:   []gatev1.Hostname{"toto.foo.com", "test.foo.com", "test.buz.com"},
			want:             []gatev1.Hostname{"test.foo.com", "toto.foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Matching wildcard subsubdomain with listener wildcard subdomain",
			listenerHostname: new(gatev1.Hostname("*.foo.com")),
			routeHostnames:   []gatev1.Hostname{"*.bar.foo.com"},
			want:             []gatev1.Hostname{"*.bar.foo.com"},
			wantOk:           true,
		},
		{
			desc:             "Duplicate intersections are returned once",
			listenerHostname: new(gatev1.Hostname("bar.foo.com")),
			routeHostnames:   []gatev1.Hostname{"*.foo.com", "bar.foo.com"},
			want:             []gatev1.Hostname{"bar.foo.com"},
			wantOk:           true,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			got, ok := findMatchingHostnames(test.listenerHostname, test.routeHostnames)
			assert.Equal(t, test.wantOk, ok)
			assert.Equal(t, test.want, got)
		})
	}
}

func Test_allowedRouteKinds(t *testing.T) {
	testCases := []struct {
		desc                string
		listener            gatev1.Listener
		supportedRouteKinds []gatev1.RouteGroupKind
		wantKinds           []gatev1.RouteGroupKind
		wantErr             bool
	}{
		{
			desc: "Empty",
		},
		{
			desc: "Empty AllowedRoutes",
			supportedRouteKinds: []gatev1.RouteGroupKind{
				{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
			},
			wantKinds: []gatev1.RouteGroupKind{
				{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
			},
		},
		{
			desc: "AllowedRoutes with unsupported Group",
			listener: gatev1.Listener{
				AllowedRoutes: &gatev1.AllowedRoutes{
					Kinds: []gatev1.RouteGroupKind{{
						Kind: kindTLSRoute, Group: new(gatev1.Group("foo")),
					}},
				},
			},
			supportedRouteKinds: []gatev1.RouteGroupKind{
				{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
			},
			wantErr: true,
		},
		{
			desc: "AllowedRoutes with nil Group",
			listener: gatev1.Listener{
				AllowedRoutes: &gatev1.AllowedRoutes{
					Kinds: []gatev1.RouteGroupKind{{
						Kind: kindTLSRoute, Group: nil,
					}},
				},
			},
			supportedRouteKinds: []gatev1.RouteGroupKind{
				{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
			},
			wantErr: true,
		},
		{
			desc: "AllowedRoutes with unsupported Kind",
			listener: gatev1.Listener{
				AllowedRoutes: &gatev1.AllowedRoutes{
					Kinds: []gatev1.RouteGroupKind{{
						Kind: "foo", Group: new(gatev1.Group(gatev1.GroupName)),
					}},
				},
			},
			supportedRouteKinds: []gatev1.RouteGroupKind{
				{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
			},
			wantErr: true,
		},
		{
			desc: "Supported AllowedRoutes",
			listener: gatev1.Listener{
				AllowedRoutes: &gatev1.AllowedRoutes{
					Kinds: []gatev1.RouteGroupKind{{
						Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName)),
					}},
				},
			},
			supportedRouteKinds: []gatev1.RouteGroupKind{
				{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
			},
			wantKinds: []gatev1.RouteGroupKind{
				{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
			},
		},
		{
			desc: "Supported AllowedRoutes with duplicates",
			listener: gatev1.Listener{
				AllowedRoutes: &gatev1.AllowedRoutes{
					Kinds: []gatev1.RouteGroupKind{
						{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
						{Kind: kindTCPRoute, Group: new(gatev1.Group(gatev1.GroupName))},
						{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
						{Kind: kindTCPRoute, Group: new(gatev1.Group(gatev1.GroupName))},
					},
				},
			},
			supportedRouteKinds: []gatev1.RouteGroupKind{
				{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
				{Kind: kindTCPRoute, Group: new(gatev1.Group(gatev1.GroupName))},
			},
			wantKinds: []gatev1.RouteGroupKind{
				{Kind: kindTLSRoute, Group: new(gatev1.Group(gatev1.GroupName))},
				{Kind: kindTCPRoute, Group: new(gatev1.Group(gatev1.GroupName))},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			got, conditions := allowedRouteKinds(0, test.listener, test.supportedRouteKinds)
			if test.wantErr {
				require.NotEmpty(t, conditions, "no conditions")
				return
			}

			require.Empty(t, conditions)
			assert.Equal(t, test.wantKinds, got)
		})
	}
}

func Test_makeListenerKey(t *testing.T) {
	testCases := []struct {
		desc        string
		listener    gatev1.Listener
		expectedKey string
	}{
		{
			desc:        "empty",
			expectedKey: "||0",
		},
		{
			desc: "listener with port, protocol and hostname",
			listener: gatev1.Listener{
				Port:     443,
				Protocol: gatev1.HTTPSProtocolType,
				Hostname: new(gatev1.Hostname("www.example.com")),
			},
			expectedKey: "HTTPS|www.example.com|443",
		},
		{
			desc: "listener with port, protocol and nil hostname",
			listener: gatev1.Listener{
				Port:     443,
				Protocol: gatev1.HTTPSProtocolType,
			},
			expectedKey: "HTTPS||443",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.expectedKey, makeListenerKey(test.listener))
		})
	}
}

func Test_referenceGrantMatchesFrom(t *testing.T) {
	testCases := []struct {
		desc           string
		referenceGrant gatev1.ReferenceGrant
		group          string
		kind           string
		namespace      string
		expectedResult bool
	}{
		{
			desc: "matches",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					From: []gatev1.ReferenceGrantFrom{
						{
							Group:     "correct-group",
							Kind:      "correct-kind",
							Namespace: "correct-namespace",
						},
					},
				},
			},
			group:          "correct-group",
			kind:           "correct-kind",
			namespace:      "correct-namespace",
			expectedResult: true,
		},
		{
			desc: "empty group matches core",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					From: []gatev1.ReferenceGrantFrom{
						{
							Group:     "",
							Kind:      "correct-kind",
							Namespace: "correct-namespace",
						},
					},
				},
			},
			group:          "core",
			kind:           "correct-kind",
			namespace:      "correct-namespace",
			expectedResult: true,
		},
		{
			desc: "wrong group",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					From: []gatev1.ReferenceGrantFrom{
						{
							Group:     "wrong-group",
							Kind:      "correct-kind",
							Namespace: "correct-namespace",
						},
					},
				},
			},
			group:          "correct-group",
			kind:           "correct-kind",
			namespace:      "correct-namespace",
			expectedResult: false,
		},
		{
			desc: "wrong kind",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					From: []gatev1.ReferenceGrantFrom{
						{
							Group:     "correct-group",
							Kind:      "wrong-kind",
							Namespace: "correct-namespace",
						},
					},
				},
			},
			group:          "correct-group",
			kind:           "correct-kind",
			namespace:      "correct-namespace",
			expectedResult: false,
		},
		{
			desc: "wrong namespace",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					From: []gatev1.ReferenceGrantFrom{
						{
							Group:     "correct-group",
							Kind:      "correct-kind",
							Namespace: "wrong-namespace",
						},
					},
				},
			},
			group:          "correct-group",
			kind:           "correct-kind",
			namespace:      "correct-namespace",
			expectedResult: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.expectedResult, referenceGrantMatchesFrom(&test.referenceGrant, test.group, test.kind, test.namespace))
		})
	}
}

func Test_referenceGrantMatchesTo(t *testing.T) {
	testCases := []struct {
		desc           string
		referenceGrant gatev1.ReferenceGrant
		group          string
		kind           string
		name           string
		expectedResult bool
	}{
		{
			desc: "matches",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					To: []gatev1.ReferenceGrantTo{
						{
							Group: "correct-group",
							Kind:  "correct-kind",
							Name:  new(gatev1.ObjectName("correct-name")),
						},
					},
				},
			},
			group:          "correct-group",
			kind:           "correct-kind",
			name:           "correct-name",
			expectedResult: true,
		},
		{
			desc: "matches without name",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					To: []gatev1.ReferenceGrantTo{
						{
							Group: "correct-group",
							Kind:  "correct-kind",
							Name:  nil,
						},
					},
				},
			},
			group:          "correct-group",
			kind:           "correct-kind",
			name:           "some-name",
			expectedResult: true,
		},
		{
			desc: "empty group matches core",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					To: []gatev1.ReferenceGrantTo{
						{
							Group: "",
							Kind:  "correct-kind",
							Name:  new(gatev1.ObjectName("correct-name")),
						},
					},
				},
			},
			group:          "core",
			kind:           "correct-kind",
			name:           "correct-name",
			expectedResult: true,
		},
		{
			desc: "wrong group",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					To: []gatev1.ReferenceGrantTo{
						{
							Group: "wrong-group",
							Kind:  "correct-kind",
							Name:  new(gatev1.ObjectName("correct-name")),
						},
					},
				},
			},
			group:          "correct-group",
			kind:           "correct-kind",
			name:           "correct-namespace",
			expectedResult: false,
		},
		{
			desc: "wrong kind",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					To: []gatev1.ReferenceGrantTo{
						{
							Group: "correct-group",
							Kind:  "wrong-kind",
							Name:  new(gatev1.ObjectName("correct-name")),
						},
					},
				},
			},
			group:          "correct-group",
			kind:           "correct-kind",
			name:           "correct-name",
			expectedResult: false,
		},
		{
			desc: "wrong name",
			referenceGrant: gatev1.ReferenceGrant{
				Spec: gatev1.ReferenceGrantSpec{
					To: []gatev1.ReferenceGrantTo{
						{
							Group: "correct-group",
							Kind:  "correct-kind",
							Name:  new(gatev1.ObjectName("wrong-name")),
						},
					},
				},
			},
			group:          "correct-group",
			kind:           "correct-kind",
			name:           "correct-name",
			expectedResult: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.expectedResult, referenceGrantMatchesTo(&test.referenceGrant, test.group, test.kind, test.name))
		})
	}
}

func Test_gatewayAddresses(t *testing.T) {
	testCases := []struct {
		desc          string
		statusAddress *StatusAddress
		paths         []string
		wantErr       require.ErrorAssertionFunc
		want          []gatev1.GatewayStatusAddress
	}{
		{
			desc:    "nothing",
			wantErr: require.NoError,
		},
		{
			desc:          "empty configuration",
			statusAddress: &StatusAddress{},
			wantErr:       require.Error,
		},
		{
			desc: "IP address",
			statusAddress: &StatusAddress{
				IP: "1.2.3.4",
			},
			wantErr: require.NoError,
			want: []gatev1.GatewayStatusAddress{
				{
					Type:  new(gatev1.IPAddressType),
					Value: "1.2.3.4",
				},
			},
		},
		{
			desc: "hostname address",
			statusAddress: &StatusAddress{
				Hostname: "foo.bar",
			},
			wantErr: require.NoError,
			want: []gatev1.GatewayStatusAddress{
				{
					Type:  new(gatev1.HostnameAddressType),
					Value: "foo.bar",
				},
			},
		},
		{
			desc: "service",
			statusAddress: &StatusAddress{
				Service: ServiceRef{
					Name:      "status-address",
					Namespace: "default",
				},
			},
			paths:   []string{"services.yml"},
			wantErr: require.NoError,
			want: []gatev1.GatewayStatusAddress{
				{
					Type:  new(gatev1.HostnameAddressType),
					Value: "foo.bar",
				},
				{
					Type:  new(gatev1.IPAddressType),
					Value: "1.2.3.4",
				},
			},
		},
		{
			desc: "missing service",
			statusAddress: &StatusAddress{
				Service: ServiceRef{
					Name:      "status-address2",
					Namespace: "default",
				},
			},
			wantErr: require.Error,
		},
		{
			desc: "service without load-balancer status",
			statusAddress: &StatusAddress{
				Service: ServiceRef{
					Name:      "whoamitcp-bar",
					Namespace: "bar",
				},
			},
			paths:   []string{"services.yml"},
			wantErr: require.NoError,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			k8sObjects, gwObjects := readResources(t, test.paths)

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				StatusAddress: test.statusAddress,
				client:        client,
			}

			got, err := p.gatewayAddresses()
			test.wantErr(t, err)

			assert.Equal(t, test.want, got)
		})
	}
}

func Test_makeGatewayStatus(t *testing.T) {
	acceptedListener := func() gatewayListener {
		return gatewayListener{Status: &gatev1.ListenerStatus{Name: "valid"}}
	}
	invalidListener := func() gatewayListener {
		return gatewayListener{Status: &gatev1.ListenerStatus{
			Name: "invalid",
			Conditions: []metav1.Condition{{
				Type:   string(gatev1.ListenerConditionAccepted),
				Status: metav1.ConditionFalse,
				Reason: string(gatev1.ListenerReasonUnsupportedProtocol),
			}},
		}}
	}
	unresolvedListener := func() gatewayListener {
		return gatewayListener{Status: &gatev1.ListenerStatus{
			Name: "unresolved",
			Conditions: []metav1.Condition{{
				Type:   string(gatev1.ListenerConditionResolvedRefs),
				Status: metav1.ConditionFalse,
				Reason: string(gatev1.ListenerReasonInvalidCertificateRef),
			}},
		}}
	}
	unresolvedListenerSetListener := func() gatewayListener {
		listener := unresolvedListener()
		listener.ListenerSet = &gatev1.ListenerSet{}
		return listener
	}
	resolvedRefsCondition := func(status metav1.ConditionStatus, reason gatev1.GatewayConditionReason, message string) metav1.Condition {
		return metav1.Condition{
			Type:               string(gatev1.GatewayConditionResolvedRefs),
			Status:             status,
			ObservedGeneration: 2,
			Reason:             string(reason),
			Message:            message,
		}
	}
	clientCertificateCondition := &metav1.Condition{
		Type:               string(gatev1.GatewayConditionResolvedRefs),
		Status:             metav1.ConditionFalse,
		ObservedGeneration: 2,
		Reason:             string(gatev1.GatewayReasonInvalidClientCertificateRef),
		Message:            "Cannot load ClientCertificateRef default/missing: getting secret: secret not found",
	}

	testCases := []struct {
		desc                       string
		infrastructure             *gatev1.GatewayInfrastructure
		listeners                  []gatewayListener
		clientCertificateCondition *metav1.Condition
		wantStatus                 metav1.ConditionStatus
		wantReason                 gatev1.GatewayConditionReason
		wantProgrammed             metav1.ConditionStatus
		wantResolvedRefs           metav1.Condition
	}{
		{
			desc:             "all listeners valid",
			listeners:        []gatewayListener{acceptedListener()},
			wantStatus:       metav1.ConditionTrue,
			wantReason:       gatev1.GatewayReasonAccepted,
			wantProgrammed:   metav1.ConditionTrue,
			wantResolvedRefs: resolvedRefsCondition(metav1.ConditionTrue, gatev1.GatewayReasonResolvedRefs, "No error found"),
		},
		{
			desc:             "one listener valid among invalid ones",
			listeners:        []gatewayListener{acceptedListener(), invalidListener()},
			wantStatus:       metav1.ConditionTrue,
			wantReason:       gatev1.GatewayReasonListenersNotValid,
			wantProgrammed:   metav1.ConditionTrue,
			wantResolvedRefs: resolvedRefsCondition(metav1.ConditionTrue, gatev1.GatewayReasonResolvedRefs, "No error found"),
		},
		{
			desc:             "no listener valid",
			listeners:        []gatewayListener{invalidListener()},
			wantStatus:       metav1.ConditionFalse,
			wantReason:       gatev1.GatewayReasonListenersNotValid,
			wantProgrammed:   metav1.ConditionFalse,
			wantResolvedRefs: resolvedRefsCondition(metav1.ConditionTrue, gatev1.GatewayReasonResolvedRefs, "No error found"),
		},
		{
			desc:      "infrastructure parametersRef",
			listeners: []gatewayListener{acceptedListener()},
			infrastructure: &gatev1.GatewayInfrastructure{
				ParametersRef: &gatev1.LocalParametersReference{
					Group: "invalid.io",
					Kind:  "InvalidParameters",
					Name:  "invalid",
				},
			},
			wantStatus:       metav1.ConditionFalse,
			wantReason:       gatev1.GatewayReasonInvalidParameters,
			wantResolvedRefs: resolvedRefsCondition(metav1.ConditionTrue, gatev1.GatewayReasonResolvedRefs, "No error found"),
		},
		{
			desc:             "one listener with unresolved references",
			listeners:        []gatewayListener{acceptedListener(), unresolvedListener()},
			wantStatus:       metav1.ConditionTrue,
			wantReason:       gatev1.GatewayReasonListenersNotValid,
			wantProgrammed:   metav1.ConditionTrue,
			wantResolvedRefs: resolvedRefsCondition(metav1.ConditionFalse, gatev1.GatewayReasonListenersNotResolved, "At least one Listener has unresolved references"),
		},
		{
			desc:             "ListenerSet listener with unresolved references",
			listeners:        []gatewayListener{acceptedListener(), unresolvedListenerSetListener()},
			wantStatus:       metav1.ConditionTrue,
			wantReason:       gatev1.GatewayReasonAccepted,
			wantProgrammed:   metav1.ConditionTrue,
			wantResolvedRefs: resolvedRefsCondition(metav1.ConditionTrue, gatev1.GatewayReasonResolvedRefs, "No error found"),
		},
		{
			desc:                       "unresolved client certificate",
			listeners:                  []gatewayListener{acceptedListener()},
			clientCertificateCondition: clientCertificateCondition,
			wantStatus:                 metav1.ConditionTrue,
			wantReason:                 gatev1.GatewayReasonAccepted,
			wantProgrammed:             metav1.ConditionTrue,
			wantResolvedRefs:           *clientCertificateCondition,
		},
		{
			desc:                       "unresolved client certificate takes precedence over the listeners",
			listeners:                  []gatewayListener{acceptedListener(), unresolvedListener()},
			clientCertificateCondition: clientCertificateCondition,
			wantStatus:                 metav1.ConditionTrue,
			wantReason:                 gatev1.GatewayReasonListenersNotValid,
			wantProgrammed:             metav1.ConditionTrue,
			wantResolvedRefs:           *clientCertificateCondition,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			p := Provider{}
			gateway := &gatev1.Gateway{
				Generation: 2,
				Spec:       gatev1.GatewaySpec{Infrastructure: test.infrastructure},
			}

			status, _ := p.makeGatewayStatus(gateway, test.listeners, test.clientCertificateCondition, nil, slices.ContainsFunc(test.listeners, func(listener gatewayListener) bool {
				return len(listener.Status.Conditions) == 0
			}))

			condition := meta.FindStatusCondition(status.Conditions, string(gatev1.GatewayConditionAccepted))
			require.NotNil(t, condition)
			assert.Equal(t, test.wantStatus, condition.Status)
			assert.Equal(t, string(test.wantReason), condition.Reason)
			assert.Len(t, status.Listeners, len(slices.DeleteFunc(slices.Clone(test.listeners), func(listener gatewayListener) bool {
				return listener.ListenerSet != nil
			})))

			resolvedRefs := meta.FindStatusCondition(status.Conditions, string(gatev1.GatewayConditionResolvedRefs))
			require.NotNil(t, resolvedRefs)
			resolvedRefs.LastTransitionTime = metav1.Time{}
			assert.Equal(t, test.wantResolvedRefs, *resolvedRefs)

			programmed := meta.FindStatusCondition(status.Conditions, string(gatev1.GatewayConditionProgrammed))
			if test.wantProgrammed == "" {
				assert.Nil(t, programmed)
				return
			}
			require.NotNil(t, programmed)
			assert.Equal(t, test.wantProgrammed, programmed.Status)
		})
	}
}

func Test_resolveFrontendValidation(t *testing.T) {
	k8sObjects, gwObjects := readResources(t, []string{"frontend_validation_ca_refs.yml"})

	kubeClient := kubefake.NewClientset(k8sObjects...)
	gwClient := newGatewaySimpleClientSet(t, gwObjects...)

	client := newClientImpl(kubeClient, gwClient)

	eventCh, err := client.WatchAll(nil, make(chan struct{}))
	require.NoError(t, err)

	if len(k8sObjects) > 0 || len(gwObjects) > 0 {
		// just wait for the first event
		<-eventCh
	}

	p := Provider{client: client}

	gateway := &gatev1.Gateway{
		Name:       "my-gateway",
		Namespace:  "default",
		Generation: 1,
	}

	testCases := []struct {
		desc       string
		validation *gatev1.FrontendTLSValidation
		expected   frontendValidation
	}{
		{
			desc:       "nil validation returns an empty result",
			validation: nil,
			expected:   frontendValidation{},
		},
		{
			desc: "a valid ConfigMap ref is resolved",
			validation: &gatev1.FrontendTLSValidation{
				Mode: gatev1.AllowValidOnly,
				CACertificateRefs: []gatev1.ObjectReference{
					{Kind: "ConfigMap", Name: "ca-configmap"},
				},
			},
			expected: frontendValidation{
				clientAuth: &tls.ClientAuth{
					CAFiles:        []types.FileOrContent{"CA1"},
					ClientAuthType: tls.RequireAndVerifyClientCert,
				},
			},
		},
		{
			desc: "AllowInsecureFallback mode maps to RequestClientCert",
			validation: &gatev1.FrontendTLSValidation{
				Mode: gatev1.AllowInsecureFallback,
				CACertificateRefs: []gatev1.ObjectReference{
					{Kind: "ConfigMap", Name: "ca-configmap"},
				},
			},
			expected: frontendValidation{
				clientAuth: &tls.ClientAuth{
					CAFiles:        []types.FileOrContent{"CA1"},
					ClientAuthType: tls.RequestClientCert,
				},
			},
		},
		{
			desc: "unsupported kind is rejected and no valid CA certificate remains",
			validation: &gatev1.FrontendTLSValidation{
				CACertificateRefs: []gatev1.ObjectReference{
					{Kind: "Service", Name: "whoami"},
				},
			},
			expected: frontendValidation{
				errCondition: &metav1.Condition{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalidCACertificateKind),
					Message:            "Unsupported CACertificateRef group/kind: /Service",
				},
			},
		},
		{
			desc: "unsupported group is rejected",
			validation: &gatev1.FrontendTLSValidation{
				CACertificateRefs: []gatev1.ObjectReference{
					{Group: "example.com", Kind: "ConfigMap", Name: "ca-configmap"},
				},
			},
			expected: frontendValidation{
				errCondition: &metav1.Condition{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalidCACertificateKind),
					Message:            "Unsupported CACertificateRef group/kind: example.com/ConfigMap",
				},
			},
		},
		{
			desc: "cross-namespace ref without a ReferenceGrant is rejected",
			validation: &gatev1.FrontendTLSValidation{
				CACertificateRefs: []gatev1.ObjectReference{
					{Kind: "ConfigMap", Name: "ca-configmap-not-granted", Namespace: new(gatev1.Namespace("other-ns"))},
				},
			},
			expected: frontendValidation{
				errCondition: &metav1.Condition{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonRefNotPermitted),
					Message:            "Cannot reference CACertificateRef: /ConfigMap: missing ReferenceGrant",
				},
			},
		},
		{
			desc: "cross-namespace ref with a matching ReferenceGrant is resolved",
			validation: &gatev1.FrontendTLSValidation{
				CACertificateRefs: []gatev1.ObjectReference{
					{Kind: "ConfigMap", Name: "ca-configmap-granted", Namespace: new(gatev1.Namespace("other-ns"))},
				},
			},
			expected: frontendValidation{
				clientAuth: &tls.ClientAuth{
					CAFiles:        []types.FileOrContent{"CA4"},
					ClientAuthType: tls.RequireAndVerifyClientCert,
				},
			},
		},
		{
			desc: "Secret ref is rejected as an unsupported kind",
			validation: &gatev1.FrontendTLSValidation{
				CACertificateRefs: []gatev1.ObjectReference{
					{Kind: "Secret", Name: "ca-secret"},
				},
			},
			expected: frontendValidation{
				errCondition: &metav1.Condition{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalidCACertificateKind),
					Message:            "Unsupported CACertificateRef group/kind: /Secret",
				},
			},
		},
		{
			desc: "unresolvable ConfigMap is rejected",
			validation: &gatev1.FrontendTLSValidation{
				CACertificateRefs: []gatev1.ObjectReference{
					{Kind: "ConfigMap", Name: "does-not-exist"},
				},
			},
			expected: frontendValidation{
				errCondition: &metav1.Condition{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalidCACertificateRef),
					Message:            `Cannot resolve configmap: default/does-not-exist: configmap "does-not-exist" not found`,
				},
			},
		},
		{
			desc: "ConfigMap missing the ca.crt key is rejected",
			validation: &gatev1.FrontendTLSValidation{
				CACertificateRefs: []gatev1.ObjectReference{
					{Kind: "ConfigMap", Name: "ca-configmap-empty"},
				},
			},
			expected: frontendValidation{
				errCondition: &metav1.Condition{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalidCACertificateRef),
					Message:            "Cannot find ca.crt: ConfigMap default/ca-configmap-empty",
				},
			},
		},
		{
			desc: "multiple refs are rejected",
			validation: &gatev1.FrontendTLSValidation{
				CACertificateRefs: []gatev1.ObjectReference{
					{Kind: "ConfigMap", Name: "ca-configmap"},
					{Kind: "ConfigMap", Name: "ca-configmap-granted", Namespace: new(gatev1.Namespace("other-ns"))},
				},
			},
			expected: frontendValidation{
				errCondition: &metav1.Condition{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalidCACertificateRef),
					Message:            "Unsupported number of CACertificateRefs: 2, only one reference is supported",
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			got := p.resolveFrontendValidation(gateway, test.validation)

			// LastTransitionTime is set with metav1.Now() and is not relevant to this test.
			if got.errCondition != nil {
				got.errCondition.LastTransitionTime = metav1.Time{}
			}

			assert.Equal(t, test.expected, got)
		})
	}
}

func Test_resolveBackendClientCertificate(t *testing.T) {
	k8sObjects, gwObjects := readResources(t, []string{"backend_client_certificate_refs.yml"})

	kubeClient := kubefake.NewClientset(k8sObjects...)
	gwClient := newGatewaySimpleClientSet(t, gwObjects...)

	client := newClientImpl(kubeClient, gwClient)

	eventCh, err := client.WatchAll(nil, make(chan struct{}))
	require.NoError(t, err)

	// just wait for the first event
	<-eventCh

	p := Provider{client: client}

	testCases := []struct {
		desc           string
		tls            *gatev1.GatewayTLSConfig
		wantCommonName string
		wantCondition  *metav1.Condition
	}{
		{
			desc: "no TLS configuration",
		},
		{
			desc: "no backend TLS configuration",
			tls:  &gatev1.GatewayTLSConfig{},
		},
		{
			desc: "no client certificate reference",
			tls:  &gatev1.GatewayTLSConfig{Backend: &gatev1.GatewayBackendTLS{}},
		},
		{
			desc: "Secret in the Gateway namespace",
			tls: &gatev1.GatewayTLSConfig{Backend: &gatev1.GatewayBackendTLS{
				ClientCertificateRef: &gatev1.SecretObjectReference{Name: "client-cert"},
			}},
			wantCommonName: "example.com",
		},
		{
			desc: "Secret with the core group",
			tls: &gatev1.GatewayTLSConfig{Backend: &gatev1.GatewayBackendTLS{
				ClientCertificateRef: &gatev1.SecretObjectReference{
					Group: new(gatev1.Group(groupCore)),
					Kind:  new(gatev1.Kind(kindSecret)),
					Name:  "client-cert",
				},
			}},
			wantCommonName: "example.com",
		},
		{
			desc: "cross-namespace Secret with a matching ReferenceGrant",
			tls: &gatev1.GatewayTLSConfig{Backend: &gatev1.GatewayBackendTLS{
				ClientCertificateRef: &gatev1.SecretObjectReference{
					Name:      "client-cert-granted",
					Namespace: new(gatev1.Namespace("other-ns")),
				},
			}},
			wantCommonName: "default-a.example.com",
		},
		{
			desc: "unsupported group",
			tls: &gatev1.GatewayTLSConfig{Backend: &gatev1.GatewayBackendTLS{
				ClientCertificateRef: &gatev1.SecretObjectReference{
					Group: new(gatev1.Group("example.com")),
					Name:  "client-cert",
				},
			}},
			wantCondition: &metav1.Condition{
				Type:               string(gatev1.GatewayConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: 1,
				Reason:             string(gatev1.GatewayReasonInvalidClientCertificateRef),
				Message:            "Unsupported ClientCertificateRef group/kind: example.com/Secret",
			},
		},
		{
			desc: "unsupported kind",
			tls: &gatev1.GatewayTLSConfig{Backend: &gatev1.GatewayBackendTLS{
				ClientCertificateRef: &gatev1.SecretObjectReference{
					Kind: new(gatev1.Kind(kindConfigMap)),
					Name: "client-cert",
				},
			}},
			wantCondition: &metav1.Condition{
				Type:               string(gatev1.GatewayConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: 1,
				Reason:             string(gatev1.GatewayReasonInvalidClientCertificateRef),
				Message:            "Unsupported ClientCertificateRef group/kind: /ConfigMap",
			},
		},
		{
			desc: "cross-namespace Secret without a ReferenceGrant",
			tls: &gatev1.GatewayTLSConfig{Backend: &gatev1.GatewayBackendTLS{
				ClientCertificateRef: &gatev1.SecretObjectReference{
					Name:      "client-cert-not-granted",
					Namespace: new(gatev1.Namespace("other-ns")),
				},
			}},
			wantCondition: &metav1.Condition{
				Type:               string(gatev1.GatewayConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: 1,
				Reason:             string(gatev1.GatewayReasonRefNotPermitted),
				Message:            "Cannot reference ClientCertificateRef other-ns/client-cert-not-granted: missing ReferenceGrant",
			},
		},
		{
			desc: "missing Secret",
			tls: &gatev1.GatewayTLSConfig{Backend: &gatev1.GatewayBackendTLS{
				ClientCertificateRef: &gatev1.SecretObjectReference{Name: "missing"},
			}},
			wantCondition: &metav1.Condition{
				Type:               string(gatev1.GatewayConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: 1,
				Reason:             string(gatev1.GatewayReasonInvalidClientCertificateRef),
				Message:            `Cannot load ClientCertificateRef default/missing: getting secret: secret "missing" not found`,
			},
		},
		{
			desc: "malformed Secret",
			tls: &gatev1.GatewayTLSConfig{Backend: &gatev1.GatewayBackendTLS{
				ClientCertificateRef: &gatev1.SecretObjectReference{Name: "malformed-cert"},
			}},
			wantCondition: &metav1.Condition{
				Type:               string(gatev1.GatewayConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: 1,
				Reason:             string(gatev1.GatewayReasonInvalidClientCertificateRef),
				Message:            "Cannot load ClientCertificateRef default/malformed-cert: validating certificate: unable to parse TLS certificate: tls: failed to find any PEM data in certificate input",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			gateway := &gatev1.Gateway{
				Name:       "my-gateway",
				Namespace:  "default",
				Generation: 1,
				Spec:       gatev1.GatewaySpec{TLS: test.tls},
			}

			cert, condition := p.resolveBackendClientCertificate(gateway)

			if condition != nil {
				// LastTransitionTime is set with metav1.Now() and is not relevant to this test.
				condition.LastTransitionTime = metav1.Time{}
			}
			assert.Equal(t, test.wantCondition, condition)

			if test.wantCommonName == "" {
				assert.Nil(t, cert)
				return
			}
			require.NotNil(t, cert)
			assert.Equal(t, test.wantCommonName, certificateCommonName(t, cert.CertFile))
		})
	}
}

func Test_loadConfigurationFromGateways_backendClientCertificate(t *testing.T) {
	testCases := []struct {
		desc        string
		path        string
		entryPoints map[string]Entrypoint
		// wantClientCertificates maps each ServersTransport to the common name of its client certificate, empty when it has none.
		wantClientCertificates map[string]string
		wantResolvedRefs       map[string]metav1.Condition
	}{
		{
			desc:        "ListenerSet listener uses its parent Gateway client certificate",
			path:        "httproute/with_listenerset_and_backend_client_certificate.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}, "web8080": {Address: ":8080"}},
			wantClientCertificates: map[string]string{
				"httproute-default-http-app-1-gw-default-my-gateway-ep-web8080-0-fdbd6d4d56f70e64fbd9-svc-default-whoami-0": "default-a.example.com",
			},
			wantResolvedRefs: map[string]metav1.Condition{
				"my-gateway": {
					Type:    string(gatev1.GatewayConditionResolvedRefs),
					Status:  metav1.ConditionTrue,
					Reason:  string(gatev1.GatewayReasonResolvedRefs),
					Message: "No error found",
				},
			},
		},
		{
			desc:        "two Gateways each present their own client certificate",
			path:        "httproute/with_two_gateways_and_backend_client_certificates.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}},
			wantClientCertificates: map[string]string{
				"httproute-default-http-app-a-gw-default-gateway-a-ep-web-0-870f09bbeaa4c1793061-svc-default-whoami-0": "default-a.example.com",
				"httproute-default-http-app-b-gw-default-gateway-b-ep-web-0-10eb9167df873fb2abc6-svc-default-whoami-0": "default-b.example.com",
			},
			wantResolvedRefs: map[string]metav1.Condition{
				"gateway-a": {
					Type:    string(gatev1.GatewayConditionResolvedRefs),
					Status:  metav1.ConditionTrue,
					Reason:  string(gatev1.GatewayReasonResolvedRefs),
					Message: "No error found",
				},
				"gateway-b": {
					Type:    string(gatev1.GatewayConditionResolvedRefs),
					Status:  metav1.ConditionTrue,
					Reason:  string(gatev1.GatewayReasonResolvedRefs),
					Message: "No error found",
				},
			},
		},
		{
			desc:        "unresolved client certificate is not used",
			path:        "httproute/with_backend_tls_policy_and_missing_client_certificate.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}},
			wantClientCertificates: map[string]string{
				"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93-svc-default-whoami-0": "",
			},
			wantResolvedRefs: map[string]metav1.Condition{
				"my-gateway": {
					Type:    string(gatev1.GatewayConditionResolvedRefs),
					Status:  metav1.ConditionFalse,
					Reason:  string(gatev1.GatewayReasonInvalidClientCertificateRef),
					Message: `Cannot load ClientCertificateRef default/missing: getting secret: secret "missing" not found`,
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			k8sObjects, gwObjects := readResources(t, []string{"services.yml", test.path})

			gwClient := newGatewaySimpleClientSet(t, gwObjects...)
			client := newClientImpl(kubefake.NewClientset(k8sObjects...), gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			// just wait for the first event
			<-eventCh

			p := Provider{
				EntryPoints: test.entryPoints,
				client:      client,
			}

			conf, statusReport, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			clientCertificates := make(map[string]string)
			for name, st := range conf.HTTP.ServersTransports {
				clientCertificates[name] = ""
				if len(st.Certificates) > 0 {
					require.Len(t, st.Certificates, 1)
					clientCertificates[name] = certificateCommonName(t, st.Certificates[0].CertFile)
				}
			}
			assert.Equal(t, test.wantClientCertificates, clientCertificates)

			statusReport.Flush(t.Context(), client)

			for gatewayName, want := range test.wantResolvedRefs {
				gateway, err := gwClient.GatewayV1().Gateways("default").Get(t.Context(), gatewayName, metav1.GetOptions{})
				require.NoError(t, err)

				resolvedRefs := meta.FindStatusCondition(gateway.Status.Conditions, string(gatev1.GatewayConditionResolvedRefs))
				require.NotNil(t, resolvedRefs)

				want.ObservedGeneration = gateway.Generation
				resolvedRefs.LastTransitionTime = metav1.Time{}
				assert.Equal(t, want, *resolvedRefs)
			}
		})
	}
}

func Test_loadGatewayListeners(t *testing.T) {
	k8sObjects, gwObjects := readResources(t, []string{"frontend_validation_ca_refs.yml"})

	kubeClient := kubefake.NewClientset(k8sObjects...)
	gwClient := newGatewaySimpleClientSet(t, gwObjects...)

	client := newClientImpl(kubeClient, gwClient)

	eventCh, err := client.WatchAll(nil, make(chan struct{}))
	require.NoError(t, err)
	<-eventCh

	p := Provider{
		EntryPoints: map[string]Entrypoint{"websecure": {Address: ":443"}},
		client:      client,
	}

	newGateway := func(caCertificateRefs []gatev1.ObjectReference) *gatev1.Gateway {
		return &gatev1.Gateway{
			Name:       "my-gateway",
			Namespace:  "default",
			Generation: 1,
			Spec: gatev1.GatewaySpec{
				TLS: &gatev1.GatewayTLSConfig{
					Frontend: &gatev1.FrontendTLSConfig{
						Default: gatev1.TLSConfig{
							Validation: &gatev1.FrontendTLSValidation{
								Mode:              gatev1.AllowValidOnly,
								CACertificateRefs: caCertificateRefs,
							},
						},
					},
				},
				Listeners: []gatev1.Listener{
					{
						Name:     "https",
						Protocol: gatev1.HTTPSProtocolType,
						Port:     443,
						TLS: &gatev1.ListenerTLSConfig{
							CertificateRefs: []gatev1.SecretObjectReference{
								{Kind: new(gatev1.Kind("Secret")), Group: new(gatev1.Group("")), Name: "tls-secret"},
							},
						},
						AllowedRoutes: &gatev1.AllowedRoutes{
							Namespaces: &gatev1.RouteNamespaces{From: new(gatev1.NamespacesFromSame)},
						},
					},
				},
			},
		}
	}

	testCases := []struct {
		desc               string
		gateway            *gatev1.Gateway
		expectedAttached   bool
		expectedConditions []metav1.Condition
	}{
		{
			desc: "fully resolved CACertificateRefs attaches the Listener with no error condition",
			gateway: newGateway([]gatev1.ObjectReference{
				{Kind: "ConfigMap", Name: "ca-configmap"},
			}),
			expectedAttached:   true,
			expectedConditions: []metav1.Condition{},
		},
		{
			desc: "multiple CACertificateRefs are rejected and the Listener is not attached",
			gateway: newGateway([]gatev1.ObjectReference{
				{Kind: "ConfigMap", Name: "ca-configmap"},
				{Kind: "ConfigMap", Name: "ca-configmap-empty"},
			}),
			expectedAttached: false,
			expectedConditions: []metav1.Condition{
				{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalidCACertificateRef),
					Message:            "Unsupported number of CACertificateRefs: 2, only one reference is supported",
				},
				{
					Type:               string(gatev1.ListenerConditionAccepted),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonNoValidCACertificate),
					Message:            "No valid CA certificate found in CACertificateRefs",
				},
				{
					Type:               string(gatev1.ListenerConditionProgrammed),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalid),
					Message:            "Invalid CA certificate configuration",
				},
			},
		},
		{
			desc: "no valid CACertificateRefs reports NoValidCACertificate and the Listener is not attached",
			gateway: newGateway([]gatev1.ObjectReference{
				{Kind: "Service", Name: "whoami"},
			}),
			expectedAttached: false,
			expectedConditions: []metav1.Condition{
				{
					Type:               string(gatev1.ListenerConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalidCACertificateKind),
					Message:            "Unsupported CACertificateRef group/kind: /Service",
				},
				{
					Type:               string(gatev1.ListenerConditionAccepted),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonNoValidCACertificate),
					Message:            "No valid CA certificate found in CACertificateRefs",
				},
				{
					Type:               string(gatev1.ListenerConditionProgrammed),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: 1,
					Reason:             string(gatev1.ListenerReasonInvalid),
					Message:            "Invalid CA certificate configuration",
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			conf := &dynamic.Configuration{TLS: &dynamic.TLSConfiguration{}}
			listeners := p.loadGatewayListeners(t.Context(), test.gateway, nil, newListenerClaims(map[gatev1.PortNumber]*tls.ClientAuth{}), conf)

			require.Len(t, listeners, 1)
			assert.Equal(t, test.expectedAttached, listeners[0].Attached)

			// LastTransitionTime is set with metav1.Now() and is not relevant to this test.
			for i := range listeners[0].Status.Conditions {
				listeners[0].Status.Conditions[i].LastTransitionTime = metav1.Time{}
			}
			assert.Equal(t, test.expectedConditions, listeners[0].Status.Conditions)
		})
	}
}

func Test_loadGatewayListeners_frontendValidationConflict(t *testing.T) {
	k8sObjects, gwObjects := readResources(t, []string{"frontend_validation_ca_refs.yml"})

	kubeClient := kubefake.NewClientset(k8sObjects...)
	gwClient := newGatewaySimpleClientSet(t, gwObjects...)

	client := newClientImpl(kubeClient, gwClient)

	eventCh, err := client.WatchAll(nil, make(chan struct{}))
	require.NoError(t, err)
	<-eventCh

	p := Provider{
		EntryPoints: map[string]Entrypoint{"websecure": {Address: ":443"}},
		client:      client,
	}

	newValidation := func(caConfigMap string, mode gatev1.FrontendValidationModeType) *gatev1.FrontendTLSValidation {
		return &gatev1.FrontendTLSValidation{
			Mode: mode,
			CACertificateRefs: []gatev1.ObjectReference{
				{Kind: "ConfigMap", Name: gatev1.ObjectName(caConfigMap)},
			},
		}
	}

	newGateway := func(name string, hostname *gatev1.Hostname, validation *gatev1.FrontendTLSValidation) *gatev1.Gateway {
		return &gatev1.Gateway{
			Name:       name,
			Namespace:  "default",
			Generation: 1,
			Spec: gatev1.GatewaySpec{
				TLS: &gatev1.GatewayTLSConfig{
					Frontend: &gatev1.FrontendTLSConfig{
						Default: gatev1.TLSConfig{
							Validation: validation,
						},
					},
				},
				Listeners: []gatev1.Listener{
					{
						Name:     "https",
						Protocol: gatev1.HTTPSProtocolType,
						Port:     443,
						Hostname: hostname,
						TLS: &gatev1.ListenerTLSConfig{
							CertificateRefs: []gatev1.SecretObjectReference{
								{Kind: new(gatev1.Kind("Secret")), Group: new(gatev1.Group("")), Name: "tls-secret"},
							},
						},
						AllowedRoutes: &gatev1.AllowedRoutes{
							Namespaces: &gatev1.RouteNamespaces{From: new(gatev1.NamespacesFromSame)},
						},
					},
				},
			},
		}
	}

	conflictedConditions := []metav1.Condition{
		{
			Type:               string(gatev1.ListenerConditionAccepted),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: 1,
			Reason:             string(listenerReasonFrontendValidationConflict),
			Message:            "An older Gateway already applies a different frontend validation on this port",
		},
		{
			Type:               string(gatev1.ListenerConditionProgrammed),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: 1,
			Reason:             string(listenerReasonFrontendValidationConflict),
			Message:            "An older Gateway already applies a different frontend validation on this port",
		},
		{
			Type:               string(gatev1.ListenerConditionConflicted),
			Status:             metav1.ConditionTrue,
			ObservedGeneration: 1,
			Reason:             string(listenerReasonFrontendValidationConflict),
			Message:            "An older Gateway already applies a different frontend validation on this port",
		},
	}

	testCases := []struct {
		desc               string
		hostname           *gatev1.Hostname
		validation         *gatev1.FrontendTLSValidation
		expectedAttached   bool
		expectedConditions []metav1.Condition
	}{
		{
			desc:               "the same validation is shared",
			validation:         newValidation("ca-configmap", gatev1.AllowValidOnly),
			expectedAttached:   true,
			expectedConditions: []metav1.Condition{},
		},
		{
			desc:               "the same validation is shared on another hostname",
			hostname:           new(gatev1.Hostname("foo.com")),
			validation:         newValidation("ca-configmap", gatev1.AllowValidOnly),
			expectedAttached:   true,
			expectedConditions: []metav1.Condition{},
		},
		{
			desc:               "a different CA is conflicted",
			validation:         newValidation("ca-configmap-2", gatev1.AllowValidOnly),
			expectedConditions: conflictedConditions,
		},
		{
			desc:               "a different CA on another hostname is conflicted",
			hostname:           new(gatev1.Hostname("foo.com")),
			validation:         newValidation("ca-configmap-2", gatev1.AllowValidOnly),
			expectedConditions: conflictedConditions,
		},
		{
			desc:               "a different mode is conflicted",
			validation:         newValidation("ca-configmap", gatev1.AllowInsecureFallback),
			expectedConditions: conflictedConditions,
		},
		{
			desc:               "no validation is conflicted",
			validation:         nil,
			expectedConditions: conflictedConditions,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			frontendValidations := map[gatev1.PortNumber]*tls.ClientAuth{}
			conf := &dynamic.Configuration{TLS: &dynamic.TLSConfiguration{}}

			gateway := newGateway("my-gateway", nil, newValidation("ca-configmap", gatev1.AllowValidOnly))
			listeners := p.loadGatewayListeners(t.Context(), gateway, nil, newListenerClaims(frontendValidations), conf)
			require.Len(t, listeners, 1)
			require.True(t, listeners[0].Attached)

			gateway = newGateway("my-other-gateway", test.hostname, test.validation)
			listeners = p.loadGatewayListeners(t.Context(), gateway, nil, newListenerClaims(frontendValidations), conf)
			require.Len(t, listeners, 1)
			assert.Equal(t, test.expectedAttached, listeners[0].Attached)

			// LastTransitionTime is set with metav1.Now() and is not relevant to this test.
			for i := range listeners[0].Status.Conditions {
				listeners[0].Status.Conditions[i].LastTransitionTime = metav1.Time{}
			}
			assert.Equal(t, test.expectedConditions, listeners[0].Status.Conditions)

			// The certificate of a conflicted listener does not reach the TLS store.
			expectedCertificates := 1
			if test.expectedAttached {
				expectedCertificates = 2
			}
			assert.Len(t, conf.TLS.Certificates, expectedCertificates)
		})
	}
}

func Test_isCrossProviderNamespaceAllowed(t *testing.T) {
	testCases := []struct {
		desc      string
		allowList []string
		namespace string
		want      bool
	}{
		{desc: "nil allowList allows any namespace", allowList: nil, namespace: "ns-a", want: true},
		{desc: "empty allowList denies every namespace", allowList: []string{}, namespace: "ns-a", want: false},
		{desc: "namespace in allowList is accepted", allowList: []string{"ns-a"}, namespace: "ns-a", want: true},
		{desc: "namespace not in allowList is rejected", allowList: []string{"ns-b"}, namespace: "ns-a", want: false},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()
			got := isCrossProviderNamespaceAllowed(test.allowList, test.namespace)
			assert.Equal(t, test.want, got)
		})
	}
}

// Test_makeListenerRouterName checks that hostnames normalizing to the same label
// do not collide, which would collapse the SNI scope they are meant to isolate.
func Test_makeListenerRouterName(t *testing.T) {
	uniq := uniqListener{epName: "web", protocol: gatev1.HTTPProtocolType}

	first := makeListenerRouterName(uniq, "*.example.com")
	second := makeListenerRouterName(uniq, "example.com")

	assert.NotEqual(t, first, second)
}

// TestCrossProviderNamespaces_HTTPRoute verifies that the
// CrossProviderNamespaces option gates `@otherProvider` TraefikService
// backendRefs declared on a Gateway HTTPRoute. The check is anchored on the
// HTTPRoute's namespace; when the route is rejected, the whole router is
// dropped from the dynamic configuration.
func TestCrossProviderNamespaces_HTTPRoute(t *testing.T) {
	testCases := []struct {
		desc                    string
		fixture                 string
		crossProviderNamespaces []string
		wantError               bool
	}{
		{desc: "nil: cross-provider TraefikService backendRefs accepted (backward compatible)", fixture: "httproute/simple_cross_provider.yml", crossProviderNamespaces: nil, wantError: false},
		{desc: "empty list: cross-provider TraefikService backendRefs are rejected, route dropped", fixture: "httproute/simple_cross_provider.yml", crossProviderNamespaces: []string{}, wantError: true},
		{desc: "namespace allowed: cross-provider TraefikService backendRefs accepted", fixture: "httproute/simple_cross_provider.yml", crossProviderNamespaces: []string{"default"}, wantError: false},
		{desc: "namespace not allowed: cross-provider TraefikService backendRefs rejected, route dropped", fixture: "httproute/simple_cross_provider.yml", crossProviderNamespaces: []string{"other"}, wantError: true},
		{desc: "namespace provided with cross-provider backendRef, route dropped", fixture: "httproute/invalid_cross_provider.yml", crossProviderNamespaces: []string{"other"}, wantError: true},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			k8sObjects, gwObjects := readResources(t, []string{"services.yml", test.fixture})

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints:             map[string]Entrypoint{"web": {Address: ":80"}},
				CrossProviderNamespaces: test.crossProviderNamespaces,
				client:                  client,
			}

			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			router, ok := conf.HTTP.Routers["httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-799bcbf0c2317c5d1c93"]
			require.True(t, ok)

			service, ok := conf.HTTP.Services[router.Service]
			require.True(t, ok)
			require.NotNil(t, service.Weighted)
			require.Len(t, service.Weighted.Services, 2)

			var hasError bool
			for _, wrrService := range service.Weighted.Services {
				// Whenever a service fails to be loaded, a placeholder service is added to the WRR to server a 500 status code.
				if wrrService.Status != nil && *wrrService.Status == http.StatusInternalServerError {
					hasError = true
					break
				}
			}

			assert.Equal(t, test.wantError, hasError)
		})
	}
}

// TestCrossProviderNamespaces_TCPRoute verifies that the option also gates
// cross-provider TraefikService backendRefs declared on a Gateway TCPRoute.
func TestCrossProviderNamespaces_TCPRoute(t *testing.T) {
	testCases := []struct {
		desc                    string
		fixture                 string
		crossProviderNamespaces []string
		wantError               bool
	}{
		{desc: "nil: cross-provider TraefikService backendRefs accepted (backward compatible)", fixture: "tcproute/simple_cross_provider.yml", crossProviderNamespaces: nil, wantError: false},
		{desc: "empty list: cross-provider TraefikService backendRefs are rejected, route dropped", fixture: "tcproute/simple_cross_provider.yml", crossProviderNamespaces: []string{}, wantError: true},
		{desc: "namespace allowed: cross-provider TraefikService backendRefs accepted", fixture: "tcproute/simple_cross_provider.yml", crossProviderNamespaces: []string{"default"}, wantError: false},
		{desc: "namespace not allowed: cross-provider TraefikService backendRefs rejected, route dropped", fixture: "tcproute/simple_cross_provider.yml", crossProviderNamespaces: []string{"other"}, wantError: true},
		{desc: "namespace provided with cross-provider backendRef, route dropped", fixture: "tcproute/invalid_cross_provider.yml", crossProviderNamespaces: []string{"other"}, wantError: true},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			k8sObjects, gwObjects := readResources(t, []string{"services.yml", test.fixture})

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints:             map[string]Entrypoint{"tcp": {Address: ":9000"}},
				CrossProviderNamespaces: test.crossProviderNamespaces,
				client:                  client,
			}

			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			router, ok := conf.TCP.Routers["tcproute-default-tcp-app-1-gw-default-my-gateway-ep-tcp-0-5840f3c3c21dcf0402da"]
			require.True(t, ok)

			service, ok := conf.TCP.Services[router.Service]
			require.True(t, ok)
			require.NotNil(t, service.Weighted)
			require.Len(t, service.Weighted.Services, 2)

			var hasError bool
			for _, wrrService := range service.Weighted.Services {
				if strings.Contains(wrrService.Name, "@") {
					continue
				}

				lbService, ok := conf.TCP.Services[wrrService.Name]
				require.True(t, ok)
				require.NotNil(t, lbService)
				require.NotNil(t, lbService.LoadBalancer)

				if len(lbService.LoadBalancer.Servers) == 0 {
					hasError = true
				}
			}

			assert.Equal(t, test.wantError, hasError)
		})
	}
}

// TestCrossProviderNamespaces_TLSRoute verifies that the option also gates
// cross-provider TraefikService backendRefs declared on a Gateway TLSRoute.
func TestCrossProviderNamespaces_TLSRoute(t *testing.T) {
	testCases := []struct {
		desc                    string
		fixture                 string
		crossProviderNamespaces []string
		wantError               bool
	}{
		{desc: "nil: cross-provider TraefikService backendRefs accepted (backward compatible)", fixture: "tlsroute/simple_cross_provider.yml", crossProviderNamespaces: nil, wantError: false},
		{desc: "empty list: cross-provider TraefikService backendRefs are rejected, route dropped", fixture: "tlsroute/simple_cross_provider.yml", crossProviderNamespaces: []string{}, wantError: true},
		{desc: "namespace allowed: cross-provider TraefikService backendRefs accepted", fixture: "tlsroute/simple_cross_provider.yml", crossProviderNamespaces: []string{"default"}, wantError: false},
		{desc: "namespace not allowed: cross-provider TraefikService backendRefs rejected, route dropped", fixture: "tlsroute/simple_cross_provider.yml", crossProviderNamespaces: []string{"other"}, wantError: true},
		{desc: "namespace provided with cross-provider backendRef, route dropped", fixture: "tlsroute/invalid_cross_provider.yml", crossProviderNamespaces: []string{"other"}, wantError: true},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			k8sObjects, gwObjects := readResources(t, []string{"services.yml", test.fixture})

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t, gwObjects...)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 || len(gwObjects) > 0 {
				// just wait for the first event
				<-eventCh
			}

			p := Provider{
				EntryPoints:             map[string]Entrypoint{"tls": {Address: ":9000"}},
				CrossProviderNamespaces: test.crossProviderNamespaces,
				client:                  client,
			}

			conf, _, err := p.loadConfigurationFromGateways(t.Context())
			assert.NoError(t, err)

			router, ok := conf.TCP.Routers["tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-efc0f8d903baa9fb6592"]
			require.True(t, ok)

			service, ok := conf.TCP.Services[router.Service]
			require.True(t, ok)
			require.NotNil(t, service.Weighted)
			require.Len(t, service.Weighted.Services, 2)

			var hasError bool
			for _, wrrService := range service.Weighted.Services {
				if strings.Contains(wrrService.Name, "@") {
					continue
				}

				lbService, ok := conf.TCP.Services[wrrService.Name]
				require.True(t, ok)
				require.NotNil(t, lbService)
				require.NotNil(t, lbService.LoadBalancer)

				if len(lbService.LoadBalancer.Servers) == 0 {
					hasError = true
				}
			}

			assert.Equal(t, test.wantError, hasError)
		})
	}
}

func findCondition(conditions []metav1.Condition, conditionType gatev1.RouteConditionType) *metav1.Condition {
	for i := range conditions {
		if conditions[i].Type == string(conditionType) {
			return &conditions[i]
		}
	}
	return nil
}

// We cannot use the gateway-api fake.NewClientset due to Gateway being pluralized as "gatewaies" instead of "gateways".
func newGatewaySimpleClientSet(t *testing.T, objects ...runtime.Object) *gatefake.Clientset {
	t.Helper()

	client := gatefake.NewSimpleClientset(objects...)
	for _, object := range objects {
		gateway, ok := object.(*gatev1.Gateway)
		if !ok {
			continue
		}

		_, err := client.GatewayV1().Gateways(gateway.Namespace).Create(t.Context(), gateway, metav1.CreateOptions{})
		require.NoError(t, err)
	}

	return client
}

func Test_loadConfigurationFromGateways_TLSCertificatesOrder(t *testing.T) {
	k8sObjects, gwObjects := readResources(t, []string{"multiple_gateways_tls_certificates.yml"})

	client := newClientImpl(kubefake.NewClientset(k8sObjects...), newGatewaySimpleClientSet(t, gwObjects...))

	eventCh, err := client.WatchAll(nil, make(chan struct{}))
	require.NoError(t, err)

	// just wait for the first event
	<-eventCh

	p := Provider{
		EntryPoints: map[string]Entrypoint{"websecure": {Address: ":443"}},
		client:      client,
	}

	// Gateways are listed from a map, so build the configuration several times to
	// make sure the certificates do not follow the iteration order.
	conf, _, err := p.loadConfigurationFromGateways(t.Context())
	require.NoError(t, err)

	var commonNames []string
	for _, cert := range conf.TLS.Certificates {
		commonNames = append(commonNames, certificateCommonName(t, cert.Certificate.CertFile))
	}

	assert.Equal(t, []string{"bar.example.com", "foo.example.com", "default-b.example.com", "default-a.example.com"}, commonNames)
}

func Test_loadConfigurationFromGateways_BackendTLSPolicyAncestorNamespace(t *testing.T) {
	testCases := []struct {
		desc        string
		path        string
		entryPoints map[string]Entrypoint
		sectionName gatev1.SectionName
	}{
		{
			desc:        "HTTPRoute",
			path:        "httproute/with_backend_tls_policy_cross_namespace.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}},
			sectionName: "http",
		},
		{
			desc:        "GRPCRoute",
			path:        "grpcroute/with_backend_tls_policy_cross_namespace.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}},
			sectionName: "web",
		},
		{
			desc:        "TLSRoute",
			path:        "tlsroute/with_backend_tls_policy_cross_namespace.yml",
			entryPoints: map[string]Entrypoint{"tls": {Address: ":9001"}},
			sectionName: "tls",
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			k8sObjects, gwObjects := readResources(t, []string{"services.yml", test.path})

			gwClient := newGatewaySimpleClientSet(t, gwObjects...)
			client := newClientImpl(kubefake.NewClientset(k8sObjects...), gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			// just wait for the first event
			<-eventCh

			p := Provider{
				EntryPoints: test.entryPoints,
				client:      client,
			}

			_, statusReport, err := p.loadConfigurationFromGateways(t.Context())
			require.NoError(t, err)

			statusReport.Flush(t.Context(), client)

			policy, err := gwClient.GatewayV1().BackendTLSPolicies("bar").Get(t.Context(), "policy-1", metav1.GetOptions{})
			require.NoError(t, err)
			require.Len(t, policy.Status.Ancestors, 1)

			assert.Equal(t, gatev1.ParentReference{
				Group:       new(gatev1.Group(groupGateway)),
				Kind:        new(gatev1.Kind(kindGateway)),
				Namespace:   new(gatev1.Namespace("default")),
				Name:        "my-gateway",
				SectionName: new(test.sectionName),
			}, policy.Status.Ancestors[0].AncestorRef)
		})
	}
}

func certificateCommonName(t *testing.T, certFile types.FileOrContent) string {
	t.Helper()

	block, _ := pem.Decode([]byte(certFile))
	require.NotNil(t, block)

	cert, err := x509.ParseCertificate(block.Bytes)
	require.NoError(t, err)

	return cert.Subject.CommonName
}

func readResources(t *testing.T, paths []string) ([]runtime.Object, []runtime.Object) {
	t.Helper()

	var k8sObjects []runtime.Object
	var gwObjects []runtime.Object
	for _, path := range paths {
		yamlContent, err := os.ReadFile(filepath.FromSlash("./fixtures/" + path))
		if err != nil {
			panic(err)
		}

		objects := k8s.MustParseYaml(yamlContent)
		for _, obj := range objects {
			switch obj.GetObjectKind().GroupVersionKind().Group {
			case "gateway.networking.k8s.io":
				gwObjects = append(gwObjects, obj)
			default:
				k8sObjects = append(k8sObjects, obj)
			}
		}
	}

	return k8sObjects, gwObjects
}

func Test_listenerSetRefsGateway(t *testing.T) {
	gw := &gatev1.Gateway{
		Name:      "my-gw",
		Namespace: "default",
	}

	testCases := []struct {
		desc     string
		ls       *gatev1.ListenerSet
		expected bool
	}{
		{
			desc: "Matching reference",
			ls: &gatev1.ListenerSet{
				Name:      "my-ls",
				Namespace: "default",
				Spec: gatev1.ListenerSetSpec{
					ParentRef: gatev1.ParentGatewayReference{
						Name: "my-gw",
					},
				},
			},
			expected: true,
		},
		{
			desc: "Wrong name",
			ls: &gatev1.ListenerSet{
				Name:      "my-ls",
				Namespace: "default",
				Spec: gatev1.ListenerSetSpec{
					ParentRef: gatev1.ParentGatewayReference{
						Name: "other-gw",
					},
				},
			},
			expected: false,
		},
		{
			desc: "Wrong namespace",
			ls: &gatev1.ListenerSet{
				Name:      "my-ls",
				Namespace: "default",
				Spec: gatev1.ListenerSetSpec{
					ParentRef: gatev1.ParentGatewayReference{
						Name:      "my-gw",
						Namespace: new(gatev1.Namespace("other")),
					},
				},
			},
			expected: false,
		},
		{
			desc: "Cross-namespace matching",
			ls: &gatev1.ListenerSet{
				Name:      "my-ls",
				Namespace: "other",
				Spec: gatev1.ListenerSetSpec{
					ParentRef: gatev1.ParentGatewayReference{
						Name:      "my-gw",
						Namespace: new(gatev1.Namespace("default")),
					},
				},
			},
			expected: true,
		},
		{
			desc: "Wrong group",
			ls: &gatev1.ListenerSet{
				Name:      "my-ls",
				Namespace: "default",
				Spec: gatev1.ListenerSetSpec{
					ParentRef: gatev1.ParentGatewayReference{
						Group: new(gatev1.Group("wrong.group")),
						Name:  "my-gw",
					},
				},
			},
			expected: false,
		},
		{
			desc: "Wrong kind",
			ls: &gatev1.ListenerSet{
				Name:      "my-ls",
				Namespace: "default",
				Spec: gatev1.ListenerSetSpec{
					ParentRef: gatev1.ParentGatewayReference{
						Kind: new(gatev1.Kind("NotAGateway")),
						Name: "my-gw",
					},
				},
			},
			expected: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, test.expected, listenerSetRefsGateway(test.ls, gw))
		})
	}
}

func Test_isListenerSetAllowed(t *testing.T) {
	testCases := []struct {
		desc       string
		gw         *gatev1.Gateway
		ls         *gatev1.ListenerSet
		namespaces []*corev1.Namespace
		expected   bool
	}{
		{
			desc: "Nil AllowedListeners",
			gw: &gatev1.Gateway{
				Name: "gw", Namespace: "default",
				Spec: gatev1.GatewaySpec{},
			},
			ls: &gatev1.ListenerSet{
				Name: "ls", Namespace: "default",
			},
			expected: false,
		},
		{
			desc: "Nil From",
			gw: &gatev1.Gateway{
				Name: "gw", Namespace: "default",
				Spec: gatev1.GatewaySpec{
					AllowedListeners: &gatev1.AllowedListeners{
						Namespaces: &gatev1.ListenerNamespaces{},
					},
				},
			},
			ls: &gatev1.ListenerSet{
				Name: "ls", Namespace: "default",
			},
			expected: false,
		},
		{
			desc: "FromNone",
			gw: &gatev1.Gateway{
				Name: "gw", Namespace: "default",
				Spec: gatev1.GatewaySpec{
					AllowedListeners: &gatev1.AllowedListeners{
						Namespaces: &gatev1.ListenerNamespaces{
							From: new(gatev1.NamespacesFromNone),
						},
					},
				},
			},
			ls: &gatev1.ListenerSet{
				Name: "ls", Namespace: "default",
			},
			expected: false,
		},
		{
			desc: "FromSame - same namespace",
			gw: &gatev1.Gateway{
				Name: "gw", Namespace: "default",
				Spec: gatev1.GatewaySpec{
					AllowedListeners: &gatev1.AllowedListeners{
						Namespaces: &gatev1.ListenerNamespaces{
							From: new(gatev1.NamespacesFromSame),
						},
					},
				},
			},
			ls: &gatev1.ListenerSet{
				Name: "ls", Namespace: "default",
			},
			expected: true,
		},
		{
			desc: "FromSame - different namespace",
			gw: &gatev1.Gateway{
				Name: "gw", Namespace: "default",
				Spec: gatev1.GatewaySpec{
					AllowedListeners: &gatev1.AllowedListeners{
						Namespaces: &gatev1.ListenerNamespaces{
							From: new(gatev1.NamespacesFromSame),
						},
					},
				},
			},
			ls: &gatev1.ListenerSet{
				Name: "ls", Namespace: "other",
			},
			expected: false,
		},
		{
			desc: "FromAll",
			gw: &gatev1.Gateway{
				Name: "gw", Namespace: "default",
				Spec: gatev1.GatewaySpec{
					AllowedListeners: &gatev1.AllowedListeners{
						Namespaces: &gatev1.ListenerNamespaces{
							From: new(gatev1.NamespacesFromAll),
						},
					},
				},
			},
			ls: &gatev1.ListenerSet{
				Name: "ls", Namespace: "any-namespace",
			},
			expected: true,
		},
		{
			desc: "FromSelector - matching namespace labels",
			gw: &gatev1.Gateway{
				Name: "gw", Namespace: "default",
				Spec: gatev1.GatewaySpec{
					AllowedListeners: &gatev1.AllowedListeners{
						Namespaces: &gatev1.ListenerNamespaces{
							From: new(gatev1.NamespacesFromSelector),
							Selector: &metav1.LabelSelector{
								MatchLabels: map[string]string{"env": "prod"},
							},
						},
					},
				},
			},
			ls: &gatev1.ListenerSet{
				Name:      "ls",
				Namespace: "default",
			},
			namespaces: []*corev1.Namespace{
				{
					Name:   "default",
					Labels: map[string]string{"env": "prod"},
				},
			},
			expected: true,
		},
		{
			desc: "FromSelector - non-matching namespace labels",
			gw: &gatev1.Gateway{
				Name: "gw", Namespace: "default",
				Spec: gatev1.GatewaySpec{
					AllowedListeners: &gatev1.AllowedListeners{
						Namespaces: &gatev1.ListenerNamespaces{
							From: new(gatev1.NamespacesFromSelector),
							Selector: &metav1.LabelSelector{
								MatchLabels: map[string]string{"env": "prod"},
							},
						},
					},
				},
			},
			ls: &gatev1.ListenerSet{
				Name:      "ls",
				Namespace: "default",
			},
			namespaces: []*corev1.Namespace{
				{
					Name:   "default",
					Labels: map[string]string{"env": "staging"},
				},
			},
			expected: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			var k8sObjects []runtime.Object
			for _, ns := range test.namespaces {
				k8sObjects = append(k8sObjects, ns)
			}

			kubeClient := kubefake.NewClientset(k8sObjects...)
			gwClient := newGatewaySimpleClientSet(t)

			client := newClientImpl(kubeClient, gwClient)

			eventCh, err := client.WatchAll(nil, make(chan struct{}))
			require.NoError(t, err)

			if len(k8sObjects) > 0 {
				<-eventCh
			}

			p := Provider{client: client}

			assert.Equal(t, test.expected, p.isListenerSetAllowed(t.Context(), test.gw, test.ls))
		})
	}
}

func Test_makeListenerSetStatus(t *testing.T) {
	testCases := []struct {
		desc                   string
		listeners              []gatewayListener
		parentAccepted         bool
		wantAccepted           bool
		wantAcceptedStatus     metav1.ConditionStatus
		wantAcceptedReason     string
		wantProgrammedStatus   metav1.ConditionStatus
		wantProgrammedReason   string
		wantListenerEntryCount int
		// wantEntryName, when set, asserts the first listener entry's name and its entry-level Programmed condition,
		// so the per-listener status is verified and not only the entry count.
		wantEntryName             gatev1.SectionName
		wantEntryProgrammedStatus metav1.ConditionStatus
		wantEntryProgrammedReason string
	}{
		{
			desc: "All listeners valid, Gateway accepted",
			listeners: []gatewayListener{
				{
					SectionName: "http",
					ListenerSet: &gatev1.ListenerSet{Namespace: "default", Name: "my-ls"},
					Status: &gatev1.ListenerStatus{
						Name:           "http",
						SupportedKinds: []gatev1.RouteGroupKind{{Kind: "HTTPRoute", Group: new(gatev1.Group(gatev1.GroupName))}},
						Conditions:     []metav1.Condition{}, // No errors.
					},
				},
			},
			parentAccepted:            true,
			wantAccepted:              true,
			wantAcceptedStatus:        metav1.ConditionTrue,
			wantAcceptedReason:        string(gatev1.ListenerSetReasonAccepted),
			wantProgrammedStatus:      metav1.ConditionTrue,
			wantProgrammedReason:      string(gatev1.ListenerSetReasonProgrammed),
			wantListenerEntryCount:    1,
			wantEntryName:             "http",
			wantEntryProgrammedStatus: metav1.ConditionTrue,
			wantEntryProgrammedReason: string(gatev1.ListenerEntryReasonProgrammed),
		},
		{
			// A parent Gateway is not accepted only when no listener serving it is valid, the ListenerSet ones included.
			desc: "Gateway not accepted",
			listeners: []gatewayListener{
				{
					SectionName: "https",
					ListenerSet: &gatev1.ListenerSet{Namespace: "default", Name: "my-ls"},
					Status: &gatev1.ListenerStatus{
						Name: "https",
						Conditions: []metav1.Condition{
							{Type: string(gatev1.ListenerConditionResolvedRefs), Status: metav1.ConditionFalse, Reason: string(gatev1.ListenerReasonInvalidCertificateRef)},
							{Type: string(gatev1.ListenerConditionProgrammed), Status: metav1.ConditionFalse, Reason: string(gatev1.ListenerReasonInvalid)},
						},
					},
				},
			},
			parentAccepted:            false,
			wantAcceptedStatus:        metav1.ConditionFalse,
			wantAcceptedReason:        string(gatev1.ListenerSetReasonParentNotAccepted),
			wantProgrammedStatus:      metav1.ConditionFalse,
			wantProgrammedReason:      "ParentNotProgrammed",
			wantListenerEntryCount:    1,
			wantEntryName:             "https",
			wantEntryProgrammedStatus: metav1.ConditionFalse,
			wantEntryProgrammedReason: string(gatev1.ListenerEntryReasonInvalid),
		},
		{
			desc: "No valid listener",
			listeners: []gatewayListener{
				{
					SectionName: "http",
					ListenerSet: &gatev1.ListenerSet{Namespace: "default", Name: "my-ls"},
					Status: &gatev1.ListenerStatus{
						Name: "http",
						Conditions: []metav1.Condition{
							{
								Type:   string(gatev1.ListenerConditionAccepted),
								Status: metav1.ConditionFalse,
								Reason: string(gatev1.ListenerReasonPortUnavailable),
							},
						},
					},
				},
			},
			parentAccepted:         true,
			wantAcceptedStatus:     metav1.ConditionFalse,
			wantAcceptedReason:     string(gatev1.ListenerSetReasonListenersNotValid),
			wantProgrammedStatus:   metav1.ConditionFalse,
			wantProgrammedReason:   string(gatev1.ListenerSetReasonListenersNotValid),
			wantListenerEntryCount: 1,
		},
		{
			desc: "At least one valid listener",
			listeners: []gatewayListener{
				{
					SectionName: "http",
					ListenerSet: &gatev1.ListenerSet{Namespace: "default", Name: "my-ls"},
					Status: &gatev1.ListenerStatus{
						Name:       "http",
						Conditions: []metav1.Condition{},
					},
				},
				{
					SectionName: "invalid",
					ListenerSet: &gatev1.ListenerSet{Namespace: "default", Name: "my-ls"},
					Status: &gatev1.ListenerStatus{
						Name: "invalid",
						Conditions: []metav1.Condition{
							{
								Type:   string(gatev1.ListenerConditionAccepted),
								Status: metav1.ConditionFalse,
								Reason: string(gatev1.ListenerReasonPortUnavailable),
							},
						},
					},
				},
			},
			parentAccepted:         true,
			wantAccepted:           true,
			wantAcceptedStatus:     metav1.ConditionTrue,
			wantAcceptedReason:     string(gatev1.ListenerSetReasonListenersNotValid),
			wantProgrammedStatus:   metav1.ConditionTrue,
			wantProgrammedReason:   string(gatev1.ListenerSetReasonProgrammed),
			wantListenerEntryCount: 2,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			status, accepted := makeListenerSetStatus(0, test.listeners, test.parentAccepted)

			assert.Equal(t, test.wantAccepted, accepted)
			assert.Len(t, status.Listeners, test.wantListenerEntryCount)

			var acceptedCond, programmedCond *metav1.Condition
			for _, c := range status.Conditions {
				switch c.Type {
				case string(gatev1.ListenerSetConditionAccepted):
					acceptedCond = &c
				case string(gatev1.ListenerSetConditionProgrammed):
					programmedCond = &c
				}
			}

			require.NotNil(t, acceptedCond)
			assert.Equal(t, test.wantAcceptedStatus, acceptedCond.Status)
			assert.Equal(t, test.wantAcceptedReason, acceptedCond.Reason)

			require.NotNil(t, programmedCond)
			assert.Equal(t, test.wantProgrammedStatus, programmedCond.Status)
			assert.Equal(t, test.wantProgrammedReason, programmedCond.Reason)

			if test.wantEntryName != "" {
				require.NotEmpty(t, status.Listeners)
				entry := status.Listeners[0]
				assert.Equal(t, test.wantEntryName, entry.Name)

				var entryProgrammed *metav1.Condition
				for _, c := range entry.Conditions {
					if c.Type == string(gatev1.ListenerEntryConditionProgrammed) {
						entryProgrammed = &c
					}
				}
				require.NotNil(t, entryProgrammed)
				assert.Equal(t, test.wantEntryProgrammedStatus, entryProgrammed.Status)
				assert.Equal(t, test.wantEntryProgrammedReason, entryProgrammed.Reason)
			}
		})
	}
}

// Test_loadListenerSetListeners covers the listener-level outcomes of merging the ListenerSet listeners into the Gateway ones:
// the conflict conditions, the precedence of a ListenerSet over its later siblings,
// and the status of the ListenerSets the Gateway does not allow.
func Test_loadListenerSetListeners(t *testing.T) {
	gateway := &gatev1.Gateway{
		Name: "my-gateway", Namespace: "default",
		Spec: gatev1.GatewaySpec{
			AllowedListeners: &gatev1.AllowedListeners{
				Namespaces: &gatev1.ListenerNamespaces{From: new(gatev1.NamespacesFromAll)},
			},
		},
	}

	listenerSet := func(name string, created time.Time, listeners ...gatev1.ListenerEntry) *gatev1.ListenerSet {
		return &gatev1.ListenerSet{
			Name:              name,
			Namespace:         "default",
			CreationTimestamp: metav1.NewTime(created),
			Spec: gatev1.ListenerSetSpec{
				ParentRef: gatev1.ParentGatewayReference{Name: "my-gateway"},
				Listeners: listeners,
			},
		}
	}

	type wantListener struct {
		name     string
		attached bool
		reason   string
	}

	testCases := []struct {
		desc             string
		allowedListeners *gatev1.AllowedListeners
		gatewayListener  *gatev1.Listener
		listenerSets     []*gatev1.ListenerSet
		wantListeners    []wantListener
		wantNotAllowed   []string
	}{
		{
			desc:             "ListenerSet not allowed by the Gateway",
			allowedListeners: &gatev1.AllowedListeners{Namespaces: &gatev1.ListenerNamespaces{From: new(gatev1.NamespacesFromNone)}},
			listenerSets: []*gatev1.ListenerSet{
				listenerSet("my-listenerset", time.Time{}, gatev1.ListenerEntry{Name: "http", Protocol: gatev1.HTTPProtocolType, Port: 80}),
			},
			wantNotAllowed: []string{"my-listenerset"},
		},
		{
			desc:            "Hostname conflict with a Gateway listener",
			gatewayListener: &gatev1.Listener{Name: "web", Protocol: gatev1.HTTPProtocolType, Port: 80},
			listenerSets: []*gatev1.ListenerSet{
				listenerSet("my-listenerset", time.Time{}, gatev1.ListenerEntry{Name: "http", Protocol: gatev1.HTTPProtocolType, Port: 80}),
			},
			wantListeners: []wantListener{{name: "http", reason: string(gatev1.ListenerReasonHostnameConflict)}},
		},
		{
			desc:            "Protocol conflict with a Gateway listener",
			gatewayListener: &gatev1.Listener{Name: "web", Protocol: gatev1.HTTPProtocolType, Port: 80},
			listenerSets: []*gatev1.ListenerSet{
				listenerSet("my-listenerset", time.Time{},
					gatev1.ListenerEntry{Name: "http", Protocol: gatev1.HTTPProtocolType, Port: 80, Hostname: new(gatev1.Hostname("foo.example.com"))},
					gatev1.ListenerEntry{Name: "tcp", Protocol: gatev1.TCPProtocolType, Port: 80},
				),
			},
			wantListeners: []wantListener{
				{name: "http", attached: true},
				{name: "tcp", reason: string(gatev1.ListenerReasonProtocolConflict)},
			},
		},
		{
			desc: "First ListenerSet wins the protocol conflict with a later sibling",
			listenerSets: []*gatev1.ListenerSet{
				listenerSet("ls-first", time.Time{}, gatev1.ListenerEntry{Name: "tcp", Protocol: gatev1.TCPProtocolType, Port: 80}),
				listenerSet("ls-second", time.Time{}, gatev1.ListenerEntry{Name: "http", Protocol: gatev1.HTTPProtocolType, Port: 80}),
			},
			wantListeners: []wantListener{
				{name: "tcp", attached: true},
				{name: "http", reason: string(gatev1.ListenerReasonProtocolConflict)},
			},
		},
		{
			desc: "First ListenerSet wins the conflict with a later sibling",
			listenerSets: []*gatev1.ListenerSet{
				listenerSet("ls-first", time.Time{}, gatev1.ListenerEntry{Name: "http", Protocol: gatev1.HTTPProtocolType, Port: 80}),
				listenerSet("ls-second", time.Time{}, gatev1.ListenerEntry{Name: "http", Protocol: gatev1.HTTPProtocolType, Port: 80}),
			},
			wantListeners: []wantListener{
				{name: "http", attached: true},
				{name: "http", reason: string(gatev1.ListenerReasonHostnameConflict)},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {
			t.Parallel()

			p := Provider{EntryPoints: map[string]Entrypoint{"web": {Address: ":80"}}}

			gateway := gateway.DeepCopy()
			if test.allowedListeners != nil {
				gateway.Spec.AllowedListeners = test.allowedListeners
			}

			claims := newListenerClaims(map[gatev1.PortNumber]*tls.ClientAuth{})
			if test.gatewayListener != nil {
				gatewayWithListener := gateway.DeepCopy()
				gatewayWithListener.Spec.Listeners = []gatev1.Listener{*test.gatewayListener}
				p.loadGatewayListeners(t.Context(), gatewayWithListener, nil, claims, &dynamic.Configuration{TLS: &dynamic.TLSConfiguration{}})
			}

			statusReport := newStatusReport()
			listeners := p.loadListenerSetListeners(t.Context(), gateway, test.listenerSets, claims, &dynamic.Configuration{TLS: &dynamic.TLSConfiguration{}}, statusReport)
			require.Len(t, listeners, len(test.wantListeners))

			require.Len(t, statusReport.listenerSets, len(test.wantNotAllowed))
			for _, name := range test.wantNotAllowed {
				status, ok := statusReport.listenerSets[ktypes.NamespacedName{Namespace: "default", Name: name}]
				require.True(t, ok)
				require.Len(t, status.Conditions, 2)
				for _, condition := range status.Conditions {
					assert.Equal(t, string(gatev1.ListenerSetReasonNotAllowed), condition.Reason)
				}
			}

			for i, want := range test.wantListeners {
				assert.Equal(t, want.name, listeners[i].SectionName)
				assert.Equal(t, want.attached, listeners[i].Attached)

				if want.reason == "" {
					assert.Empty(t, listeners[i].Status.Conditions)
					continue
				}

				require.NotEmpty(t, listeners[i].Status.Conditions)
				for _, condition := range listeners[i].Status.Conditions {
					assert.Equal(t, want.reason, condition.Reason)
				}
			}
		})
	}
}

```

#### pkg/tlsroute.go

```go
package gateway

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/provider"
	"github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	"k8s.io/utils/ptr"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func (p *Provider) loadTLSRoutes(ctx context.Context, gateways []gatewayWithListeners, conf *dynamic.Configuration, statusReport *statusReport) {
	logger := log.Ctx(ctx)
	routes, err := p.client.ListTLSRoutes()
	if err != nil {
		logger.Error().Err(err).Msgf("Unable to list TLSRoute")
		return
	}

	// The provider loads the routes in the order that the specification gives,
	// to make a decision between the routes that match a connection equally well.
	slices.SortStableFunc(routes, func(a, b *gatev1.TLSRoute) int { return compareRoutes(a, b) })

	served := make(servedRules)

	for _, route := range routes {
		logger := log.Ctx(ctx).With().
			Str("tls_route", route.Name).
			Str("namespace", route.Namespace).Logger()

		routeParentRefs := matchingGatewayListenersForParentRef(gateways, route.Namespace, route.Spec.ParentRefs)
		if len(routeParentRefs) == 0 {
			continue
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

				if accepted && !allowRoute(listener, route.Namespace, kindTLSRoute) {
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

				if accepted {
					listener.Status.AttachedRoutes++
				}

				// The ResolvedRefs condition must be reported for every parentRef,
				// even when the route does not attach to the listener.
				routerConfs, condition := p.loadTLSRoute(match.GatewayName, match.GatewayNamespace, listener, route, hostnames, statusReport)
				if resolvedRefCondition == nil || resolvedRefCondition.Status == metav1.ConditionTrue {
					resolvedRefCondition = new(condition)
				}

				if accepted && listener.Attached {
					for _, rc := range routerConfs {
						router := rc.Conf.TCP.Routers[rc.Name]
						if servedBy, alreadyServed := served.register(rc.Name, router.EntryPoints, router.Rule); alreadyServed {
							logger.Warn().Msgf("Traefik does not create router %q, because router %q serves the rule %q on the same entry points", rc.Name, servedBy, router.Rule)
							continue
						}

						mergeTCPConfiguration(rc.Conf, conf)
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

			statusReport.RecordTLSRouteStatus(ktypes.NamespacedName{Namespace: route.Namespace, Name: route.Name}, gatev1.RouteParentStatus{
				ParentRef:      match.ParentRef,
				ControllerName: controllerName,
				Conditions:     parentStatusConditions,
			})
		}

		// When there is at least one TLS listener, we add a default deny-all route to avoid accepting traffic for undefined hosts.
		// Note that when there is HTTPS listeners this will predate the traffic and reject the connection to undefined hosts instead of returning a 404.
		if len(conf.TCP.Routers) > 0 {
			conf.TCP.Routers["deny-unknown-host"] = &dynamic.TCPRouter{
				Rule:     "HostSNI(`*`) && !ALPN(`h2`) && !ALPN(`http/1.1`)",
				Priority: 1,
				Service:  "deny-unknown-host",
				TLS:      &dynamic.RouterTCPTLSConfig{},
			}
			conf.TCP.Services["deny-unknown-host"] = &dynamic.TCPService{
				LoadBalancer: &dynamic.TCPServersLoadBalancer{},
			}
		}
	}
}

// tlsRouteRouter is a router of a TLSRoute, with the configuration of its resources.
type tlsRouteRouter struct {
	Name string
	Conf *dynamic.Configuration
}

func (p *Provider) loadTLSRoute(gatewayName, gatewayNamespace string, listener gatewayListener, route *gatev1.TLSRoute, hostnames []gatev1.Hostname, statusReport *statusReport) ([]tlsRouteRouter, metav1.Condition) {
	var routers []tlsRouteRouter

	condition := metav1.Condition{
		Type:               string(gatev1.RouteConditionResolvedRefs),
		Status:             metav1.ConditionTrue,
		ObservedGeneration: route.Generation,
		LastTransitionTime: metav1.Now(),
		Reason:             string(gatev1.RouteConditionResolvedRefs),
	}

	for ri, routeRule := range route.Spec.Rules {
		if len(routeRule.BackendRefs) == 0 {
			// Should not happen due to validation.
			// https://github.com/kubernetes-sigs/gateway-api/blob/v0.4.0/apis/v1alpha2/tlsroute_types.go#L120
			continue
		}

		rule, priority := hostSNIRule(hostnames)
		router := dynamic.TCPRouter{
			RuleSyntax:  "default",
			Rule:        rule,
			Priority:    priority,
			EntryPoints: []string{listener.EPName},
			TLS: &dynamic.RouterTCPTLSConfig{
				Passthrough: listener.TLS != nil && listener.TLS.Mode != nil && *listener.TLS.Mode == gatev1.TLSModePassthrough,
			},
		}

		// The rule tells apart the routers of the listeners sharing the entry point,
		// as they differ by the hostnames they match.
		routerName := makeRouterName(strings.ToLower(kindTLSRoute), rule, route.Namespace, route.Name, gatewayNamespace, gatewayName, listener.EPName, ri)

		routerConf := &dynamic.Configuration{
			TCP: &dynamic.TCPConfiguration{
				Routers:           make(map[string]*dynamic.TCPRouter),
				Middlewares:       make(map[string]*dynamic.TCPMiddleware),
				Services:          make(map[string]*dynamic.TCPService),
				ServersTransports: make(map[string]*dynamic.TCPServersTransport),
			},
		}

		if len(routeRule.BackendRefs) == 1 && isInternalService(routeRule.BackendRefs[0]) {
			if !isCrossProviderNamespaceAllowed(p.CrossProviderNamespaces, route.Namespace) {
				condition = metav1.Condition{
					Type:               string(gatev1.RouteConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: route.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.RouteReasonRefNotPermitted),
					Message:            fmt.Sprintf("Cannot load TLSRoute BackendRef %s: internal service reference is not allowed: TLSRoute namespace %q is not in crossProviderNamespaces", routeRule.BackendRefs[0].Name, route.Namespace),
				}

				continue
			}

			router.Service = string(routeRule.BackendRefs[0].Name)
			routerConf.TCP.Routers[routerName] = &router
			routers = append(routers, tlsRouteRouter{Name: routerName, Conf: routerConf})
			continue
		}

		var serviceCondition *metav1.Condition
		router.Service, serviceCondition = p.loadTLSWRRService(gatewayName, gatewayNamespace, listener, routerConf, routerName, routeRule.BackendRefs, route, statusReport)
		if serviceCondition != nil {
			condition = *serviceCondition
		}

		routerConf.TCP.Routers[routerName] = &router
		routers = append(routers, tlsRouteRouter{Name: routerName, Conf: routerConf})
	}

	return routers, condition
}

// loadTLSWRRService is generating a WRR service, even when there is only one target.
func (p *Provider) loadTLSWRRService(gatewayName, gatewayNamespace string, listener gatewayListener, conf *dynamic.Configuration, routerName string, backendRefs []gatev1.BackendRef, route *gatev1.TLSRoute, statusReport *statusReport) (string, *metav1.Condition) {
	name := routerName + "-wrr"
	if _, ok := conf.TCP.Services[name]; ok {
		return name, nil
	}

	var wrr dynamic.TCPWeightedRoundRobin
	var condition *metav1.Condition
	for bi, backendRef := range backendRefs {
		svcName, svc, errCondition := p.loadTLSService(gatewayName, gatewayNamespace, listener, conf, routerName, route, bi, backendRef, statusReport)
		weight := new(int(ptr.Deref(backendRef.Weight, 1)))

		if errCondition != nil {
			condition = errCondition

			errName := routerName + "-err-lb"
			conf.TCP.Services[errName] = &dynamic.TCPService{
				LoadBalancer: &dynamic.TCPServersLoadBalancer{
					Servers: []dynamic.TCPServer{},
				},
			}

			wrr.Services = append(wrr.Services, dynamic.TCPWRRService{
				Name:   errName,
				Weight: weight,
			})
			continue
		}

		if svc != nil {
			conf.TCP.Services[svcName] = svc
		}

		wrr.Services = append(wrr.Services, dynamic.TCPWRRService{
			Name:   svcName,
			Weight: weight,
		})
	}

	conf.TCP.Services[name] = &dynamic.TCPService{Weighted: &wrr}
	return name, condition
}

func (p *Provider) loadTLSService(gatewayName, gatewayNamespace string, listener gatewayListener, conf *dynamic.Configuration, routerName string, route *gatev1.TLSRoute, backendIndex int, backendRef gatev1.BackendRef, statusReport *statusReport) (string, *dynamic.TCPService, *metav1.Condition) {
	kind := ptr.Deref(backendRef.Kind, kindService)

	group := groupCore
	if backendRef.Group != nil && *backendRef.Group != "" {
		group = string(*backendRef.Group)
	}

	namespace := route.Namespace
	if backendRef.Namespace != nil && *backendRef.Namespace != "" {
		namespace = string(*backendRef.Namespace)

		if strings.Contains(string(backendRef.Name), "@") {
			return provider.Normalize(fmt.Sprintf("%s-svc-%s-%s-%d", routerName, namespace, backendRef.Name, backendIndex)), nil, &metav1.Condition{
				Type:               string(gatev1.RouteConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: route.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.RouteReasonRefNotPermitted),
				Message:            fmt.Sprintf("Cannot load TLSRoute BackendRef %s/%s/%s/%s: namespace is not allowed with a cross-provider reference", group, kind, namespace, backendRef.Name),
			}
		}
	}

	serviceName := fmt.Sprintf("%s-svc-%s-%s-%d", routerName, namespace, backendRef.Name, backendIndex)

	if err := p.isReferenceGranted(kindTLSRoute, route.Namespace, group, string(kind), string(backendRef.Name), namespace); err != nil {
		return serviceName, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot load TLSRoute BackendRef %s/%s/%s/%s: %s", group, kind, namespace, backendRef.Name, err),
		}
	}

	if group != groupCore || kind != kindService {
		name, err := p.loadTCPBackendRef(route.Namespace, backendRef)
		if err != nil {
			return serviceName, nil, &metav1.Condition{
				Type:               string(gatev1.RouteConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: route.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.RouteReasonInvalidKind),
				Message:            fmt.Sprintf("Cannot load TLSRoute BackendRef %s/%s/%s/%s: %s", group, kind, namespace, backendRef.Name, err),
			}
		}

		return name, nil, nil
	}

	port := ptr.Deref(backendRef.Port, gatev1.PortNumber(0))
	if port == 0 {
		return serviceName, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonUnsupportedProtocol),
			Message:            fmt.Sprintf("Cannot load TLSRoute BackendRef %s/%s/%s/%s port is required", group, kind, namespace, backendRef.Name),
		}
	}

	lb, st, errCondition := p.loadTLSServers(gatewayName, gatewayNamespace, namespace, route, backendRef, listener, statusReport)
	if errCondition != nil {
		return serviceName, nil, errCondition
	}

	if st != nil {
		lb.ServersTransport = serviceName
		conf.TCP.ServersTransports[serviceName] = st
	}

	return serviceName, &dynamic.TCPService{LoadBalancer: lb}, nil
}

func (p *Provider) loadTLSServers(gatewayName, gatewayNamespace, namespace string, route *gatev1.TLSRoute, backendRef gatev1.BackendRef, listener gatewayListener, statusReport *statusReport) (*dynamic.TCPServersLoadBalancer, *dynamic.TCPServersTransport, *metav1.Condition) {
	backendAddresses, svcPort, err := p.getBackendAddresses(namespace, backendRef)
	if err != nil {
		return nil, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.GetGeneration(),
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonBackendNotFound),
			Message:            fmt.Sprintf("Cannot load TLSRoute BackendRef %s/%s: %s", namespace, backendRef.Name, err),
		}
	}

	backendTLSPolicies, err := p.client.ListBackendTLSPoliciesForService(namespace, string(backendRef.Name))
	if err != nil {
		return nil, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot list BackendTLSPolicies for Service %s/%s: %s", namespace, string(backendRef.Name), err),
		}
	}

	// Sort BackendTLSPolicies by creation timestamp, then by name to match the BackendTLSPolicy requirements.
	slices.SortStableFunc(backendTLSPolicies, func(a, b *gatev1.BackendTLSPolicy) int {
		cmpTime := a.CreationTimestamp.Time.Compare(b.CreationTimestamp.Time)
		if cmpTime == 0 {
			return strings.Compare(a.Name, b.Name)
		}
		return cmpTime
	})

	var serversTransport *dynamic.TCPServersTransport
	for _, policy := range backendTLSPolicies {
		for _, targetRef := range policy.Spec.TargetRefs {
			// Skip targetRefs that doesn't match the backendRef,
			// since a BackendTLSPolicy can select multiple services.
			if targetRef.Name != backendRef.Name {
				continue
			}
			// Skip the targetRef if the sectionName doesn't match the backendRef port.
			if targetRef.SectionName != nil && svcPort.Name != string(*targetRef.SectionName) {
				continue
			}

			policyAncestorStatus := gatev1.PolicyAncestorStatus{
				AncestorRef: gatev1.ParentReference{
					Group:       new(gatev1.Group(groupGateway)),
					Kind:        new(gatev1.Kind(kindGateway)),
					Namespace:   new(gatev1.Namespace(gatewayNamespace)),
					Name:        gatev1.ObjectName(gatewayName),
					SectionName: new(gatev1.SectionName(listener.SectionName)),
				},
				ControllerName: controllerName,
			}

			// Multiple BackendTLSPolicies can match the same service port, meaning that there is a conflict.
			if serversTransport != nil {
				policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions,
					metav1.Condition{
						Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
						Status:             metav1.ConditionFalse,
						ObservedGeneration: policy.Generation,
						LastTransitionTime: metav1.Now(),
						Reason:             string(gatev1.BackendTLSPolicyReasonResolvedRefs),
					},
					metav1.Condition{
						Type:               string(gatev1.PolicyConditionAccepted),
						Status:             metav1.ConditionFalse,
						ObservedGeneration: policy.Generation,
						LastTransitionTime: metav1.Now(),
						Reason:             string(gatev1.PolicyReasonConflicted),
					},
				)

				statusReport.RecordBackendTLSPolicyStatus(ktypes.NamespacedName{Namespace: policy.Namespace, Name: policy.Name}, policyAncestorStatus)

				continue
			}

			var resolvedRefCondition metav1.Condition
			serversTransport, resolvedRefCondition = p.loadTCPServersTransport(namespace, policy, listener.BackendClientCertificate)

			policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions, resolvedRefCondition)
			if resolvedRefCondition.Status == metav1.ConditionFalse {
				policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions, metav1.Condition{
					Type:               string(gatev1.PolicyConditionAccepted),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.BackendTLSPolicyReasonNoValidCACertificate),
				})
			} else {
				policyAncestorStatus.Conditions = append(policyAncestorStatus.Conditions, metav1.Condition{
					Type:               string(gatev1.PolicyConditionAccepted),
					Status:             metav1.ConditionTrue,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.PolicyReasonAccepted),
				})
			}

			statusReport.RecordBackendTLSPolicyStatus(ktypes.NamespacedName{Namespace: policy.Namespace, Name: policy.Name}, policyAncestorStatus)

			// When something went wrong during the loading of a ServersTransport,
			// we stop here and return a route condition error.
			if resolvedRefCondition.Status == metav1.ConditionFalse {
				return nil, nil, &metav1.Condition{
					Type:               string(gatev1.RouteConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: route.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.RouteReasonRefNotPermitted),
					Message:            fmt.Sprintf("Cannot apply BackendTLSPolicy for Service %s/%s: %s", namespace, string(backendRef.Name), resolvedRefCondition.Message),
				}
			}
		}
	}

	if svcPort.Protocol != corev1.ProtocolTCP {
		return nil, nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: route.GetGeneration(),
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonUnsupportedProtocol),
			Message:            fmt.Sprintf("Cannot load TLSRoute BackendRef %s/%s: only TCP protocol is supported", namespace, backendRef.Name),
		}
	}

	lb := &dynamic.TCPServersLoadBalancer{}

	for _, ba := range backendAddresses {
		lb.Servers = append(lb.Servers, dynamic.TCPServer{
			// TODO determine whether the servers needs TLS, from the port?
			Address: net.JoinHostPort(ba.IP, strconv.Itoa(int(ba.Port))),
		})
	}
	return lb, serversTransport, nil
}

func (p *Provider) loadTCPServersTransport(namespace string, policy *gatev1.BackendTLSPolicy, clientCertificate *tls.Certificate) (*dynamic.TCPServersTransport, metav1.Condition) {
	st := &dynamic.TCPServersTransport{
		TLS: &dynamic.TLSClientConfig{
			ServerName: string(policy.Spec.Validation.Hostname),
		},
	}
	if clientCertificate != nil {
		st.TLS.Certificates = tls.Certificates{*clientCertificate}
	}

	if len(policy.Spec.Validation.SubjectAltNames) > 0 {
		// Per the Gateway API specification the Hostname should only be used for authentication
		// and not for certificate validation. Thus, if SubjectAltNames is specified, we ignore
		// the Hostname validation by setting the InsecureSkipVerify option to true.
		st.TLS.InsecureSkipVerify = true

		for _, san := range policy.Spec.Validation.SubjectAltNames {
			switch san.Type {
			case gatev1.URISubjectAltNameType:
				st.TLS.PeerCertSANs = append(st.TLS.PeerCertSANs, tls.SAN{
					Type:  tls.SANURIType,
					Value: string(san.URI),
				})
			case gatev1.HostnameSubjectAltNameType:
				st.TLS.PeerCertSANs = append(st.TLS.PeerCertSANs, tls.SAN{
					Type:  tls.SANDNSNameType,
					Value: string(san.Hostname),
				})
			default:
				return nil, metav1.Condition{
					Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.BackendTLSPolicyReasonInvalidKind),
					Message:            fmt.Sprintf("Unsupported SubjectAltName type %q; only URI and Hostname types are supported", san.Type),
				}
			}
		}
	}

	if policy.Spec.Validation.WellKnownCACertificates != nil {
		return st, metav1.Condition{
			Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
			Status:             metav1.ConditionTrue,
			ObservedGeneration: policy.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.BackendTLSPolicyReasonResolvedRefs),
		}
	}

	for _, caCertRef := range policy.Spec.Validation.CACertificateRefs {
		if (caCertRef.Group != "" && caCertRef.Group != groupCore) || (caCertRef.Kind != kindConfigMap && caCertRef.Kind != kindSecret) {
			return nil, metav1.Condition{
				Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: policy.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.BackendTLSPolicyReasonInvalidKind),
				Message:            "Only ConfigMaps and Secrets are supported",
			}
		}

		var caCRT string
		switch caCertRef.Kind {
		case kindConfigMap:
			configmap, err := p.client.GetConfigMap(namespace, string(caCertRef.Name))
			if err != nil {
				return nil, metav1.Condition{
					Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.BackendTLSPolicyReasonInvalidCACertificateRef),
					Message:            fmt.Sprintf("getting configmap %s/%s: %s", namespace, string(caCertRef.Name), err),
				}
			}
			caCRT = configmap.Data["ca.crt"]
		case kindSecret:
			secret, err := p.client.GetSecret(namespace, string(caCertRef.Name))
			if err != nil {
				return nil, metav1.Condition{
					Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: policy.Generation,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.BackendTLSPolicyReasonInvalidCACertificateRef),
					Message:            fmt.Sprintf("getting secret %s/%s: %s", namespace, string(caCertRef.Name), err),
				}
			}
			caCRT = string(secret.Data["ca.crt"])
		}

		if caCRT == "" {
			return nil, metav1.Condition{
				Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
				Status:             metav1.ConditionFalse,
				ObservedGeneration: policy.Generation,
				LastTransitionTime: metav1.Now(),
				Reason:             string(gatev1.BackendTLSPolicyReasonInvalidCACertificateRef),
				Message:            fmt.Sprintf("%s %s/%s does not have a ca.crt", caCertRef.Kind, namespace, string(caCertRef.Name)),
			}
		}

		st.TLS.RootCAs = append(st.TLS.RootCAs, types.FileOrContent(caCRT))
	}

	return st, metav1.Condition{
		Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
		Status:             metav1.ConditionTrue,
		ObservedGeneration: policy.Generation,
		LastTransitionTime: metav1.Now(),
		Reason:             string(gatev1.BackendTLSPolicyReasonResolvedRefs),
	}
}

func hostSNIRule(hostnames []gatev1.Hostname) (string, int) {
	var priority int

	rules := make([]string, 0, len(hostnames))
	uniqHostnames := map[gatev1.Hostname]struct{}{}

	for _, hostname := range hostnames {
		if len(hostname) == 0 {
			continue
		}

		host := string(hostname)
		wildcard := strings.Count(host, "*")

		thisPriority := len(hostname) - wildcard

		if priority < thisPriority {
			priority = thisPriority
		}

		if _, exists := uniqHostnames[hostname]; exists {
			continue
		}

		uniqHostnames[hostname] = struct{}{}

		if wildcard == 0 {
			rules = append(rules, fmt.Sprintf("HostSNI(%q)", host))
			continue
		}

		// A wildcard label matches one or more labels: *.com matches both
		// example.com and www.example.com.
		host = strings.Replace(regexp.QuoteMeta(host), `\*\.`, `[a-z0-9-\.]+\.`, 1)
		rules = append(rules, fmt.Sprintf("HostSNIRegexp(%q)", fmt.Sprintf("^%s$", host)))
	}

	if len(hostnames) == 0 || len(rules) == 0 {
		return `HostSNI("*")`, 0
	}

	return strings.Join(rules, " || "), priority
}

```

Answer with one JSON object and nothing else:
{"findings": [{"file": "...", "line": 0, "problem": "...", "fix": "..."}]}
Give each finding a concrete file, line and fix. Leave out vague remarks.
An empty list is a valid answer.
