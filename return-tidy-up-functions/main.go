package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func Setup() (*os.File, func(), error) {
	teardown := func() {}
	f, err := ioutil.TempFile(os.TempDir(), "test")
	if err != nil {
		return nil, teardown, err
	}
	teardown = func() {
		f.Close()
		os.RemoveAll(f.Name())
	}
	return f, teardown, nil
}

func TestSomething(t *testing.T) {
	f, teardown, err := Setup()
	defer teardown()
	if err != nil {
		t.Error(err)
	}
	log.Println("TODO: do something with", f.Name())
}

func main() {
	f, teardown, err := Setup()
	if err != nil {
		log.Fatalln(err)
	}
	defer teardown()
	log.Println("TODO: do something with", f.Name())
}
