package model

type Appointment struct {
	ID                int64  `json:"id"`
	VisitorID         int64  `json:"visitor_id"`
	VisitorName       string `json:"visitor_name"`
	VisitorCompany    string `json:"visitor_company"`
	EmployeeID        int64  `json:"employee_id"`
	EmployeeName      string `json:"employee_name"`
	Department        string `json:"department"`
	Purpose           string `json:"purpose"`
	Status            string `json:"status"`
	VisitDate         string `json:"visit_date"`
	CheckIn          string `json:"check_in,omitempty"`
	CheckOut         string `json:"check_out,omitempty"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}