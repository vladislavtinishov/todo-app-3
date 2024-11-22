package middlewares

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo_app_3/utils"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func Authenticated(c *gin.Context) {
	header := c.GetHeader(authHeader)
	if header == "" {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) == 1 {
		utils.ErrorResponse(c, http.StatusUnauthorized, "Empty access token")
		return
	}

	claims, err := utils.ParseJWT(headerParts[1])

	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	userId := claims.UserId

	c.Set(userCtx, userId)
}

func GetUserId(c *gin.Context) (uint, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(uint)

	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "user is of invalid type")
		return 0, errors.New("user is of invalid type")
	}

	return idInt, nil
}
