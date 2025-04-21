package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	type Context struct {
		Err         string
		ErrExists   bool
		OTPRequired bool
		Username    string
		Password    string
	}
	c := Context{}

	if r.Method == "POST" {
		username := r.PostFormValue("username")
		username = strings.TrimSpace(strings.ToLower(username))
		password := r.PostFormValue("password")

		session, _ := uadmin.Login(r, username, password)
		uadmin.Trail(uadmin.DEBUG, "Login: %s", username)
		if session == nil || !session.User.Active {
			c.ErrExists = true
			c.Err = "Invalid username/password or inactive user"
		} else {
			cookie, _ := r.Cookie("session")
			if cookie == nil {
				cookie = &http.Cookie{}
			}
			cookie.Name = "session"
			cookie.Value = session.Key
			cookie.Path = "/"
			cookie.SameSite = http.SameSiteStrictMode
			http.SetCookie(w, cookie)
			uadmin.Trail(uadmin.DEBUG, "Your login credentials are valid.")
			http.Redirect(w, r, "/home/", http.StatusSeeOther)
		}
	}

	uadmin.RenderHTML(w, r, "templates/login.html", c)
}
