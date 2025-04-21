package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/login/")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

	session := uadmin.IsAuthenticated(r)
	if session == nil {
		LoginHandler(w, r)
		return
	}

	HomeHandler(w, r, session)
	return
}
