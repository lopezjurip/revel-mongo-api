package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mrpatiwi/revel-mongo-api/app/database"
	"github.com/mrpatiwi/revel-mongo-api/app/models"
	"github.com/revel/revel"
	"gopkg.in/mgo.v2/bson"
)

/*
Books controller
*/
type Books struct {
	*revel.Controller
}

/*
Index of all books
*/
func (c Books) Index() revel.Result {
	results := []models.Book{}
	if err := database.Books.Find(bson.M{}).All(&results); err != nil {
		// Internal Server Error
		log.Fatal(err)
	}
	return c.RenderJson(results)
}

/*
Show particular book
*/
func (c Books) Show(id string) revel.Result {
	result := models.Book{}
	if !bson.IsObjectIdHex(id) {
		c.Response.Status = http.StatusBadRequest
		return c.RenderText("id is not valid")

	} else if obj := bson.ObjectIdHex(id); !obj.Valid() {
		c.Response.Status = http.StatusBadRequest
		return c.RenderText("id is not valid")

	} else if err := database.Books.Find(bson.M{"_id": obj}).One(&result); err != nil {
		// Internal Server Error
		log.Fatal(err)
	}
	return c.RenderJson(result)
}

/*
Create book
*/
func (c Books) Create() revel.Result {
	book := &models.Book{}
	if body, err := ioutil.ReadAll(c.Request.Body); err != nil {
		return c.RenderText("bad request")

	} else if err := json.Unmarshal(body, book); err != nil {
		return c.RenderText("could not parse request")

	} else if err := database.Books.Insert(book); err != nil {
		// Internal Server Error
		log.Fatal(err)
		c.Response.Status = http.StatusInternalServerError
		return c.RenderText("could not be saved")
	}
	c.Response.Status = http.StatusCreated
	return c.RenderJson(book)
}
