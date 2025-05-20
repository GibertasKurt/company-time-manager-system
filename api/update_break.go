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
	clockhistory := models.ClockHistory{}
	employee := models.Employee{}
	session := uadmin.IsAuthenticated(r)
	var breakAux = r.FormValue("data_value")

	uadmin.Get(&employee, "user_id = ?", session.UserID)
	if employee.ID == 0 {
		fmt.Println("Invalid! EmployeeID is ", employee.ID)
		return
	} else {
		fmt.Println("Session UserID: ", session.UserID)
		fmt.Println("Employee UserID: ", employee.UserID)
		uadmin.Get(&clockhistory, "employee_id = ? AND clock_out IS NULL AND break_start IS NULL", session.ID)
	}
	if employee.UserID != session.UserID {
		fmt.Println("Session and Employee UserID are the NOT same! ")
		fmt.Println("Session UserID: ", session.UserID)
		fmt.Println("Employee UserID: ", employee.UserID)
		return
	}

	currentTime := time.Now()
	// formattedTime := currentTime.Format("01/02/2006 03:04 PM") // 12 hour format
	// formattedTime := currentTime.Format("01/02/2006 15:04 -0700 -0700") // 24 hour with time zone +0800
	// formattedTime := currentTime.Format("01/02/2006 15:04:05") // 24 hour with seconds
	// formattedTime := currentTime.Format("01/02/2006 15:04") // 24 hour, no seconds
	formattedTime := currentTime.Format("01/02/2006 15:04 -0700 -0700") // 24 hour with time zone +0800
	switch breakAux {
	case "breakStart":

		fmt.Println("Case breakStart executed")
		fmt.Println("Current Date and Time: ", formattedTime)
		clockhistory := []models.ClockHistory{}
		uadmin.AdminPage("id", false, 0, 1, &clockhistory, "employee_id = ?", employee.ID)

		for _, t := range clockhistory {
			t.BreakStart = &currentTime
			uadmin.Save(&t)
			uadmin.Trail(uadmin.DEBUG, "Break Start - ID: %d, Break Start Time: %v\n", t.ID, t.BreakStart)
		}
	case "breakEnd":

		fmt.Println("Case breakEnd executed")
		fmt.Println("Current Date and Time: ", formattedTime)
		clockhistory := []models.ClockHistory{}
		uadmin.AdminPage("id", false, 0, 1, &clockhistory, "employee_id = ?", employee.ID)

		for _, t := range clockhistory {
			t.BreakEnd = &currentTime
			uadmin.Save(&t)
			uadmin.Trail(uadmin.DEBUG, "Break Start - ID: %d, Break Start Time: %v\n", t.ID, t.BreakStart)
		}
	default:
		fmt.Println("Error: Invalid data-value! Thrown: ", breakAux)
		return
	}
}
