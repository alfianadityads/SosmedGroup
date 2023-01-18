package contents

import (
	"mime/multipart"

	"github.com/labstack/echo/v4"
)

type CoreContent struct {
	ID           uint
	Content      string `validate:"required" json:"content" from:"content"`
	ContentImage string `json:"content_image" from:"content_image"`
	CreateAt     string
	OwnerName    string
	NumbComment  uint
	Users        CoreUser
}

type CoreUser struct {
	ID       uint
	UserName string
	Name     string
}

type ContentHandler interface {
	AddContent() echo.HandlerFunc
	// UpdateContent() echo.HandlerFunc
	// DeleteContent() echo.HandlerFunc
	// DetailContent() echo.HandlerFunc
	AllContent() echo.HandlerFunc
}

type ContentService interface {
	AddContent(token interface{}, fileHeader multipart.FileHeader, newContent CoreContent) (CoreContent, error)
	// UpdateContent(token interface{}, contentID uint, updateContent CoreContent) (CoreContent, error)
	// DeleteContent(token interface{}, contentID uint) error
	// DetailContent(token interface{}) ([]CoreContent, error)
	AllContent() ([]CoreContent, error)
}

type ContentData interface {
	AddContent(userID uint, newContent CoreContent) (CoreContent, error)
	// UpdateContent(userID uint, contentID uint, updateContent CoreContent) (CoreContent, error)
	// DeleteContent(userID uint, contentID uint) error
	// DetailContent(contentID uint) ([]CoreContent, error)
	AllContent() ([]CoreContent, error)
}
