package service

import (
	"errors"
	"strings"

	"visitor_management_backend/internal/app/fcmtoken/model"
	"visitor_management_backend/internal/app/fcmtoken/port"
)

type fcmTokenService struct {
	repository port.FCMTokenRepository
}

func NewFCMTokenService(
	repository port.FCMTokenRepository,
) *fcmTokenService {
	return &fcmTokenService{
		repository: repository,
	}
}

// ========================================
// REGISTER / SAVE TOKEN
// ========================================

func (s *fcmTokenService) Register(
	token *model.FCMToken,
) error {

	if token == nil {
		return errors.New("fcm token tidak boleh kosong")
	}

	token.InstallationID = strings.TrimSpace(token.InstallationID)
	token.RegistrationToken = strings.TrimSpace(token.RegistrationToken)
	token.DeviceType = strings.TrimSpace(token.DeviceType)
	token.DeviceName = strings.TrimSpace(token.DeviceName)

	if token.EmployeeID <= 0 {
		return errors.New("employee_id tidak valid")
	}

	if token.RegistrationToken == "" {
		return errors.New("registration token kosong")
	}

	if token.DeviceType == "" {
		token.DeviceType = "web"
	}

	// ===================================================
	// CEK BERDASARKAN EMPLOYEE
	// SATU EMPLOYEE = SATU TOKEN AKTIF
	// ===================================================

	existingTokens, err := s.repository.GetByEmployeeID(token.EmployeeID)
	if err != nil {
		return err
	}

	// Jika sudah ada → UPDATE token lama
	if len(existingTokens) > 0 {

		token.ID = existingTokens[0].ID

		// gunakan installation id lama
		token.InstallationID = existingTokens[0].InstallationID

		return s.repository.Update(token)
	}

	// Jika belum ada → INSERT
	return s.repository.Create(token)
}

// ========================================
// GET TOKEN EMPLOYEE
// ========================================

func (s *fcmTokenService) GetByEmployeeID(
	employeeID int64,
) ([]model.FCMToken, error) {

	if employeeID <= 0 {
		return nil, errors.New("employee_id tidak valid")
	}

	return s.repository.GetByEmployeeID(employeeID)
}

// ========================================
// DELETE TOKEN
// ========================================

func (s *fcmTokenService) Delete(
	id int64,
) error {

	if id <= 0 {
		return errors.New("id token tidak valid")
	}

	return s.repository.Delete(id)
}