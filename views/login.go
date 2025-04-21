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
	if r.Method == http.MethodPost {
		c.Username = r.FormValue("username")
		c.Username = strings.TrimSpace(strings.ToLower(c.Username))
		c.Password = r.FormValue("password")
		session := uadmin.Login2FA(r, c.Username, c.Password, "")
		if session == nil || !session.User.Active {
			c.ErrExists = true
			c.Err = "Invalid username or password or inactive user"
		} else {
			c.Err = "Username and password is invalid"
		}
		c.Err = "Invalid username or password"
		c.ErrExists = true
		// Testing for inputs
		uadmin.Trail(uadmin.DEBUG, "Username: %s", c.Username)
		uadmin.Trail(uadmin.DEBUG, "Password: %s", c.Password)
		uadmin.Trail(uadmin.DEBUG, "Session: %s", session)
		return
	}
}
