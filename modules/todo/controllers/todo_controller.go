package todocontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	todomodels "todo_app_3/modules/todo/models"
	todoservices "todo_app_3/modules/todo/services"
	"todo_app_3/utils"
)

type TodoController struct {
	service *todoservices.TodoService
}

func NewTodoController(service *todoservices.TodoService) *TodoController {
	return &TodoController{service}
}

func (td *TodoController) Index(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "UserId is undefined")
		return
	}

	var filter todomodels.Search

	if err := c.Bind(&filter); err != nil {
		return
	}

	var pagination utils.Pagination

	if err := c.Bind(&pagination); err != nil {
		return
	}

	todos, err := td.service.GetAll(userId, filter, &pagination)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessListResponse(c, todos, pagination)
}

func (td *TodoController) Create(c *gin.Context) {
	var body todomodels.Todo

	err := c.Bind(&body)
	if err != nil {
		return
	}

	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "UserId is undefined")
		return
	}

	body.UserID = userId

	todo, err := td.service.Create(body)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}

func (td *TodoController) Show(c *gin.Context) {
	id, err := utils.GetIdFromParam(c)
	if err != nil {
		utils.ErrorEntityIdResponse(c)
		return
	}

	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "UserId is undefined")
		return
	}

	todo, err := td.service.Find(id, userId)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "todo is not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}

func (td *TodoController) Update(c *gin.Context) {
	id, err := utils.GetIdFromParam(c)
	if err != nil {
		utils.ErrorEntityIdResponse(c)
		return
	}

	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "UserId is undefined")
		return
	}

	var body todomodels.Todo
	err = c.Bind(&body)
	if err != nil {
		return
	}

	todo, err := td.service.Update(id, userId, body)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusUnauthorized, "todo is not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"todo": todo,
	})
}

func (td *TodoController) Delete(c *gin.Context) {
	id, err := utils.GetIdFromParam(c)
	if err != nil {
		utils.ErrorEntityIdResponse(c)
		return
	}

	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "UserId is undefined")
		return
	}

	err = td.service.Delete(id, userId)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "todo is not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": true,
	})
}
