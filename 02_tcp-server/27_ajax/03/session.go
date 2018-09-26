package main

import (
	"fmt"
	"net/http"
	"time"

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
	c.MaxAge = sessionLength
	http.SetCookie(res, c)

	var u user

	if s, ok := dbSessions[c.Value]; ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
		u = dbUsers[s.un]
	}
	return u
}

func alreadyLoggedIn(res http.ResponseWriter, req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := dbSessions[c.Value]
	if ok {
		s.lastActivity = time.Now()
		dbSessions[c.Value] = s
	}
	_, ok = dbUsers[s.un]
	c.MaxAge = sessionLength
	http.SetCookie(res, c)
	return ok
}

func cleanSessions() {
	fmt.Println("BEFOR CLEAN")
	showSessions()
	for k, v := range dbSessions {
		if time.Now().Sub(v.lastActivity) > (time.Second * 30) {
			delete(dbSessions, k)
		}
	}
	dbSessionCleaned = time.Now()
	fmt.Println("AFTER CLEAN")
	showSessions()
}

func showSessions() {
	fmt.Println("*****")
	for k, v := range dbSessions {
		fmt.Println(k, v.un)
	}
	fmt.Println()
}
