package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
)

func main() {
	// get an engine instance
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		response, err := http.Get("http://gin-receiver:3000/animals")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the CRUD server"})
			return
		}
		defer response.Body.Close()

		body, err := io.ReadAll(response.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read body from the CRUD server"})
			return
		}

		c.Data(http.StatusOK, "application/json", body)
	})

	// run the server
	err := r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
