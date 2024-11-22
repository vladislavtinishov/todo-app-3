package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"todo_app_3/middlewares"
	todocontrollers "todo_app_3/modules/todo/controllers"
	todoservices "todo_app_3/modules/todo/services"
	usercontrollers "todo_app_3/modules/users/controllers"
	userservices "todo_app_3/modules/users/services"
)

func NewHandler(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userService := userservices.NewUserService(db)
	todoService := todoservices.NewTodoService(db)

	authController := usercontrollers.NewAuthController(userService)
	todoController := todocontrollers.NewTodoController(todoService)

	r.POST("/auth/sign-up", authController.SignUp)
	r.POST("/auth/sign-in", authController.SignIn)

	authenticated := r.Group("/", middlewares.Authenticated)
	{
		authenticated.GET("/todos", todoController.Index)
		authenticated.GET("/todos/:id", todoController.Show)
		authenticated.POST("/todos", todoController.Create)
		authenticated.PUT("/todos/:id", todoController.Update)
		authenticated.DELETE("/todos/:id", todoController.Delete)
	}

	return r
}
