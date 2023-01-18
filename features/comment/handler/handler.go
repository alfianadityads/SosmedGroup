package handler

import (
	"net/http"
	"sosmedapps/features/comment"
	"strconv"

	"github.com/labstack/echo/v4"
)

type commentController struct {
	srv comment.CommentService
}

func New(ch comment.CommentService) comment.CommentHandler {
	return &commentController{
		srv: ch,
	}
}

// NewComment implements comment.CommentHandler
func (cc *commentController) NewComment() echo.HandlerFunc {
	return func(c echo.Context) error {
		cID := c.Param("id")
		contentID, _ := strconv.Atoi(cID)
		input := NewComment{}
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{"message": "comment cannot allowed empty"})
		}
		input.ContentID = uint(contentID)
		res, err := cc.srv.NewComment(c.Get("user"), *RequstToCore(input))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"message": "internal server error"})
		}
		return c.JSON(http.StatusCreated, map[string]interface{}{
			"data":    res,
			"message": "success creating comment",
		})
	}
}

// Delete implements comment.CommentHandler
func (cc *commentController) Delete() echo.HandlerFunc {
	panic("unimplemented")
}
