package drivers

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	User string
	Pass string
	Host string
	Port string
	DB   string
}

func (m *Mysql) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", m.User, m.Pass, m.Host, m.Port, m.DB)
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}
