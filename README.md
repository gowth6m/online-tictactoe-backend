# Online TicTacToe Backend

## Description

This is a backend for the Online TicTacToe project - you can play TicTacToe with your friends online by inviting them to your game. It's built with Go Gin-gonic, MongoDB driver and Pusher. Uses MongoDB to store game objects, users, etc and Pusher for push notifications. There is also a Cron job that runs every 48 hours to reset games database. The architecture is a serverless approach using Vercel's serverless functions.

## How to run server

1. Clone the repository
2. Run `go mod download` to download the dependencies
3. Run `go run cmd/server/main.go` or `make server` to start the server
4. The server will be running on `localhost:8080`

## How to generate swagger documentation

1. Run `make swagger`
2. The doc should be on `localhost:8080/swagger`

<!-- ## Endpoints -->

<!-- View the API documentation [here](https://todo) -->