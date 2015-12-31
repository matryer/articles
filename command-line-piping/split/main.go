package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	var (
		delimiter = flag.String("delimiter", ",", "character to split by")
	)
	flag.Parse()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		println(strings.Split(s.Text(), *delimiter)...)
	}
}

func println(s ...string) {
	for _, ss := range s {
		fmt.Println(ss)
	}
}
