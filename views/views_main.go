package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	page := strings.TrimPrefix(r.URL.Path, "/")
	page = strings.TrimSuffix(page, "/")

	session := uadmin.IsAuthenticated(r)
	if session == nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return
	}

	c := map[string]interface{}{}
	uadmin.Trail(uadmin.DEBUG, "page: %v", page)
	switch page {
	case "home":
		c = HomeHandler(w, r)
		page = "home"
	default:
		page = "login"
	}

	c["Page"] = page
	uadmin.Trail(uadmin.DEBUG, page)
	uadmin.RenderHTML(w, r, fmt.Sprintf("./templates/%v.html", page), c)
}
