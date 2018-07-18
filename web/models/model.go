package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// Model is forum-api base model
type Model struct {
	ID        string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate sets Model's primary key value to uuid
func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	id, err := uuid.NewV4()

	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return err
	}

	scope.SetColumn("ID", id)
	return nil
}
