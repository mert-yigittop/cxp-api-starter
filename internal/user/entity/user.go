package entity

import (
	"github.com/mert-yigittop/cxp-api-starter/internal/todo/entity"
	"github.com/mert-yigittop/cxp-api-starter/internal/user/dto"
	"github.com/mert-yigittop/cxp-api-starter/pkg/constant"
)

type User struct {
	constant.DefaultModel
	Username string        `gorm:"unique;not null"`
	Email    string        `gorm:"index;unique;not null"`
	Password string        `gorm:"not null"`
	IsActive bool          `gorm:"default:true"`
	Todos    []entity.Todo `gorm:"foreignKey:UserID"`
}

func New(data dto.RegisterRequest) User {
	return User{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
	}
}
