#!make

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

run:
	go run cmd/app/main.go

dev:
	docker compose -f docker-compose.dev.yaml up --build --renew-anon-volumes --abort-on-container-exit --exit-code-from lazts-website; \
	docker compose -f docker-compose.dev.yaml down

release:
	docker compose -f docker-compose.yaml down
	docker compose -f docker-compose.yaml up --build -d