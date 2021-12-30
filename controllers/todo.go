package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thrashzone13/golang-todo/forms"
	"github.com/thrashzone13/golang-todo/models"
)

type ToDoController struct{}

func (cntrlr ToDoController) Index(c *gin.Context) {

}

func (cntrlr ToDoController) Create(c *gin.Context) {

	var data forms.ToDoData

	if c.BindJSON(&data) != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Provide relevant fields"})
		c.Abort()
		return
	}

	err := models.ToDoModel{}.Create(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Problem creating todo"})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New todo added"})
}

func (cntrlr ToDoController) Update(c *gin.Context) {

}

func (cntrlr ToDoController) Delete(c *gin.Context) {

}
