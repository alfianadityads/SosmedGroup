package services

import (
	"errors"
	"sosmedapps/helper"
	"sosmedapps/mocks"
	"testing"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	repo := mocks.NewContentData(t)

	t.Run("success delete content", func(t *testing.T) {
		repo.On("DeleteContent", uint(1), uint(1)).Return(nil).Once()

		srv := New(repo)
		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteContent(pToken, 1)
		assert.Nil(t, err)
		repo.AssertExpectations(t)
	})


	// t.Run("data not found", func(t *testing.T) {
	// 	repo.On("DeleteContent", uint(7), uint(1)).Return(errors.New("data not found")).Once()

	// 	srv := New(repo)

	// 	_, token := helper.GenerateToken(7)
	// 	pToken := token.(*jwt.Token)
	// 	pToken.Valid = true
	// 	err := srv.DeleteContent(pToken, 1)
	// 	assert.NotNil(t, err)
	// 	assert.ErrorContains(t, err, "query error")
	// 	repo.AssertExpectations(t)
	// })

	t.Run("internal server error", func(t *testing.T) {
		repo.On("DeleteContent", uint(1), uint(1)).Return(errors.New("server error")).Once()
		srv := New(repo)

		_, token := helper.GenerateToken(1)
		pToken := token.(*jwt.Token)
		pToken.Valid = true
		err := srv.DeleteContent(pToken, 1)
		assert.NotNil(t, err)
		assert.ErrorContains(t, err, "server error")
		repo.AssertExpectations(t)
	})
}