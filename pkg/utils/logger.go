package utils

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

// LogRequestBody is a middleware function to log the request body before passing it to the next middleware/handler
func LogRequestBody(c *gin.Context) {
	var bodyBytes []byte
	if c.Request.Body != nil {
		bodyBytes, _ = io.ReadAll(c.Request.Body)
	}
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	var bodyMap map[string]interface{}
	if err := json.Unmarshal(bodyBytes, &bodyMap); err == nil {
		bodyString, _ := json.Marshal(bodyMap)
		log.Println("Request body --> ", string(bodyString))
	} else {
		log.Println("Error reading request body for logging")
	}

	c.Next()
}
