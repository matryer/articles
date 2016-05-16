package commentsapp

import (
	"io"
	"net/http"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Our app is serving requests")
	})
}
