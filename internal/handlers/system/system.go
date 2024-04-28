package system

import (
	"lazts/internal/modules/http"
	"lazts/internal/services/markdown"
	"lazts/internal/services/templ"
)

type handler struct {
	hs templ.Servicer
	ms markdown.Servicer
}

func New(m http.Module, hs templ.Servicer, ms markdown.Servicer) {
	h := &handler{hs, ms}
	h.setRouter(m)
}

func (h *handler) setRouter(m http.Module) {
	m.Register("GET /markdown", h.Markdown)
	m.Register("GET /static/", h.Static)
}
