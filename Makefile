#!make

ifeq ($(OS),Windows_NT)
  uname_S := Windows
else
  uname_S := $(shell uname -s)
endif

fmt:
	go fmt ./...

test:
	go test -v ./...

test-int:
	docker compose -f docker-compose.test.yaml up --build --abort-on-container-exit --exit-code-from lazts-website; \
	docker compose -f docker-compose.test.yaml down

cover:
	go test -cover -coverprofile=report.out -v ./...
	go tool cover -html=report.out -o coverage.html

swag:
	swag fmt
	swag init

init:
	go install github.com/cosmtrek/air@latest

dev:
ifeq ($(uname_S), Windows)
	air -c .air.win.toml
else
	air -c .air.toml
endif

run:
	make css-minify
	go run cmd/app/main.go

docker-dev:
	docker compose -f docker-compose.dev.yaml up --build --renew-anon-volumes --abort-on-container-exit --exit-code-from lazts-website; \
	docker compose -f docker-compose.dev.yaml down

docker-release:
	docker compose -f docker-compose.yaml up --build -d

css-watch:
	npx tailwindcss -i ./static/css/tailwind.css -o ./static/css/app.css --watch

css-minify:
	npx tailwindcss -i ./static/css/tailwind.css -o ./static/css/app.css --minify
