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
	clockhistories := []models.ClockHistory{} // For filtering
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
	// formattedTime := currentTime.Format("01/02/2006 03:04 PM") // 12 hour format
	formattedTime := currentTime.Format("01/02/2006 15:04 -0700 -0700") // 24 hour with time zone +0800
	switch clockAux {
	case "clockIn":
		fmt.Println("Case clockIn executed")
		uadmin.Filter(&clockhistories, "employee_id = ? AND clock_out IS NULL AND clock_in IS NOT NULL AND break_end IS NOT NULL", employee.ID)
		if len(clockhistories) > 0 { // Cannot clock in if already clocked in
			uadmin.Trail(uadmin.ERROR, "Clock in Failed: Cannot clock in, already clocked in or on break.")
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
			uadmin.Filter(&clockhistories, "employee_id = ? AND clock_in IS NULL", employee.ID)
			if len(clockhistories) > 0 { // Cannot clock out if not clocked in
				uadmin.Trail(uadmin.ERROR, "Clock In Failed.")
			}
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
		uadmin.Filter(&clockhistories, "employee_id = ? AND (clock_in IS NULL OR (break_start IS NOT NULL AND break_end IS NULL))", employee.ID)
		if len(clockhistories) > 0 { // Cannot clock out if not clocked in
			uadmin.Trail(uadmin.ERROR, "Clock out Failed: Not clocked in or on break.")
			return
		}
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
		uadmin.Filter(&clockhistories, "employee_id = ? AND clock_in IS NULL", employee.ID)
		if len(clockhistories) > 0 { // Cannot start break if not clocked in
			uadmin.Trail(uadmin.ERROR, "Break start Failed: not clocked in.")
			return
		}
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
		uadmin.Filter(&clockhistories, "employee_id = ? AND (clock_in IS NULL OR break_start IS NULL) AND clock_out IS NULL", employee.ID)
		if len(clockhistories) > 0 { // Cannot end break if not on break
			uadmin.Trail(uadmin.ERROR, "Break end Failed: Cannot end break, not on break or not clocked in.")
			return
		}
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
