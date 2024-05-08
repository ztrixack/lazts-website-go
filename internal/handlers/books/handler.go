package books

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
	m.Register("GET /books", h.Page)
	m.Register("GET /books/filter", h.Filter)
	m.Register("GET /books/list", h.List)
}
