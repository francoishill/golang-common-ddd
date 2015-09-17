package Authentication

import (
	"net/http"
)

type AuthenticationService interface {
	AuthenticateUserFromRequest(r *http.Request) bool
	GetStoredUserOfRequest(r *http.Request) interface{}
}
