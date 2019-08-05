package handler

import (
	"gopkg.in/mgo.v2"
)

type (
	// Handler - Handle with DB connection
	Handler struct {
		DB *mgo.Session
	}
)
