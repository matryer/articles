package main

import (
	"errors"
	"log"
)

func main() {
	var err error
	defer func() {
		if err != nil {
			log.Fatalln(err)
		}
	}()

	defer log.Println("close A")
	defer log.Println("close B")

	log.Fatalln("fail")

	defer log.Println("close C")

	err = errors.New("error")

}
