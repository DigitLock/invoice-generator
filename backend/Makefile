include .env
export

DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

.PHONY: run build migrate-up migrate-down sqlc-generate test

run:
	go run ./cmd/server/

build:
	go build -o bin/server ./cmd/server/

migrate-up:
	migrate -path database/migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path database/migrations -database "$(DB_URL)" down 1

sqlc-generate:
	sqlc generate

test:
	go test ./... -v
