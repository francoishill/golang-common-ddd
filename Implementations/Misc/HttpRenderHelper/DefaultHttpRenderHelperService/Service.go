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

func New() HttpRenderHelperService {
	tmpRenderer := render.New(render.Options{
		IndentJSON:    true,
		Extensions:    []string{".gohtml", ".gotmpl"},
		IsDevelopment: true,
	})
	return &service{
		tmpRenderer,
	}
}
