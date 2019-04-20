package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	DefaultPort = 8080
)

type args struct {
	port     uint16
	certPath string
	keyPath  string
	url      string
	skipSSL  bool
}

func (o *args) String() string {
	return fmt.Sprintf("-cert=%s -key=%s -url=%s -port=%d --skipSSL=%v",
		o.certPath, o.keyPath, o.url, o.port, o.skipSSL)
}

func parse() *args {
	certPath := flag.String("cert", "", "Path to certificate file")
	keyPath := flag.String("key", "", "Path to private key file")
	url := flag.String("url", "", "Kyma URL")
	port := flag.Uint("port", DefaultPort, "Port number in range 0-65535")
	skipSSL := flag.Bool("skip-ssl", false, "Skip SSL verification")

	flag.Parse()

	if len(*certPath) == 0 {
		printErrorAndExit("Specify path to certificate using -cert parameter")
	}

	if len(*keyPath) == 0 {
		printErrorAndExit("Specify path to certificate key using -key parameter")
	}

	if len(*url) == 0 {
		printErrorAndExit("Specify URL to Kyma instance using -url parameter")
	}

	return &args{
		certPath: *certPath,
		keyPath:  *keyPath,
		port:     uint16(*port),
		url:      *url,
		skipSSL:  *skipSSL,
	}
}

func printErrorAndExit(err string) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
