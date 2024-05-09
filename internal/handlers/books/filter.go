package books

import (
	"net/http"
)

func (h *handler) Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	search := r.URL.Query().Get("search")
	catalog := r.URL.Query().Get("catalog")
	status := r.URL.Query().Get("status")

	if err := h.book.RenderFilter(w, search, catalog, status); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
