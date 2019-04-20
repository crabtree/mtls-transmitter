package proxyhandler

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func Create(url string, cert tls.Certificate, skipSSL bool) http.HandlerFunc {
	tlsConfig := tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: skipSSL,
	}

	transport := &http.Transport{
		TLSClientConfig: &tlsConfig,
	}

	proxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "https"
			r.URL.Host = url
			r.Host = url
		},
		Transport: transport,
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Proxing request to " + r.URL.String())
		proxy.ServeHTTP(w, r)
	})
}
