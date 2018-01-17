// Steve Phillips / elimisteve
// 2017.03.20

package frame

import (
	"fmt"
	"net/http"

	"github.com/cryptag/gosecure/set"
)

func GetHandler(domain, port string) func(h http.Handler) http.Handler {
	var maybeColonPort string
	if port != "" {
		maybeColonPort = ":" + port
	}

	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("X-Frame-Options", fmt.Sprintf("ALLOW-FROM https://%s%s", domain, maybeColonPort))
			h.ServeHTTP(w, req)
		})
	}
}

// SameOriginHandler is middleware that sets the header `X-Frame-Options: SAMEORIGIN`.
func SameOriginHandler(h http.Handler) http.Handler {
	return set.Header(h, "X-Frame-Options", "SAMEORIGIN")
}

// DenyHandler is middleware that sets the header `X-Frame-Options: DENY`.
func DenyHandler(h http.Handler) http.Handler {
	return set.Header(h, "X-Frame-Options", "DENY")
}
