package main

import (
	"lazts/internal/handlers/books"
	"lazts/internal/handlers/home"
	"lazts/internal/handlers/notes"
	"lazts/internal/handlers/system"
	"lazts/internal/handlers/vacations"
	"lazts/internal/modules/http"
	"lazts/internal/modules/http/middlewares"
	"lazts/internal/modules/log"
	"lazts/internal/modules/markdown"
	"lazts/internal/services/book"
	"lazts/internal/services/note"
	"lazts/internal/services/page"
	"lazts/internal/services/vacation"
	"lazts/internal/services/watermark"

	"lazts/internal/modules/imaging"
)

func main() {
	log := log.New(log.Config())
	server := http.New(http.Config())
	server.Use(middlewares.Logger(log), middlewares.Compressor())
	markdown := markdown.New(markdown.Config())
	img := imaging.New(imaging.Config())

	pager := page.New(log, markdown)
	booker := book.New(log)
	vacationer := vacation.New(log, markdown)
	noter := note.New(log, markdown)
	watermarker := watermark.New(watermark.Config(), log, img)

	home.New(server, pager, booker, vacationer, noter)
	books.New(server, pager, booker)
	vacations.New(server, pager, vacationer)
	notes.New(log, server, pager, noter)
	system.New(server, pager, watermarker)

	log.Fields("port", server.Config().Port).I("starting server")
	err := server.Serve()
	if err != nil {
		log.Err(err).I("failed to start server")
	}
}
