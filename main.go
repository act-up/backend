package main

import (
            "fmt"
            "net/http"
            "github.com/gin-gonic/gin"
)

func main() {

    // Declare a new router
    r := gin.Default()

    // Define GET route to / endpoint
    r.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H{"data": "hello world"})
    })

    db := models.SetupModels()

    // Listen and serve on localhost:8080
    r.Run()

}

func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Hello World!")
}
