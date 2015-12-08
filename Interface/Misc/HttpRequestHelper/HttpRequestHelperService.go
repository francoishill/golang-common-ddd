package HttpRequestHelper

import (
	"net/http"
)

type HttpRequestHelperService interface {
	OptionalQueryValue(r *http.Request, keyName string) RequestValue
	MustQueryValue(r *http.Request, keyName string) RequestValue

	OptionalUrlParamValue(r *http.Request, paramName string) RequestValue
	MustUrlParamValue(r *http.Request, paramName string) RequestValue

	DecodeJsonRequest(r *http.Request, destination interface{})

	SaveToRequestContext(r *http.Request, key, val interface{})
	LoadFromRequestContext(r *http.Request, key interface{}) (interface{}, bool)
}
