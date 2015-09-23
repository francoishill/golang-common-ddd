package Errors

import (
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors/ClientError"
)

type ErrorsService interface {
	CreateClientError(statusCode int, statusText string) *ClientError
	CreateClientError_Fmt(statusCode int, statusText string, statusTextArgs ...interface{}) *ClientError
}
