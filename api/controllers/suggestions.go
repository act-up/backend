// controllers/issues.go

package controllers

import (
    "net/http"
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    "backend/api/models"
)

func GetSuggestions(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var active_issues []models.Issue
    err := models.GetIssues(db, &active_issues)

    if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, active_issues)
	}

    //db.Find(&active_issues)
    //c.JSON(http.StatusOK, gin.H{"data": active_issues})

}
