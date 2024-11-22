package userservices

import (
	"gorm.io/gorm"
	usermodels "todo_app_3/modules/users/models"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) Create(body usermodels.CreateUser, passwordHash string) (usermodels.User, error) {
	user := usermodels.User{
		Name:         body.Name,
		Login:        body.Login,
		PasswordHash: passwordHash,
	}

	result := s.db.Create(&user)

	return user, result.Error
}

func (s *UserService) FindByLogin(login string) (usermodels.User, error) {
	var user usermodels.User
	result := s.db.First(&user, "login=?", login)

	return user, result.Error
}
