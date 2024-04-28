package http

import (
	"os"
	"strconv"
)

type config struct {
	Port string
}

func Config() *config {
	port := os.Getenv("PORT")
	if port == "" || !isValidPort(port) {
		port = "8080"
	}

	return &config{
		Port: port,
	}
}

func isValidPort(port string) bool {
	_, err := strconv.Atoi(port)
	return err == nil && port != ""
}
