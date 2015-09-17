package BasicHttpAuthenticationService

import (
	"encoding/base64"
	. "github.com/francoishill/golang-common-ddd/Interface/Security/Authentication"
	"github.com/gorilla/context"
	"net/http"
	"strings"
)

const (
	cCONTEXT_USER_KEY = "auth-user"
)

type service struct {
	userWithEmailAndPasswordLocator UserWithEmailAndPasswordLocator
}

func (s *service) getBasicAuthCredentials(r *http.Request) (successfullyExtractedDetails bool, username, password string) {
	authorizationArray := r.Header["Authorization"]
	if len(authorizationArray) == 0 {
		return false, "", ""
	}

	authorization := strings.TrimSpace(authorizationArray[0])
	credentials := strings.Split(authorization, " ")

	if len(credentials) != 2 || !strings.EqualFold(credentials[0], "Basic") {
		return false, "", ""
	}

	authstr, err := base64.StdEncoding.DecodeString(credentials[1])
	if err != nil {
		return false, "", ""
	}

	usernameAndPassword := strings.Split(string(authstr), ":")
	if len(usernameAndPassword) != 2 {
		return false, "", ""
	}

	return true, usernameAndPassword[0], usernameAndPassword[1]
}

func (s *service) AuthenticateUserFromRequest(r *http.Request) bool {
	successfullyExtractedDetails, email, password := s.getBasicAuthCredentials(r)
	if !successfullyExtractedDetails {
		return false
	}

	usr := s.userWithEmailAndPasswordLocator.FindUserWithEmailAndPassword(email, password)
	if usr == nil {
		return false
	}

	context.Set(r, cCONTEXT_USER_KEY, usr)
	return true
}

func (s *service) GetStoredUserOfRequest(r *http.Request) interface{} {
	if usr, ok := context.GetOk(r, cCONTEXT_USER_KEY); !ok {
		return nil
	} else {
		return usr
	}
}

func New(userWithEmailAndPasswordLocator UserWithEmailAndPasswordLocator) AuthenticationService {
	return &service{
		userWithEmailAndPasswordLocator,
	}
}
