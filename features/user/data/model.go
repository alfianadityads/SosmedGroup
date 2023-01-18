package data

import (
	cmData "sosmedapps/features/comment/data"
	cData "sosmedapps/features/contents/data"
	"sosmedapps/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Bio      string
	Image    string
	UserName string
	Password string
	Content  []cData.Content
	Comment  []cmData.Comment
}

func DataToCore(data User) user.Core {
	return user.Core{
		ID:       data.ID,
		Name:     data.Name,
		Email:    data.Email,
		Bio:      data.Bio,
		Image:    data.Image,
		UserName: data.UserName,
		Password: data.Password,
		Content:  data.Content,
		Comment:  data.Comment,
	}
}

func CoreToData(core user.Core) User {
	return User{
		Model:    gorm.Model{ID: core.ID},
		Name:     core.Name,
		Email:    core.Email,
		Bio:      core.Bio,
		Image:    core.Image,
		UserName: core.UserName,
		Password: core.Password,
		Content:  core.Content,
		Comment:  core.Comment,
	}
}
