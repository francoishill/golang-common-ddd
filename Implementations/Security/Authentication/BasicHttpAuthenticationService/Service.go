package BasicHttpBaseAuthenticationService

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
	BaseAuthUserHelperService
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

func (s *service) finishHandlerByVerifyingUser(w http.ResponseWriter, user BaseUser) {
	w.WriteHeader(http.StatusOK)
}

func (s *service) BaseLoginHandler(w http.ResponseWriter, r *http.Request) {
	email, username, password := s.getBasicAuthCredentials(r)

	usr := s.BaseAuthUserHelperService.BaseVerifyAndGetUserFromCredentials(email, username, password)
	if usr == nil {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442934196] User does not exist"))
	}

	s.finishHandlerByVerifyingUser(w, usr)
}

func (s *service) BaseRegisterHandler(w http.ResponseWriter, r *http.Request) {
	email, username, password := s.getBasicAuthCredentials(r)

	usr := s.BaseAuthUserHelperService.BaseVerifyAndGetUserFromCredentials(email, username, password)
	if usr == nil {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442934196] User does not exist"))
	}

	s.finishHandlerByVerifyingUser(w, usr)
}

func (s *service) BaseLogoutHandler(w http.ResponseWriter, r *http.Request) {
	//We do nothing currently need to clear anything, we have not saved it to cookies
	w.WriteHeader(http.StatusOK)
}

func (s *service) BaseAuthenticateUserFromRequest(r *http.Request) BaseUser {
	email, username, password := s.getBasicAuthCredentials(r)

	usr := s.BaseAuthUserHelperService.BaseVerifyAndGetUserFromCredentials(email, username, password)
	if usr == nil {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442894606] User does not exist"))
	}

	return usr
}

func (s *service) BaseSaveUserInRequest(r *http.Request, user BaseUser) {
	s.HttpRequestHelperService.SaveToRequestContext(r, cCONTEXT_USER_KEY, user)
}

func (s *service) BaseGetUserFromRequest(r *http.Request) BaseUser {
	usr, ok := s.HttpRequestHelperService.LoadFromRequestContext(r, cCONTEXT_USER_KEY)
	if !ok {
		panic(s.ErrorsService.CreateClientError(http.StatusInternalServerError, "[1442936125] Context does not contain user"))
	}

	if authUsr, ok := usr.(BaseUser); !ok {
		panic(s.ErrorsService.CreateClientError(http.StatusInternalServerError, "[1442892567] Invalid user value"))
	} else {
		return authUsr
	}
}

func New(errorsService ErrorsService, httpRequestHelperService HttpRequestHelperService, baseAuthUserHelperService BaseAuthUserHelperService) BaseAuthenticationService {
	return &service{
		errorsService,
		httpRequestHelperService,
		baseAuthUserHelperService,
	}
}
