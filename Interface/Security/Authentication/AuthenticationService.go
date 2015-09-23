package Authentication

import (
	"net/http"
)

type AuthenticationService interface {
	LoginHandler(w http.ResponseWriter, r *http.Request)
	RegisterHandler(w http.ResponseWriter, r *http.Request)
	LogoutHandler(w http.ResponseWriter, r *http.Request)
	AuthenticateUserFromRequest(r *http.Request) AuthUser

	SaveUserInRequest(r *http.Request, user AuthUser)
	GetBaseUserFromRequest(r *http.Request) AuthUser
}
