package db

import (
	"strings"

	"../web/models"
	"github.com/jinzhu/gorm"
)

// Seed adds sample data to database
func Seed(db *gorm.DB) error {
	tx := db.Begin()

	// add users to database
	createUsers(tx)

	// add categories to database
	createCategories(tx)

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

func createCategories(db *gorm.DB) error {
	categories := []string{
		"Help/Questions",
		"Code Review",
		"Releases",
		"Discussions",
		"Golang News",
		"Confs & Meet Ups",
		"Show & Tell",
		"Announcements",
		"Sponsor Messages",
		"Learning Resources",
		"Site Feedback",
		"Jobs",
	}

	for _, c := range categories {
		code := strings.ToUpper(c)
		category := models.Category{
			Name:     c,
			Code:     strings.Replace(code, " ", "-", -1),
			IsActive: true,
		}

		if db.Where("code = ?", category.Code).First(&category).RecordNotFound() {
			// create category. rollback if an error occurs
			if err := db.Create(&category).Error; err != nil {
				db.Rollback()
				return err
			}
		}
	}

	return nil
}
