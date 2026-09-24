package service

import (
	"visitor_management_backend/internal/app/visitor/model"
	"visitor_management_backend/internal/app/visitor/port"
)

type visitorService struct {
	repository port.VisitorRepository
}

func NewVisitorService(repository port.VisitorRepository) port.VisitorService {
	return &visitorService{
		repository: repository,
	}
}

// =============================
// GET ALL
// =============================

func (s *visitorService) GetAll() ([]model.Visitor, error) {
	return s.repository.GetAll()
}

// =============================
// GET BY ID
// =============================

func (s *visitorService) GetByID(id int64) (*model.Visitor, error) {
	return s.repository.GetByID(id)
}

// =============================
// CREATE
// =============================

func (s *visitorService) Create(visitor *model.Visitor) error {
	return s.repository.Create(visitor)
}

// =============================
// UPDATE
// =============================

func (s *visitorService) Update(visitor *model.Visitor) error {
	return s.repository.Update(visitor)
}

// =============================
// UPDATE STATUS
// =============================

func (s *visitorService) UpdateStatus(
	id int64,
	status string,
) error {
	return s.repository.UpdateStatus(id, status)
}

// =============================
// DELETE
// =============================

func (s *visitorService) Delete(id int64) error {
	return s.repository.Delete(id)
}