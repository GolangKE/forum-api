package models

// Category model
type Category struct {
	*Base
	Name        string
	Code        string
	Description string
	IsActive    bool
}
