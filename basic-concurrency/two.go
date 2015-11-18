package main

import "fmt"

func main() {
	fmt.Println("start")
	go doSomething()
	fmt.Println("end")
}

func doSomething() {
	fmt.Println("do something")
}
