package main

import (
	"github.com/kurt/company-time-manager-system/models"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Departments{},
	)
	uadmin.RegisterInlines(models.Departments{}, map[string]string{
		"uadmin.User": "EmployeesID",
	})
	uadmin.SiteName = "Company Time Management System"
	uadmin.StartServer()
}
