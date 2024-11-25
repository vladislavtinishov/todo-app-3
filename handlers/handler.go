package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"todo_app_3/middlewares"
	todocontrollers "todo_app_3/modules/todo/controllers"
	todoservices "todo_app_3/modules/todo/services"
	todostatuscontrollers "todo_app_3/modules/todostatus/controllers"
	todostatusservices "todo_app_3/modules/todostatus/services"
	usercontrollers "todo_app_3/modules/users/controllers"
	userservices "todo_app_3/modules/users/services"
)

func NewHandler(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	userService := userservices.NewUserService(db)
	todoService := todoservices.NewTodoService(db)
	todoStatusService := todostatusservices.NewTodoStatusService(db)

	authController := usercontrollers.NewAuthController(userService)
	todoController := todocontrollers.NewTodoController(todoService)
	todoStatusController := todostatuscontrollers.NewTodoStatusController(todoStatusService)

	r.POST("/auth/sign-up", authController.SignUp)
	r.POST("/auth/sign-in", authController.SignIn)

	authenticated := r.Group("/", middlewares.Authenticated)
	{
		authenticated.GET("/todos", todoController.Index)
		authenticated.GET("/todos/:id", todoController.Show)
		authenticated.POST("/todos", todoController.Create)
		authenticated.PUT("/todos/:id", todoController.Update)
		authenticated.DELETE("/todos/:id", todoController.Delete)

		authenticated.GET("/status", todoStatusController.GetAll)
		authenticated.GET("/status/:id", todoStatusController.Show)
		authenticated.POST("/status", todoStatusController.Create)
		authenticated.PUT("/status/:id", todoStatusController.Update)
		authenticated.DELETE("/status/:id", todoStatusController.Delete)
	}

	return r
}
