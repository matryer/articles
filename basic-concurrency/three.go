package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("start")
	wg.Add(1)
	go doSomething()
	fmt.Println("end")
	wg.Wait()
}

func doSomething() {
	fmt.Println("do something")
	wg.Done()
}
