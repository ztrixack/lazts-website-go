package system

import "net/http"

func (h *handler) Icons(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fs := http.FileServer(http.Dir("web/static/icons"))
	http.StripPrefix("/static/icons", fs).ServeHTTP(w, r)
}
