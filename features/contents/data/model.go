package data

import (
	"sosmedapps/features/comment/data"

	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Content      string
	ContentImage string
	NumbComment  string
	UserID       uint
	Comment      []data.Comment
}

// type Content struct {
// 	gorm.Model
// 	Content      string
// 	ContentImage string
// 	Owner        string
// 	UserID       uint
// 	Create_at    string
// 	NumbComment  string
// 	User         User
// }

type User struct {
	gorm.Model
	Name        string
	UserName    string
	Email       string
	Image       string
	Password    string
	Bio         string
	DateOfBirth string
	Content     []Content
}

// func (data *Content) ContentToCore() contents.CoreContent {
// 	return contents.CoreContent{
// 		ID:           data.ID,
// 		Content:      data.Content,
// 		ContentImage: data.ContentImage,
// 		Create_at:    data.Create_at,
// 		Owner:        data.Owner,
// 		UserID:       data.User.ID,
// 		NumbComment:  data.NumbComment,
// 		Users: contents.CoreUser{
// 			ID:       data.User.ID,
// 			Name:     data.User.Name,
// 			UserName: data.User.UserName,
// 		},
// 	}
// }

// func CoreToData(data contents.CoreContent) Content {
// 	return Content{
// 		Model:        gorm.Model{ID: data.ID},
// 		Content:      data.Content,
// 		ContentImage: data.ContentImage,
// 		UserID:       data.UserID,
// 	}
// }

// func ToCoreContent(data []Content) []contents.CoreContent {
// 	var tmp []contents.CoreContent
// 	for _, v := range data {
// 		tmp = append(tmp, v.ContentToCore())
// 	}
// 	return tmp
// }
