package user

import (
	"github.com/labstack/echo/v4"
)

// perjanjian kontrak
type Core struct {
	ID          uint
	Name        string
	Email       string
	Bio         string
	Image       string
	UserName    string
	Password    string
	DateOfBirth string
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
	Profile() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type UserService interface {
	Register(newUser Core) (Core, error)
	Login(username, password string) (string, Core, error)
	Profile(userToken interface{}) (Core, error)
	Update(userToken interface{}, updateData Core) (Core, error)
	Delete(userToken interface{}) error
}

type UserData interface {
	Register(newUser Core) (Core, error)
	Login(username string) (Core, error)
	Profile(id int) (Core, error)
	Update(id int, updateData Core) (Core, error)
	Delete(id int) error
}
