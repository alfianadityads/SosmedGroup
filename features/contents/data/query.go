package data

import (
	"errors"
	"fmt"
	"log"
	"sosmedapps/features/contents"

	"gorm.io/gorm"
)

type contentQry struct {
	db *gorm.DB
}

func New(db *gorm.DB) contents.ContentData {
	return &contentQry{
		db: db,
	}
}

func (cq *contentQry) AddContent(userID uint, newContent contents.CoreContent) (contents.CoreContent, error) {
	cnv := CoreToData(newContent)
	cnv.UserID = uint(userID)
	err := cq.db.Create(&cnv).Error
	if err != nil {
		return contents.CoreContent{}, err
	}
	newContent.ID = cnv.ID

	return newContent, nil
}

// AllContent implements contents.ContentData
func (cq *contentQry) AllContent() ([]contents.CoreContent, error) {
	res := []Content{}
	err := cq.db.Find(&res).Error
	if err != nil {
		log.Println("query error", err.Error())
		return []contents.CoreContent{}, errors.New("server error")
	}
	hasil := []contents.CoreContent{}
	for i := 0; i < len(res); i++ {
		hasil = append(hasil, ContentToCore(res[i]))
		qry := User{}
		err := cq.db.Where("id=?", res[i].UserID).First(&qry).Error
		if err != nil {
			log.Println("no data found")
			return []contents.CoreContent{}, errors.New("data not found")
		}
		hasil[i].Users.Name = qry.Name
		hasil[i].Users.UserName = qry.UserName
		hasil[i].CreateAt = fmt.Sprintf("%d - %s - %d", res[i].CreatedAt.Day(), res[i].CreatedAt.Month(), res[i].CreatedAt.Year())
	}
	return hasil, nil
}

// DetailContent implements contents.ContentData
func (cq *contentQry) DetailContent(contentID uint) (contents.CoreContent, error) {
	res := Content{}
	err := cq.db.Where("id=?", contentID).First(&res).Error
	if err != nil {
		log.Println("no data found")
		return contents.CoreContent{}, errors.New("data not found")
	}
	qry := User{}
	err = cq.db.Where("id=?", res.UserID).First(&qry).Error
	if err != nil {
		log.Println("no data found")
		return contents.CoreContent{}, errors.New("data not found")
	}
	hasil := ContentToCore(res)
	hasil.Users.Name = qry.Name
	hasil.Users.UserName = qry.UserName
	hasil.CreateAt = fmt.Sprintf("%d - %s - %d", res.CreatedAt.Day(), res.CreatedAt.Month(), res.CreatedAt.Year())
	return hasil, nil
}

// UpdateContent implements contents.ContentData
func (cq *contentQry) UpdateContent(userID uint, contentID uint, content string) (string, error) {
	vld := Content{}
	err := cq.db.Where("user_id=? AND id=?", userID, contentID).First(&vld).Error
	if err != nil {
		log.Println("content not found", err.Error())
		return "", errors.New("content cannot edited")
	}
	res := Content{}
	res.Content = content
	qry := cq.db.Where("user_id=? AND id=?", userID, contentID).Updates(&res)
	if qry.RowsAffected <= 0 {
		log.Println("update error : no rows affected")
		return "", errors.New("update error : no rows updated")
	}
	err = qry.Error
	if err != nil {
		log.Println("update error")
		return "", errors.New("query error,update fail")
	}
	return content, nil
}

// DeleteContent implements contents.ContentData
func (cq *contentQry) DeleteContent(userID uint, contentID uint) error {
	//cek apakah content yang akan dihapus milik user yang akan menghapus
	vld := Content{}
	err := cq.db.Where("user_id=? AND id=?", userID, contentID).First(&vld).Error
	if err != nil {
		log.Println("content not found", err.Error())
		return errors.New("content cannot deleted")
	}
	// ok hapus
	qry := cq.db.Delete(&Content{}, contentID)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}
	err = qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete content fail")
	}
	return nil
}
