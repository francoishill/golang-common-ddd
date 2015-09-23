package Jwt

import (
	"time"
)

type JwtHelperService interface {
	IsInLoggedOutList(token string) bool
	AddTokenToLoggedOutList(token string, tokenExpiry time.Time)
}
