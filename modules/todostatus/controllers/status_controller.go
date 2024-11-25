package todostatuscontrollers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	todostatusmodels "todo_app_3/modules/todostatus/models"
	todostatusservices "todo_app_3/modules/todostatus/services"
	"todo_app_3/utils"
)

type TodoStatusController struct {
	service *todostatusservices.TodoStatusService
}

func NewTodoStatusController(service *todostatusservices.TodoStatusService) *TodoStatusController {
	return &TodoStatusController{service}
}

func (td *TodoStatusController) GetAll(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "UserId is undefined")
		return
	}

	data, err := td.service.GetAll(userId)

	utils.SuccessResponse(c, data)
}

func (td *TodoStatusController) Show(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := utils.GetIdFromParam(c)
	if err != nil {
		utils.ErrorEntityIdResponse(c)
		return
	}

	data, err := td.service.Find(userId, id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusNotFound, "status is not found")
		return
	}

	utils.SuccessResponse(c, data)
}

func (td *TodoStatusController) Create(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "UserId is undefined")
		return
	}

	var status todostatusmodels.TodoStatus

	if err := c.Bind(&status); err != nil {
		return
	}

	data, err := td.service.Create(userId, status)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, data)
}

func (td *TodoStatusController) Update(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := utils.GetIdFromParam(c)
	if err != nil {
		utils.ErrorEntityIdResponse(c)
		return
	}

	var body todostatusmodels.TodoStatus
	if err := c.Bind(&body); err != nil {
		return
	}

	data, err := td.service.Update(userId, id, body)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		utils.ErrorResponse(c, http.StatusNotFound, "status is not found")
		return
	}

	utils.SuccessResponse(c, data)
}

func (td *TodoStatusController) Delete(c *gin.Context) {
	userId, err := utils.GetUserId(c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	id, err := utils.GetIdFromParam(c)
	if err != nil {
		utils.ErrorEntityIdResponse(c)
		return
	}

	err = td.service.Delete(userId, id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, map[string]bool{
		"success": true,
	})
}
