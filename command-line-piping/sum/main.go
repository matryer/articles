package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	sum := 0
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		val, err := strconv.Atoi(s.Text())
		if err != nil {
			io.WriteString(os.Stderr, "(ignoring) expected whole numbers: "+err.Error())
			continue
		}
		sum += val
	}
	fmt.Println(sum)
}
