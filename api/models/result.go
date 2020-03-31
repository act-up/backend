// models/result.go

package models

import (
            "github.com/jinzhu/gorm"
)

type Result struct {
    ID uint `json:"issue_id" gorm:"primary_key"`
    NumForms string `json:"num_forms_sent"`
    NumEmails string `json:"num_emails_sent"`
    NumDMs string `json:"num_twitter_dms_sent"`
    Resolution string `json:"resolution"`
}

// Return table name
func (i *Result) TableName() string {
	return "results"
}

// Get results of an issue
func GetAResult(db *gorm.DB, result *Result, id string) (err error) {
	if err := db.Where("id = ?", id).First(result).Error; err != nil {
		return err
	}
	return nil
}


// Update results for an issue
func UpdateAResult(db *gorm.DB, result *Result, id string) (err error) {
	db.Save(result)
	return nil
}