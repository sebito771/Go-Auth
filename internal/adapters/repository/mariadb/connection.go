package mariadb

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func NewMariaDBRepo(user, password, host, dbname string) (*MariaQueries, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", 
		user, password, host, dbname)

	
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return &MariaQueries{db: db}, nil
}