package BasicHttpAuthenticationService

import (
	"encoding/base64"
	"net/http"
	"strings"

	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors"
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/HttpRequestHelper"
	. "github.com/francoishill/golang-common-ddd/Interface/Security/Authentication"
)

type service struct {
	ErrorsService
	HttpRequestHelperService
	AuthUserHelperService
}

const (
	cCONTEXT_USER_KEY = "auth-user-basic"
)

func (s *service) getBasicAuthCredentials(r *http.Request) (email, username, password string) {
	authorizationArray := r.Header["Authorization"]
	if len(authorizationArray) == 0 {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "Unable to extract credentials 1"))
	}

	authorization := strings.TrimSpace(authorizationArray[0])
	credentials := strings.Split(authorization, " ")

	if len(credentials) != 2 || !strings.EqualFold(credentials[0], "Basic") {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "Unable to extract credentials 2"))
	}

	authstr, err := base64.StdEncoding.DecodeString(credentials[1])
	if err != nil {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "Unable to extract credentials 3"))
	}

	usernameAndPassword := strings.Split(string(authstr), ":")
	if len(usernameAndPassword) != 2 {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "Unable to extract credentials 4"))
	}

	email = usernameAndPassword[0]
	username = usernameAndPassword[0]
	password = usernameAndPassword[1]
	return
}

func (s *service) finishHandlerByVerifyingUser(w http.ResponseWriter, user AuthUser) {
	w.WriteHeader(http.StatusOK)
}

func (s *service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	email, username, password := s.getBasicAuthCredentials(r)

	usr := s.AuthUserHelperService.VerifyAndGetUserFromCredentials(email, username, password)
	if usr == nil {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442934196] User does not exist"))
	}

	s.finishHandlerByVerifyingUser(w, usr)
}

func (s *service) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	email, username, password := s.getBasicAuthCredentials(r)

	usr := s.AuthUserHelperService.VerifyAndGetUserFromCredentials(email, username, password)
	if usr == nil {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442934196] User does not exist"))
	}

	s.finishHandlerByVerifyingUser(w, usr)
}

func (s *service) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//We do nothing currently need to clear anything, we have not saved it to cookies
	w.WriteHeader(http.StatusOK)
}

func (s *service) AuthenticateUserFromRequest(r *http.Request) AuthUser {
	email, username, password := s.getBasicAuthCredentials(r)

	usr := s.AuthUserHelperService.VerifyAndGetUserFromCredentials(email, username, password)
	if usr == nil {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442894606] User does not exist"))
	}

	return usr
}

func (s *service) SaveUserInRequest(r *http.Request, user AuthUser) {
	s.HttpRequestHelperService.SaveToRequestContext(r, cCONTEXT_USER_KEY, user)
}

func (s *service) GetBaseUserFromRequest(r *http.Request) AuthUser {
	usr, ok := s.HttpRequestHelperService.LoadFromRequestContext(r, cCONTEXT_USER_KEY)
	if !ok {
		panic(s.ErrorsService.CreateClientError(http.StatusInternalServerError, "[1442936125] Context does not contain user"))
	}

	if authUsr, ok := usr.(AuthUser); !ok {
		panic(s.ErrorsService.CreateClientError(http.StatusInternalServerError, "[1442892567] Invalid user value"))
	} else {
		return authUsr
	}
}

func New(errorsService ErrorsService, httpRequestHelperService HttpRequestHelperService, authUserHelperService AuthUserHelperService) AuthenticationService {
	return &service{
		errorsService,
		httpRequestHelperService,
		authUserHelperService,
	}
}
