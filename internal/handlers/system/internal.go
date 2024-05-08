package system

import "net/http"

func (h *handler) serveStaticFiles(directory string) http.Handler {
	return http.FileServer(http.Dir(directory))
}

func (h *handler) Javascript() http.Handler {
	fs := h.serveStaticFiles("static/js")
	return http.StripPrefix("/static/js", fs)
}

func (h *handler) CSS() http.Handler {
	fs := h.serveStaticFiles("static/css")
	return http.StripPrefix("/static/css", fs)
}

func (h *handler) Images(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fs := http.FileServer(http.Dir("static/images"))
	http.StripPrefix("/static/images", fs).ServeHTTP(w, r)
}

func (h *handler) NoteContent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fs := http.FileServer(http.Dir("contents/notes"))
	http.StripPrefix("/static/notes", fs).ServeHTTP(w, r)
}

func (h *handler) Error(w http.ResponseWriter, r *http.Request, err error) {
	http.Error(w, "Resource not allowed", http.StatusMethodNotAllowed)
}
