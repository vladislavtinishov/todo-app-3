package todostatusmodels

import (
	"gorm.io/gorm"
	"time"
	usermodels "todo_app_3/modules/users/models"
)

type TodoStatus struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	UserID    uint           `json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title     string         `json:"title" gorm:"type:varchar(100)"`

	User usermodels.User `gorm:"foreignKey:ID;references:UserID" json:"-"`
}
