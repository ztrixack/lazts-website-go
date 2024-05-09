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
	"lazts/internal/services/book"
	"lazts/internal/services/note"
	"lazts/internal/services/page"
	"lazts/internal/services/vacation"
	"lazts/pkg/logger"
)

func main() {
	log := logger.NewZerolog(logger.Config())
	server := http.New(http.Config())
	server.Use(middlewares.Logger(log), middlewares.Compressor())
	markdown := md.New(md.Config())

	pager := page.New(log, markdown)
	booker := book.New(log)
	vacationer := vacation.New(log, markdown)
	noter := note.New(log, markdown)

	home.New(server, pager, booker, vacationer, noter)
	books.New(server, pager, booker)
	vacations.New(server, pager, vacationer)
	notes.New(server, pager, noter)
	system.New(server, pager)

	log.Fields("port", server.Config().Port).I("starting server")
	err := server.Serve()
	if err != nil {
		log.Err(err).I("failed to start server")
	}
}
