package drivers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	User string
	Pass string
	Host string
	Port string
	DB   string
}

func (m *Postgres) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", m.Host, m.Port, m.User, m.Pass, m.DB)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
