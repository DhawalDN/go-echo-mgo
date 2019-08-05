package main

import (
	"echo-vue-mongod/handler"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
)

func main() {
	e := echo.New()

	// Database connection
	db, err := mgo.Dial("localhost:27017")
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Initialize handler
	h := &handler.Handler{DB: db}

	// Routes

	e.POST("/signup", h.Signup)
	e.POST("/login", h.Login)
	e.POST("/fetch", h.FetchStudents)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
