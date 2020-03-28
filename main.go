package main

import (
            "net/http"
            "github.com/gin-gonic/gin"
            "backend/api/models"
            "backend/api/controllers"
)

func main() {

    // Declare a new router
    r := gin.Default()

    // Define GET route to / endpoint
    r.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{"data": "hello world"})
    })

    db := models.SetupModels()

    r.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    // List all issues in database
    r.GET("/active_issues", controllers.ListIssues)

    // Listen and serve on localhost:8080
    r.Run()

}
