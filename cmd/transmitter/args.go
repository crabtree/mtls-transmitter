package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
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
	silent   bool
}

func (o *args) String() string {
	return fmt.Sprintf("cert=%s key=%s url=%s port=%d skip-ssl=%v silent=%v",
		o.certPath, o.keyPath, o.url, o.port, o.skipSSL, o.silent)
}

func parse() *args {
	certPath := flag.String("cert", "", "Path to certificate file [CERT]")
	keyPath := flag.String("key", "", "Path to private key file [KEY]")
	url := flag.String("url", "", "Target URL [URL]")
	port := flag.Uint("port", DefaultPort, "Port number in range 0-65535 [PORT]")
	skipSSL := flag.Bool("skip-ssl", false, "Skip SSL verification [SKIP_SSL=true]")
	silent := flag.Bool("silent", false, "Do not log proxied requests [SILENT=true]")

	flag.Parse()

	// Check for configuration from environment.
	// - Attempts to make flags authoritative. It gets tricky when there is a
	//   real default set.
	if len(*certPath) == 0 {
		envCert := os.Getenv("CERT")
		certPath = &envCert
	}

	if len(*keyPath) == 0 {
		envKey := os.Getenv("KEY")
		keyPath = &envKey
	}

	if len(*url) == 0 {
		envUrl := os.Getenv("URL")
		url = &envUrl
	}

	if *port != uint(DefaultPort) {
		if envPort := os.Getenv("PORT"); len(envPort) != 0 {
			p, e := strconv.ParseUint(envPort, 10, 32)
			if e != nil {
				log.Fatalf("Fatal: %v\n", e)
			}
			*port = uint(p)
		}
	}

	if !*skipSSL {
		if os.Getenv("SKIP_SSL") == "true" {
			*skipSSL = true
		}
	}

	if !*silent {
		if os.Getenv("SILENT") == "true" {
			*silent = true
		}
	}

	// Validate required or exit.
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
		silent:   *silent,
	}
}

func printErrorAndExit(err string) {
	log.Fatalln(err)
}
