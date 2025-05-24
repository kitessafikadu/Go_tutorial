package main

import (
	// "fmt"
	"go_crud/initializers"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)

func init(){
	// Load environment variables from .env file
initializers.LoadEnvVariables()
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