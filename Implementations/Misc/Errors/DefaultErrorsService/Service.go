package DefaultErrorsService

import (
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors"
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors/ClientError"
)

type service struct{}

func (s *service) CreateClientError(statusCode int, statusText string) *ClientError {
	return &ClientError{statusCode, statusText}
}

func New() ErrorsService {
	return &service{}
}
