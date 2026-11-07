package gateway

import (
	"fmt"

	"github.com/traefik/traefik/v3/pkg/tls"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

// loadGatewayBackendTLSClientCertificates returns the client certificate the Gateway presents to the backends
// a BackendTLSPolicy dials with TLS (GEP-3155), or the Gateway ResolvedRefs condition explaining why it cannot be resolved.
//
func (p *Provider) loadGatewayBackendTLSClientCertificates(gateway *gatev1.Gateway) (tls.Certificates, *metav1.Condition) {
	if gateway.Spec.TLS == nil || gateway.Spec.TLS.Backend == nil || gateway.Spec.TLS.Backend.ClientCertificateRef == nil {
		return nil, nil
	}

	ref := gateway.Spec.TLS.Backend.ClientCertificateRef

	group := string(ptr.Deref(ref.Group, ""))
	kind := string(ptr.Deref(ref.Kind, kindSecret))
	if (group != "" && group != groupCore) || kind != kindSecret {
		return nil, &metav1.Condition{
			Type:               string(gatev1.GatewayConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.GatewayReasonInvalidClientCertificateRef),
			Message:            fmt.Sprintf("Unsupported ClientCertificateRef group/kind: %s/%s", group, kind),
		}
	}

	namespace := string(ptr.Deref(ref.Namespace, gatev1.Namespace(gateway.Namespace)))
	if err := p.isReferenceGranted(kindGateway, gateway.Namespace, groupCore, kindSecret, string(ref.Name), namespace); err != nil {
		return nil, &metav1.Condition{
			Type:               string(gatev1.GatewayConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.GatewayReasonRefNotPermitted),
			Message:            fmt.Sprintf("Cannot reference ClientCertificateRef %s/%s: %s", namespace, ref.Name, err),
		}
	}

	cert, err := p.getTLSCert(ref.Name, namespace)
	if err != nil {
		return nil, &metav1.Condition{
			Type:               string(gatev1.GatewayConditionResolvedRefs),
			Status:             metav1.ConditionFalse,
			ObservedGeneration: gateway.Generation,
			LastTransitionTime: metav1.Now(),
			Reason:             string(gatev1.GatewayReasonInvalidClientCertificateRef),
			Message:            fmt.Sprintf("Cannot load ClientCertificateRef %s/%s: %s", namespace, ref.Name, err),
		}
	}

	return tls.Certificates{{CertFile: cert.CertFile, KeyFile: cert.KeyFile}}, nil
}
