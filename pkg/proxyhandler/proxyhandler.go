package proxyhandler

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
)

func Create(url string, cert tls.Certificate, skipSSL bool, silent bool) http.HandlerFunc {
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
		if !silent {
			log.Printf("path=%s\n", r.URL.String())
		}
		proxy.ServeHTTP(w, r)
	})
}
