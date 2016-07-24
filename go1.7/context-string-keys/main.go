package main

import (
	"context"
	"log"
)

type contextKey string

var (
	contextKeyOne = contextKey("key")
	contextKeyTwo = contextKey("key")
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextKeyOne, "one")
	ctx = context.WithValue(ctx, contextKeyTwo, "two")
	log.Println(" first key", ctx.Value(contextKeyOne))
	log.Println("second key", ctx.Value(contextKeyTwo))
}
