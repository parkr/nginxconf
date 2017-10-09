package nginxconf

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var exampleStaticSiteConfig = &SiteConfiguration{
	Domain:      "static.example.com",
	AltDomains:  []string{"www.example.com"},
	Template:    StaticSite,
	SSL:         true,
	SSLProvider: LetsEncrypt{},
	Webroot:     "/var/www/static",
}

var exampleProxySiteConfig = &SiteConfiguration{
	Domain:      "proxy.example.com",
	AltDomains:  []string{"www.example.com"},
	Template:    ProxySite,
	SSL:         true,
	SSLProvider: LetsEncrypt{},
	ProxyPort:   1234,
}

func TestHttpToHttpsRedirectTmpl(t *testing.T) {
	buf := &bytes.Buffer{}
	err := httpToHttpsRedirectTmpl.Execute(buf, exampleStaticSiteConfig)
	assert.NoError(t, err)
	assert.NotEmpty(t, strings.TrimSpace(buf.String()))
}

func TestSSLConfigTmpl(t *testing.T) {
	buf := &bytes.Buffer{}
	err := sslConfigTmpl.Execute(buf, exampleStaticSiteConfig)
	assert.NoError(t, err)
	assert.NotEmpty(t, strings.TrimSpace(buf.String()))
}

func TestRootLocationTryFilesTmpl(t *testing.T) {
	buf := &bytes.Buffer{}
	err := rootLocationTmpl.Execute(buf, exampleStaticSiteConfig)
	assert.NoError(t, err)
	assert.NotEmpty(t, strings.TrimSpace(buf.String()))
}

func TestMediaLocationTmpl(t *testing.T) {
	buf := &bytes.Buffer{}
	assert.NoError(t, mediaLocationTmpl.Execute(buf, exampleStaticSiteConfig))
	assert.NotEmpty(t, strings.TrimSpace(buf.String()))
}

func TestSiteConfigTmpl(t *testing.T) {
	buf := &bytes.Buffer{}
	err := siteConfigTmpl.Execute(buf, exampleStaticSiteConfig)
	assert.NoError(t, err)
	assert.NotEmpty(t, strings.TrimSpace(buf.String()))
}
