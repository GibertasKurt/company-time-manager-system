package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	c := map[string]interface{}{}
	session := uadmin.IsAuthenticated(r)
	if session != nil {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
	if r.Method == "POST" {
		username := r.PostFormValue("username")
		username = strings.TrimSpace(strings.ToLower(username))
		password := r.PostFormValue("password")

		session, _ := uadmin.Login(r, username, password)
		uadmin.Trail(uadmin.DEBUG, "Login: %s", username)
		if session != nil {
			http.SetCookie(w, &http.Cookie{
				Name:     "session",
				Value:    session.Key,
				Path:     "/",
				SameSite: http.SameSiteStrictMode,
			})
			http.Redirect(w, r, "/home", http.StatusSeeOther)
		} else {
			c["Err"] = "Invalid username or password"
		}
	}

	uadmin.RenderHTML(w, r, "templates/login.html", c)
}
