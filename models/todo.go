package models

import (
	"github.com/thrashzone13/golang-todo/forms"
	"gopkg.in/mgo.v2/bson"
)

type ToDo struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	IsDone      bool          `json:"is_done" bson:"is_done"`
}

type ToDoModel struct{}

func (model ToDoModel) Create(data forms.ToDoData) error {

	collection := dbConnect.Use("todos")

	err := collection.Insert(bson.M{
		"title":       data.Title,
		"description": data.Description,
		"is_done":     data.IsDone,
	})

	return err
}

func (model ToDoModel) Delete(id int) {

}
