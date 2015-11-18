package main

import "fmt"

func main() {
	fmt.Println("start")
	doSomething()
	fmt.Println("end")
}

func doSomething() {
	fmt.Println("do something")
}
