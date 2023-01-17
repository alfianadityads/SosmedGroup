package data

import (
	"errors"
	"log"
	"sosmedapps/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

// Register implements user.UserData
func (uq *userQuery) Register(newUser user.Core) (user.Core, error) {
	// validasi cek duplicate email
	dupeEmail := CoreToData(newUser)
	err := uq.db.Where("email = ?", newUser.Email).First(&dupeEmail).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("email duplicated")
	}
	// validasi cek duplicate username
	dupeUN := CoreToData(newUser)
	err = uq.db.Where("user_name = ?", newUser.UserName).First(&dupeUN).Error
	if err == nil {
		log.Println("duplicated")
		return user.Core{}, errors.New("username duplicated")
	}
	// proses query
	qry := CoreToData(newUser)
	err = uq.db.Create(&qry).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("query error")
	}
	newUser.ID = qry.ID
	return newUser, nil

}

// Login implements user.UserData
func (uq *userQuery) Login(username string) (user.Core, error) {
	qry := User{}
	err := uq.db.Where("email = ? OR user_name = ?", username, username).First(&qry).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("query error")
	}
	return DataToCore(qry), nil
}

// Update implements user.UserData
func (uq *userQuery) Update(id int, updateData user.Core) (user.Core, error) {

	if updateData.Email != "" {
		// Proses validasi cek duplicate email
		dupe := CoreToData(updateData)
		err := uq.db.Where("email = ?", dupe.Email).First(&dupe).Error
		if err == nil {
			log.Println("duplicated")
			return user.Core{}, errors.New("email duplicated")
		}
	}
	if updateData.UserName != "" {
		// Proses validasi cek duplicate username
		dupe := CoreToData(updateData)
		err := uq.db.Where("user_name = ?", dupe.UserName).First(&dupe).Error
		if err == nil {
			log.Println("duplicated")
			return user.Core{}, errors.New("username duplicated")
		}
	}
	data := CoreToData(updateData)
	qry := uq.db.Where("id = ?", id).Updates(&data)
	if qry.RowsAffected <= 0 {
		log.Println("update error : no rows affected")
		return user.Core{}, errors.New("update error : no rows updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("update error")
		return user.Core{}, errors.New("query error,update fail")
	}
	return DataToCore(data), nil
}

// Profile implements user.UserData
func (uq *userQuery) Profile(id int) (user.Core, error) {
	panic("unimplemented")
}

// Delete implements user.UserData
func (uq *userQuery) Delete(id int) error {
	qry := uq.db.Delete(&User{}, id)
	rowAffect := qry.RowsAffected
	if rowAffect <= 0 {
		log.Println("no data processed")
		return errors.New("no user has delete")
	}
	err := qry.Error
	if err != nil {
		log.Println("delete query error", err.Error())
		return errors.New("delete account fail")
	}
	return nil
}

// UploadImg implements user.UserData
func (uq *userQuery) UploadImg(id int, newImage string) error {
	data := User{}
	data.Image = newImage
	qry := uq.db.Where("id = ?", id).Updates(&data)
	if qry.RowsAffected <= 0 {
		log.Println("update error : no rows affected")
		return errors.New("update error : no rows updated")
	}
	err := qry.Error
	if err != nil {
		log.Println("upload error")
		return errors.New("query error,upload image fail")
	}
	return nil
}
