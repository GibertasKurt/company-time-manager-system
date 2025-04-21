package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type Departments struct {
	uadmin.Model
	Name          string
	Description   string
	Employees     []uadmin.User `uadmin:"list_exclude" gorm:"many2many:-"`
	EmployeesList string        `uadmin:"read_only"`
	EmployeesID   uint
}

func (i *Departments) Save() {
	employeesList := []string{}
	for c := range i.Employees {
		employeesList = append(employeesList, i.Employees[c].Username)
	}
	joinList := strings.Join(employeesList, ", ")
	// fmt.Println("EmployeesList: ", joinList)
	i.EmployeesList = joinList
	uadmin.Save(i)
}
