# Online TicTacToe Backend

## Description

This is a backend for the Online TicTacToe project - you can play TicTacToe with your friends online by inviting them to your game. Can set up a game with board of any size and any winning condition in a row! The game is turn-based and you can play with multiple friends at the same time. The game is real-time and you will get a notification when it's your turn to play. Uses DFS to check for winning conditions.

It's built with Go, Gin-Gonic, MongoDB driver and Pusher. Uses MongoDB to store game objects, users, etc and Pusher for push notifications. There is also a Cron job that runs at 05:00 AM (UTC), every 2 days to reset games database. The architecture is a serverless approach using Vercel's serverless functions.

## How to run server

1. Clone the repository
2. Run `go mod download` to download the dependencies
3. Run `go run cmd/server/main.go` or `make server` to start the server
4. The server will be running on `localhost:8080`

## How to generate swagger documentation

1. Run `make swagger`
2. The doc should be on `localhost:8080/swagger`

## Swagger documentation

View the API documentation [here](https://api-online-tictactoe.vercel.app/swagger/index.html)

## Preview

![Preview](https://github.com/gowth6m/online-tictactoe-frontend/blob/main/public/showcase1.png)

![Preview](https://github.com/gowth6m/online-tictactoe-frontend/blob/main/public/showcase2.png)

![Preview](https://github.com/gowth6m/online-tictactoe-frontend/blob/main/public/showcase3.png)