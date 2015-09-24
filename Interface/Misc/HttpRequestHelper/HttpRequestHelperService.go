package HttpRequestHelper

import (
	"net/http"
)

type HttpRequestHelperService interface {
	GetRequiredUrlQueryValue_String(r *http.Request, keyName string) string
	GetRequiredUrlQueryValue_Int64(r *http.Request, keyName string) int64

	GetRequiredUrlParamValue_String(r *http.Request, paramName string) string
	GetRequiredUrlParamValue_Int64(r *http.Request, paramName string) int64

	DecodeJsonRequest(r *http.Request, destination interface{})

	SaveToRequestContext(r *http.Request, key, val interface{})
	LoadFromRequestContext(r *http.Request, key interface{}) (interface{}, bool)
}
