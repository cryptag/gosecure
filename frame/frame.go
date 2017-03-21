// Steve Phillips / elimisteve
// 2017.03.20

package frame

import (
	"net/http"

	"github.com/cryptag/gosecure/set"
)

// GetHandler is middleware that sets the header `X-Frame-Options: SAMEORIGIN`.
func GetHandler(h http.Handler) http.Handler {
	return set.Header(h, "X-Frame-Options", "SAMEORIGIN")
}
