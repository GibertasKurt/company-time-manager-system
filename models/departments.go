package models

import (
	"github.com/uadmin/uadmin"
)

type Departments struct {
	uadmin.Model
	Name        string
	Description string
}

func (d *Departments) String() string {
	return d.Name
}
