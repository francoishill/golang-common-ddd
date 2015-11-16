package DefaultJwtHelperService

import (
	"time"

	. "github.com/francoishill/golang-common-ddd/Interface/Security/Authentication/Jwt"
	. "github.com/francoishill/golang-common-ddd/Interface/Storage/KeyValueStorage"
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors"
	. "github.com/francoishill/golang-common-ddd/Interface/Logger"
	"strings"
)

type service struct {
	logger Logger
	errors ErrorsService
	KeyValueStorage
}

const (
	KEYVAL_TOKEN_PREFIX = "LOGGED_OUT_TOKEN"
	cEXPIRE_OFFSET_SECONDS = 3600
)

func (s *service) IsInLoggedOutList(token string) bool {
	fullKeyName := KEYVAL_TOKEN_PREFIX + token
	_, err := s.KeyValueStorage.Get(fullKeyName)
	if strings.Contains(strings.ToLower(err.Error()), "no connection could be made") ||
	strings.Contains(strings.ToLower(err.Error()), "existing connection was forcibly closed") {
		s.logger.Critical("Unable to connect to Redis server, error: %s", err.Error())
		panic(s.errors.CreateHttpStatusClientError_InternalServerError("[1447694837] Unable to connect to an internal service. Please try again in a few minutes."))
	} else if err.Error() == "redis: nil" {
		return false
	} else if err == nil {
		return true
	} else {
		s.logger.Critical("Unexpected error with Redis GET, error: %s", err.Error())
		panic(s.errors.CreateHttpStatusClientError_InternalServerError("[1447694838] Unexpected error with an internal service. Please try again in a few minutes."))
	}
}

func (s *service) AddTokenToLoggedOutList(token string, tokenExpiry time.Time) {
	diff := tokenExpiry.Sub(time.Now())

	var valueExpiryDuration time.Duration
	if diff > 0 {
		valueExpiryDuration = time.Duration(diff.Seconds() + float64(cEXPIRE_OFFSET_SECONDS)) * time.Second
	} else {
		valueExpiryDuration = cEXPIRE_OFFSET_SECONDS * time.Second
	}

	fullKeyName := KEYVAL_TOKEN_PREFIX + token
	s.KeyValueStorage.MustSet(fullKeyName, token, valueExpiryDuration)
}

func New(logger Logger, errors ErrorsService, keyValueStorage KeyValueStorage) JwtHelperService {
	return &service{
		logger,
		errors,
		keyValueStorage,
	}
}
