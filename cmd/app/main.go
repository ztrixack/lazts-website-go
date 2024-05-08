package main

import (
	"lazts/internal/handlers/books"
	"lazts/internal/handlers/home"
	"lazts/internal/handlers/notes"
	"lazts/internal/handlers/system"
	"lazts/internal/handlers/vacations"
	"lazts/internal/modules/http"
	"lazts/internal/modules/http/middlewares"
	"lazts/internal/modules/md"
	"lazts/internal/services/templ"
	"lazts/pkg/logger"
)

func main() {
	log := logger.NewZerolog(logger.Config())
	server := http.New(http.Config())
	server.Use(middlewares.Logger(log), middlewares.Compressor())
	markdown := md.New(md.Config())

	ht := templ.New(log, markdown)

	home.New(server, ht)
	vacations.New(server, ht)
	books.New(server, ht)
	notes.New(server, ht)
	system.New(server, ht)

	log.I("starting server")
	err := server.Serve()
	if err != nil {
		log.Err(err).I("failed to start server")
	}
}
