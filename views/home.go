package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type Context struct {
		User string
	}
	c := Context{}
	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}
	c.User = session.User.FirstName + " " + session.User.LastName
	// c.User = session.User.Username
	uadmin.RenderHTML(w, r, "templates/home.html", c)
	return
}
