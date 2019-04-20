package proxyserver

import "net/http"

func NewServer(mux *http.ServeMux, handler http.Handler) {
	mux.Handle("/", handler)
}
