// config/db.go

package config

import (
            "fmt"
            "github.com/jinzhu/gorm"
            _ "github.com/jinzhu/gorm/dialects/postgres"
            //"github.com/GoogleCloudPlatform/cloudsql-proxy/cmd/cloud_sql_proxy"
            "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
            "github.com/act-up/backend/api/models"
)

var db *gorm.DB

func SetupDB() *gorm.DB {

    dsn := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable",
                   "cellular-codex-273108:us-central1:postgres-actup",
                   "advocacy_db",
                   "postgres",
                   "postgres")

    db, err := sql.Open("cloudsqlpostgres", dsn)

    //db, err := gorm.Open("postgres", "dbname=advocacy_db user=postgres password=postgres host=/cloudsql/35.223.164.101 sslmode=disable")

    if err != nil {
        fmt.Printf("Error: %s\n", err)
        panic("Failed to connect to database!")
    }

    db.AutoMigrate(&models.Issue{})

    return db
}
