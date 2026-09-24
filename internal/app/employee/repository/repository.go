package repository

import (
	"database/sql"

	"visitor_management_backend/internal/app/employee/model"
	"visitor_management_backend/internal/app/employee/port"
)

type employeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(
	db *sql.DB,
) port.EmployeeRepository {
	return &employeeRepository{
		db: db,
	}
}

// ======================================================
// GET ALL
// ======================================================

func (r *employeeRepository) GetAll() ([]model.Employee, error) {
	query := `
		SELECT
			id,
			full_name,
			email,
			role,
			department,
			position,
			phone,
			status,
			created_at,
			updated_at
		FROM employee
		ORDER BY id ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	employees := make([]model.Employee, 0)

	for rows.Next() {
		var employee model.Employee

		err := rows.Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Email,
			&employee.Role,
			&employee.Department,
			&employee.Position,
			&employee.Phone,
			&employee.Status,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		employees = append(
			employees,
			employee,
		)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

// ======================================================
// LOGIN
// ======================================================

func (r *employeeRepository) Login(
	email string,
) (*model.Employee, error) {
	query := `
		SELECT
			id,
			full_name,
			email,
			password,
			role,
			department,
			position,
			phone,
			status,
			created_at,
			updated_at
		FROM employee
		WHERE email = ?
	`

	var employee model.Employee

	err := r.db.QueryRow(
		query,
		email,
	).Scan(
		&employee.ID,
		&employee.FullName,
		&employee.Email,
		&employee.Password,
		&employee.Role,
		&employee.Department,
		&employee.Position,
		&employee.Phone,
		&employee.Status,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &employee, nil
}

// ======================================================
// GET BY ID
// ======================================================

func (r *employeeRepository) GetByID(
	id int64,
) (*model.Employee, error) {
	query := `
		SELECT
			id,
			full_name,
			email,
			role,
			department,
			position,
			phone,
			status,
			created_at,
			updated_at
		FROM employee
		WHERE id = ?
	`

	var employee model.Employee

	err := r.db.QueryRow(
		query,
		id,
	).Scan(
		&employee.ID,
		&employee.FullName,
		&employee.Email,
		&employee.Role,
		&employee.Department,
		&employee.Position,
		&employee.Phone,
		&employee.Status,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &employee, nil
}

// ======================================================
// CREATE
// ======================================================

func (r *employeeRepository) Create(
	employee *model.Employee,
) error {
	query := `
		INSERT INTO employee
		(
			full_name,
			email,
			password,
			role,
			department,
			position,
			phone,
			status,
			created_at,
			updated_at
		)
		VALUES
		(
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			CURRENT_TIMESTAMP,
			CURRENT_TIMESTAMP
		)
	`

	_, err := r.db.Exec(
		query,
		employee.FullName,
		employee.Email,
		employee.Password,
		employee.Role,
		employee.Department,
		employee.Position,
		employee.Phone,
		employee.Status,
	)

	return err
}

// ======================================================
// UPDATE EMPLOYEE
// ======================================================
//
// Update dari halaman Hosts / Employees.
// Password TIDAK diubah di sini.
//

func (r *employeeRepository) Update(
	employee *model.Employee,
) error {
	query := `
		UPDATE employee
		SET
			full_name = ?,
			email = ?,
			role = ?,
			department = ?,
			position = ?,
			phone = ?,
			status = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		employee.FullName,
		employee.Email,
		employee.Role,
		employee.Department,
		employee.Position,
		employee.Phone,
		employee.Status,
		employee.ID,
	)

	return err
}

// ======================================================
// DELETE
// ======================================================

func (r *employeeRepository) Delete(
	id int64,
) error {
	query := `
		DELETE FROM employee
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		id,
	)

	return err
}

// ======================================================
// REGISTER
// ======================================================

func (r *employeeRepository) Register(
	employee *model.Employee,
) error {
	return r.Create(employee)
}

// ======================================================
// UPDATE PROFILE
// ======================================================
//
// Digunakan oleh:
// PUT /api/profile
//
// Hanya Full Name dan Email yang diubah.
//

func (r *employeeRepository) UpdateProfile(
	id int64,
	fullName string,
	email string,
) error {
	query := `
		UPDATE employee
		SET
			full_name = ?,
			email = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		fullName,
		email,
		id,
	)

	return err
}

// ======================================================
// GET PASSWORD BY ID
// ======================================================
//
// Password hash hanya digunakan di backend untuk
// verifikasi password lama.
//
// Tidak pernah dikirim ke frontend.
//

func (r *employeeRepository) GetPasswordByID(
	id int64,
) (string, error) {
	query := `
		SELECT password
		FROM employee
		WHERE id = ?
	`

	var passwordHash string

	err := r.db.QueryRow(
		query,
		id,
	).Scan(&passwordHash)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}

		return "", err
	}

	return passwordHash, nil
}

// ======================================================
// UPDATE PASSWORD
// ======================================================
//
// HANYA SATU method UpdatePassword.
// Jangan membuat method kedua dengan nama sama.
//

func (r *employeeRepository) UpdatePassword(
	id int64,
	hashedPassword string,
) error {
	query := `
		UPDATE employee
		SET
			password = ?,
			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	result, err := r.db.Exec(
		query,
		hashedPassword,
		id,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
