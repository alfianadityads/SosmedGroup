package services

import (
	"errors"
	"log"
	"sosmedapps/features/comment"
	"sosmedapps/helper"
)

type commentServiceCase struct {
	qry comment.CommentData
}

func New(cd comment.CommentData) comment.CommentService {
	return &commentServiceCase{
		qry: cd,
	}
}

// NewComment implements comment.CommentService
func (css *commentServiceCase) NewComment(token interface{}, NewComment comment.Core) (comment.Core, error) {
	id := helper.ExtractToken(token)
	res, err := css.qry.NewComment(id, NewComment)
	if err != nil {
		log.Println("query error", err.Error())
		return comment.Core{}, errors.New("server error, cannot query data")
	}
	return res, nil
}

// Delete implements comment.CommentService
func (css *commentServiceCase) Delete(token interface{}) error {
	panic("unimplemented")
}
