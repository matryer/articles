package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		max = flag.Int("max", 21, "maximum number")
	)
	flag.Parse()
	f := fib()
	for {
		i := f()
		if i > *max {
			break
		}
		fmt.Println(i)
	}
}

// fib returns a function that returns
// successive Fibonacci numbers.
// The state is stored in closures.
func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
