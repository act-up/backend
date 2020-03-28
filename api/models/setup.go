// models/setup.go

package models

import (
            "fmt"
            "github.com/jinzhu/gorm"
            _ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
    db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=advocacy_db sslmode=disable")

    if err != nil {
        fmt.Printf("Error: %s\n", err)
        panic("Failed to connect to database!")
    }

    return db
}
