package home

import (
	"net/http"
)

func (h *handler) Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	err := h.hs.Render(w, "home")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
