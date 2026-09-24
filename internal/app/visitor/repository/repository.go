package repository

import (
	"database/sql"

	"visitor_management_backend/internal/app/visitor/model"
	"visitor_management_backend/internal/app/visitor/port"
)

type visitorRepository struct {
	db *sql.DB
}

func NewVisitorRepository(db *sql.DB) port.VisitorRepository {
	return &visitorRepository{
		db: db,
	}
}

// ======================================================
// GET ALL VISITORS
// ======================================================

func (r *visitorRepository) GetAll() ([]model.Visitor, error) {
	query := `
	SELECT
		v.id,
		v.full_name,
		v.email,
		v.phone,
		v.company,
		v.host_employee_id,

		COALESCE(e.full_name, '') AS host_employee_name,
		COALESCE(e.department, 'General') AS department,

		COALESCE(vr.purpose, '') AS purpose,
		COALESCE(
			DATE_FORMAT(vr.visit_date, '%Y-%m-%d'),
			''
		) AS visit_date,

		CASE
			WHEN vr.status = 'requested'  THEN 'pending'
			WHEN vr.status = 'approved'   THEN 'approved'
			WHEN vr.status = 'checked_in' THEN 'checked-in'
			WHEN vr.status = 'completed'  THEN 'checked-out'
			WHEN vr.status = 'rejected'   THEN 'rejected'
			ELSE COALESCE(v.status, 'pending')
		END AS status,

		v.created_at,
		v.updated_at

	FROM visitor v

	LEFT JOIN employee e
		ON e.id = v.host_employee_id

	LEFT JOIN visitor_request vr
		ON vr.id = (
			SELECT MAX(vr2.id)
			FROM visitor_request vr2
			WHERE vr2.visitor_id = v.id
		)

	ORDER BY v.id DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	visitors := make([]model.Visitor, 0)

	for rows.Next() {
		var visitor model.Visitor
		var host sql.NullInt64

		err := rows.Scan(
			&visitor.ID,
			&visitor.FullName,
			&visitor.Email,
			&visitor.Phone,
			&visitor.Company,
			&host,

			&visitor.HostEmployeeName,
			&visitor.Department,

			&visitor.Purpose,
			&visitor.VisitDate,

			&visitor.Status,

			&visitor.CreatedAt,
			&visitor.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		if host.Valid {
			visitor.HostEmployeeID = &host.Int64
		}

		visitors = append(visitors, visitor)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return visitors, nil
}

// ======================================================
// GET VISITOR BY ID
// ======================================================

func (r *visitorRepository) GetByID(id int64) (*model.Visitor, error) {
	query := `
	SELECT
		v.id,
		v.full_name,
		v.email,
		v.phone,
		v.company,
		v.host_employee_id,

		COALESCE(e.full_name, ''),
		COALESCE(e.department, 'General'),

		COALESCE(vr.purpose, ''),
		COALESCE(
			DATE_FORMAT(vr.visit_date, '%Y-%m-%d'),
			''
		),

		CASE
			WHEN vr.status = 'requested'  THEN 'pending'
			WHEN vr.status = 'approved'   THEN 'approved'
			WHEN vr.status = 'checked_in' THEN 'checked-in'
			WHEN vr.status = 'completed'  THEN 'checked-out'
			WHEN vr.status = 'rejected'   THEN 'rejected'
			ELSE COALESCE(v.status, 'pending')
		END AS status,

		v.created_at,
		v.updated_at

	FROM visitor v

	LEFT JOIN employee e
		ON e.id = v.host_employee_id

	LEFT JOIN visitor_request vr
		ON vr.id = (
			SELECT MAX(vr2.id)
			FROM visitor_request vr2
			WHERE vr2.visitor_id = v.id
		)

	WHERE v.id = ?
	`

	var visitor model.Visitor
	var host sql.NullInt64

	err := r.db.QueryRow(query, id).Scan(
		&visitor.ID,
		&visitor.FullName,
		&visitor.Email,
		&visitor.Phone,
		&visitor.Company,
		&host,

		&visitor.HostEmployeeName,
		&visitor.Department,

		&visitor.Purpose,
		&visitor.VisitDate,

		&visitor.Status,

		&visitor.CreatedAt,
		&visitor.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if host.Valid {
		visitor.HostEmployeeID = &host.Int64
	}

	return &visitor, nil
}

// ======================================================
// CREATE VISITOR
// ======================================================

func (r *visitorRepository) Create(visitor *model.Visitor) error {
	query := `
	INSERT INTO visitor
	(
		full_name,
		email,
		phone,
		company,
		host_employee_id,
		status,
		created_at,
		updated_at
	)
	VALUES
	(
		?,?,?,?,?,
		'pending',
		CURRENT_TIMESTAMP,
		CURRENT_TIMESTAMP
	)
	`

	result, err := r.db.Exec(
		query,
		visitor.FullName,
		visitor.Email,
		visitor.Phone,
		visitor.Company,
		visitor.HostEmployeeID,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	visitor.ID = id
	visitor.Status = "pending"

	return nil
}

// ======================================================
// UPDATE VISITOR
// ======================================================

func (r *visitorRepository) Update(visitor *model.Visitor) error {
	query := `
	UPDATE visitor
	SET
		full_name = ?,
		email = ?,
		phone = ?,
		company = ?,
		host_employee_id = ?,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		visitor.FullName,
		visitor.Email,
		visitor.Phone,
		visitor.Company,
		visitor.HostEmployeeID,
		visitor.ID,
	)

	return err
}

// ======================================================
// UPDATE STATUS
// ======================================================

func (r *visitorRepository) UpdateStatus(id int64, status string) error {
	query := `
	UPDATE visitor
	SET
		status = ?,
		updated_at = CURRENT_TIMESTAMP
	WHERE id = ?
	`

	_, err := r.db.Exec(query, status, id)

	return err
}

// ======================================================
// DELETE VISITOR
// ======================================================

func (r *visitorRepository) Delete(id int64) error {
	_, err := r.db.Exec(
		`DELETE FROM visitor WHERE id = ?`,
		id,
	)

	return err
}