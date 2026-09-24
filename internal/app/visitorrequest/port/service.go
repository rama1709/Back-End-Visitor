package port

import "visitor_management_backend/internal/app/visitorrequest/model"

type VisitorRequestService interface {
	// Read
	GetAll() ([]model.VisitorRequest, error)
	GetByID(id int64) (*model.VisitorRequest, error)

	// CRUD
	Create(visitorRequest *model.VisitorRequest) error
	Update(visitorRequest *model.VisitorRequest) error
	Delete(id int64) error

	// Workflow Appointment
	Approve(id int64) error
	Reject(id int64) error

	// Attendance
	CheckIn(id int64) error
	CheckOut(id int64) error
}