package JWTAuthenticationService

import (
	"encoding/json"
	jwt "github.com/dgrijalva/jwt-go"
	. "github.com/francoishill/golang-common-ddd/Interface/Logger"
	"net/http"
	"time"

	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors"
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/HttpRequestHelper"
	. "github.com/francoishill/golang-common-ddd/Interface/Security/Authentication"
	. "github.com/francoishill/golang-common-ddd/Interface/Security/Authentication/Jwt"
)

type service struct {
	Logger

	privateKeyBytes         []byte
	publicKeyBytes          []byte
	jwtExpirationDeltaHours int

	ErrorsService
	HttpRequestHelperService
	JwtHelperService

	AuthUserHelperService
}

const (
	cCONTEXT_USER_KEY = "auth-user-jwt"
)

type authUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type tokenAuthenticationWrapper struct {
	Token string `json:"token" form:"token"`
}

func (s *service) jwtKeyFuncGetBytes(tok *jwt.Token) (interface{}, error) {
	return s.publicKeyBytes, nil
}

func (s *service) getRequestCredentials(r *http.Request) (email, username, password string) {
	tmpUser := new(authUser)
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&tmpUser)
	if err != nil {
		panic(s.ErrorsService.CreateClientError(http.StatusBadRequest, "[1442894150] Cannot find credentials"))
	}
	return tmpUser.Email, tmpUser.Username, tmpUser.Password
}

func (s *service) generateToken(user AuthUser) string {
	//This signing method must be the one that we validate below by calling `if _, ok := token.Method.(*jwt.SigningMethod`
	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["exp"] = time.Now().Add(time.Hour * time.Duration(s.jwtExpirationDeltaHours)).Unix()
	token.Claims["iat"] = time.Now().Unix()
	token.Claims["sub"] = user.UUID()
	tokenString, err := token.SignedString(s.privateKeyBytes)
	if err != nil {
		s.Logger.Error("Cannot generate JWT token: %s", err.Error())
		panic(s.ErrorsService.CreateClientError(http.StatusBadRequest, "[1442894957] Token error"))
	}
	return tokenString
}

func (s *service) getResponseBytesFromToken(token string) []byte {
	response, err := json.Marshal(&tokenAuthenticationWrapper{token})
	if err != nil {
		s.Logger.Error("Cannot marshal JWT token: %s", err.Error())
		panic(s.ErrorsService.CreateClientError(http.StatusBadRequest, "[1442895069] Token error"))
	}
	return response
}

func (s *service) getAndValidateJwtTokenFromRequest(r *http.Request) *jwt.Token {
	token, err := jwt.ParseFromRequest(r, s.jwtKeyFuncGetBytes)

	switch errType := err.(type) {
	case *jwt.ValidationError:
		switch errType.Errors {
		case jwt.ValidationErrorExpired:
			panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "Token expired"))

		default:
			s.Logger.Error("Unable to parse token: %s", err.Error())
			panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442894612] Invalid token"))
		}
	case nil:
		//This validation must match up with the one above in call to `jwt.New(jwt.GetSigningMethod(`
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442894613] Invalid token"))
		}

		if !token.Valid || s.JwtHelperService.IsInLoggedOutList(token.Raw) {
			panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442894611] Invalid token"))
		}

		return token
	default: // something else went wrong
		if err == jwt.ErrNoTokenInRequest {
			panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442947885] Token missing"))
		}

		s.Logger.Error("Error casting token error type: %s", errType)
		panic(s.ErrorsService.CreateClientError(http.StatusInternalServerError, "[1442936515] Invalid token"))
	}
}

func (s *service) finishHandlerByVerifyingUser(w http.ResponseWriter, user AuthUser) {
	token := s.generateToken(user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(s.getResponseBytesFromToken(token))
}

func (s *service) LoginHandler(w http.ResponseWriter, r *http.Request) {
	email, username, password := s.getRequestCredentials(r)
	usr := s.AuthUserHelperService.VerifyAndGetUserFromCredentials(email, username, password)
	if usr == nil {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442894607] User does not exist"))
	}

	s.finishHandlerByVerifyingUser(w, usr)
}

func (s *service) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	email, username, password := s.getRequestCredentials(r)
	usr := s.AuthUserHelperService.RegisterUser(email, username, password)
	s.finishHandlerByVerifyingUser(w, usr)
}

func (s *service) getTokenExpiry(token *jwt.Token) time.Time {
	timestamp := token.Claims["exp"]
	if validity, ok := timestamp.(float64); ok {
		return time.Unix(int64(validity), 0)
	} else {
		return time.Now()
	}
}

func (s *service) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	token := s.getAndValidateJwtTokenFromRequest(r)
	s.JwtHelperService.AddTokenToLoggedOutList(token.Raw, s.getTokenExpiry(token))
}

func (s *service) AuthenticateUserFromRequest(r *http.Request) AuthUser {
	token := s.getAndValidateJwtTokenFromRequest(r)

	usr := s.AuthUserHelperService.GetUserWithUUID(token.Claims["sub"])
	if usr == nil {
		panic(s.ErrorsService.CreateClientError(http.StatusUnauthorized, "[1442894608] User does not exist"))
	}

	return usr
}

func (s *service) SaveUserInRequest(r *http.Request, user AuthUser) {
	s.HttpRequestHelperService.SaveToRequestContext(r, cCONTEXT_USER_KEY, user)
}

func (s *service) GetUserFromRequest(r *http.Request) AuthUser {
	usr, ok := s.HttpRequestHelperService.LoadFromRequestContext(r, cCONTEXT_USER_KEY)
	if !ok {
		panic(s.ErrorsService.CreateClientError(http.StatusInternalServerError, "[1442936126] Context does not contain user"))
	}

	if authUsr, ok := usr.(AuthUser); !ok {
		panic(s.ErrorsService.CreateClientError(http.StatusInternalServerError, "[1442892581] Invalid user value"))
	} else {
		return authUsr
	}
}

func New(logger Logger, privateKeyBytes, publicKeyBytes []byte, jwtExpirationDeltaHours int, errorsService ErrorsService, httpRequestHelperService HttpRequestHelperService, jwtHelperService JwtHelperService, authUserHelperService AuthUserHelperService) AuthenticationService {
	return &service{
		logger,
		privateKeyBytes,
		publicKeyBytes,
		jwtExpirationDeltaHours,
		errorsService,
		httpRequestHelperService,
		jwtHelperService,
		authUserHelperService,
	}
}
