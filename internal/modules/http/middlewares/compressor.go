package middlewares

import (
	"compress/gzip"
	"io"
	mux "lazts/internal/modules/http"
	"net/http"
	"strings"
)

func Compressor() mux.MiddlewareFunc {
	return func(next http.Handler) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				gz := gzip.NewWriter(w)
				w.Header().Set("Content-Encoding", "gzip")
				next.ServeHTTP(&compressResponseWriter{w, gz}, r)
				gz.Close()
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}

type compressResponseWriter struct {
	http.ResponseWriter
	writer io.Writer
}

func (crw *compressResponseWriter) Write(data []byte) (int, error) {
	return crw.writer.Write(data)
}
