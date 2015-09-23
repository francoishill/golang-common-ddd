package DefaultJwtHelperService

import (
	"time"

	. "github.com/francoishill/golang-common-ddd/Interface/Security/Authentication/Jwt"
	. "github.com/francoishill/golang-common-ddd/Interface/Storage/KeyValueStorage"
)

type service struct {
	KeyValueStorage
}

const (
	KEYVAL_TOKEN_PREFIX    = "LOGGED_OUT_TOKEN"
	cEXPIRE_OFFSET_SECONDS = 3600
)

func (s *service) IsInLoggedOutList(token string) bool {
	fullKeyName := KEYVAL_TOKEN_PREFIX + token
	_, err := s.KeyValueStorage.Get(fullKeyName)
	return err == nil
}

func (s *service) AddTokenToLoggedOutList(token string, tokenExpiry time.Time) {
	diff := tokenExpiry.Sub(time.Now())

	var valueExpiryDuration time.Duration
	if diff > 0 {
		valueExpiryDuration = time.Duration(diff.Seconds() + float64(cEXPIRE_OFFSET_SECONDS))
	} else {
		valueExpiryDuration = time.Duration(cEXPIRE_OFFSET_SECONDS)
	}

	fullKeyName := KEYVAL_TOKEN_PREFIX + token
	s.KeyValueStorage.MustSet(fullKeyName, token, valueExpiryDuration)
}

func New(keyValueStorage KeyValueStorage) JwtHelperService {
	return &service{
		keyValueStorage,
	}
}
