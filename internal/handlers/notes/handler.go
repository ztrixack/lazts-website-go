package notes

import (
	"lazts/internal/modules/http"
	"lazts/internal/modules/log"
	"lazts/internal/services/note"
	"lazts/internal/services/page"
)

type handler struct {
	logger log.Moduler
	pager  page.Servicer
	noter  note.Servicer
}

func New(l log.Moduler, m http.Moduler, ps page.Servicer, ns note.Servicer) {
	h := &handler{l, ps, ns}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Moduler) {
	// page
	m.Register("GET /notes/", h.Page)

	// partials
	m.Register("GET /_notes/tags", h.Tags)
	m.Register("GET /_notes/list", h.List)
	m.Register("GET /_notes/count", h.Count)
	m.Register("GET /_notes/headers/", h.Header)
	m.Register("GET /_notes/contents/", h.Content)
}
