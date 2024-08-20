run:
	go run ./cmd/server/main.go

dbrun:
	cd ./docs && docker-compose up -d