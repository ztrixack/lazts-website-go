package http

import (
	"os"
	"strconv"
)

const DEFAULT_PORT = "8080"

type config struct {
	Host string
	Port string
}

func Config() *config {
	host, ok := os.LookupEnv("HTTP_HOST")
	if !ok {
		host = "localhost"
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" || !isValidPort(port) {
		port = DEFAULT_PORT
	}

	return &config{
		Host: host,
		Port: port,
	}
}

func isValidPort(port string) bool {
	_, err := strconv.Atoi(port)
	return err == nil && port != ""
}
