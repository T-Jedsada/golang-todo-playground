package todo

import (
	"log"

	"gopkg.in/mgo.v2"
)

var session *mgo.Session

type Context struct {
	MongoSession *mgo.Session
}

func (c *Context) Close() {
	c.MongoSession.Close()
}

func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB("demo").C(name)
}

func NewContext() *Context {
	session := getSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}

func getSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial("mongo:27017")
		if err != nil {
			log.Fatal("Could not connect to mongo: ", err.Error())
		}
		if err != nil {
			log.Fatalf("[GetSession]: %s\n", err)
		}
	}
	return session
}
