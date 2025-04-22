package controllers

import (
	"net/http"

	"github.com/kurt/company-time-manager-system/models"
	"github.com/uadmin/uadmin"
)

type ClockHistoryViewData struct {
	ClockHistories []ClockHistoryRow
}

type ClockHistoryRow struct {
	ClockIn    string
	ClockOut   string
	BreakStart string
	BreakEnd   string
	TotalHours float64
	Department string
	Employee   string
}

func ClockHistoryHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all ClockHistory records
	var clockHistories []models.ClockHistory
	uadmin.All(&clockHistories)

	// Prepare data for the frontend
	var rows []ClockHistoryRow
	for _, record := range clockHistories {
		rows = append(rows, ClockHistoryRow{
			ClockIn:    record.ClockIn.Format("2006-01-02 15:04:05"),
			ClockOut:   record.ClockOut.Format("2006-01-02 15:04:05"),
			BreakStart: record.BreakStart.Format("2006-01-02 15:04:05"),
			BreakEnd:   record.BreakEnd.Format("2006-01-02 15:04:05"),
			TotalHours: record.TotalHours,
			Department: record.Department.Name, // Assuming `Departments` has a `Name` field
			Employee:   record.Employee,
		})
	}

	// Pass data to the template
	data := ClockHistoryViewData{
		ClockHistories: rows,
	}

	uadmin.RenderHTML(w, r, "templates/home.html", data)
}
