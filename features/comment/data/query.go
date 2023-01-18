package data

import (
	"errors"
	"fmt"
	"log"
	"sosmedapps/features/comment"

	"gorm.io/gorm"
)

type commentQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) comment.CommentData {
	return &commentQuery{
		db: db,
	}
}

// NewComment implements comment.CommentData
func (cq *commentQuery) NewComment(id int, newComment comment.Core) (comment.Core, error) {
	data := CoreToData(newComment)
	data.UserID = uint(id)
	err := cq.db.Create(&data).Error
	if err != nil {
		log.Println("query error", err.Error())
		return comment.Core{}, errors.New("server error")
	}
	usrQry := User{}
	err = cq.db.Where("id = ?", id).First(&usrQry).Error
	if err != nil {
		log.Println("query error", err.Error())
		return comment.Core{}, errors.New("server error")
	}
	newComment.OwnerName = usrQry.UserName
	newComment.CreateAt = fmt.Sprintf("%d - %s - %d", data.CreatedAt.Day(), data.CreatedAt.Month(), data.CreatedAt.Year())
	return newComment, nil

}

// Delete implements comment.CommentData
func (cq *commentQuery) Delete(commentID int) error {
	qry := cq.db.Delete(&Comment{}, commentID)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no comment has delete")
	}
	err := qry.Error
	if err != nil {
		log.Println("query error", err.Error())
		return errors.New("delete comment fail")
	}
	return nil
}
