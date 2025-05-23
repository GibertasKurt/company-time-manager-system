package api

import (
	"net/http"
	"strings"
)

func APIHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "/registration") {
		RegistrationHandler(w, r)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/updateclock") {
		UpdateClockAPIHandler(w, r)
		return
	}
}
