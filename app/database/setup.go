package database

import "gopkg.in/mgo.v2"

/*
Database session
*/
var Session *mgo.Session

/*
Book's model connection
*/
var Books *mgo.Collection

/*
Init database
*/
func Init(uri, dbname string) error {
	session, err := mgo.Dial(uri)
	if err != nil {
		return err
	}

	// See https://godoc.org/labix.org/v2/mgo#Session.SetMode
	session.SetMode(mgo.Monotonic, true)

	// Expose session and models
	Session = session
	Books = session.DB(dbname).C("books")

	return nil
}
