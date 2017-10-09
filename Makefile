PKG=github.com/parkr/nginxconf

all: build test

build:
	go install $(PKG)/...

test:
	go test $(PKG)/...

run: build
	nginx-conf-gen \
	  -domain=static.example.com \
	  -static -ssl \
	  -altDomains="www.example.com" \
	  -webroot="/var/www/octocat/example-website" > static.conf
	nginx-conf-gen \
	  -domain=proxy.example.com \
	  -proxy -ssl \
	  -altDomains="www2.example.com" \
	  -port=1313 > proxy.conf
	nginx-mimes-gen > mimetypes.conf
