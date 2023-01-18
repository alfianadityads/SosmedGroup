package comment

import "github.com/labstack/echo/v4"

type Core struct {
	ID        uint
	Comment   string
	OwnerName string
	ContentID uint
	CreateAt  string
}

type CommentHandler interface {
	NewComment() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type CommentService interface {
	NewComment(token interface{}, NewComment Core) (Core, error)
	Delete(token interface{}) error
}

type CommentData interface {
	NewComment(userID int, newComment Core) (Core, error)
	Delete(commentID int) error
}
