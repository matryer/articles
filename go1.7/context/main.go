package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func main() {
	greeter := &Greeter{Format: "Hello %s"}
	http.Handle("/", greeter.Handler(http.HandlerFunc(handle)))
	http.ListenAndServe(":8080", nil)
}

func handle(w http.ResponseWriter, req *http.Request) {
	greeter := GetGreeter(req)
	io.WriteString(w, greeter.Greet(req.URL.Query().Get("name")))
}

var ContextKeyGreeter = struct{}{}

type Greeter struct {
	Format string
}

func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf(g.Format, name)
}

func (g *Greeter) Handler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), ContextKeyGreeter, g)
		h.ServeHTTP(w, req.WithContext(ctx))
	})
}

func GetGreeter(req *http.Request) *Greeter {
	return req.Context().Value(ContextKeyGreeter).(*Greeter)
}
