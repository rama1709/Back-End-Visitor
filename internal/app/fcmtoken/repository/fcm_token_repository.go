package repository

import (
	"database/sql"

	"visitor_management_backend/internal/app/fcmtoken/model"
	"visitor_management_backend/internal/app/fcmtoken/port"
)

type fcmTokenRepository struct {
	db *sql.DB
}

func NewFCMTokenRepository(db *sql.DB) port.FCMTokenRepository {
	return &fcmTokenRepository{
		db: db,
	}
}

// ========================================
// CREATE (UPSERT)
// ========================================

func (r *fcmTokenRepository) Create(token *model.FCMToken) error {

	query := `
	INSERT INTO fcm_tokens (
		employee_id,
		installation_id,
		registration_token,
		device_type,
		device_name
	)
	VALUES (?, ?, ?, ?, ?)
	ON DUPLICATE KEY UPDATE
		registration_token = VALUES(registration_token),
		device_type = VALUES(device_type),
		device_name = VALUES(device_name),
		updated_at = CURRENT_TIMESTAMP
	`

	_, err := r.db.Exec(
		query,
		token.EmployeeID,
		token.InstallationID,
		token.RegistrationToken,
		token.DeviceType,
		token.DeviceName,
	)

	return err
}

// ========================================
// GET BY INSTALLATION ID
// ========================================

func (r *fcmTokenRepository) GetByInstallationID(
	installationID string,
) (*model.FCMToken, error) {

	query := `
	SELECT
		id,
		employee_id,
		installation_id,
		registration_token,
		device_type,
		device_name,
		created_at,
		updated_at
	FROM fcm_tokens
	WHERE installation_id = ?
	LIMIT 1
	`

	var token model.FCMToken

	err := r.db.QueryRow(query, installationID).Scan(
		&token.ID,
		&token.EmployeeID,
		&token.InstallationID,
		&token.RegistrationToken,
		&token.DeviceType,
		&token.DeviceName,
		&token.CreatedAt,
		&token.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &token, nil
}

// ========================================
// GET BY EMPLOYEE ID
// ========================================

func (r *fcmTokenRepository) GetByEmployeeID(
	employeeID int64,
) ([]model.FCMToken, error) {

	query := `
	SELECT
		id,
		employee_id,
		installation_id,
		registration_token,
		device_type,
		device_name,
		created_at,
		updated_at
	FROM fcm_tokens
	WHERE employee_id = ?
	ORDER BY updated_at DESC
	`

	rows, err := r.db.Query(query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []model.FCMToken

	for rows.Next() {
		var token model.FCMToken

		err := rows.Scan(
			&token.ID,
			&token.EmployeeID,
			&token.InstallationID,
			&token.RegistrationToken,
			&token.DeviceType,
			&token.DeviceName,
			&token.CreatedAt,
			&token.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

// ========================================
// UPDATE
// ========================================

func (r *fcmTokenRepository) Update(
	token *model.FCMToken,
) error {

	query := `
	UPDATE fcm_tokens
	SET
		employee_id = ?,
		registration_token = ?,
		device_type = ?,
		device_name = ?,
		updated_at = CURRENT_TIMESTAMP
	WHERE installation_id = ?
	`

	_, err := r.db.Exec(
		query,
		token.EmployeeID,
		token.RegistrationToken,
		token.DeviceType,
		token.DeviceName,
		token.InstallationID,
	)

	return err
}

// ========================================
// DELETE
// ========================================

func (r *fcmTokenRepository) Delete(id int64) error {

	query := `
	DELETE FROM fcm_tokens
	WHERE id = ?
	`

	_, err := r.db.Exec(query, id)

	return err
}