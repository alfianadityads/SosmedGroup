package services

import (
	"errors"
	"log"
	"mime/multipart"
	"sosmedapps/features/user"
	"sosmedapps/helper"
	"strings"
)

type userServiceCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userServiceCase{
		qry: ud,
	}
}

// Register implements user.UserService
func (usc *userServiceCase) Register(newUser user.Core) (user.Core, error) {
	hashingPassword, err := helper.GeneratePassword(newUser.Password)
	// Validasi
	if newUser.Email == "" || newUser.UserName == "" || newUser.Password == "" || newUser.Name == "" {
		return user.Core{}, errors.New("data not allowed empty")
	}
	if err != nil {
		return user.Core{}, errors.New("cannot hashing password")
	}
	newUser.Password = hashingPassword
	res, err := usc.qry.Register(newUser)
	if err != nil {
		log.Println(err.Error())
		if strings.Contains(err.Error(), "email duplicated") {
			return user.Core{}, errors.New("email already registered")
		} else if strings.Contains(err.Error(), "username duplicated") {
			return user.Core{}, errors.New("username already registered")
		} else {
			return user.Core{}, errors.New("internal server error")
		}
	}
	return res, nil
}

// Login implements user.UserService
func (usc *userServiceCase) Login(username string, password string) (string, user.Core, error) {
	// Validasi
	if username == "" || password == "" {
		return "", user.Core{}, errors.New("username or password not allowed empty")
	}
	res, err := usc.qry.Login(username)
	if err != nil {
		log.Println("query error", err.Error())
		return "", user.Core{}, errors.New("data not found")
	}

	err = helper.ComparePassword(res.Password, password)
	if err != nil {
		log.Println("password compare error", err.Error())
		return "", user.Core{}, errors.New("password not matched")
	}
	tokenGen, _ := helper.GenerateToken(int(res.ID))
	return tokenGen, res, nil
}

// Update implements user.UserService
func (usc *userServiceCase) Update(formHeader multipart.FileHeader, userToken interface{}, updateData user.Core) (user.Core, error) {
	id := helper.ExtractToken(userToken)
	if id <= 0 {
		return user.Core{}, errors.New("data not found")
	}
	// Validasi password
	if updateData.Password != "" {
		hashingPassword, err := helper.GeneratePassword(updateData.Password)
		if err != nil {
			return user.Core{}, errors.New("update password error")
		}
		updateData.Password = hashingPassword
	}
	//-------------ImageProses--------------
	//validasi size
	if formHeader.Size > 500000 {
		return user.Core{}, errors.New("size error")
	}
	//get file from header to check type
	formFile, err := formHeader.Open()
	if err != nil {
		return user.Core{}, errors.New("error open formheader")
	}
	// Validasi Type
	if !helper.TypeFile(formFile) {
		return user.Core{}, errors.New("file type error")
	}
	defer formFile.Close()
	formFile, _ = formHeader.Open()
	uploadUrl, err := helper.NewMediaUpload().FileUpload(helper.File{File: formFile})
	if err != nil {
		return user.Core{}, errors.New("server error")
	}
	updateData.Image = uploadUrl

	// -----------------Input data to query----------------
	res, err := usc.qry.Update(id, updateData)
	if err != nil {
		log.Println("query error", err.Error())
		if strings.Contains(err.Error(), "email duplicated") {
			return user.Core{}, errors.New("email already used")
		} else if strings.Contains(err.Error(), "username duplicated") {
			return user.Core{}, errors.New("username already used")
		} else {
			return user.Core{}, errors.New("query error, update fail")
		}
	}

	return res, nil
}

// Profile implements user.UserService
func (usc *userServiceCase) Profile(userToken interface{}) (user.Core, error) {
	panic("unimplemented")
}

// Delete implements user.UserService
func (usc *userServiceCase) Delete(userToken interface{}) error {
	id := helper.ExtractToken(userToken)
	if id <= 0 {
		return errors.New("id not found, server error")
	}
	err := usc.qry.Delete(id)
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("query error, delete account fail")
	}
	return nil
}
