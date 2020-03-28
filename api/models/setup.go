// models/setup.go

package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func SetupModels() *gorm.DB {
  db, err := gorm.Open("postgres", "user=antonellawilby dbname=advocacy_db")

  if err != nil {
    panic("Failed to connect to database!")
  }

  db.AutoMigrate(&Issue{})

  return db
}
