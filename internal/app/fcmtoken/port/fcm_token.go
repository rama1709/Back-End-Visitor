package port

import (
	"visitor_management_backend/internal/app/fcmtoken/model"
)

type FCMTokenRepository interface {
	Create(
		token *model.FCMToken,
	) error

	GetByInstallationID(
		installationID string,
	) (*model.FCMToken, error)

	GetByEmployeeID(
		employeeID int64,
	) ([]model.FCMToken, error)

	Update(
		token *model.FCMToken,
	) error

	Delete(
		id int64,
	) error
}

type FCMTokenService interface {
	Register(
		token *model.FCMToken,
	) error

	GetByEmployeeID(
		employeeID int64,
	) ([]model.FCMToken, error)

	Delete(
		id int64,
	) error
}
