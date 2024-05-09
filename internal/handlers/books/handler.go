package books

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/book"
	"lazts/internal/services/page"
)

type handler struct {
	page page.Servicer
	book book.Servicer
}

func New(m http.Module, page page.Servicer, book book.Servicer) {
	h := &handler{page, book}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Module) {
	// page
	m.Register("GET /books", h.Page)

	// partials
	m.Register("GET /_books/filter", h.Filter)
	m.Register("GET /_books/list", h.List)
}
