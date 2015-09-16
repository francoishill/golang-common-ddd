package HttpRenderHelper

import (
	"net/http"
)

type HttpRenderHelperService interface {
	RenderJson(w http.ResponseWriter, data interface{})
}
