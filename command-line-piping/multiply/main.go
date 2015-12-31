package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var (
		by = flag.Int("by", 2, "number to multiply by")
	)
	flag.Parse()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			io.WriteString(os.Stderr, "(ignoring) expected whole numbers: "+err.Error())
			continue
		}
		fmt.Println(val * *by)
	}
}
