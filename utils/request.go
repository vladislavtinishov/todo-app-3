package utils

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"todo_app_3/config"
)

func GetUserId(c *gin.Context) (uint, error) {
	id, ok := c.Get(config.USER_CTX)

	if !ok {
		ErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(uint)

	if !ok {
		ErrorResponse(c, http.StatusInternalServerError, "user is of invalid type")
		return 0, errors.New("user is of invalid type")
	}

	return idInt, nil
}
