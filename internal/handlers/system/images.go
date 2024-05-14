package system

import (
	"image/jpeg"
	"lazts/internal/utils"
	"net/http"
	"strings"
)

func (h *handler) Images(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if !strings.HasSuffix(r.URL.Path, ".jpeg") {
		fs := http.FileServer(http.Dir("static/images"))
		http.StripPrefix("/static/images", fs).ServeHTTP(w, r)
		return
	}

	imagePath := utils.GetStaticDir("images", r.URL.Path)
	img, err := h.ws.LoadImage(imagePath)
	if err != nil {
		http.Error(w, "Failed to load image", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/jpeg")

	jpeg.Encode(w, img, nil)
}
