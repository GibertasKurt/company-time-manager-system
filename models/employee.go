package models

import "github.com/uadmin/uadmin"

type Employee struct {
	uadmin.Model
	User          uadmin.User
	UserID        uint
	Name          string
	Departments   Departments
	DepartmentsID uint
	IsActive      bool
}
