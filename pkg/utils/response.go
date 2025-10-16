package utils

import "github.com/gin-gonic/gin"

// ErrorResponse merepresentasikan struktur response error
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse merepresentasikan struktur response sukses dengan message
type SuccessResponse struct {
	Message string `json:"message"`
}

// RespondError mengirimkan response error dengan format standar
func RespondError(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, ErrorResponse{Error: message})
}

// RespondSuccess mengirimkan response sukses dengan data
func RespondSuccess(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// RespondMessage mengirimkan response sukses dengan message
func RespondMessage(c *gin.Context, statusCode int, message string) {
	c.JSON(statusCode, SuccessResponse{Message: message})
}
