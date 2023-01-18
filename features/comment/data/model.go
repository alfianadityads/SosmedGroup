package data

import (
	"sosmedapps/features/comment"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Comment   string
	UserID    uint
	ContentID uint
}

type Content struct {
	gorm.Model
	Content      string
	ContentImage string
	NumbComment  string
	UserID       User
	Comment      []Comment
}

type User struct {
	gorm.Model
	Name     string
	Email    string
	Bio      string
	Image    string
	UserName string
	Password string
	Content  []Content
	Comment  []Comment
}

func DataToCore(data Comment) comment.Core {
	return comment.Core{
		ID:        data.ID,
		Comment:   data.Comment,
		ContentID: data.ContentID,
	}
}

func CoreToData(core comment.Core) Comment {
	return Comment{
		Model:     gorm.Model{ID: core.ID},
		Comment:   core.Comment,
		ContentID: core.ContentID,
	}
}
