package DefaultHttpRequestHelperService

import (
	"encoding/json"
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"

	. "github.com/francoishill/golang-common-ddd/Interface/Misc/HttpRequestHelper"
)

type service struct{}

func parseInt64(s string) int64 {
	intVal, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return intVal
}

func parseFloat64(s string) float64 {
	floatVal, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return floatVal
}

func (s *service) OptionalQueryValue(r *http.Request, keyName string) RequestValue {
	val := strings.TrimSpace(r.URL.Query().Get(keyName))
	return NewRequestValue(val != "", RequestValue_Query, val)
}
func (s *service) MustQueryValue(r *http.Request, keyName string) RequestValue {
	val := s.OptionalQueryValue(r, keyName)
	if !val.HasValue() {
		panic(keyName + " cannot be found from URL query values")
	}
	return val
}

func (s *service) OptionalUrlParamValue(r *http.Request, paramName string) RequestValue {
	vars := mux.Vars(r)
	paramValue, varFound := vars[paramName]
	if !varFound {
		return NewRequestValue(false, RequestValue_Param, "")
	}
	trimmedVal := strings.TrimSpace(paramValue)
	return NewRequestValue(trimmedVal != "", RequestValue_Param, trimmedVal)
}
func (s *service) MustUrlParamValue(r *http.Request, paramName string) RequestValue {
	paramValue := s.OptionalUrlParamValue(r, paramName)
	if !paramValue.HasValue() {
		panic(paramName + " cannot be found from URL route params")
	}
	return paramValue
}

func (s *service) DecodeJsonRequest(r *http.Request, destination interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(destination)
	CheckError(err)
}

func (s *service) SaveToRequestContext(r *http.Request, key, val interface{}) {
	context.Set(r, key, val)
}

func (s *service) LoadFromRequestContext(r *http.Request, key interface{}) (interface{}, bool) {
	return context.GetOk(r, key)
}

func New() HttpRequestHelperService {
	return &service{}
}
