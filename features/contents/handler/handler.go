package handler

import (
	"net/http"
	"sosmedapps/features/contents"
	"sosmedapps/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type contentController struct {
	srv contents.ContentService
}

func New(cs contents.ContentService) contents.ContentHandler {
	return &contentController{
		srv: cs,
	}
}

// AddContent implements contents.ContentHandler
func (cc *contentController) AddContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		formHeader, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		}
		input := AddContentRequest{}
		err = c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		// Proses Input Ke Service
		res, err := cc.srv.AddContent(c.Get("user"), *formHeader, *RequstToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "only jpg or png file can be upload"})
			} else if strings.Contains(err.Error(), "size") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "max file size is 500KB"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success creating new contents",
		})
	}
}

// AllContent implements contents.ContentHandler
func (cc *contentController) AllContent() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := cc.srv.AllContent()
		if err != nil {
			if strings.Contains(err.Error(), "type") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "only jpg or png file can be upload"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"message": "success",
		})

	}
}
