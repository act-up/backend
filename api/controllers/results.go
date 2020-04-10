// controllers/results.go

package controllers

import (
    "net/http"
    "github.com/jinzhu/gorm"
    "github.com/gin-gonic/gin"
    "github.com/act-up/backend/api/models"
)

// Get all results
func GetResults(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

    var results []models.Result
    err := models.GetAllResults(db, &results)

    if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, results)
	}

    //db.Find(&active_issues)
    //c.JSON(http.StatusOK, gin.H{"data": active_issues})

}

// Get results for an issue by ID
func GetResult(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)
	id := c.Params.ByName("id")

	var result models.Result
	err := models.GetAResult(db, &result, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, result)
	}
}


// Update an issue results record
func UpdateResult(c *gin.Context) {
    db := c.MustGet("db").(*gorm.DB)

	var result models.Result
	id := c.Params.ByName("id")
	err := models.GetAResult(db, &result, id)

	if err != nil {
		c.JSON(http.StatusNotFound, result)
	}

	c.BindJSON(&result)
	err = models.UpdateAResult(db, &result, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
