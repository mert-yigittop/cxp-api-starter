package entity

import (
	"github.com/mert-yigittop/cxp-api-starter/pkg/constant"
)

type Todo struct {
	constant.DefaultModel
	Content   string `json:"content",gorm:"not null"`
	Completed bool   `json:"completed",gorm:"default:false"`
	UserID    uint   `json:"userId",gorm:"not null"`
}

func New(content string, userID uint) Todo {
	return Todo{
		Content: content,
		UserID:  userID,
	}
}
