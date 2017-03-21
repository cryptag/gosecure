// Steve Phillips / elimisteve
// 2017.03.16

package csp

import (
	"fmt"
	"net/http"
)

func GetHandler(domain string) func(h http.Handler) http.Handler {
	return GetCustomHandler(domain, domain)
}

func GetCustomHandler(domain, apiDomain string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Security-Policy", fmt.Sprintf("default-src 'none'; script-src %s; style-src %[1]s; img-src %[1]s; connect-src %s; child-src 'self'", domain, apiDomain))
			h.ServeHTTP(w, req)
		})
	}
}
