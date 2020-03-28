// controllers/issues.go

package controllers

import (
    "net/http"
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    "backend/api/models"
)

func ListIssues(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var active_issues []models.Issue
    db.Find(&active_issues)

    c.JSON(http.StatusOK, gin.H{"data": active_issues})

}
