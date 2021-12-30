package db

import (
	"gopkg.in/mgo.v2"
	// "os"
	"time"
)

type DBConnection struct {
	session *mgo.Session
}

// NewConnection handles connecting to a mongo database
func NewConnection() (conn *DBConnection) {

	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		// Address if its a local db then the value host=localhost
		Addrs: []string{"todo_mongo"},
		// Timeout when a failure to connect to db
		Timeout: 10 * time.Second,
		// Database name
		Database: "test",
		// Database credentials if your db is protected
		// Username: "db",
		// Password: "dbpass",
	})

	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	conn = &DBConnection{session}
	return conn
}

// Use handles connect to a certain collection
func (conn *DBConnection) Use(tableName string) (collection *mgo.Collection) {
	// This returns method that interacts with a specific collection and table
	return conn.session.DB("").C(tableName)
}

// Close handles closing a database connection
func (conn *DBConnection) Close() {
	// This closes the connection
	conn.session.Close()
}
