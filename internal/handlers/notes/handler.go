package notes

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/note"
	"lazts/internal/services/page"
)

type handler struct {
	pager page.Servicer
	noter note.Servicer
}

func New(m http.Module, ps page.Servicer, ns note.Servicer) {
	h := &handler{ps, ns}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Module) {
	// page
	m.Register("GET /notes/", h.Page)

	// partials
	m.Register("GET /_notes/tags", h.Tags)
	m.Register("GET /_notes/list", h.List)
	m.Register("GET /_notes/count", h.Count)
	m.Register("GET /_notes/headers/", h.Header)
	m.Register("GET /_notes/contents/", h.Content)
}
