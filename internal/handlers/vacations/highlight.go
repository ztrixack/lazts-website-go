package vacations

import (
	"net/http"
)

func (h *handler) Highlight(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	
	if err := h.vacationer.RenderHighlight(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
