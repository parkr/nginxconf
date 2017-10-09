package main

// nginx-conf-gen -domain=heh.blah.com \
//  -static -ssl -letsencrypt > sites-available/heh.blah.com

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/parkr/nginxconf"
)

func fail(msg string) {
	fmt.Println(fail)
	os.Exit(1)
}

func main() {
	domain := flag.String("domain", "", "Host name for the config")
	altDomainsList := flag.String("altDomains", "", "Domains to redirect to the new domain")
	ssl := flag.Bool("ssl", true, "Write SSL configuration")
	static := flag.Bool("static", false, "Write configuration for a static site")
	webroot := flag.String("webroot", "", "Root on the filesystem for static sites")
	proxy := flag.Bool("proxy", false, "Write configuration for a proxy site")
	proxyPort := flag.Int("port", -1, "Port to proxy to")
	flag.Parse()

	if *static && *proxy {
		fmt.Println("fatal: cannot mix static & proxy types")
		os.Exit(1)
	}

	var altDomains []string
	if *altDomainsList != "" {
		altDomains = strings.Split(*altDomainsList, ",")
	}

	var conf *nginxconf.SiteConfiguration
	if *static {
		conf = &nginxconf.SiteConfiguration{
			Domain:      *domain,
			AltDomains:  altDomains,
			Template:    nginxconf.StaticSite,
			Webroot:     *webroot,
			SSL:         *ssl,
			SSLProvider: nginxconf.LetsEncrypt{},
		}
	} else if *proxy {
		conf = &nginxconf.SiteConfiguration{
			Domain:      *domain,
			AltDomains:  altDomains,
			Template:    nginxconf.ProxySite,
			ProxyPort:   *proxyPort,
			SSL:         *ssl,
			SSLProvider: nginxconf.LetsEncrypt{},
		}
	} else {
		fail("fatal: specify -static or -proxy")
	}

	nginxconf.PrintConfiguration(os.Stdout, conf)
}
