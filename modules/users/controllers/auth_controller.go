package usercontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"todo_app_3/common/utils"
	usermodels "todo_app_3/modules/users/models"
	userservices "todo_app_3/modules/users/services"
)

type AuthController struct {
	service *userservices.UserService
}

const InvalidLoginError = "Invalid login or password"

func NewAuthController(service *userservices.UserService) *AuthController {
	return &AuthController{service: service}
}

func (uc *AuthController) SignUp(c *gin.Context) {
	var body usermodels.CreateUser

	err := c.Bind(&body)
	if err != nil {
		return
	}

	passwordHash := utils.GenerateHash(body.Password)

	user, err := uc.service.Create(body, passwordHash)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	accessToken, err := utils.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": accessToken,
	})
}

func (uc *AuthController) SignIn(c *gin.Context) {
	var body usermodels.SignIn

	err := c.Bind(&body)
	if err != nil {
		return
	}

	user, err := uc.service.FindByLogin(body.Login)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": InvalidLoginError,
		})
		return
	}

	passwordHash := utils.GenerateHash(body.Password)

	if user.PasswordHash != passwordHash {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": InvalidLoginError,
		})
		return
	}

	accessToken, err := utils.GenerateToken(user.ID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": InvalidLoginError,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user":  user,
		"token": accessToken,
	})
}
