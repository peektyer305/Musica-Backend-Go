
POSTGRES_USER     := akira
POSTGRES_PASSWORD := newpassword
POSTGRES_HOST     := localhost# ローカル実行時は localhost、コンテナ内実行時は "postgres"
POSTGRES_PORT     := 5432
POSTGRES_DB       := app

# ホスト側 goose 実行用の接続文字列
DATABASE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOST):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

MIGRATION_DIR := internal/infrastructure/postgre/migrations

.PHONY: migrate-up migrate-down migrate-status
 
.PHONY: run

run:
	@echo "→ running server (local mode)"
	@DB_HOST=localhost go run ./cmd/server/main.go

dbrun:
	cd ./docs && docker compose up -d

dbdown:
	cd ./docs && docker compose down

dbstart:
	cd ./docs && docker compose start

dbstop:
	cd ./docs && docker compose stop

migrate-up:
	@echo "→ running migrations up"
	@GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING="$(DATABASE_URL)" \
	goose -dir $(MIGRATION_DIR) up

# １つだけ上げる
migrate-up-by-one:
	@echo "→ running migrations up-by-one"
	@GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING="$(DATABASE_URL)" \
	goose -dir $(MIGRATION_DIR) up-by-one

# 最新のものだけやり直す
migrate-redo:
	@echo "→ re-running latest migration"
	@GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING="$(DATABASE_URL)" \
	goose -dir $(MIGRATION_DIR) redo

# ステータス確認
migrate-status:
	@echo "→ migration status"
	@GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING="$(DATABASE_URL)" \
	goose -dir $(MIGRATION_DIR) status

# 全ダウン
migrate-down:
	@echo "→ migrating all the way down"
	@GOOSE_DRIVER=postgres \
	GOOSE_DBSTRING="$(DATABASE_URL)" \
	goose -dir $(MIGRATION_DIR) reset
go-install:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/cespare/reflex@latest
	go install github.com/go-delve/delve/cmd/dlv@latest
	go install honnef.co/go/tools/cmd/staticcheck@latest
	go install github.com/pressly/goose/v3/cmd/goose@latest

wire-gen:
	wire ./di/container.go