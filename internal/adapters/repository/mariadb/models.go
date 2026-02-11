package mariadb

import (
	"Auth/internal/domain/user"
	"time"
)

type MariaModel struct {
	Id    uint      `gorm:"primaryKey:autoIncrement"`
	Email string    `gorm:"uniqueIndex:not null;size(255)"`
	Password string `gorm:"not null"`
	Role  string    `gorm:"default:'user';size(50)"`
	Created_at time.Time `gorm:"column:created_at;autoCreateTime"`
	Update_at time.Time `gorm:"column:updated_at;autoUpdateTime"`
}


func(m *MariaModel) TableName() string {
	return "users"
}

func ToDomain(ma *MariaModel) *user.User{
    return user.Restore(
	int64(ma.Id),
	ma.Email,
	ma.Role,
	ma.Password)
}

func  FromDomain(u *user.User)*MariaModel{
	return &MariaModel{
		Id: uint(u.GetId()),
		Email: u.Email(),
		Role: string(u.Role()),
		Password: u.Password(),

	}
}
