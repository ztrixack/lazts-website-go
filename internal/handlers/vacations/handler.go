package vacations

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/page"
	"lazts/internal/services/vacation"
)

type handler struct {
	pager      page.Servicer
	vacationer vacation.Servicer
}

func New(m http.Module, pager page.Servicer, vacationer vacation.Servicer) {
	h := &handler{pager, vacationer}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Module) {
	// page
	m.Register("GET /vacations/", h.Page)

	// partials
	m.Register("GET /_vacations/highlight", h.Highlight)
	m.Register("GET /_vacations/list", h.List)
	m.Register("GET /_vacations/contents/", h.Content)
}
