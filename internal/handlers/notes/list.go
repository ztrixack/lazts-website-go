package notes

import (
	"net/http"
)

func (h *handler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := h.noter.RenderList(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
