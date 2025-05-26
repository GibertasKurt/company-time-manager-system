package models

import (
	"math"
	"time"

	"github.com/uadmin/uadmin"
)

type ClockHistory struct {
	uadmin.Model
	Employee   Employee
	EmployeeID uint
	ClockIn    time.Time
	ClockOut   *time.Time
	BreakStart *time.Time
	BreakEnd   *time.Time
	TotalHours float64 `uadmin:"readonly"`
	Current    bool
}

func (c *ClockHistory) String() string {
	uadmin.Preload(c)
	return c.Employee.FirstName + " " + c.Employee.LastName
}

func (c *ClockHistory) Save() {
	if c.ClockOut != nil { // breakStart - breakEnd
		duration := c.ClockOut.Sub(c.ClockIn).Hours()
		c.TotalHours = math.Round(duration*100) / 100
	}
	uadmin.Save(c)
}
