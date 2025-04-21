package views

import (
	"net/http"
	"strings"

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
	if r.Method == "POST" {
		username := r.FormValue("username")
		username = strings.TrimSpace(strings.ToLower(username))
		password := r.FormValue("password")
		session := uadmin.Login2FA(r, username, password, "")

		if session == nil || !session.User.Active {
			c.ErrExists = true
			c.Err = "Invalid username or password or user is inactive"
		}
		cookie, _ := r.Cookie("session")
		if cookie == nil {
			cookie = &http.Cookie{}
		}
		cookie.Name = "session"
		cookie.Value = session.Key
		cookie.Path = "/"
		cookie.SameSite = http.SameSiteStrictMode
		http.SetCookie(w, cookie)

		uadmin.Trail(uadmin.DEBUG, "Username: %s", username)
		uadmin.Trail(uadmin.DEBUG, "Password: %s", password)
	}
}
