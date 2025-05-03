package models

import "github.com/uadmin/uadmin"

type Employee struct {
	uadmin.Model
	User          uadmin.User
	UserID        uint
	FirstName     string
	LastName      string
	Email         string
	Departments   Departments
	DepartmentsID uint
	IsActive      bool
}

func (e *Employee) String() string {
	return e.FirstName + " " + e.LastName
}
