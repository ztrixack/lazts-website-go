package system

import (
	"lazts/internal/modules/http"
	"lazts/internal/modules/http/middlewares"
	"lazts/internal/services/page"
)

type handler struct {
	hs page.Servicer
}

func New(m http.Module, hs page.Servicer) {
	h := &handler{hs}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Module) {
	minify := middlewares.Minify()
	m.Register("GET /static/icons/", h.Icons)
	m.Register("GET /static/images/", h.Images)
	m.Handle("GET /static/js/", minify.MiddlewareWithError(h.Javascript(), h.Error))
	m.Handle("GET /static/css/", minify.MiddlewareWithError(h.CSS(), h.Error))
	m.Register("GET /manifest.json", h.StaticFile)
	m.Register("GET /service-worker.js", h.StaticFile)
}
