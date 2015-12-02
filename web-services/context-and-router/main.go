package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
)

func main() {

	r := mux.NewRouter()

	r.Path("/something").Methods("GET").Handler(WithContext(context.Background(), ContextHandlerFunc(handleSomething)))

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalln(err)
	}
}

/*
	Architecture
*/

func WithContext(ctx context.Context, h ContextHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTPContext(ctx, w, r)
	})
}

type ContextHandler interface {
	ServeHTTPContext(context.Context, http.ResponseWriter, *http.Request)
}

type ContextHandlerFunc func(context.Context, http.ResponseWriter, *http.Request)

func (f ContextHandlerFunc) ServeHTTPContext(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	f(ctx, w, r)
}

/*
	Middleware
*/

/*
	Handlers
*/

func handleSomething(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	log.Println("handle something with context", ctx)
}
