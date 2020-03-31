package main

import (
            "backend/api/config"
            //"backend/api/models"
)

func main() {

    // Set up database connection
    db := config.SetupDB()

    // Set up routes
    r := config.SetupRoutes(db)


    // Listen and serve on localhost:8080
    r.Run()

}
