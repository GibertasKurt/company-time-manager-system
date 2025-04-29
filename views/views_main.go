package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) { //needed prefix initialized from main.go
	page := strings.TrimPrefix(r.URL.Path, "/")

	session := uadmin.IsAuthenticated(r)
	c := map[string]interface{}{}

	if session == nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
	}
	fmt.Println(uadmin.DEBUG, "page: ", page)
	switch page {
	case "home": //Name of HTML
		c = HomeHandler(w, r)
	default:
		page = "ascz"
	}
	//uadmin.Trail(uadmin.DEBUG, page)
	c["Page"] = page
	uadmin.Trail(uadmin.DEBUG, page)
	uadmin.RenderHTML(w, r, fmt.Sprintf("./templates/%v.html", page), c)
	// Rendering(w, r, page, c)

}

func Rendering(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	path := "./templates/" + page + ".html"
	uadmin.RenderHTML(w, r, path, context)
}
