package DefaultErrorsService

import (
	"fmt"
	. "github.com/francoishill/golang-common-ddd/Interface/Misc/Errors"
)

type defaultErrorService struct{}

type tmpClientError struct {
	Error interface{}
}

func (d *defaultErrorService) PanicClientErrorLocal(e interface{}) {
	panic(&tmpClientError{e})
}

func (d *defaultErrorService) PanicClientErrorLocal_FormattedString(format string, args ...interface{}) {
	d.PanicClientErrorLocal(fmt.Errorf(format, args...))
}

func New() ErrorsService {
	return &defaultErrorService{}
}
