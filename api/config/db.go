// config/db.go

package config

import (
            "fmt"
            "github.com/jinzhu/gorm"
            _ "github.com/jinzhu/gorm/dialects/postgres"
            "backend/api/models"
)

var db *gorm.DB

func SetupDB() *gorm.DB {
    db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=advocacy_db sslmode=disable")

    if err != nil {
        fmt.Printf("Error: %s\n", err)
        panic("Failed to connect to database!")
    }

    db.AutoMigrate(&models.Issue{})

    return db
}
