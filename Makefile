#!make

ifeq ($(OS),Windows_NT)
  uname_S := Windows
else
  uname_S := $(shell uname -s)
endif

PHONY: setup-pre-commit
setup-pre-commit:
	@echo "Setting up pre-commit..."
	./scripts/setup-pre-commit.sh

PHONY: setup-air
setup-air:
	@echo "Setting up air..."
	go install github.com/cosmtrek/air@latest

PHONY: setup-env
setup-env:
	@echo "Setting up env..."
	cp .env.template .env

PHONY: test
test:
	@echo "Running tests..."
	go test -v ./...

PHONY: test-it
test-it:
	@echo "Running integration tests..."
	go test -v -run "Test.*IT" -tags=integration ./...

PHONY: coverage
coverage:
	go test -cover -coverprofile=report.out -v ./...
	go tool cover -html=report.out -o coverage.html

PHONY: swagger
swagger:
	swag fmt
	swag init

PHONY: dev
dev:
	@echo "Running the development server..."
ifeq ($(uname_S), Windows)
	air -c .air.win.toml
else
	air -c .air.toml
endif

.PHONY: run
run:
	@echo "Running the server..."
	go run main.go

.PHONY: css
css:
	npx tailwindcss -i ./static/css/tailwind.css -o ./static/css/app.css --watch

.PHONY: css-minify
css-minify:
	npx tailwindcss -i ./static/css/tailwind.css -o ./static/css/app.css --minify
