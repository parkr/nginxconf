# nginxconf

A set of utilities for generating nginx configurations.

To install, simply run:

```console
$ go get github.com/parkr/nginxconf/...
```

## nginx-conf-gen

This command generates nginx `server` blocks. It can be used to generate
one of three different types of sites:

1. Static site (serve files from a directory)
2. Proxy site (act as a reverse proxy to another server running on the same host)
3. Redirect site (redirect any incoming traffic to another domain)

By default, it generates strong a strong SSL configuration based on Let's
Encrypt conventions. It does not generate certificates for you, however.

These all print to stdout. You can redirect the contents to a file in your
nginx installation's `sites-available` directory.

### 1. Static site

To generate a static site, you'll need a directory with your static site to
serve from. Then, simply run:

```console
$ nginx-conf-gen -domain="example.org" \
    -static -webroot=/var/www/html
```

This will serve the contents of `/var/www/html` at `https://example.org`.

### 2. Proxy site

To generate a proxy site, you'll need to know what port to forward the
traffic to.

```console
$ nginx-conf-gen -domain="example.org" \
    -proxy -port=8080
```

This will serve all traffic from your server running at `localhost:8080` at `https://example.org`.

### 3. Redirect site

To generate a proxy site, you'll need a schema and host to redirect traffic
to, e.g. `https://example.co`.

```console
$ nginx-conf-gen -domain="example.org" \
    -redirect="https://example.co"
```

This will redirect all traffic to `example.org` to `https://example.co`.

## nginx-mimes-gen

This command pulls down the MIME types from [mime-db](https://github.com/jshttp/mime-db)
and generates an nginx `types` block. It is quite an extensive list but is
useful for serving static sites properly.

To run this command, you will need to know where your nginx configuration
lives. On my servers, it usually lives in `/opt/nginx/conf`. There's a file
called `mime.types` which I just overwrite:

```console
$ nginx-mimes-gen > /opt/nginx/conf/mime.types
```

In the main `http` block in your `nginx.conf` file, ensure the line
`include mime.types;` exists. If it's not there, add it and reload nginx.
Voilà!

## LICENSE

Code released under The MIT License. See LICENSE.md for details.
