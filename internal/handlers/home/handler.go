package home

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/book"
	"lazts/internal/services/note"
	"lazts/internal/services/page"
	"lazts/internal/services/vacation"
)

type handler struct {
	pager      page.Servicer
	booker     book.Servicer
	vacationer vacation.Servicer
	noter      note.Servicer
}

func New(m http.Module, ps page.Servicer, bs book.Servicer, vs vacation.Servicer, ns note.Servicer) {
	h := &handler{ps, bs, vs, ns}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Module) {
	// page
	m.Register("GET /", h.Home)

	// partials
	m.Register("GET /_home/blackhole", h.Blackhole)
	m.Register("GET /_home/cloud", h.Cloud)
	m.Register("GET /_home/books", h.Book)
	m.Register("GET /_home/vacations", h.Vacations)
	m.Register("GET /_home/notes", h.Notes)
}
