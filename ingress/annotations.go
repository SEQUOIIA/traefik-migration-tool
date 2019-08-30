package ingress

const (
	annotationKubernetesIngressClass                    = "kubernetes.io/ingress.class"
	annotationKubernetesAuthRealm                       = "ingress.kubernetes.io/auth-realm"
	annotationKubernetesAuthType                        = "ingress.kubernetes.io/auth-type"
	annotationKubernetesAuthSecret                      = "ingress.kubernetes.io/auth-secret"
	annotationKubernetesAuthHeaderField                 = "ingress.kubernetes.io/auth-header-field"
	annotationKubernetesAuthForwardResponseHeaders      = "ingress.kubernetes.io/auth-response-headers"
	annotationKubernetesAuthRemoveHeader                = "ingress.kubernetes.io/auth-remove-header"
	annotationKubernetesAuthForwardURL                  = "ingress.kubernetes.io/auth-url"
	annotationKubernetesAuthForwardTrustHeaders         = "ingress.kubernetes.io/auth-trust-headers"
	annotationKubernetesAuthForwardTLSSecret            = "ingress.kubernetes.io/auth-tls-secret"
	annotationKubernetesAuthForwardTLSInsecure          = "ingress.kubernetes.io/auth-tls-insecure"
	annotationKubernetesRewriteTarget                   = "ingress.kubernetes.io/rewrite-target"
	annotationKubernetesWhiteListSourceRange            = "ingress.kubernetes.io/whitelist-source-range"
	annotationKubernetesWhiteListUseXForwardedFor       = "ingress.kubernetes.io/whitelist-x-forwarded-for"
	annotationKubernetesPreserveHost                    = "ingress.kubernetes.io/preserve-host"
	annotationKubernetesPassTLSCert                     = "ingress.kubernetes.io/pass-tls-cert" // Deprecated
	annotationKubernetesPassTLSClientCert               = "ingress.kubernetes.io/pass-client-tls-cert"
	annotationKubernetesFrontendEntryPoints             = "ingress.kubernetes.io/frontend-entry-points"
	annotationKubernetesPriority                        = "ingress.kubernetes.io/priority"
	annotationKubernetesCircuitBreakerExpression        = "ingress.kubernetes.io/circuit-breaker-expression"
	annotationKubernetesLoadBalancerMethod              = "ingress.kubernetes.io/load-balancer-method"
	annotationKubernetesAffinity                        = "ingress.kubernetes.io/affinity"
	annotationKubernetesSessionCookieName               = "ingress.kubernetes.io/session-cookie-name"
	annotationKubernetesRuleType                        = "ingress.kubernetes.io/rule-type"
	annotationKubernetesRedirectEntryPoint              = "ingress.kubernetes.io/redirect-entry-point"
	annotationKubernetesRedirectPermanent               = "ingress.kubernetes.io/redirect-permanent"
	annotationKubernetesRedirectRegex                   = "ingress.kubernetes.io/redirect-regex"
	annotationKubernetesRedirectReplacement             = "ingress.kubernetes.io/redirect-replacement"
	annotationKubernetesMaxConnAmount                   = "ingress.kubernetes.io/max-conn-amount"
	annotationKubernetesMaxConnExtractorFunc            = "ingress.kubernetes.io/max-conn-extractor-func"
	annotationKubernetesRateLimit                       = "ingress.kubernetes.io/rate-limit"
	annotationKubernetesErrorPages                      = "ingress.kubernetes.io/error-pages"
	annotationKubernetesBuffering                       = "ingress.kubernetes.io/buffering"
	annotationKubernetesResponseForwardingFlushInterval = "ingress.kubernetes.io/responseforwarding-flushinterval"
	annotationKubernetesAppRoot                         = "ingress.kubernetes.io/app-root"
	annotationKubernetesServiceWeights                  = "ingress.kubernetes.io/service-weights"
	annotationKubernetesRequestModifier                 = "ingress.kubernetes.io/request-modifier"

	annotationKubernetesSSLForceHost            = "ingress.kubernetes.io/ssl-force-host"
	annotationKubernetesSSLRedirect             = "ingress.kubernetes.io/ssl-redirect"
	annotationKubernetesHSTSMaxAge              = "ingress.kubernetes.io/hsts-max-age"
	annotationKubernetesHSTSIncludeSubdomains   = "ingress.kubernetes.io/hsts-include-subdomains"
	annotationKubernetesCustomRequestHeaders    = "ingress.kubernetes.io/custom-request-headers"
	annotationKubernetesCustomResponseHeaders   = "ingress.kubernetes.io/custom-response-headers"
	annotationKubernetesAllowedHosts            = "ingress.kubernetes.io/allowed-hosts"
	annotationKubernetesProxyHeaders            = "ingress.kubernetes.io/proxy-headers"
	annotationKubernetesSSLTemporaryRedirect    = "ingress.kubernetes.io/ssl-temporary-redirect"
	annotationKubernetesSSLHost                 = "ingress.kubernetes.io/ssl-host"
	annotationKubernetesSSLProxyHeaders         = "ingress.kubernetes.io/ssl-proxy-headers"
	annotationKubernetesHSTSPreload             = "ingress.kubernetes.io/hsts-preload"
	annotationKubernetesForceHSTSHeader         = "ingress.kubernetes.io/force-hsts"
	annotationKubernetesFrameDeny               = "ingress.kubernetes.io/frame-deny"
	annotationKubernetesCustomFrameOptionsValue = "ingress.kubernetes.io/custom-frame-options-value"
	annotationKubernetesContentTypeNosniff      = "ingress.kubernetes.io/content-type-nosniff"
	annotationKubernetesBrowserXSSFilter        = "ingress.kubernetes.io/browser-xss-filter"
	annotationKubernetesCustomBrowserXSSValue   = "ingress.kubernetes.io/custom-browser-xss-value"
	annotationKubernetesContentSecurityPolicy   = "ingress.kubernetes.io/content-security-policy"
	annotationKubernetesPublicKey               = "ingress.kubernetes.io/public-key"
	annotationKubernetesReferrerPolicy          = "ingress.kubernetes.io/referrer-policy"
	annotationKubernetesIsDevelopment           = "ingress.kubernetes.io/is-development"
	annotationKubernetesProtocol                = "ingress.kubernetes.io/protocol"
)

// TODO [breaking] remove label support
// var compatibilityMapping = map[string]string{
// 	annotationKubernetesPreserveHost:             "traefik.frontend.passHostHeader",
// 	annotationKubernetesPassTLSCert:              "traefik.frontend.passTLSCert",
// 	annotationKubernetesFrontendEntryPoints:      "traefik.frontend.entryPoints",
// 	annotationKubernetesPriority:                 "traefik.frontend.priority",
// 	annotationKubernetesCircuitBreakerExpression: "traefik.backend.circuitbreaker",
// 	annotationKubernetesLoadBalancerMethod:       "traefik.backend.loadbalancer.method",
// 	annotationKubernetesAffinity:                 "traefik.backend.loadbalancer.stickiness",
// 	annotationKubernetesSessionCookieName:        "traefik.backend.loadbalancer.stickiness.cookieName",
// 	annotationKubernetesRuleType:                 "traefik.frontend.rule.type",
// 	annotationKubernetesRedirectEntryPoint:       "traefik.frontend.redirect.entrypoint",
// 	annotationKubernetesRedirectRegex:            "traefik.frontend.redirect.regex",
// 	annotationKubernetesRedirectReplacement:      "traefik.frontend.redirect.replacement",
// }