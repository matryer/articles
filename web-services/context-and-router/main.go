package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func main() {

	r := mux.NewRouter()

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (f ContextHandlerFunc) ServeHTTPContext(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	f(ctx, w, r)
}
