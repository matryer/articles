package main

import (
	"context"
	"errors"
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
	http.Handle("/", WithAuth(db, http.HandlerFunc(handle)))
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	token, err := GetAuthToken(r)
	// TODO: do something with the token
}

type contextKey struct {
	name string
}

func (c *contextKey) String() string {
	return "context value " + c.name
}

var contextKeyAuthToken = &contextKey{"authtoken"}

// GetSomething gets the something for the given Request.
// Handlers must be wrapped with WithSomething.
func GetAuthToken(r *http.Request) (string, error) {
	authTokenVal := r.Context().Value(contextKeyAuthToken)
	authTokenStr, ok := authTokenVal.(string)
	if !ok {
		return "", errors.New("missing authorization token")
	}
	return authTokenStr, nil
}
, e
func WithAuth(db *db.Session, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbsession := db.Copy()
		defer dbsession.Close()
		authtoken, err := verifyAuthHeader(dbsession, r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyAuthToken, authtoken)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
