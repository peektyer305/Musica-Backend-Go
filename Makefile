run:
	go run ./cmd/server/main.go

dbrun:
	cd ./docs && docker compose up -d

dbdown:
	cd ./docs && docker compose down

dbstart:
	cd ./docs && docker compose start

dbstop:
	cd ./docs && docker compose stop

migrate-up:
	goose -dir ./internal/infrastructure/postgre/migrations up

migrate-status:
	goose -dir ./internal/infrastructure/postgre/migrations status

migrate-down:
	goose -dir ./internal/infrastructure/postgre/migrations down

migrate-reset:
	goose -dir ./internal/infrastructure/postgre/migrations reset
go-install:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/cespare/reflex@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

wire-gen:
	wire ./di/container.go