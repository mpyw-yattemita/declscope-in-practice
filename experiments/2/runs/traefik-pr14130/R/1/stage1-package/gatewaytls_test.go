package gateway

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
	"github.com/traefik/traefik/v3/pkg/tls"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kubefake "k8s.io/client-go/kubernetes/fake"
	gatev1 "sigs.k8s.io/gateway-api/apis/v1"
)

func Test_loadConfigurationFromGateways_backendClientCertificate(t *testing.T) {
	// The client certificates are identified by their common name:
	// "example.com" for the listener certificate, and "client-b.example.com" for the other one.
	testCases := []struct {
		desc        string
		path        string
		entryPoints map[string]Entrypoint
		// wantServersTransports and wantTCPServersTransports hold the client certificates of each ServersTransport, nil when none is expected.
		wantServersTransports    map[string][]string
		wantTCPServersTransports map[string][]string
		// wantResolvedRefs holds the expected ResolvedRefs condition of each Gateway, by name.
		wantResolvedRefs map[string]metav1.Condition
	}{
		{
			desc:        "HTTPRoute",
			path:        "httproute/with_backend_client_certificate.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}},
			wantServersTransports: map[string][]string{
				"httproute-default-http-app-1-gw-default-my-gateway-ep-web-0-a621494c27c089eff189-svc-default-whoami-0": {"example.com"},
			},
			wantResolvedRefs: map[string]metav1.Condition{
				"my-gateway": resolvedRefs(),
			},
		},
		{
			desc:        "GRPCRoute",
			path:        "grpcroute/with_backend_client_certificate.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}},
			wantServersTransports: map[string][]string{
				"grpcroute-default-grpc-app-1-gw-default-my-gateway-ep-web-0-0abda411573f19f36296-svc-default-whoami-0": {"example.com"},
			},
			wantResolvedRefs: map[string]metav1.Condition{
				"my-gateway": resolvedRefs(),
			},
		},
		{
			desc:        "TLSRoute",
			path:        "tlsroute/with_backend_client_certificate.yml",
			entryPoints: map[string]Entrypoint{"tls": {Address: ":9001"}},
			wantTCPServersTransports: map[string][]string{
				"tlsroute-default-tls-app-1-gw-default-my-gateway-ep-tls-0-5c999a999d12ac0531c5-svc-default-whoami-0": {"client-b.example.com"},
			},
			wantResolvedRefs: map[string]metav1.Condition{
				"my-gateway": resolvedRefs(),
			},
		},
		{
			desc:        "ListenerSet listener",
			path:        "httproute/with_backend_client_certificate_listenerset.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}, "web8080": {Address: ":8080"}},
			wantServersTransports: map[string][]string{
				"httproute-default-http-app-1-gw-default-my-gateway-ep-web8080-0-b6e15c94e4c622c7df2f-svc-default-whoami-0": {"example.com"},
			},
			wantResolvedRefs: map[string]metav1.Condition{
				"my-gateway": resolvedRefs(),
			},
		},
		{
			desc:        "two Gateways",
			path:        "httproute/with_backend_client_certificate_two_gateways.yml",
			entryPoints: map[string]Entrypoint{"web": {Address: ":80"}, "web8080": {Address: ":8080"}},
			wantServersTransports: map[string][]string{
				"httproute-default-http-app-1-gw-default-gateway-a-ep-web-0-2b4a349f68e029368baa-svc-default-whoami-0":     {"example.com"},
				"httproute-default-http-app-1-gw-default-gateway-b-ep-web8080-0-fe643e6a954d4337e9b9-svc-default-whoami-0": {"client-b.example.com"},
			},
			wantResolvedRefs: map[string]metav1.Condition{
				"gateway-a": resolvedRefs(),
				"gateway-b": resolvedRefs(),
			},
		},
		{
			desc: "references",
			path: "httproute/with_backend_client_certificate_refs.yml",
			entryPoints: map[string]Entrypoint{
				"unsupported-kind":  {Address: ":8001"},
				"unsupported-group": {Address: ":8002"},
				"not-permitted":     {Address: ":8003"},
				"missing-secret":    {Address: ":8004"},
				"malformed-secret":  {Address: ":8005"},
				"granted":           {Address: ":8006"},
			},
			wantServersTransports: map[string][]string{
				"httproute-default-http-app-1-gw-default-granted-ep-granted-0-5502c52b662b9d02b22e-svc-default-whoami-0":                     {"example.com"},
				"httproute-default-http-app-1-gw-default-unsupported-kind-ep-unsupported-kind-0-fa1dcd52178947a9cd1b-svc-default-whoami-0":   nil,
				"httproute-default-http-app-1-gw-default-unsupported-group-ep-unsupported-group-0-74377071f9fb9f91ae7e-svc-default-whoami-0": nil,
				"httproute-default-http-app-1-gw-default-not-permitted-ep-not-permitted-0-4f0246b12b1fab4f1d5e-svc-default-whoami-0":         nil,
				"httproute-default-http-app-1-gw-default-missing-secret-ep-missing-secret-0-1fb83d17faf417416a77-svc-default-whoami-0":       nil,
				"httproute-default-http-app-1-gw-default-malformed-secret-ep-malformed-secret-0-da917b36feaa42d54221-svc-default-whoami-0":   nil,
			},
			wantResolvedRefs: map[string]metav1.Condition{
				"granted":           resolvedRefs(),
				"unsupported-kind":  unresolvedRefs(gatev1.GatewayReasonInvalidClientCertificateRef, "Unsupported ClientCertificateRef group/kind: /ConfigMap"),
				"unsupported-group": unresolvedRefs(gatev1.GatewayReasonInvalidClientCertificateRef, "Unsupported ClientCertificateRef group/kind: example.com/Secret"),
				"not-permitted":     unresolvedRefs(gatev1.GatewayReasonRefNotPermitted, "Cannot reference ClientCertificateRef bar/client-cert: missing ReferenceGrant"),
				"missing-secret":    unresolvedRefs(gatev1.GatewayReasonInvalidClientCertificateRef, `Cannot load ClientCertificateRef default/missing: getting secret: secret "missing" not found`),
				"malformed-secret":  unresolvedRefs(gatev1.GatewayReasonInvalidClientCertificateRef, "Cannot load ClientCertificateRef default/malformed: getting certificate blocks: secret default/malformed is missing the following TLS data entries: tls.key"),
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

			assert.Equal(t, test.wantServersTransports, clientCertificatesByTransport(t, conf.HTTP.ServersTransports, func(st *dynamic.ServersTransport) tls.Certificates {
				return st.Certificates
			}))
			assert.Equal(t, test.wantTCPServersTransports, clientCertificatesByTransport(t, conf.TCP.ServersTransports, func(st *dynamic.TCPServersTransport) tls.Certificates {
				return st.TLS.Certificates
			}))

			statusReport.Flush(t.Context(), client)

			for name, want := range test.wantResolvedRefs {
				gateway, err := gwClient.GatewayV1().Gateways("default").Get(t.Context(), name, metav1.GetOptions{})
				require.NoError(t, err)

				condition := meta.FindStatusCondition(gateway.Status.Conditions, string(gatev1.GatewayConditionResolvedRefs))
				require.NotNil(t, condition, name)

				condition.LastTransitionTime = metav1.Time{}
				assert.Equal(t, want, *condition, name)
			}
		})
	}
}

