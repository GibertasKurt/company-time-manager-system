package views

import (
	"net/http"

	"github.com/gibertaskurt/company-time-manager-system/models"
	"github.com/uadmin/uadmin"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	c := map[string]interface{}{}

	department := []models.Departments{}
	uadmin.All(&department)
	c["Departments"] = department

	uadmin.RenderHTML(w, r, "templates/register.html", c)
	// Render the registration form

}
