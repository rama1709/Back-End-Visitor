package service

import (
	"fmt"
	"log"
	"time"

	firebaseConfig "visitor_management_backend/firebase"

	visitorPort "visitor_management_backend/internal/app/visitor/port"
	"visitor_management_backend/internal/app/visitorrequest/model"
	"visitor_management_backend/internal/app/visitorrequest/port"
)

type visitorRequestService struct {
	repository  port.VisitorRequestRepository
	visitorRepo visitorPort.VisitorRepository
	firebase    *firebaseConfig.Firebase
}

func NewVisitorRequestService(
	repository port.VisitorRequestRepository,
	visitorRepo visitorPort.VisitorRepository,
	firebaseClient *firebaseConfig.Firebase,
) port.VisitorRequestService {
	return &visitorRequestService{
		repository:  repository,
		visitorRepo: visitorRepo,
		firebase:    firebaseClient,
	}
}

// ======================================================
// GET ALL
// ======================================================

func (s *visitorRequestService) GetAll() ([]model.VisitorRequest, error) {
	return s.repository.GetAll()
}

// ======================================================
// GET BY ID
// ======================================================

func (s *visitorRequestService) GetByID(id int64) (*model.VisitorRequest, error) {
	return s.repository.GetByID(id)
}

// ======================================================
// CREATE APPOINTMENT
// ======================================================

func (s *visitorRequestService) Create(request *model.VisitorRequest) error {
	request.Status = "requested"

	if err := s.repository.Create(request); err != nil {
		return err
	}

	log.Printf("Appointment %d created", request.ID)

	if s.firebase != nil {
		data, _ := s.repository.GetByID(request.ID)

		if data != nil {
			_ = s.firebase.CreateNotification(
				firebaseConfig.Notification{
					Type:  "appointment",
					Title: "New Appointment",
					Message: fmt.Sprintf(
						"%s memiliki appointment dengan %s pada %s %s",
						data.VisitorName,
						data.EmployeeName,
						data.VisitDate,
						data.VisitTime,
					),
					Read:      false,
					CreatedAt: time.Now(),
				},
			)
		}
	}

	return nil
}

// ======================================================
// UPDATE APPOINTMENT
// ======================================================

func (s *visitorRequestService) Update(request *model.VisitorRequest) error {

	if err := s.repository.Update(request); err != nil {
		return err
	}

	// Sinkron status ke tabel visitor
	visitorStatus := "pending"

	switch request.Status {
	case "requested":
		visitorStatus = "pending"

	case "approved":
		visitorStatus = "approved"

	case "checked_in":
		visitorStatus = "checked_in"

	case "completed":
		visitorStatus = "checked_out"

	case "rejected":
		visitorStatus = "rejected"
	}

	if err := s.visitorRepo.UpdateStatus(
		request.VisitorID,
		visitorStatus,
	); err != nil {
		log.Println("Sync visitor status:", err)
	}

	// Firebase Notification
	if s.firebase != nil {

		title := ""
		message := ""

		switch request.Status {

		case "approved":
			title = "Appointment Approved"
			message = "Visitor telah disetujui"

		case "rejected":
			title = "Appointment Rejected"
			message = "Visitor ditolak"

		case "checked_in":
			title = "Visitor Checked In"
			message = "Visitor telah check in"

		case "completed":
			title = "Visit Completed"
			message = "Visitor telah check out"
		}

		if title != "" {
			_ = s.firebase.CreateNotification(
				firebaseConfig.Notification{
					Type:      "appointment_status",
					Title:     title,
					Message:   message,
					Read:      false,
					CreatedAt: time.Now(),
				},
			)
		}
	}

	log.Printf(
		"Appointment %d updated -> %s",
		request.ID,
		request.Status,
	)

	return nil
}

// ======================================================
// APPROVE
// ======================================================

func (s *visitorRequestService) Approve(id int64) error {

	req, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}

	if req == nil {
		return fmt.Errorf("appointment not found")
	}

	req.Status = "approved"

	return s.Update(req)
}

// ======================================================
// REJECT
// ======================================================

func (s *visitorRequestService) Reject(id int64) error {

	req, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}

	if req == nil {
		return fmt.Errorf("appointment not found")
	}

	req.Status = "rejected"

	return s.Update(req)
}

// ======================================================
// CHECK IN
// ======================================================

func (s *visitorRequestService) CheckIn(id int64) error {

	req, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}

	if req == nil {
		return fmt.Errorf("appointment not found")
	}

	now := time.Now()

	req.Status = "checked_in"
	req.CheckIn = &now

	// Update appointment
	if err := s.repository.Update(req); err != nil {
		return err
	}

	// Sinkron ke visitor
	if err := s.visitorRepo.UpdateStatus(
		req.VisitorID,
		"checked_in",
	); err != nil {
		return err
	}

	return nil
}

// ======================================================
// CHECK OUT
// ======================================================

func (s *visitorRequestService) CheckOut(id int64) error {

	req, err := s.repository.GetByID(id)
	if err != nil {
		return err
	}

	if req == nil {
		return fmt.Errorf("appointment not found")
	}

	now := time.Now()

	req.Status = "completed"
	req.CheckOut = &now

	// Update appointment
	if err := s.repository.Update(req); err != nil {
		return err
	}

	// Sinkron ke visitor
	if err := s.visitorRepo.UpdateStatus(
		req.VisitorID,
		"completed",
	); err != nil {
		return err
	}

	return nil
}

// ======================================================
// DELETE
// ======================================================

func (s *visitorRequestService) Delete(id int64) error {
	return s.repository.Delete(id)
}