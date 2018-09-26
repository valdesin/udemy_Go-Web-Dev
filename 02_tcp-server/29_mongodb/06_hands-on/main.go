package main

import (
	"net/http"

	"./controllers"
	"./models"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user/", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe(":9090", r)
}

func getSession() map[string]models.User {
	return make(map[string]models.User)
}
