package main

import (
	"net/http"

	"github.com/kurt/company-time-manager-system/models"
	"github.com/kurt/company-time-manager-system/views"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Database = &uadmin.DBSettings{
		Type:     "mysql",
		Name:     "company_time_manager_system",
		User:     "root",
		Password: "root",
		Host:     "localhost",
		Port:     3306,
	}
	uadmin.Register(
		models.Departments{},
		models.ClockHistory{},
		models.Employee{},
	)
	uadmin.SiteName = "Company Time Management System"
	uadmin.RootURL = "/admin/"
	http.HandleFunc("/", uadmin.Handler(views.MainHandler))
	http.HandleFunc("/login/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))
	http.HandleFunc("/register/", uadmin.Handler(views.LogoutHandler))
	uadmin.StartServer()
}
