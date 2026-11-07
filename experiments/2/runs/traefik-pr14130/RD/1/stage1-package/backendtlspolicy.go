package gateway

import (
	"fmt"
	"slices"
	"strings"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ktypes "k8s.io/apimachinery/pkg/types"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

// loadBackendTLSPolicyClientConfig returns the TLS configuration to dial the given Service port with,
// set by the oldest BackendTLSPolicy targeting it, and records the status of the policies targeting it.
// It returns nil when no BackendTLSPolicy targets the port, as the backend is then not dialed with TLS,
// and the route ResolvedRefs condition to report when the policy cannot be applied.
//
func (p *Provider) loadBackendTLSPolicyClientConfig(gatewayName, gatewayNamespace string, listener gatewayListener, namespace string, serviceName gatev1.ObjectName, svcPort corev1.ServicePort, routeGeneration int64, statusReport *statusReport) (*dynamic.TLSClientConfig, *metav1.Condition) {
	backendTLSPolicies, err := p.client.ListBackendTLSPoliciesForService(namespace, string(serviceName))
	if err != nil {
		return nil, &metav1.Condition{
			Type:               string(gatev1.RouteConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: routeGeneration,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.RouteReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot list BackendTLSPolicies for Service %s/%s: %s", namespace, string(serviceName), err),
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

	var clientConfig *dynamic.TLSClientConfig
	for _, policy := range backendTLSPolicies {
		for _, targetRef := range policy.Spec.TargetRefs {
			// Skip targetRefs that doesn't match the backendRef,
			// since a BackendTLSPolicy can select multiple services.
			if targetRef.Name != serviceName {
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
			if clientConfig != nil {
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
			clientConfig, resolvedRefCondition = p.buildBackendTLSPolicyClientConfig(namespace, policy)

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

			// When something went wrong during the loading of the TLS configuration,
			// we stop here and return a route condition error.
			if resolvedRefCondition.Status == metav1.ConditionFalse {
				return nil, &metav1.Condition{
					Type:               string(gatev1.RouteConditionResolvedRefs),
					Status:             metav1.ConditionFalse,
					ObservedGeneration: routeGeneration,
					LastTransitionTime: metav1.Now(),
					Reason:             string(gatev1.RouteReasonRefNotPermitted),
					Message:            fmt.Sprintf("Cannot apply BackendTLSPolicy for Service %s/%s: %s", namespace, string(serviceName), resolvedRefCondition.Message),
				}
			}
		}
	}

	// The client certificate is set on the Gateway, but only presented to the backends a BackendTLSPolicy dials with TLS.
	if clientConfig != nil {
		clientConfig.Certificates = listener.BackendClientCertificates
	}

	return clientConfig, nil
}

// serversTransportForBackendTLSPolicy returns the HTTP ServersTransport dialing with the TLS configuration of a BackendTLSPolicy.
//
func serversTransportForBackendTLSPolicy(clientConfig *dynamic.TLSClientConfig) *dynamic.ServersTransport {
	if clientConfig == nil {
		return nil
	}

	return &dynamic.ServersTransport{
		ServerName:         clientConfig.ServerName,
		InsecureSkipVerify: clientConfig.InsecureSkipVerify,
		RootCAs:            clientConfig.RootCAs,
		Certificates:       clientConfig.Certificates,
		PeerCertSANs:       clientConfig.PeerCertSANs,
	}
}

func (p *Provider) buildBackendTLSPolicyClientConfig(namespace string, policy *gatev1.BackendTLSPolicy) (*dynamic.TLSClientConfig, metav1.Condition) {
	clientConfig := &dynamic.TLSClientConfig{
		ServerName: string(policy.Spec.Validation.Hostname),
	}

	if len(policy.Spec.Validation.SubjectAltNames) > 0 {
		// Per the Gateway API specification the Hostname should only be used for authentication
		// and not for certificate validation. Thus, if SubjectAltNames is specified, we ignore
		// the Hostname validation by setting the InsecureSkipVerify option to true.
		clientConfig.InsecureSkipVerify = true

		for _, san := range policy.Spec.Validation.SubjectAltNames {
			switch san.Type {
			case gatev1.URISubjectAltNameType:
				clientConfig.PeerCertSANs = append(clientConfig.PeerCertSANs, tls.SAN{
					Type:  tls.SANURIType,
					Value: string(san.URI),
				})
			case gatev1.HostnameSubjectAltNameType:
				clientConfig.PeerCertSANs = append(clientConfig.PeerCertSANs, tls.SAN{
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
		return clientConfig, metav1.Condition{
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

		clientConfig.RootCAs = append(clientConfig.RootCAs, types.FileOrContent(caCRT))
	}

	return clientConfig, metav1.Condition{
		Type:               string(gatev1.BackendTLSPolicyConditionResolvedRefs),
		Status:             metav1.ConditionTrue,
		ObservedGeneration: policy.Generation,
		LastTransitionTime: metav1.Now(),
		Reason:             string(gatev1.BackendTLSPolicyReasonResolvedRefs),
	}
}
