package middlewares

import (
	mux "lazts/internal/modules/http"
	"lazts/internal/modules/log"
	"net/http"
	"time"
)

func Logger(log log.Moduler) mux.MiddlewareFunc {
	return func(next http.Handler) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// wr := &loggableResponseWriter{w, http.StatusOK, time.Now()}
			next.ServeHTTP(w, r)

			// log.Fields(
			// 	"remote-addr", r.RemoteAddr,
			// 	"host", r.Host,
			// 	"method", r.Method,
			// 	"uri", r.RequestURI,
			// 	"status", wr.statusCode,
			// 	"latency", time.Since(wr.time),
			// ).I("Request completed")
		})
	}
}

type loggableResponseWriter struct {
	http.ResponseWriter
	statusCode int
	time       time.Time
}

func (lrw *loggableResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}
