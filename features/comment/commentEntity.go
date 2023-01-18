package comment

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	Comment   string
	ContentID uint
	CreateAt  string
	User      UserCore
}

type UserCore struct {
	ID       uint
	UserName string
	Name     string
}

type CommentHandler interface {
	NewComment() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetCom() echo.HandlerFunc
}

type CommentService interface {
	NewComment(token interface{}, contentID uint, NewComment string) (Core, error)
	Delete(token interface{}, commentID uint) error
	GetCom() ([]Core, error)
}

type CommentData interface {
	NewComment(userID int, contentID uint, newComment string) (Core, error)
	Delete(userID uint, commentID uint) error
	GetCom() ([]Core, error)
}
