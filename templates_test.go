package nginxconf

import (
	"bytes"
	"net/url"
	"strings"
	"testing"
	"text/template"

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

var exampleRedirectSiteConfig = &SiteConfiguration{
	Domain:      "redirect.example.com",
	AltDomains:  []string{"www.example.com"},
	Template:    RedirectSite,
	SSL:         true,
	SSLProvider: LetsEncrypt{},
	RedirectURL: &url.URL{
		Scheme: "https",
		Host:   "example.org",
		Path:   "/foobar",
	},
}

var tests = map[string]*SiteConfiguration{
	"static site":   exampleStaticSiteConfig,
	"proxy site":    exampleProxySiteConfig,
	"redirect site": exampleRedirectSiteConfig,
}

func testTmplExecutes(t *testing.T, tmpl *template.Template) {
	for testName, testConfig := range tests {
		testName := testName     // capture range variable
		testConfig := testConfig // capture range variable
		t.Run(testName, func(t *testing.T) {
			t.Parallel()

			buf := &bytes.Buffer{}
			err := tmpl.Execute(buf, testConfig)
			assert.NoError(t, err)
			assert.NotEmpty(t, strings.TrimSpace(buf.String()))
		})
	}
}

func TestHttpToHttpsRedirectTmpl(t *testing.T) {
	testTmplExecutes(t, httpToHttpsRedirectTmpl)
}

func TestSSLConfigTmpl(t *testing.T) {
	testTmplExecutes(t, sslConfigTmpl)
}

func TestRootLocationTryFilesTmpl(t *testing.T) {
	testTmplExecutes(t, rootLocationTmpl)
}

func TestMediaLocationTmpl(t *testing.T) {
	testTmplExecutes(t, mediaLocationTmpl)
}

func TestSiteConfigTmpl(t *testing.T) {
	testTmplExecutes(t, siteConfigTmpl)
}
