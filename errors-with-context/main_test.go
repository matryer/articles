package main

import (
	"os"
	"testing"
)

func TestWriteManyFiles(t *testing.T) {

	if err := os.RemoveAll("./test"); err != nil {
		t.Error("remove test folder:", err)
	}

	obj0 := "a string"
	obj1 := map[string]interface{}{"an object": true}
	obj2 := struct {
		Other map[string]interface{} `json:"other"`
	}{
		Other: obj1,
	}

	if err := Save("./test/file.json", obj0, obj1, obj2); err != nil {
		t.Error("failed to save:", err)
	}

}
