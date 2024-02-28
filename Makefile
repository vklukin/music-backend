tidy:
	go mod tidy

start:
	ENV_FILE=./cmd/web-app/.env go run ./cmd/web-app/main.go