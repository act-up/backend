package main

import (
    "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
  db, err := gorm.Open("postgres", "user=postgres password=postgres host=/cloudsql/actup-273804:us-central1:postgres-actup dbname=advocacy_db sslmode=disable")

  if err != nil {
      fmt.Printf("Error: %s\n", err)
      panic("Failed to connect to database!")
  } else {
      fmt.Printf("connected!")
  }


  defer db.Close()
}
