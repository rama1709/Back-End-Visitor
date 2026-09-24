package model

import "time"

type Visitor struct {
	ID       int64  `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Company  string `json:"company"`

	// Host
	HostEmployeeID   *int64 `json:"host_employee_id,omitempty"`
	HostEmployeeName string `json:"host_employee_name"`
	Department       string `json:"department"`

	// Appointment (diambil dari visitor_request)
	Purpose   string `json:"purpose"`
	VisitDate string `json:"visit_date"`

	// Status Visitor
	Status string `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}