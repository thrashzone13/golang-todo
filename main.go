package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/thrashzone13/golang-todo/controllers"
	"log"
	"net/http"
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
		// v1.GET("/todos", controllers.ToDoController{}.Index)
		v1.POST("/todos", controllers.ToDoController{}.Create)
		// v1.PUT("/todos/:id", controllers.ToDoController{}.Update)
		// v1.DELETE("/todos/:id", controllers.ToDoController{}.Delete)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
	})

	router.Run()
}
