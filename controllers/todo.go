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

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	if err := new(models.ToDoModel).Create(data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New todo added"})
}

func (cntrlr ToDoController) Update(c *gin.Context) {

}

func (cntrlr ToDoController) Delete(c *gin.Context) {

}
