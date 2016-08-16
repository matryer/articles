package main

import "log"

func main() {
	for i := 0; i < 20; i++ {
		func() {
			defer db()()
			if i == 10 {
				test()
			}
		}()
	}
}

func test() {
	log.Println("----------")
	defer log.Println("----------")
	log.Println("issue update request...")
	log.Println("update request failed")
}

func db() func() {
	log.Println("open database connection")
	log.Println("connecting to session...")
	return func() {
		log.Println("closing session")
		log.Println("close database connection")
	}
}
