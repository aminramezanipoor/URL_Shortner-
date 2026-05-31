run:
	go run cmd/server/main.go

docker:
	docker compose up -d

test:
	go test ./...