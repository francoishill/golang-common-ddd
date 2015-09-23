package Authentication

import (
	"net/http"
)

type BaseAuthenticationService interface {
	BaseLoginHandler(w http.ResponseWriter, r *http.Request)
	BaseRegisterHandler(w http.ResponseWriter, r *http.Request)
	BaseLogoutHandler(w http.ResponseWriter, r *http.Request)
	BaseAuthenticateUserFromRequest(r *http.Request) BaseUser

	BaseSaveUserInRequest(r *http.Request, user BaseUser)
	BaseGetUserFromRequest(r *http.Request) BaseUser
}