// clientCertificatesByTransport returns the common names of the client certificates of each ServersTransport, nil when there is none.
func clientCertificatesByTransport[T any](t *testing.T, serversTransports map[string]T, certificates func(T) tls.Certificates) map[string][]string {
	t.Helper()

	var commonNamesByTransport map[string][]string
	for name, st := range serversTransports {
		var commonNames []string
		for _, certificate := range certificates(st) {
			commonNames = append(commonNames, certificateCommonName(t, certificate.CertFile))
		}

		if commonNamesByTransport == nil {
			commonNamesByTransport = make(map[string][]string)
		}
		commonNamesByTransport[name] = commonNames
	}

	return commonNamesByTransport
}

func resolvedRefs() metav1.Condition {
	return metav1.Condition{
		Type:    string(gatev1.GatewayConditionResolvedRefs),
		Status:  metav1.ConditionTrue,
		Reason:  string(gatev1.GatewayReasonResolvedRefs),
		Message: conditionNoErrorMessage,
	}
}

func unresolvedRefs(reason gatev1.GatewayConditionReason, message string) metav1.Condition {
	return metav1.Condition{
		Type:    string(gatev1.GatewayConditionResolvedRefs),
		Status:  metav1.ConditionFalse,
		Reason:  string(reason),
		Message: message,
	}
}
