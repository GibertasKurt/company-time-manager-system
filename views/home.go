package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	type Context struct {
		User string
	}
	c := Context{}
	c.User = session.User.Username
	uadmin.RenderHTML(w, r, "templates/home.html", c)
	return
}
