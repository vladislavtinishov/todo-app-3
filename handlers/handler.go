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
		todos := authenticated.Group("/todos")
		{
			statuses := todos.Group("/statuses")
			{
				statuses.GET("/", todoStatusController.GetAll)
				statuses.GET("/:id", todoStatusController.Show)
				statuses.POST("", todoStatusController.Create)
				statuses.PUT("/:id", todoStatusController.Update)
				statuses.DELETE("/:id", todoStatusController.Delete)
			}

			todos.GET("/", todoController.Index)
			todos.GET("/:id", todoController.Show)
			todos.POST("/", todoController.Create)
			todos.PUT("/:id", todoController.Update)
			todos.DELETE("/:id", todoController.Delete)
		}
	}

	return r
}
