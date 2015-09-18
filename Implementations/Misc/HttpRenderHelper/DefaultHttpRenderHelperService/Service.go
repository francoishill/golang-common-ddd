package DefaultHttpRenderHelperService

import (
	"github.com/unrolled/render"
	"net/http"

	. "github.com/francoishill/golang-common-ddd/Interface/Misc/HttpRenderHelper"
)

type service struct {
	renderer *render.Render
}

func (s *service) RenderJson(w http.ResponseWriter, data interface{}) {
	s.renderer.JSON(w, http.StatusOK, data)
}

func New(isDevelopment, indentJSON bool, templateExtensions []string) HttpRenderHelperService {
	tmpRenderer := render.New(render.Options{
		IsDevelopment: isDevelopment,
		IndentJSON:    indentJSON,
		Extensions:    templateExtensions,
	})
	return &service{
		tmpRenderer,
	}
}
