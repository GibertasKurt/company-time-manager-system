package views

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gibertaskurt/company-time-manager-system/models"
	"github.com/uadmin/uadmin"
)

func formatTime(t time.Time) string {
	return t.Format("15:04") // 24-hour format
}

func HomeHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	clockhistory := []models.ClockHistory{}
	employee := models.Employee{}
	session := uadmin.IsAuthenticated(r)
	recentLogin := models.ClockHistory{}
	fmt.Println("Recent Login:", recentLogin)
	uadmin.Get(&employee, "user_id = ?", session.UserID)
	if employee.ID == 0 { // Filters table by employee ID
		uadmin.All(&clockhistory)
	} else {
		uadmin.Filter(&clockhistory, "employee_id = ?", employee.ID)
	}
	if session == nil {
		http.Redirect(w, r, "/login/", http.StatusSeeOther)
		return nil
	}
	isAdmin := session.User.Admin

	for i := range clockhistory {
		uadmin.Preload(&clockhistory[i], "Employee")
		uadmin.Preload(&clockhistory[i].Employee, "Departments")
	}

	uadmin.Get(&recentLogin, "employee_id = ? AND clock_out IS NULL", employee.ID)
	c := map[string]interface{}{
		"IsAdmin":        isAdmin,
		"ClockHistories": clockhistory,
		"formatTime":     formatTime,
		"Username":       session.User.FirstName + " " + session.User.LastName,
		"EmployeeID":     employee.ID,
		"Recent":         recentLogin.ID,
	}
	return c
}
