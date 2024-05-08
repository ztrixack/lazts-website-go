package notes

import (
	"lazts/internal/modules/http"
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
	m.Register("GET /notes/", h.Page)

	m.Register("GET /_notes/tags", h.Tags)
	m.Register("GET /_notes/list", h.List)
	m.Register("GET /_notes/contents/", h.Content)
}
