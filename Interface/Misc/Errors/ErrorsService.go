package Errors

import (
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors/ClientError"
)

type ErrorsService interface {
	CreateClientError(statusCode int, statusText string) *ClientError

	CreateHttpStatusClientError_BadRequest(statusText string) *ClientError
	CreateHttpStatusClientError_Unauthorized(statusText string) *ClientError
	CreateHttpStatusClientError_PaymentRequired(statusText string) *ClientError
	CreateHttpStatusClientError_Forbidden(statusText string) *ClientError
	CreateHttpStatusClientError_NotFound(statusText string) *ClientError
	CreateHttpStatusClientError_MethodNotAllowed(statusText string) *ClientError
	CreateHttpStatusClientError_NotAcceptable(statusText string) *ClientError
	CreateHttpStatusClientError_ProxyAuthRequired(statusText string) *ClientError
	CreateHttpStatusClientError_RequestTimeout(statusText string) *ClientError
	CreateHttpStatusClientError_Conflict(statusText string) *ClientError
	CreateHttpStatusClientError_Gone(statusText string) *ClientError
	CreateHttpStatusClientError_LengthRequired(statusText string) *ClientError
	CreateHttpStatusClientError_PreconditionFailed(statusText string) *ClientError
	CreateHttpStatusClientError_RequestEntityTooLarge(statusText string) *ClientError
	CreateHttpStatusClientError_RequestURITooLong(statusText string) *ClientError
	CreateHttpStatusClientError_UnsupportedMediaType(statusText string) *ClientError
	CreateHttpStatusClientError_RequestedRangeNotSatisfiable(statusText string) *ClientError
	CreateHttpStatusClientError_ExpectationFailed(statusText string) *ClientError
	CreateHttpStatusClientError_Teapot(statusText string) *ClientError
	CreateHttpStatusClientError_InternalServerError(statusText string) *ClientError
	CreateHttpStatusClientError_NotImplemented(statusText string) *ClientError
	CreateHttpStatusClientError_BadGateway(statusText string) *ClientError
	CreateHttpStatusClientError_ServiceUnavailable(statusText string) *ClientError
	CreateHttpStatusClientError_GatewayTimeout(statusText string) *ClientError
	CreateHttpStatusClientError_HTTPVersionNotSupported(statusText string) *ClientError
}
