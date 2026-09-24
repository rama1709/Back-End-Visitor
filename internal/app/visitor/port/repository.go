package port

import "visitor_management_backend/internal/app/visitor/model"

type VisitorRepository interface {
	// Read
	GetAll() ([]model.Visitor, error)
	GetByID(id int64) (*model.Visitor, error)

	// CRUD
	Create(visitor *model.Visitor) error
	Update(visitor *model.Visitor) error
	Delete(id int64) error

	// Sinkronisasi status
	UpdateStatus(id int64, status string) error
}