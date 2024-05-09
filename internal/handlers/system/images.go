package system

import "net/http"

func (h *handler) Images(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fs := http.FileServer(http.Dir("static/images"))
	http.StripPrefix("/static/images", fs).ServeHTTP(w, r)
}
