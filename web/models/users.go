package models

// User model, table: users
type User struct {
	*Model
	Username       string `gorm:"column:username;unique_index"`
	FullName       string `gorm:"column:full_name"`
	Email          string `gorm:"column:email;unique_index"`
	PasswordDigest string `gorm:"column:password_digest"`
}
