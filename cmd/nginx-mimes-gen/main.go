package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/parkr/nginxconf/mimetypes"
)

func main() {
	flag.Parse()

	types, err := mimetypes.Fetch()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: failed fetching mime types: %+v\n", err)
		os.Exit(1)
	}

	err = mimetypes.PrintConfiguration(os.Stdout, types)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: failed printing mime types: %+v\n", err)
		os.Exit(1)
	}
}
