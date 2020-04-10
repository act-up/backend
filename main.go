package main

import (
            "github.com/act-up/backend/api/config"
)

func main() {

    // Set up database connection
    db := config.SetupDB()

    // Set up routes
    r := config.SetupRoutes(db)

    // Listen and serve on localhost:8080
    r.Run()

}
