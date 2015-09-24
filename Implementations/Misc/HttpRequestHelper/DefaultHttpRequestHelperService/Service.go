package DefaultHttpRequestHelperService

import (
	"encoding/json"
	. "github.com/francoishill/golang-web-dry/errors/checkerror"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

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

func (s *service) GetRequiredUrlQueryValue_String(r *http.Request, keyName string) string {
	val := r.URL.Query().Get(keyName)
	if val == "" {
		panic(keyName + " cannot be found from URL")
	}
	return val
}

func (s *service) GetRequiredUrlQueryValue_Int64(r *http.Request, keyName string) int64 {
	return parseInt64(s.GetRequiredUrlQueryValue_String(r, keyName))
}

func (s *service) GetRequiredUrlParamValue_String(r *http.Request, paramName string) string {
	vars := mux.Vars(r)
	paramValue, varFound := vars[paramName]
	if !varFound {
		panic(paramName + " cannot be found from URL")
	}
	return paramValue
}

func (s *service) GetRequiredUrlParamValue_Int64(r *http.Request, paramName string) int64 {
	return parseInt64(s.GetRequiredUrlParamValue_String(r, paramName))
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
