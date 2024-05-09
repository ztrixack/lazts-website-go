package books

import (
	"net/http"
)

func (h *handler) Page(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/books" {
		http.NotFound(w, r)
		return
	}

	if err := h.page.Render(w, "books"); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
