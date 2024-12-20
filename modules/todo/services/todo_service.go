package todoservices

import (
	"gorm.io/gorm"
	"strings"
	"todo_app_3/common/utils"
	todomodels "todo_app_3/modules/todo/models"
)

type TodoService struct {
	db *gorm.DB
}

func NewTodoService(db *gorm.DB) *TodoService {
	return &TodoService{db: db}
}

func (s *TodoService) Create(body todomodels.Todo) (todomodels.Todo, error) {
	result := s.db.Preload("User").Preload("Status").Create(&body)

	return body, result.Error
}

func (s *TodoService) Find(id, userId uint) (todomodels.Todo, error) {
	var todo todomodels.Todo

	result := s.db.Preload("Todo").Preload("Status").Preload("User").First(&todo, "id = ? and user_id = ?", id, userId)

	if result.Error != nil {
		return todo, result.Error
	}

	return todo, nil
}

func (s *TodoService) GetAll(userId uint, filter todomodels.Search, pagination *utils.Pagination) ([]todomodels.Todo, error) {
	var todos []todomodels.Todo

	query := s.db.Model(&todomodels.Todo{})
	query = query.Preload("User")

	query = query.Where("parent_id is null AND user_id = ?", userId)

	if title := filter.Title; title != "" {
		query = query.Where("lower(title) like ?", "%"+strings.ToLower(title)+"%")
	}

	if description := filter.Description; description != "" {
		query = query.Where("lower(description) like ?", "%"+strings.ToLower(description)+"%")
	}

	if statusId := filter.StatusID; statusId > 0 {
		query = query.Where("status_id = ?", statusId)
	}

	query, err := pagination.Apply(query)

	if err != nil {
		return []todomodels.Todo{}, err
	}

	result := query.Find(&todos)

	return todos, result.Error
}

func (s *TodoService) Update(id, userId uint, body todomodels.Todo) (todomodels.Todo, error) {
	if _, err := s.Find(id, userId); err != nil {
		return todomodels.Todo{}, err
	}

	body.ID = id
	if err := s.db.Model(&body).Updates(body).Error; err != nil {
		return todomodels.Todo{}, err
	}

	return s.Find(id, userId)
}

func (s *TodoService) Delete(id, userId uint) error {
	var todo todomodels.Todo
	var err error

	if todo, err = s.Find(id, userId); err != nil {
		return err
	}

	if err := s.db.Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}
