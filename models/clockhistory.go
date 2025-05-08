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

	uadmin.Save(c)
}

// func (c *ClockHistory) FindCurrent() {
// 	uadmin.Filter(c, "employee_id = ? AND clock_out IS NULL", c.EmployeeID)
// 	if c.ID == 0 {
// 		c.ClockIn = time.Now()
// 		c.Current = true
// 	} else {
// 		c.ClockOut = nil
// 	}
// 	uadmin.Save(c)
// }
