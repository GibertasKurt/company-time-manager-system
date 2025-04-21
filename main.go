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
	)
	uadmin.SiteName = "Company Time Management System"
	uadmin.StartServer()
	http.HandleFunc("/login/", uadmin.Handler(views.MainHandler))

}
