package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gibertaskurt/company-time-manager-system/models"
	"github.com/uadmin/uadmin"
)

func UpsBreakHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpsBreakHandler called")
	clockhistory := []models.ClockHistory{}

	// Switch case for data-value is breakStart or breakEnd
	// switch _ := _, _ {
	// case "breakStart":
	// case "breakEnd":
	// }

	results := []map[string]interface{}{}

	uadmin.AdminPage("id", false, 0, 5, &clockhistory, "")

	for _, t := range clockhistory {
		now := time.Now()
		t.BreakStart = &now
		t.BreakEnd = &now

		uadmin.Preload(&t, "Employee")
		uadmin.Preload(&t.Employee, "Departments")

		results = append(results, map[string]interface{}{
			"ID":         t.ID,
			"EmployeeID": t.Employee.ID,
			"FirstName":  t.Employee.FirstName,
			"LastName":   t.Employee.LastName,
			"Department": t.Employee.Departments.Name,
			"ClockIn":    t.ClockIn,
			"ClockOut":   t.ClockOut,
			"BreakIn":    t.BreakStart,
			"BreakOut":   t.BreakEnd,
		})
	}
	fmt.Println("results: ", results)
	uadmin.ReturnJSON(w, r, results)
}
