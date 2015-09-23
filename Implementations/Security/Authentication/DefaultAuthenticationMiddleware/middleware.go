package DefaultAuthenticationMiddleware

import (
	"net/http"

	. "github.com/francoishill/golang-common-ddd/Interface/Security/Authentication"
)

type service struct {
	BaseAuthenticationService
}

func (s *service) CheckAuthentication(w http.ResponseWriter, r *http.Request) {
	usr := s.BaseAuthenticationService.BaseAuthenticateUserFromRequest(r)
	s.BaseAuthenticationService.BaseSaveUserInRequest(r, usr)
}

func New(authService BaseAuthenticationService) AuthenticationMiddleware {
	return &service{
		authService,
	}
}
