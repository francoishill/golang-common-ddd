package DefaultErrorsService

import (
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors"
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors/ClientError"
	"net/http"
	"strings"
)

type service struct{}

func (s *service) CreateClientError(statusCode int, statusText string) *ClientError {
	return &ClientError{statusCode, statusText}
}

func (s *service) createClientErrorWithDefaultStatusTextIfEmpty(statusCode int, statusText string) *ClientError {
	if strings.TrimSpace(statusText) == "" {
		return s.CreateClientError(statusCode, http.StatusText(statusCode))
	}
	return s.CreateClientError(statusCode, statusText)
}

func (s *service) CreateHttpStatusClientError_BadRequest(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusBadRequest, statusText)
}

func (s *service) CreateHttpStatusClientError_Unauthorized(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusUnauthorized, statusText)
}

func (s *service) CreateHttpStatusClientError_PaymentRequired(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusPaymentRequired, statusText)
}

func (s *service) CreateHttpStatusClientError_Forbidden(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusForbidden, statusText)
}

func (s *service) CreateHttpStatusClientError_NotFound(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusNotFound, statusText)
}

func (s *service) CreateHttpStatusClientError_MethodNotAllowed(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusMethodNotAllowed, statusText)
}

func (s *service) CreateHttpStatusClientError_NotAcceptable(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusNotAcceptable, statusText)
}

func (s *service) CreateHttpStatusClientError_ProxyAuthRequired(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusProxyAuthRequired, statusText)
}

func (s *service) CreateHttpStatusClientError_RequestTimeout(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusRequestTimeout, statusText)
}

func (s *service) CreateHttpStatusClientError_Conflict(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusConflict, statusText)
}

func (s *service) CreateHttpStatusClientError_Gone(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusGone, statusText)
}

func (s *service) CreateHttpStatusClientError_LengthRequired(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusLengthRequired, statusText)
}

func (s *service) CreateHttpStatusClientError_PreconditionFailed(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusPreconditionFailed, statusText)
}

func (s *service) CreateHttpStatusClientError_RequestEntityTooLarge(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusRequestEntityTooLarge, statusText)
}

func (s *service) CreateHttpStatusClientError_RequestURITooLong(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusRequestURITooLong, statusText)
}

func (s *service) CreateHttpStatusClientError_UnsupportedMediaType(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusUnsupportedMediaType, statusText)
}

func (s *service) CreateHttpStatusClientError_RequestedRangeNotSatisfiable(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusRequestedRangeNotSatisfiable, statusText)
}

func (s *service) CreateHttpStatusClientError_ExpectationFailed(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusExpectationFailed, statusText)
}

func (s *service) CreateHttpStatusClientError_Teapot(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusTeapot, statusText)
}

func (s *service) CreateHttpStatusClientError_InternalServerError(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusInternalServerError, statusText)
}

func (s *service) CreateHttpStatusClientError_NotImplemented(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusNotImplemented, statusText)
}

func (s *service) CreateHttpStatusClientError_BadGateway(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusBadGateway, statusText)
}

func (s *service) CreateHttpStatusClientError_ServiceUnavailable(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusServiceUnavailable, statusText)
}

func (s *service) CreateHttpStatusClientError_GatewayTimeout(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusGatewayTimeout, statusText)
}

func (s *service) CreateHttpStatusClientError_HTTPVersionNotSupported(statusText string) *ClientError {
	return s.createClientErrorWithDefaultStatusTextIfEmpty(http.StatusHTTPVersionNotSupported, statusText)
}

func New() ErrorsService {
	return &service{}
}
