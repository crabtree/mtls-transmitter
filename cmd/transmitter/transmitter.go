package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crabtree/mtls-transmitter/pkg/certificates"
	"github.com/crabtree/mtls-transmitter/pkg/proxyhandler"
	"github.com/crabtree/mtls-transmitter/pkg/proxyserver"
)

func main() {
	opts := parse()

	cert, err := certificates.Load(opts.certPath, opts.keyPath)
	if err != nil {
		printErrorAndExit("Error when loading cerificate, " + err.Error())
	}

	handler := proxyhandler.Create(opts.url, cert, opts.skipSSL)

	proxyserver.NewServer(http.DefaultServeMux, handler)

	fmt.Println("mtls-transmitter is listening on port:", opts.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", opts.port), nil))
}
