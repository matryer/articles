package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

}

func Save(path string, vs ...interface{}) error {
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil {
		return err
	}
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	encoder := json.NewEncoder(f)
	for i, v := range vs {
		if err := encoder.Encode(v); err != nil {
			return &errSave{i, err}
		}
	}
	return nil
}

type errSave struct {
	i   int
	err error
}

func (e *errSave) Error() string {
	return fmt.Sprintf("save: object %d: %s", e.i, e.err)
}

type errWithContext struct {
	err error
	msg string
}

func (e *errWithContext) Error() string {
	return e.msg + ": " + e.err.Error()
}
