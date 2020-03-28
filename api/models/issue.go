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
