package system

import (
	"lazts/internal/modules/http"
	"lazts/internal/modules/http/middlewares"
	"lazts/internal/services/page"
	"lazts/internal/services/watermark"
)

type handler struct {
	hs page.Servicer
	ws watermark.Servicer
}

func New(m http.Moduler, hs page.Servicer, ws watermark.Servicer) {
	h := &handler{hs, ws}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	minify := middlewares.Minify()
	m.Register("GET /static/icons/", h.Icons)
	m.Register("GET /static/images/", h.Images)
	m.Register("GET /static/contents/", h.ImageContents)
	m.Handle("GET /static/js/", minify.MiddlewareWithError(h.Javascript(), h.Error))
	m.Handle("GET /static/css/", minify.MiddlewareWithError(h.CSS(), h.Error))
	m.Register("GET /manifest.json", h.StaticFile)
	m.Register("GET /service-worker.js", h.StaticFile)
}
