package home

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
	m.Register("GET /", h.Home)
	m.Register("GET /home/books", h.Book)
	m.Register("GET /home/cloud", h.Cloud)
	m.Register("GET /home/blackhole", h.Blackhole)
	m.Register("GET /home/vacations", h.Vacations)
	m.Register("GET /home/notes", h.Notes)
}
