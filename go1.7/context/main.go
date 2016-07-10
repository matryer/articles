package main

import (
	"context"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"
)

func main() {
	db, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	http.Handle("/", WithDBSession(db, http.HandlerFunc(handle)))
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	session := GetDBSession(r)
	// TODO: do something with the database session
}

type contextKey struct {
	name string
}

func (c *contextKey) String() string {
	return "context value " + c.name
}

var contextKeyDBSession = &contextKey{"database-session"}

// GetDBSession gets the mgo.Session for the given Request.
// Handlers must be wrapped with WithDBSession.
func GetDBSession(r *http.Request) *mgo.Session {
	return r.Context().Value(contextKeyDBSession).(*mgo.Session)
}

func WithDBSession(db *mgo.Session, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbSessionCopy := db.Copy()
		defer dbSessionCopy.Close()
		ctx := context.WithValue(r.Context(), contextKey{""}, dbSessionCopy)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
