package DefaultAuthenticationMiddleware

import (
	"net/http"

	. "github.com/francoishill/golang-common-ddd/Interface/Security/Authentication"
)

type service struct {
	AuthenticationService
}

func (s *service) CheckAuthentication(w http.ResponseWriter, r *http.Request) {
	usr := s.AuthenticationService.AuthenticateUserFromRequest(r)
	s.AuthenticationService.SaveUserInRequest(r, usr)
}

func New(authService AuthenticationService) AuthenticationMiddleware {
	return &service{
		authService,
	}
}
