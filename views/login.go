package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	type Context struct {
		Err       string
		ErrExists bool
		Username  string
		Password  string
	}
	c := Context{}
	uadmin.RenderHTML(w, r, "templates/login.html", c)
}
