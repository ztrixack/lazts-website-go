package main

import (
	"lazts/internal/handlers/system"
	"lazts/internal/handlers/web"
	"lazts/internal/modules/http"
	"lazts/internal/modules/http/middlewares"
	"lazts/internal/services/markdown"
	"lazts/internal/services/templ"
	"lazts/pkg/logger"
)

func main() {
	log := logger.NewZerolog(logger.Config())
	server := http.New(http.Config())
	server.Use(middlewares.Logger(log), middlewares.Compressor())

	ht := templ.New(log)
	md := markdown.New(log)

	system.New(server, ht, md)
	web.New(server, ht)

	log.I("starting server")
	err := server.Serve()
	if err != nil {
		log.Err(err).I("failed to start server")
	}
}
