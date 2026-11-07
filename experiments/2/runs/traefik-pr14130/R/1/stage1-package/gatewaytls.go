package gateway

import (
	"fmt"

	"github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

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

// resolveBackendClientCertificate returns the client certificate the Gateway presents to its backends, nil when it references none,
// or the Gateway ResolvedRefs condition explaining why the reference cannot be resolved.
// Only the Core support is implemented: a reference to a Secret.
func (p *Provider) resolveBackendClientCertificate(gateway *gatev1.Gateway) (*tls.Certificate, *metav1.Condition) {
	if gateway.Spec.TLS == nil || gateway.Spec.TLS.Backend == nil || gateway.Spec.TLS.Backend.ClientCertificateRef == nil {
		return nil, nil
	}

	ref := gateway.Spec.TLS.Backend.ClientCertificateRef

	group := ptr.Deref(ref.Group, "")
	kind := ptr.Deref(ref.Kind, kindSecret)
	if (group != "" && group != groupCore) || kind != kindSecret {
		return nil, new(makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionFalse, gatev1.GatewayReasonInvalidClientCertificateRef,
			fmt.Sprintf("Unsupported ClientCertificateRef group/kind: %s/%s", group, kind)))
	}

	namespace := string(ptr.Deref(ref.Namespace, gatev1.Namespace(gateway.Namespace)))
	if err := p.isReferenceGranted(kindGateway, gateway.Namespace, groupCore, kindSecret, string(ref.Name), namespace); err != nil {
		return nil, new(makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionFalse, gatev1.GatewayReasonRefNotPermitted,
			fmt.Sprintf("Cannot reference ClientCertificateRef %s/%s: %s", namespace, ref.Name, err)))
	}

	cert, err := p.getTLSCert(ref.Name, namespace)
	if err != nil {
		return nil, new(makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionFalse, gatev1.GatewayReasonInvalidClientCertificateRef,
			fmt.Sprintf("Cannot load ClientCertificateRef %s/%s: %s", namespace, ref.Name, err)))
	}

	return &cert.Certificate, nil
}

func makeGatewayResolvedRefsCondition(generation int64, status metav1.ConditionStatus, reason gatev1.GatewayConditionReason, message string) metav1.Condition {
	return metav1.Condition{
		Type:               string(gatev1.GatewayConditionResolvedRefs),
		Status:             status,
		ObservedGeneration: generation,
		LastTransitionTime: metav1.Now(),
		Reason:             string(reason),
		Message:            message,
	}
}

// gatewayResolvedRefsCondition returns the Gateway ResolvedRefs condition. The backend client certificate reference,
// a Gateway level reference, takes precedence over the references of the listeners the Gateway declares itself.
func gatewayResolvedRefsCondition(gateway *gatev1.Gateway, listenerStatuses []gatev1.ListenerStatus, backendClientCertificateErr *metav1.Condition) metav1.Condition {
	if backendClientCertificateErr != nil {
		return *backendClientCertificateErr
	}

	for _, listenerStatus := range listenerStatuses {
		if meta.IsStatusConditionFalse(listenerStatus.Conditions, string(gatev1.ListenerConditionResolvedRefs)) {
			return makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionFalse, gatev1.GatewayReasonListenersNotResolved, "At least one Listener has unresolved references")
		}
	}

	return makeGatewayResolvedRefsCondition(gateway.Generation, metav1.ConditionTrue, gatev1.GatewayReasonResolvedRefs, conditionNoErrorMessage)
}
