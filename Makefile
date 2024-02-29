tidy:
	go mod tidy

mod:
	go mod download
	
start:
	go run ./cmd/web-app/main.go
