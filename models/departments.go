package models

import "github.com/uadmin/uadmin"

type Departments struct {
	uadmin.Model
	Name        string
	Description string
	Employees   uadmin.User
	EmployeesID uint
}
