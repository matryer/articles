package decoding

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cheekybits/is"
)

func TestHandleCreateUser(t *testing.T) {
	is := is.New(t)

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/user", strings.NewReader(`{
		"email":"something@me.com"
	}`))
	is.NoErr(err)

	handleCreateUser(w, r)
	is.Equal(w.Code, http.StatusCreated)

}
