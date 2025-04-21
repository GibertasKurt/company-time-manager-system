package views

import (
	"net/http"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")
	LoginHandler(w, r)
}
