package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var sum, vals int
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			io.WriteString(os.Stderr, "(ignoring) expected whole numbers: "+err.Error())
			continue
		}
		vals++
		sum += val
	}
	fmt.Println(sum / vals)
}
