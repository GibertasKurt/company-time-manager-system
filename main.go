package main

import (
	"net/http"

	"github.com/kurt/company-time-manager-system/models"
	"github.com/kurt/company-time-manager-system/views"
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
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))
	http.HandleFunc("/home/", uadmin.Handler(views.HomeHandler))
	// http.HandleFunc("/clockhistory/", controllers.ClockHistoryHandler)
	uadmin.StartServer()
}
