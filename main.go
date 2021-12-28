package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thrashzone13/golang-todo/controllers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file found!")
	}
}

func main() {

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", controllers.PingController{}.Ping)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	router.Run()
}
