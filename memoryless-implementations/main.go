package main

import (
	"encoding/json"

	"io"
)

func main() {

}

type Codec interface {
	Encode(w io.Writer, v interface{}) error
	Decode(r io.Reader, v interface{}) error
}

type jsonCodec struct{}

func (jsonCodec) Encode(w io.Writer, v interface{}) error {
	return json.NewEncoder(w).Encode(v)
}

func (jsonCodec) Decode(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

var JSON Codec = jsonCodec{}
