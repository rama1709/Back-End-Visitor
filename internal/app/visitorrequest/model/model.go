package model

import "time"

type VisitorRequest struct {
	ID         int64 `json:"id"`
	VisitorID  int64 `json:"visitor_id"`
	EmployeeID int64 `json:"employee_id"`

	// ==========================
	// Visitor
	// ==========================

	VisitorName    string `json:"visitor_name"`
	VisitorCompany string `json:"visitor_company"`

	// ==========================
	// Host
	// ==========================

	EmployeeName string `json:"employee_name"`
	Department   string `json:"department"`

	// ==========================
	// Appointment
	// ==========================

	Purpose string `json:"purpose"`
	Status  string `json:"status"`

	VisitDate string `json:"visit_date"`
	VisitTime string `json:"visit_time"`

	DurationMinutes int    `json:"duration_minutes"`
	MeetingRoom     string `json:"meeting_room"`

	Notes string `json:"notes"`

	// ==========================
	// Attendance
	// ==========================

	CheckIn  *time.Time `json:"check_in,omitempty"`
	CheckOut *time.Time `json:"check_out,omitempty"`

	// ==========================
	// Timestamp
	// ==========================

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}