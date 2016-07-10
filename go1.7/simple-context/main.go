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
	http.Handle("/", WithAuth(http.HandlerFunc(handle)))
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, r *http.Request) {
	token := GetAuthToken(r)
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
func GetAuthToken(r *http.Request) string {
	return r.Context().Value(contextKeyAuthToken).(string)
}

func WithAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authtoken, err := verifyAuthHeader(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), contextKeyAuthToken, authtoken)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
