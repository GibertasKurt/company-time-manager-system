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
	employee := models.Employee{}
	session := uadmin.IsAuthenticated(r)
	var breakAux = r.FormValue("data_value")

	currentTime := time.Now()
	// formattedTime := currentTime.Format("01/02/2006 03:04 PM") 12 hour format
	// formattedTime := currentTime.Format("01/02/2006 15:04") 24 hour, no seconds
	formattedTime := currentTime.Format("01/02/2006 15:04:05")
	switch breakAux {
	case "breakStart":
		fmt.Println("Case breakStart executed")
		fmt.Println("Current Date and Time: ", formattedTime)
	case "breakEnd":
		fmt.Println("Case breakEnd executed")
		fmt.Println("Current Date and Time: ", formattedTime)
	default:
		fmt.Println("Error: Invalid data-value! Thrown: ", breakAux)
		return
	}

	uadmin.Get(&employee, "user_id = ?", session.UserID)

	if employee.ID == 0 {
		fmt.Println("Invalid! EmployeeID is ", employee.ID)
		return
	} else {
		uadmin.Filter(&clockhistory, "clock_out IS NULL AND break_start IS NULL")
		// fmt.Println("Clock history filetered: ", clockhistory)
	}
	if employee.UserID == session.UserID {
		fmt.Println("EUREKA! EXCELSIOR!")
	} else {
		fmt.Println("Session UserID: ", session.UserID)
		fmt.Println("Employee ID: ", employee.UserID)
	}

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
	// fmt.Println("results: ", results)
	uadmin.ReturnJSON(w, r, results)
}
