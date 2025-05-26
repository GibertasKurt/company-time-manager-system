package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gibertaskurt/company-time-manager-system/models"
	"github.com/uadmin/uadmin"
)

func UpdateClockAPIHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateClockHandler called")
	clockhistory := models.ClockHistory{}
	employee := models.Employee{}
	session := uadmin.IsAuthenticated(r)
	var clockAux = r.FormValue("data_value")

	uadmin.Get(&employee, "user_id = ?", session.UserID)
	if employee.ID == 0 {
		fmt.Println("Invalid! EmployeeID is ", employee.ID)
		return
	} else {
		fmt.Println("Session UserID: ", session.UserID)
		fmt.Println("Employee UserID: ", employee.UserID)
		uadmin.Get(&clockhistory, "employee_id = ? AND clock_out IS NULL", session.ID)
	}
	if employee.UserID != session.UserID {
		fmt.Println("Session and Employee UserID are the NOT same! ")
		fmt.Println("Session UserID: ", session.UserID)
		fmt.Println("Employee UserID: ", employee.UserID)
		return
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("01/02/2006 03:04 PM") // 12 hour format
	// formattedTime := currentTime.Format("01/02/2006 15:04 -0700 -0700") // 24 hour with time zone +0800
	switch clockAux {
	case "clockIn":
		fmt.Println("Case clockIn executed")
		clockhistories := []models.ClockHistory{}
		uadmin.Filter(&clockhistories, "employee_id = ? AND clock_out IS NULL", employee.ID)
		if len(clockhistories) > 0 {
			uadmin.Trail(uadmin.DEBUG, "You are already clocked in.")
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"message": "You are already clocked in",
			})
			return
		}
		fmt.Println("Current Date and Time: ", formattedTime)

		clockhistory.ClockIn = currentTime
		clockhistory.EmployeeID = employee.ID
		clockhistory.BreakStart = nil
		clockhistory.BreakEnd = nil
		clockhistory.ClockOut = nil
		err := uadmin.Save(&clockhistory)
		if err != nil {
			uadmin.ReturnJSON(w, r, map[string]interface{}{
				"status":  "error",
				"message": "Failed to save clock history",
			})
			return
		}
		uadmin.Trail(uadmin.DEBUG, "Clock In - ID: %d, Clock In Time: %v\n", clockhistory.ID, clockhistory.ClockIn)
		uadmin.ReturnJSON(w, r, map[string]interface{}{
			"status":  "success",
			"message": "Clock In successful",
			"clockIn": clockhistory.ClockIn,
		})
	case "clockOut":
		fmt.Println("Case clockOut executed")
		fmt.Println("Current Date and Time: ", formattedTime)
		clockhistory := []models.ClockHistory{}
		uadmin.AdminPage("id", false, 0, 1, &clockhistory, "employee_id = ?", employee.ID)
		for _, t := range clockhistory {
			t.ClockOut = &currentTime
			t.Save()
			uadmin.Trail(uadmin.DEBUG, "Clock Out - ID: %d, Clock Out Time: %v\n", t.ID, t.ClockOut)
		}
	case "breakStart":
		fmt.Println("Case breakStart executed")
		fmt.Println("Current Date and Time: ", formattedTime)
		clockhistory := []models.ClockHistory{}
		uadmin.AdminPage("id", false, 0, 1, &clockhistory, "employee_id = ?", employee.ID)
		for _, t := range clockhistory {
			t.BreakStart = &currentTime
			t.Save()
			uadmin.Trail(uadmin.DEBUG, "Break Start - ID: %d, Break Start Time: %v\n", t.ID, t.BreakStart)
		}
	case "breakEnd":
		fmt.Println("Case breakEnd executed")
		fmt.Println("Current Date and Time: ", formattedTime)
		clockhistory := []models.ClockHistory{}
		uadmin.AdminPage("id", false, 0, 1, &clockhistory, "employee_id = ?", employee.ID)
		for _, t := range clockhistory {
			t.BreakEnd = &currentTime
			t.Save()
			uadmin.Trail(uadmin.DEBUG, "Break End - ID: %d, Break End Time: %v\n", t.ID, t.BreakStart)
		}
	default:
		fmt.Println("Error: Invalid data-value! Thrown: ", clockAux)
		return
	}
}
