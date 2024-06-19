package api

import (
	"github.com/gin-gonic/gin"
)

// --------------- Response structures ---------------

type ResponseData struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data,omitempty"`
	Error   *[]FieldError `json:"error,omitempty"`
}

type FieldError struct {
	Field   *string `json:"field"`
	Message string  `json:"message"`
}

// --------------- Response functions ---------------

func ApiResponse(c *gin.Context, statusCode int, message string, data interface{}, errors *[]FieldError) {
	response := ResponseData{
		Status:  statusCode,
		Message: message,
		Data:    data,
		Error:   errors,
	}

	c.JSON(statusCode, response)
}

func Error(c *gin.Context, statusCode int, message string, errors *[]FieldError) {
	ApiResponse(c, statusCode, message, nil, errors)
	c.Abort()
}

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	ApiResponse(c, statusCode, message, data, nil)
}
