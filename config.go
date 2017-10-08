package nginxconf

type SiteConfigurationTemplateType string

var (
	StaticSite SiteConfigurationTemplateType = "static"
	ProxySite  SiteConfigurationTemplateType = "proxy"
)

type SiteConfigurationSSLProvider interface {
	SSLCertificatePath(domain string) string
	SSLCertificateKeyPath(domain string) string
}

type SiteConfiguration struct {
	// The domain, required
	Domain string

	// The template type, required
	Template SiteConfigurationTemplateType

	// Whether to write SSL
	SSL bool

	// The SSL provider
	SSLProvider SiteConfigurationSSLProvider

	// Proxy port, required for proxy type
	ProxyPort int
}
