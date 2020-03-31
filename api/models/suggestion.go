// models/suggestion.go

package models

import (
            "github.com/jinzhu/gorm"
)

type Suggestion struct {
    ID uint `json:"suggestion_id" gorm:"primary_key"`
    ContactName string `json:"contact_name"`
    ContactEmail string `json:"contact_email"`
    Target string `json:"issue_target"`
    Description string `json:"issue_description"`
    Comments string `json:"other_comments"`
}

// Return table name
func (i *Suggestion) TableName() string {
	return "suggested"
}

// Create a suggested issue
func CreateASuggestion(db *gorm.DB, suggestion *Suggestion) (err error) {
	if err = db.Create(suggestion).Error; err != nil {
		return err
	}
	return nil
}
