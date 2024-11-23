package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"todo_app_3/config"
	"todo_app_3/utils"
)

const (
	authHeader = "Authorization"
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

	c.Set(config.USER_CTX, userId)
}
