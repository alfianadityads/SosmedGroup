package services

import (
	"errors"
	"sosmedapps/features/user"
	"sosmedapps/helper"
	"sosmedapps/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("success creating account", func(t *testing.T) {
		inputData := user.Core{Name: "Alif", Email: "alif@example.com", UserName: "alif123", Password: "alif342"}
		resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@example.com", UserName: "alif123"}
		repo.On("Register", mock.Anything).Return(resData, nil).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		assert.Equal(t, resData.UserName, res.UserName)
		repo.AssertExpectations(t)
	})

	t.Run("internal server error", func(t *testing.T) {
		inputData := user.Core{Name: "Alif", Email: "alif@example.com", UserName: "alif123", Password: "alif342"}
		resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@example.com", UserName: "alif123"}
		repo.On("Register", mock.Anything).Return(resData, errors.New("internal server error")).Once()
		srv := New(repo)
		res, err := srv.Register(inputData)
		assert.NotNil(t, err)
		assert.Equal(t, uint(0), res.ID)
		assert.ErrorContains(t, err, "server")
		repo.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("login success", func(t *testing.T) {
		inputEmail := "alif@example.com"
		hashed, _ := helper.GeneratePassword("alif342")
		resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@example.com", UserName: "alif123", Password: hashed}

		repo.On("Login", inputEmail).Return(resData, nil)

		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "alif342")
		assert.Nil(t, err)
		assert.NotEmpty(t, token)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	// 
	t.Run("not found", func(t *testing.T) {
		inputEmail := "alif@example.com"

		repo.On("Login", inputEmail).Return(user.Core{}, errors.New("not found"))

		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "alif342")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Empty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})

	// wrong password
	t.Run("wrong password", func(t *testing.T) {
		inputEmail := "alif@example.com"
		hashed, _ := helper.GeneratePassword("alif342")
		resData := user.Core{ID: uint(1), Name: "Alif", Email: "alif@example.com", UserName: "alif123", Password: hashed}
		repo.On("Login", inputEmail).Return(resData, nil)

		srv := New(repo)
		token, res, err := srv.Login(inputEmail, "alif342")
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "wrong password")
		assert.NotEmpty(t, token)
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestProfile(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("succes show profile", func(t *testing.T) {
		resData := user.Core{ID: 1, Name: "Alif", Email: "alif@example.com", UserName: "alif123"}

		repo.On("Profile", int(1)).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateToken(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Profile(pToken)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("jwt not found", func(t *testing.T) {
		srv := New(repo)

		_, token := helper.GenerateToken(1)

		res, err := srv.Profile(token)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "not found")
		assert.Equal(t, 0, res.ID)
	})

	// internal server error
	t.Run("internal server error", func(t *testing.T) {
		repo.On("Profile", mock.Anything).Return(user.Core{}, errors.New("internal server error")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		res, err := srv.Profile(pToken)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	repo := mocks.NewUserData(t)

	t.Run("success updating account", func(t *testing.T) {
		inputData := user.Core{Name: "Alif", Email: "alif@example.com", UserName: "alif123"}

		hash, _ := helper.GeneratePassword("alfian1221")
		resData := user.Core{ID: uint(1),Name: "Alif", Email: "alif@example.com", UserName: "alif123", Password: hash}
		repo.On("Update", uint(1), inputData).Return(resData, nil).Once()

		srv := New(repo)

		_, token := helper.GenerateToken(1)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, inputData)
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, res.ID)
		repo.AssertExpectations(t)
	})

	t.Run("Not found", func(t *testing.T) {
		inputData := user.Core{Name: "Alif", Email: "alif@example.com", UserName: "alif123"}
		repo.On("Update", uint(2), inputData).Return(user.Core{}, errors.New("not found")).Once()

		srv := New(repo)

		_, token := helper.GenerateToken(2)

		pToken := token.(*jwt.Token)
		pToken.Valid = true

		res, err := srv.Update(pToken, inputData)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "tidak ditemukan")
		assert.Equal(t, uint(0), res.ID)
		repo.AssertExpectations(t)
	})
}