package repository

import (
	"database/sql"
	"time"

	"visitor_management_backend/internal/app/visitorrequest/model"
	"visitor_management_backend/internal/app/visitorrequest/port"
)

type visitorRequestRepository struct {
	db *sql.DB
}

func NewVisitorRequestRepository(db *sql.DB) port.VisitorRequestRepository {
	return &visitorRequestRepository{
		db: db,
	}
}

// ======================================================
// GET ALL
// ======================================================

func (r *visitorRequestRepository) GetAll() ([]model.VisitorRequest, error) {
	query := `
	SELECT
		vr.id,
		vr.visitor_id,
		vr.employee_id,

		v.full_name,
		COALESCE(v.company, ''),

		e.full_name,
		COALESCE(e.department, ''),

		vr.purpose,
		vr.status,

		DATE_FORMAT(vr.visit_date, '%Y-%m-%d'),
		COALESCE(TIME_FORMAT(vr.visit_time, '%H:%i'), ''),

		COALESCE(vr.duration_minutes, 30),
		COALESCE(vr.meeting_room, ''),

		COALESCE(vr.notes, ''),

		vr.check_in,
		vr.check_out,

		vr.created_at,
		vr.updated_at

	FROM visitor_request vr

	INNER JOIN visitor v
		ON vr.visitor_id = v.id

	INNER JOIN employee e
		ON vr.employee_id = e.id

	ORDER BY vr.created_at DESC, vr.id DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	requests := make([]model.VisitorRequest, 0)

	for rows.Next() {
		var req model.VisitorRequest

		err := rows.Scan(
			&req.ID,
			&req.VisitorID,
			&req.EmployeeID,

			&req.VisitorName,
			&req.VisitorCompany,

			&req.EmployeeName,
			&req.Department,

			&req.Purpose,
			&req.Status,

			&req.VisitDate,
			&req.VisitTime,

			&req.DurationMinutes,
			&req.MeetingRoom,

			&req.Notes,

			&req.CheckIn,
			&req.CheckOut,

			&req.CreatedAt,
			&req.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		requests = append(requests, req)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return requests, nil
}

// ======================================================
// GET BY ID
// ======================================================

func (r *visitorRequestRepository) GetByID(id int64) (*model.VisitorRequest, error) {
	query := `
	SELECT
		vr.id,
		vr.visitor_id,
		vr.employee_id,

		v.full_name,
		COALESCE(v.company, ''),

		e.full_name,
		COALESCE(e.department, ''),

		vr.purpose,
		vr.status,

		DATE_FORMAT(vr.visit_date, '%Y-%m-%d'),
		COALESCE(TIME_FORMAT(vr.visit_time, '%H:%i'), ''),

		COALESCE(vr.duration_minutes, 30),
		COALESCE(vr.meeting_room, ''),

		COALESCE(vr.notes, ''),

		vr.check_in,
		vr.check_out,

		vr.created_at,
		vr.updated_at

	FROM visitor_request vr

	INNER JOIN visitor v
		ON vr.visitor_id = v.id

	INNER JOIN employee e
		ON vr.employee_id = e.id

	WHERE vr.id = ?
	`

	var req model.VisitorRequest

	err := r.db.QueryRow(query, id).Scan(
		&req.ID,
		&req.VisitorID,
		&req.EmployeeID,

		&req.VisitorName,
		&req.VisitorCompany,

		&req.EmployeeName,
		&req.Department,

		&req.Purpose,
		&req.Status,

		&req.VisitDate,
		&req.VisitTime,

		&req.DurationMinutes,
		&req.MeetingRoom,

		&req.Notes,

		&req.CheckIn,
		&req.CheckOut,

		&req.CreatedAt,
		&req.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &req, nil
}

// ======================================================
// CREATE
// ======================================================

func (r *visitorRequestRepository) Create(req *model.VisitorRequest) error {
	query := `
	INSERT INTO visitor_request
	(
		visitor_id,
		employee_id,
		purpose,
		status,
		visit_date,
		visit_time,
		duration_minutes,
		meeting_room,
		notes,
		created_at,
		updated_at
	)
	VALUES
	(
		?,?,?,?,?,?,?,?,?,
		CURRENT_TIMESTAMP,
		CURRENT_TIMESTAMP
	)
	`

	result, err := r.db.Exec(
		query,
		req.VisitorID,
		req.EmployeeID,
		req.Purpose,
		"requested",
		req.VisitDate,
		req.VisitTime,
		req.DurationMinutes,
		req.MeetingRoom,
		req.Notes,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	req.ID = id
	req.Status = "requested"

	return nil
}

// ======================================================
// UPDATE
// ======================================================

func (r *visitorRequestRepository) Update(req *model.VisitorRequest) error {
	query := `
	UPDATE visitor_request
	SET
		employee_id = ?,
		purpose = ?,
		status = ?,
		visit_date = ?,
		visit_time = ?,
		duration_minutes = ?,
		meeting_room = ?,
		notes = ?,
		check_in = ?,
		check_out = ?,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		req.EmployeeID,
		req.Purpose,
		req.Status,
		req.VisitDate,
		req.VisitTime,
		req.DurationMinutes,
		req.MeetingRoom,
		req.Notes,
		req.CheckIn,
		req.CheckOut,
		req.ID,
	)

	return err
}

// ======================================================
// APPROVE
// ======================================================

func (r *visitorRequestRepository) Approve(id int64) error {
	_, err := r.db.Exec(`
		UPDATE visitor_request
		SET
			status = 'approved',
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`, id)

	return err
}

// ======================================================
// REJECT
// ======================================================

func (r *visitorRequestRepository) Reject(id int64) error {
	_, err := r.db.Exec(`
		UPDATE visitor_request
		SET
			status = 'rejected',
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`, id)

	return err
}

// ======================================================
// CHECK IN
// ======================================================

func (r *visitorRequestRepository) CheckIn(id int64) error {
	now := time.Now()

	_, err := r.db.Exec(`
		UPDATE visitor_request
		SET
			status = 'checked_in',
			check_in = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`, now, id)

	return err
}

// ======================================================
// CHECK OUT
// ======================================================

func (r *visitorRequestRepository) CheckOut(id int64) error {
	now := time.Now()

	_, err := r.db.Exec(`
		UPDATE visitor_request
		SET
			status = 'completed',
			check_out = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`, now, id)

	return err
}

// ======================================================
// DELETE
// ======================================================

func (r *visitorRequestRepository) Delete(id int64) error {
	_, err := r.db.Exec(`
		DELETE FROM visitor_request
		WHERE id = ?
	`, id)

	return err
}