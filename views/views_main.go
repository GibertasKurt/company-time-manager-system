package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {

	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}

	uadmin.RenderHTML(w, r, "templates/home.html", nil)
}
