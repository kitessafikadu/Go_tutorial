package main

import (
	"fmt"
  "net/http"

  "github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init(){
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
	  fmt.Println("Error loading .env file")
	}
}

func main() {
  r := gin.Default()
  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}