package todomodels

import (
	"gorm.io/gorm"
	"time"
	statusmodels "todo_app_3/modules/todostatus/models"
	usermodels "todo_app_3/modules/users/models"
)

type Todo struct {
	ID          uint           `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Title       string         `json:"title" gorm:"type:varchar(255)"`
	Description *string        `json:"description"`
	ParentID    *uint          `json:"parent_id"`
	StatusID    uint           `json:"status_id"`
	UserID      uint           `json:"-"`

	User   usermodels.User         `gorm:"foreignKey:ID;references:UserID" json:"creator"`
	Todo   []Todo                  `gorm:"foreignKey:ParentID;references:ID" json:"todos"`
	Status statusmodels.TodoStatus `gorm:"foreignKey:ID;references:StatusID" json:"status"`
}
