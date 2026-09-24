package port

import "visitor_management_backend/internal/app/visitor/model"

type VisitorService interface {
	GetAll() ([]model.Visitor, error)
	GetByID(id int64) (*model.Visitor, error)

	Create(visitor *model.Visitor) error
	Update(visitor *model.Visitor) error
	Delete(id int64) error

	// Sinkronisasi status visitor
	UpdateStatus(id int64, status string) error
}