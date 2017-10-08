package main

// nginx-conf-gen -domain=heh.blah.com \
//  -static -ssl -letsencrypt > sites-available/heh.blah.com

import (
	"flag"
	"fmt"
	"os"

	"github.com/parkr/nginxconf"
)

func fail(msg string) {
	fmt.Println(fail)
	os.Exit(1)
}

func main() {
	domain := flag.String("domain", "", "Host name for the config")
	ssl := flag.Bool("ssl", true, "Write SSL configuration")
	sslProvider := flag.String("sslProvider", "letsencrypt", "Write SSL configuration to use Let's Encrypt")
	static := flag.Bool("static", true, "Write configuration for a static site")
	proxy := flag.Bool("proxy", false, "Write configuration for a proxy site")
	proxyPort := flag.Int("port", -1, "Port to proxy to")
	flag.Parse()

	if *static && *proxy {
		fmt.Println("fatal: cannot mix static & proxy types")
		os.Exit(1)
	}

	var conf *nginxconf.SiteConfiguration
	if *static {
		conf = &nginxconf.SiteConfiguration{
			Domain:      *domain,
			Template:    "static",
			SSL:         *ssl,
			SSLProvider: *sslProvider,
		}
	} else if *proxy {
		conf = &nginxconf.SiteConfiguration{
			Domain:      *domain,
			Template:    "proxy",
			ProxyPort:   *proxyPort,
			SSL:         *ssl,
			SSLProvider: *sslProvider,
		}
	} else {
		fail("fatal: specify -static or -proxy")
	}

	nginxconf.PrintConfiguration(os.Stdout, conf)
}
