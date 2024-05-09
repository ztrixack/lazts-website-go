package http

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Module interface {
	Config() config
	Register(path string, handler http.HandlerFunc)
	Handle(path string, handler http.Handler)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Serve() error
}

type module struct {
	config *config
	*router
}

func New(config *config) *module {
	return &module{
		config: config,
		router: NewRouter(),
	}
}

func (m *module) Config() config {
	return *m.config
}

func (m *module) Register(path string, handler http.HandlerFunc) {
	m.router.HandleFunc(path, handler)
}

func (m *module) Handle(path string, handler http.Handler) {
	m.router.Handle(path, handler)
}

func (m *module) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}

func (m *module) Serve() error {
	server := &http.Server{
		Addr:    ":" + m.config.Port,
		Handler: m.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			fmt.Printf("Error starting server: %s\n", err)
		}
	}()

	<-stop
	fmt.Println("Received stop signal. Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server: %s\n", err)
		return err
	}

	fmt.Println("Server shutdown completed")
	return nil
}

func (m *module) Use(ms ...MiddlewareFunc) {
	m.router.Use(ms...)
}
