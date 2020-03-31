// config/routes.go

package config

import (
            "github.com/gin-gonic/gin"
            "github.com/jinzhu/gorm"
            "backend/api/controllers"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {

    // Declare a new router
    r := gin.Default()

    r.Use(func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    })

    v1 := r.Group("/")
    {
        // List all issues in database
        v1.GET("/active_issues", controllers.GetIssues)
    }

    return r

}
