package home

import (
	"net/http"
)

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if err := h.pager.Render(w, "home"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
