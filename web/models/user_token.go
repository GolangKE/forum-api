package models

import (
	"time"

	"../../config"
)

// UserToken model
type UserToken struct {
	*Base
	Token     string `gorm:"column:token;unique_index"`
	ExpiresAt time.Time
	IsValid   bool
	User      User
	UserID    string `gorm:"type:uuid"`
}

func (ut UserToken) createNew() (interface{}, error) {
	db := config.Repo()
	err := db.Create(&ut).Error
	return ut, err
}

func (ut UserToken) findOne(condition interface{}) (interface{}, error) {
	db := config.Repo()
	err := db.Where(condition).First(&ut).Error
	return ut, err
}

func (ut *UserToken) FindOne() (interface{}, error) {
	db := config.Repo()
	err := db.Where(ut).First(ut).Error
	return ut, err
}

func (ut *UserToken) InvalidateToken() error {
	db := config.Repo()
	err := db.Model(ut).Update("is_valid", false).Error
	return err
}
