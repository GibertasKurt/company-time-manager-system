package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gibertaskurt/company-time-manager-system/models"
	"github.com/uadmin/uadmin"
)

func ClockOutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ClockOutHandler called")
	clockhistory := models.ClockHistory{}
	employee := models.Employee{}
	session := uadmin.IsAuthenticated(r)
	var clockOutAux = r.FormValue("data_value")

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
	formattedTime := currentTime.Format("01/02/2006 15:04 -0700 -0700") // 24 hour with time zone +0800
	switch clockOutAux {
	case "clockOut":
		fmt.Println("Case clockOut executed")
		fmt.Println("Current Date and Time: ", formattedTime)
		clockhistory := []models.ClockHistory{}
		uadmin.AdminPage("id", false, 0, 1, &clockhistory, "employee_id = ?", employee.ID)

		for _, t := range clockhistory {
			t.ClockOut = &currentTime
			// uadmin.Save(&t)
			t.Save()
			uadmin.Trail(uadmin.DEBUG, "Clock Out - ID: %d, Clock Out Time: %v\n", t.ID, t.ClockOut)
		}
	default:
		fmt.Println("Invalid case for clock out")
	}

}
