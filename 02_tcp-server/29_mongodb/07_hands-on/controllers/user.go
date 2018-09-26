package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/julienschmidt/httprouter"
	"github.com/satori/go.uuid"
)

type UserController struct {
	session map[string]models.User
}

func NewUserController(s map[string]models.User) *UserController {
	return &UserController{s}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}
	id := p.ByName("id")

	u, ok := uc.session[id]
	if !ok {
		w.WriteHeader(404)
		return
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}
	json.NewDecoder(r.Body).Decode(&u)

	uid, _ := uuid.NewV4()
	u.Id = uid.String()
	uc.session[u.Id] = u
	models.StoreUsers(uc.session)
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	_, ok := uc.session[id]
	if !ok {
		w.WriteHeader(404)
		return
	}

	delete(uc.session, id)
	models.StoreUsers(uc.session)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Delete user", id, "\n")
}
