package data

// type contentQry struct {
// 	db *gorm.DB
// }

// func NewCont(db *gorm.DB) contents.ContentData {
// 	return &contentQry{
// 		db: db,
// 	}
// }

// func (cq *contentQry) AddContent(userID uint, newContent contents.CoreContent) (contents.CoreContent, error) {
// 	cnv := CoreToData(newContent)
// 	cnv.UserID = uint(userID)
// 	err := cq.db.Create(&cnv).Error
// 	if err != nil {
// 		return contents.CoreContent{}, err
// 	}
// 	newContent.ID = cnv.ID

// 	return newContent, nil
// }

// func (cq *contentQry) UpdateContent(userID uint, contentID uint, updateContent contents.CoreContent) (contents.CoreContent, error) {

// 	cnv := CoreToData(updateContent)
// 	qry := cq.db.Where("user_id = ? AND id = ?", userID, contentID).Updates(&cnv)

// 	affrows := qry.RowsAffected
// 	if affrows <= 0 {
// 		log.Println("no rows affected")
// 		return contents.CoreContent{}, errors.New("no content has been changed")
// 	}

// 	err := qry.Error
// 	if err != nil {
// 		log.Println("update query error", err.Error())
// 		return contents.CoreContent{}, err
// 	}
// 	cToCore := cnv.ContentToCore()
// 	return cToCore, nil
// }

// func (cq *contentQry) DeleteContent(userID uint, contentID uint) error {
// 	qry := cq.db.Where("id = ? AND id_user = ?", contentID, userID).Delete(&Content{})

// 	affrows := qry.RowsAffected
// 	if affrows >= 0 {
// 		log.Println("now rows affected")
// 		return errors.New("no content has been deleted")
// 	}

// 	err := qry.Error
// 	if err != nil {
// 		log.Println("delete query error")
// 		return errors.New("can't delete content")
// 	}
// 	return nil
// }

// func (cq *contentQry) DetailContent(contentID uint) ([]contents.CoreContent, error) {
// 	tmp := []Content{}
// 	if err := cq.db.Preload("User").Where("id = ?", contentID).First(&tmp); err != nil {
// 		log.Println("get by ID content querry error", err.Error)
// 		return ToCoreContent(tmp), err.Error
// 	}
// 	tmp2 := ToCoreContent(tmp)
// 	return tmp2, nil
// }

// func (cq *contentQry) AllContent() ([]contents.CoreContent, error)
