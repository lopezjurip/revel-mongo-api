package models

import "gopkg.in/mgo.v2/bson"

/*
Book model
*/
type Book struct {
	ID    bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string        `json:"title" bson:"title"`
	Pages int           `json:"pages" bson:"pages"`
}
