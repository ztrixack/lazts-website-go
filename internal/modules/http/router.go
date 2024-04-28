package http

import (
	"net/http"
)

type MiddlewareFunc func(http.Handler) http.HandlerFunc

type router struct {
	mux         *http.ServeMux
	middlewares []MiddlewareFunc
}

func NewRouter() *router {
	return &router{
		mux:         http.NewServeMux(),
		middlewares: make([]MiddlewareFunc, 0),
	}
}

func (r *router) Handle(pattern string, handler http.Handler) {
	r.mux.Handle(pattern, handler)
}

func (r *router) HandleFunc(pattern string, handler http.HandlerFunc) {
	r.mux.HandleFunc(pattern, handler)
}

func (r *router) Use(middlewares ...MiddlewareFunc) {
	r.middlewares = append(r.middlewares, middlewares...)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	finalHandler := r.mux.ServeHTTP
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		finalHandler = r.middlewares[i](http.HandlerFunc(finalHandler))
	}
	finalHandler(w, req)
}
