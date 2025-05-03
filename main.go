package main

import (
	"net/http"

	"github.com/gibertaskurt/company-time-manager-system/api"
	"github.com/gibertaskurt/company-time-manager-system/models"
	"github.com/gibertaskurt/company-time-manager-system/views"
	"github.com/uadmin/uadmin"
)

func main() {

	uadmin.Register(
		models.Departments{},
		models.ClockHistory{},
		models.Employee{},
	)
	uadmin.SiteName = "Company Time Management System"
	uadmin.RootURL = "/admin/"
	http.HandleFunc("/", uadmin.Handler(views.MainHandler))
	http.HandleFunc("/api/", uadmin.Handler(api.APIHandler))
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))
	http.HandleFunc("/register/", uadmin.Handler(views.RegisterHandler))
	uadmin.StartServer()
}
