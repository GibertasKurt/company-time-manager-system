package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type ClockHistory struct {
	uadmin.Model
	Employee   Employee
	EmployeeID uint
	ClockIn    time.Time
	ClockOut   time.Time
	BreakStart time.Time
	BreakEnd   time.Time
	TotalHours float64 `uadmin:"readonly"`
}

func (c *ClockHistory) String() string {
	uadmin.Preload(c, "Employee")
	return c.Employee.Name
}

func (c *ClockHistory) Save() {
	totalWork := c.ClockOut.Sub(c.ClockIn)

	breakDuration := c.BreakEnd.Sub(c.BreakStart)

	netWorkDuration := totalWork - breakDuration

	c.TotalHours = netWorkDuration.Hours()

	uadmin.Save(c)
}
