package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type ClockHistory struct {
	uadmin.Model
	ClockIn       time.Time
	ClockOut      time.Time
	BreakStart    time.Time
	BreakEnd      time.Time
	TotalHours    float64
	DepartmentsID uint
	Departments   Departments
	EmployeeID    uint
	Employee      string
}

func (c *ClockHistory) String() string {
	return c.Employee
}
