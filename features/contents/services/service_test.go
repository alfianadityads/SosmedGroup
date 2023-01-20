package services

import (
	"errors"
	"sosmedapps/features/contents"
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


func TestUpdateContent(t *testing.T) {
	repo := mocks.NewContentData(t)
	inputData := contents.CoreContent{ID: uint(1), Content: "hello everyone"}
	resData := contents.CoreContent{ID: uint(1),Content: "hello everybody"}

	t.Run("success update content", func(t *testing.T) {
		repo.On("UpdateContent", uint(1), uint(1), inputData).Return(resData, nil).Once()
		srv := New(repo)
		_, tokenIDUser := helper.GenerateToken(1)
		userID := tokenIDUser.(*jwt.Token)
		userID.Valid = true
		_, err := srv.UpdateContent(userID, uint(1), "hello everyone")
		assert.Nil(t, err)
		assert.Equal(t, resData.ID, userID)
		repo.AssertExpectations(t)
	})
}