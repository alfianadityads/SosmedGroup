package helper

import (
	"sosmedapps/features/comment/data"

	"gorm.io/gorm"
)

type UserStruct struct {
	gorm.Model
	Name     string
	Email    string
	Bio      string
	Image    string
	UserName string
	Password string
	Content  []data.Content `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
