package views

import (
	"net/http"

	"github.com/kurt/company-time-manager-system/models"
	"github.com/uadmin/uadmin"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	clockhistory := []models.ClockHistory{}
	uadmin.All(&clockhistory)

	for i := range clockhistory {
		uadmin.Preload(&clockhistory[i], "Departments")
	}
	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}
	c := map[string]interface{}{
		"ClockHistories": clockhistory,
		"Username":       session.User.FirstName + " " + session.User.LastName,
		// "Username":       session.User.Username,
	}
	uadmin.RenderHTML(w, r, "templates/home.html", c)
}
