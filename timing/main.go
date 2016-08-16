package main

import (
	"log"
	"time"
)

func main() {
	FunkyFunc()
}

func StartTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		d := time.Now().Sub(t)
		log.Println(name, "took", d)
	}
}

func FunkyFunc() {
	stop := StartTimer("FunkyFunc")
	defer stop()

	time.Sleep(1 * time.Second)
}

type Item interface {
	Name() string
	Ext() string
}

type MD5 interface {
	MD5() string
}

func Filename(i Item) string {
	if m, ok := i.(MD5); ok {
		return m.MD5() + "." + i.Ext()
	}
	return i.Name() + "." + i.Ext()
}
