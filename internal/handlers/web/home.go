package web

import (
	"html/template"
	"net/http"
)

type PageData struct {
	Title   string
	Content template.HTML
}

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := h.hs.Render(w, "home", "lazts")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
