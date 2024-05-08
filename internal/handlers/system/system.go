package system

import (
	"lazts/internal/modules/http"
	"lazts/internal/modules/http/middlewares"
	"lazts/internal/services/templ"
)

type handler struct {
	hs templ.Servicer
}

func New(m http.Module, hs templ.Servicer) {
	h := &handler{hs}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Module) {
	minify := middlewares.Minify()
	m.Register("GET /static/images/", h.Images)
	m.Register("GET /static/notes/", h.NoteContent)
	m.Handle("GET /static/js/", minify.MiddlewareWithError(h.Javascript(), h.Error))
	m.Handle("GET /static/css/", minify.MiddlewareWithError(h.CSS(), h.Error))
}
