package main

import (
	"log"
	"sosmedapps/config"
	uData "sosmedapps/features/user/data"
	uHdl "sosmedapps/features/user/handler"
	uSrv "sosmedapps/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	cfg := config.InitConfig()
	db := config.InitDB(*cfg)

	// gunakan migrate
	// config.Migrate(db)

	// gunakan New dari method user
	usrData := uData.New(db)
	usrSrv := uSrv.New(usrData)
	usrHdl := uHdl.New(usrSrv)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, error=${error}\n",
	}))

	// USER
	e.POST("/register", usrHdl.Register())
	e.POST("/login", usrHdl.Login())
	e.PUT("/users", usrHdl.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/users", usrHdl.Delete(), middleware.JWT([]byte(config.JWTKey)))

	// ========== Run Program ===========
	if err := e.Start(":8000"); err != nil {
		log.Println(err.Error())
	}
}
