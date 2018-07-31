package models

import (
	"../../config"
	"golang.org/x/crypto/bcrypt"
)

// User model, table: users
type User struct {
	*Base
	Username       string `gorm:"column:username;unique_index"`
	FullName       string `gorm:"column:full_name"`
	Email          string `gorm:"column:email;unique_index"`
	PasswordDigest string `gorm:"column:password_digest"`
}

func (u User) createNew() (interface{}, error) {
	db := config.Repo()
	err := db.Create(&u).Error
	return u, err
}

func (u User) findOne(condition interface{}) (interface{}, error) {
	db := config.Repo()
	err := db.Where(condition).First(&u).Error
	return u, err
}

// HashPassword generates encrypted password from
// password string and assigns to PasswordDigest
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	u.PasswordDigest = string(bytes)
	return err
}

// CheckPassword compares given password against
// the user's hashed PasswordDigest
func (u *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(password))
	return err
}
