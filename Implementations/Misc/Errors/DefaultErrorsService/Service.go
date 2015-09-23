package DefaultErrorsService

import (
	"fmt"

	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors"
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors/ClientError"
)

type service struct{}

func (s *service) CreateClientError(statusCode int, statusText string) *ClientError {
	return &ClientError{statusCode, statusText}
}

func (s *service) CreateClientError_Fmt(statusCode int, statusText string, statusTextArgs ...interface{}) *ClientError {
	return s.CreateClientError(statusCode, fmt.Sprintf(statusText, statusTextArgs...))
}

func New() ErrorsService {
	return &service{}
}
