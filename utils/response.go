package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

func ErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func SuccessListResponse(c *gin.Context, data interface{}, pagination Pagination) {
	c.JSON(http.StatusOK, gin.H{
		"data":       data,
		"pagination": pagination,
	})
}
