package views

import (
	"net/http"

	"github.com/uadmin/uadmin"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request, session *uadmin.Session) {
	uadmin.Logout(r)
	for _, cookie := range r.Cookies() {
		c := &http.Cookie{
			Name:   cookie.Name,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(w, c)
	}
	http.Redirect(w, r, "/login/", http.StatusSeeOther)
}
