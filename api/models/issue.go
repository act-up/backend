// models/issue.go

package models

import (
            "github.com/jinzhu/gorm"
)

type Issue struct {
    ID uint `json:"issue_id" gorm:"primary_key"`
    Name string `json:"issue_name"`
    Target string `json:"issue_target"`
    Description string `json:"issue_description"`
    URL string `json:"contact_url"`
    Twitter string `json:"twitter_username"`
    Email string `json:"contact_email"`
    Letter string `json:"form_letter"`
    DM string `json:"twitter_dm_form"`
}

// Return table name 
func (i *Issue) TableName() string {
	return "active_issues"
}

// Get all issues in database
func GetIssues(db *gorm.DB, issues *[]Issue) (err error) {
	if err = db.Find(issues).Error; err != nil {
		return err
	}
	return nil
}

// Get an issue
func GetAnIssue(db *gorm.DB, issue *Issue, id string) (err error) {
	if err := db.Where("id = ?", id).First(issue).Error; err != nil {
		return err
	}
	return nil
}
