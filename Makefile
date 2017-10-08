PKG=github.com/parkr/nginxconf

all: build test

build:
	go install $(PKG)/...

test:
	go test $(PKG)/...
