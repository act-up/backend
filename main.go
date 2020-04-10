package main

import (
    "fmt"
    "github.com/act-up/api/config"
)

func main() {

    // Set up database connection
    db := config.SetupDB()

    fmt.Println("Connected to database.")

    // Set up routes
    r := config.SetupRoutes(db)

    // Listen and serve on localhost:8080
    r.Run()

    // Clean up
    db.Close()

}
