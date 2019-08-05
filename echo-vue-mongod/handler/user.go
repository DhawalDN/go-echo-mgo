package handler

import (
	"echo-vue-mongod/model"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func (h *Handler) Signup(c echo.Context) (err error) {
	// Bind
	u := &model.Student{ID: bson.NewObjectId()}
	if err = c.Bind(u); err != nil {
		return
	}

	// Validate
	if u.Email == "" || u.Password == "" || u.Gender == "" {
		return &echo.HTTPError{Code: http.StatusBadRequest, Message: "invalid email or password"}
	}

	// Save user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("sc").C("students").Insert(u); err != nil {
		return
	}

	return c.JSON(http.StatusCreated, u)
}

func (h *Handler) Login(c echo.Context) (err error) {
	// Bind
	u := new(model.Student)
	if err = c.Bind(u); err != nil {
		return
	}

	// Find user
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("sc").C("students").
		Find(bson.M{"email": u.Email, "password": u.Password}).One(u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email or password"}
		}
		return
	}

	u.Password = "" // Don't send password
	return c.JSON(http.StatusOK, u)
}

func (h *Handler) FetchStudents(c echo.Context) (err error) {
	u := []model.Student{}
	db := h.DB.Clone()
	defer db.Close()
	if err = db.DB("sc").C("students").Find(nil).All(&u); err != nil {
		if err == mgo.ErrNotFound {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Something Wrong"}
		}
		return
	}
	return c.JSON(http.StatusOK, u)
}
