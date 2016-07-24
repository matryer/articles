package main

import (
	"context"
	"log"
)

type contextKeyA string
type contextKeyB string

var (
	contextKeyOne = contextKeyA("key")
	contextKeyTwo = contextKeyB("key")
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, contextKeyOne, "one")
	ctx = context.WithValue(ctx, contextKeyTwo, "two")
	log.Println(" first key", ctx.Value(contextKeyOne))
	log.Println("second key", ctx.Value(contextKeyTwo))
}

// Output:
// 2016/07/24 14:44:40  first key two
// 2016/07/24 14:44:40 second key two
