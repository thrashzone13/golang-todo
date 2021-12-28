package models

import "gopkg.in/mgo.v2/bson"

type ToDo struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string        `json:"name" bson:"name"`
	Description string        `json:"description" bson:"description"`
	IsDone      bool          `json:"is_done" bson:"is_done"`
}

type ToDoModel struct{}

func (model ToDoModel) Add() {

}
