// models/issue.go

package models

import (
            "github.com/jinzhu/gorm"
)

type Issue struct {
    ID uint `json:"id" gorm:"primary_key"`
    IssueName string `json:"issue_name" binding:"required"`
    IssueTarget string `json:"issue_target"`
    IssueDescription string `json:"issue_description"`
    ContactURL string `json:"contact_url"`
    TwitterUsername string `json:"twitter_username"`
    ContactEmail string `json:"contact_email"`
    FormLetter string `json:"form_letter"`
    TwitterDMForm string `json:"twitter_dm_form"`
}

// Return table name
func (i *Issue) TableName() string {
	return "active_issues"
}

// Get all issues in database
func GetAllIssues(db *gorm.DB, issues *[]Issue) (err error) {
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
