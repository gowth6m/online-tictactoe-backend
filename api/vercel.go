package handler

import (
	"net/http"
	"online-tictactoe/pkg/serverless"
)

// Entry point for the Vercel serverless function.
func Handler(w http.ResponseWriter, r *http.Request) {
	router, cleanup := serverless.Initialize()
	defer cleanup()
	router.ServeHTTP(w, r)
}
