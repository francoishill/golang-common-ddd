package Authentication

import (
	"net/http"
)

type AuthenticationMiddleware interface {
	CheckAuthentication(w http.ResponseWriter, r *http.Request)
}
