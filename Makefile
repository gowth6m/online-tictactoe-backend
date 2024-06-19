server:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

swagger:
	swag init -g cmd/server/main.go