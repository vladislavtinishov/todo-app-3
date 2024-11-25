package todostatusservices

import (
	"gorm.io/gorm"
	todostatusmodels "todo_app_3/modules/todostatus/models"
)

type TodoStatusService struct {
	db *gorm.DB
}

func NewTodoStatusService(db *gorm.DB) *TodoStatusService {
	return &TodoStatusService{db: db}
}

func (s *TodoStatusService) GetAll(userId uint) ([]todostatusmodels.TodoStatus, error) {
	var data []todostatusmodels.TodoStatus

	query := s.db.Model(&todostatusmodels.TodoStatus{}).Where("user_id = ?", userId)

	result := query.Find(&data)

	return data, result.Error
}

func (s *TodoStatusService) Create(userId uint, status todostatusmodels.TodoStatus) (todostatusmodels.TodoStatus, error) {
	status.UserID = userId

	result := s.db.Create(&status)

	return status, result.Error
}

func (s *TodoStatusService) Update(userId, id uint, body todostatusmodels.TodoStatus) (todostatusmodels.TodoStatus, error) {
	if _, err := s.Find(userId, id); err != nil {
		return todostatusmodels.TodoStatus{}, err
	}

	body.ID = id
	if err := s.db.Model(&body).Updates(body).Error; err != nil {
		return todostatusmodels.TodoStatus{}, err
	}

	return s.Find(userId, id)
}

func (s *TodoStatusService) Find(userId, id uint) (todostatusmodels.TodoStatus, error) {
	var status todostatusmodels.TodoStatus

	result := s.db.First(&status, "id = ? and user_id = ?", id, userId)

	if result.Error != nil {
		return status, result.Error
	}

	return status, nil
}

func (s *TodoStatusService) Delete(userId, id uint) error {
	var todo todostatusmodels.TodoStatus
	var err error

	if todo, err = s.Find(userId, id); err != nil {
		return err
	}

	if err := s.db.Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}
