package nginxconf

import (
	"fmt"
	"net/url"
)

type SiteConfigurationTemplateType string

var (
	StaticSite   SiteConfigurationTemplateType = "static"
	ProxySite    SiteConfigurationTemplateType = "proxy"
	RedirectSite SiteConfigurationTemplateType = "redirect"
)

type SiteConfigurationSSLProvider interface {
	SSLCertificatePath(domain string) string
	SSLCertificateKeyPath(domain string) string
	SSLTrustedCertificatePath(domain string) string
}

type SiteConfiguration struct {
	// The domain, required
	Domain string

	// Alternative domains we should handle but redirect
	AltDomains []string

	// The template type, required
	Template SiteConfigurationTemplateType

	// Whether to write SSL
	SSL bool

	// The SSL provider
	SSLProvider SiteConfigurationSSLProvider

	// Proxy port, required for proxy type
	ProxyPort int

	// Root directory
	Webroot string

	// Redirect URL, required for redirect type
	RedirectURL *url.URL
}

func (c *SiteConfiguration) IsStatic() bool {
	return c.Template == StaticSite
}

func (c *SiteConfiguration) IsProxy() bool {
	return c.Template == ProxySite
}

func (c *SiteConfiguration) IsRedirect() bool {
	return c.Template == RedirectSite
}

func (c *SiteConfiguration) ProxyURL() string {
	return fmt.Sprintf("http://localhost:%d", c.ProxyPort)
}

func (c *SiteConfiguration) RedirectSchemeHost() string {
	if c.RedirectURL == nil {
		return "https://$host"
	}
	return fmt.Sprintf("%s://%s", c.RedirectURL.Scheme, c.RedirectURL.Host)
}

func (c *SiteConfiguration) SSLCertificatePath() string {
	return c.SSLProvider.SSLCertificatePath(c.Domain)
}

func (c *SiteConfiguration) SSLCertificateKeyPath() string {
	return c.SSLProvider.SSLCertificateKeyPath(c.Domain)
}

func (c *SiteConfiguration) SSLTrustedCertificatePath() string {
	return c.SSLProvider.SSLTrustedCertificatePath(c.Domain)
}
