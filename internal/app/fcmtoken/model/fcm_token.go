package model

import "time"

type FCMToken struct {
	ID                int64     `json:"id"`
	EmployeeID        int64     `json:"employee_id"`
	InstallationID    string    `json:"installation_id"`
	RegistrationToken string    `json:"registration_token"`
	DeviceType        string    `json:"device_type"`
	DeviceName        string    `json:"device_name"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
