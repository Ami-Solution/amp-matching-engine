package daos

import (
	"github.com/Proofsuite/amp-matching-engine/app"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Database struct contains the pointer to mgo.session
// It is a wrapper over mgo to help utilize mgo connection pool
type Database struct {
	session *mgo.Session
}

// Global instance of Database struct for singleton use
var db *Database

// InitSession initializes a new session with mongodb
func InitSession() error {
	if db == nil {
		db1, err := mgo.Dial(app.Config.DSN)
		if err != nil {
			return err
		}
		db = &Database{db1}
	}
	return nil
}

// Create is a wrapper for mgo.Insert function.
// It creates a copy of session initialized, sends query over this session
// and returns the session to connection pool
func (d *Database) Create(dbName, collection string, data ...interface{}) (err error) {
	sc := d.session.Copy()
	defer sc.Close()

	err = sc.DB(dbName).C(collection).Insert(data...)
	return
}

// GetByID is a wrapper for mgo.FindId function.
// It creates a copy of session initialized, sends query over this session
// and returns the session to connection pool
func (d *Database) GetByID(dbName, collection string, id bson.ObjectId, response interface{}) (err error) {
	sc := d.session.Copy()
	defer sc.Close()

	err = sc.DB(dbName).C(collection).FindId(id).One(response)
	return
}

// Get is a wrapper for mgo.Find function.
// It creates a copy of session initialized, sends query over this session
// and returns the session to connection pool
func (d *Database) Get(dbName, collection string, query interface{}, offset, limit int, response interface{}) (err error) {
	sc := d.session.Copy()
	defer sc.Close()

	err = sc.DB(dbName).C(collection).Find(query).Skip(offset).Limit(limit).All(response)
	return
}

// GetWithSort is a wrapper for mgo.Find function with SORT function in pipeline.
// It creates a copy of session initialized, sends query over this session
// and returns the session to connection pool
func (d *Database) GetWithSort(dbName, collection string, query interface{}, sort []string, offset, limit int, response interface{}) (err error) {
	sc := d.session.Copy()
	defer sc.Close()

	err = sc.DB(dbName).C(collection).Find(query).Sort(sort...).Skip(offset).Limit(limit).All(response)
	return
}

// Update is a wrapper for mgo.Update function.
// It creates a copy of session initialized, sends query over this session
// and returns the session to connection pool
func (d *Database) Update(dbName, collection string, query interface{}, update interface{}) (err error) {
	sc := d.session.Copy()
	defer sc.Close()

	err = sc.DB(dbName).C(collection).Update(query, update)
	return
}

// Aggregate is a wrapper for mgo.Pipe function.
// It is used to make mongo aggregate pipeline queries
// It creates a copy of session initialized, sends query over this session
// and returns the session to connection pool
func (d *Database) Aggregate(dbName, collection string, query []bson.M) (response []interface{}, err error) {
	sc := d.session.Copy()
	defer sc.Close()
	err = sc.DB(dbName).C(collection).Pipe(query).All(&response)
	return
}
