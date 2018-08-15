package main

import (
	"net/http"

	"github.com/satori/go.uuid"
)

func getUser(res http.ResponseWriter, req *http.Request) user {
	c, err := req.Cookie("session")
	if err != nil {
		uId, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: uId.String(),
		}
	}
	http.SetCookie(res, c)

	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLogIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
