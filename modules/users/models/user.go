package usermodels

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Name         string         `json:"name" gorm:"type:varchar(100)"`
	Login        string         `json:"login" gorm:"type:varchar(100);unique"`
	PasswordHash string         `json:"-" gorm:"type:varchar(100)"`
}
