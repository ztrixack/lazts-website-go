package books

import (
	"net/http"
)

func (h *handler) Page(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/books" {
		http.NotFound(w, r)
		return
	}

	err := h.hs.Render(w, "books")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
