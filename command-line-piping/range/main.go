package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		from = flag.Int("from", 1, "start number")
		to   = flag.Int("to", 10, "end number")
		step = flag.Int("step", 1, "step interval")
	)
	flag.Parse()
	for i := *from; i <= *to; i += *step {
		fmt.Println(i)
	}
}
