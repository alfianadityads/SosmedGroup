package handler

import (
	"io"
	"net/http"
	"sosmedapps/features/user"
	"sosmedapps/helper"
	"strings"

	"github.com/labstack/echo/v4"
)

type userController struct {
	srv user.UserService
}

func New(us user.UserService) user.UserHandler {
	return &userController{
		srv: us,
	}
}

// Register implements user.UserHandler
func (uc *userController) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := RegisterRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		res, err := uc.srv.Register(*RequstToCore(input))
		if err != nil {
			if strings.Contains(err.Error(), "email") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already registered"})
			} else if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "all input must fill"})
			} else if strings.Contains(err.Error(), "username") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "username already registered"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success creating account",
		})

	}
}

// Login implements user.UserHandler
func (uc *userController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		input := LoginRequest{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
		}
		tokenGen, res, err := uc.srv.Login(input.UserName, input.Password)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "account not registered"})
			} else if strings.Contains(err.Error(), "empty") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "username or password not allowed empty"})
			} else if strings.Contains(err.Error(), "password") {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong password"})
			} else {
				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
			}
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"data":    res,
			"token":   tokenGen,
			"message": "login success",
		})
	}
}

// Profile implements user.UserHandler
func (uc *userController) Profile() echo.HandlerFunc {
	panic("unimplemented")
}

// Delete implements user.UserHandler
func (uc *userController) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := uc.srv.Delete(c.Get("user"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"message": "internal server error, account fail to delete",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "deleting account successful",
		})
	}
}

// Update implements user.UserHandler
func (uc *userController) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		formHeader, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		}
		//validasi size
		if formHeader.Size > 500000 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "max file size is 500KB"})
		}
		//get file from header
		formFile, err := formHeader.Open()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
		}
		// Validasi
		fileByte, _ := io.ReadAll(formFile)
		fileType := http.DetectContentType(fileByte)
		// log.Println(formHeader.Filename, fileType, formHeader.Size)
		if fileType == "image/png" || fileType == "image/jpeg" {

			uploadUrl, err := helper.NewMediaUpload().FileUpload(helper.File{File: formFile})
			if err != nil {
				return c.JSON(http.StatusInternalServerError, helper.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       &echo.Map{"data": err.Error()},
				})
			}
			input := UpdateRequest{}
			err = c.Bind(&input)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "wrong input format"})
			}
			input.Image = uploadUrl
			// Proses Input Ke Service
			res, err := uc.srv.Update(c.Get("user"), *RequstToCore(input))
			if err != nil {
				if strings.Contains(err.Error(), "email") {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "email already used"})
				} else if strings.Contains(err.Error(), "username") {
					return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "username already used"})
				} else {
					return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
				}
			}
			return c.JSON(http.StatusCreated, map[string]interface{}{
				"data":    res,
				"message": "success updating account",
			})

		} else {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "only jpg or png file can be upload"})
		}
	}
}

// // UploadImg implements user.UserHandler
// func (uc *userController) UploadImg() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		formHeader, err := c.FormFile("file")
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.MediaDto{
// 				StatusCode: http.StatusInternalServerError,
// 				Message:    "error",
// 				Data:       &echo.Map{"data": "Select a file to upload"},
// 			})
// 		}
// 		//validasi size
// 		if formHeader.Size > 500000 {
// 			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "max file size is 500KB"})
// 		}
// 		//get file from header
// 		formFile, err := formHeader.Open()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, helper.MediaDto{
// 				StatusCode: http.StatusInternalServerError,
// 				Message:    "error",
// 				Data:       &echo.Map{"data": err.Error()},
// 			})
// 		}
// 		// Validasi
// 		fileByte, _ := io.ReadAll(formFile)
// 		fileType := http.DetectContentType(fileByte)
// 		// log.Println(formHeader.Filename, fileType, formHeader.Size)
// 		if fileType == "image/png" || fileType == "image/jpeg" {

// 			uploadUrl, err := helper.NewMediaUpload().FileUpload(helper.File{File: formFile})
// 			if err != nil {
// 				return c.JSON(http.StatusInternalServerError, helper.MediaDto{
// 					StatusCode: http.StatusInternalServerError,
// 					Message:    "error",
// 					Data:       &echo.Map{"data": err.Error()},
// 				})
// 			}
// 			err = uc.srv.UploadImg(c.Get("user"), uploadUrl)
// 			if err != nil {
// 				return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error, upload image fail"})
// 			}
// 			return c.JSON(http.StatusOK, helper.MediaDto{
// 				StatusCode: http.StatusOK,
// 				Message:    "success",
// 				Data:       &echo.Map{"data": uploadUrl},
// 			})
// 		} else {
// 			//type file bukan jpeg/png
// 			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "only jpg or png file can be upload"})
// 		}
// 	}
// }
