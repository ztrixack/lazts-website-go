package vacations

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
	m.Register("GET /vacations/", h.Page)

	m.Register("GET /_vacations/highlight", h.Highlight)
	m.Register("GET /_vacations/list", h.List)
	m.Register("GET /_vacations/contents/", h.Content)
}
