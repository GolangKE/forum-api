package db

import (
	"../web/models"
	"github.com/jinzhu/gorm"
)

// Seed adds sample data to database
func Seed(db *gorm.DB) error {
	tx := db.Begin()

	// add users to database
	createUsers(tx)

	return tx.Commit().Error
}

func createUsers(db *gorm.DB) error {
	user := models.User{
		Email:    "jim@mail.com",
		Username: "jimmy",
	}

	if db.Where("email = ?", user.Email).First(&user).RecordNotFound() {
		// set user password
		user.HashPassword("123456")

		// create users. rollback if an error occurs
		if err := db.Create(&user).Error; err != nil {
			db.Rollback()
			return err
		}
	}

	return nil
}
