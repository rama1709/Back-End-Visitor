package service

import (
	"errors"
	"strings"

	"visitor_management_backend/internal/app/employee/model"
	"visitor_management_backend/internal/app/employee/port"
	"visitor_management_backend/internal/helper"
)

type employeeService struct {
	repository port.EmployeeRepository
}

func NewEmployeeService(
	repository port.EmployeeRepository,
) port.EmployeeService {
	return &employeeService{
		repository: repository,
	}
}

// ======================================================
// GET ALL EMPLOYEES
// ======================================================

func (s *employeeService) GetAll() ([]model.Employee, error) {
	return s.repository.GetAll()
}

// ======================================================
// GET EMPLOYEE BY ID
// ======================================================

func (s *employeeService) GetByID(
	id int64,
) (*model.Employee, error) {
	if id <= 0 {
		return nil, errors.New(
			"ID employee tidak valid",
		)
	}

	return s.repository.GetByID(id)
}

// ======================================================
// CREATE EMPLOYEE
// ======================================================

func (s *employeeService) Create(
	employee *model.Employee,
) error {
	if employee == nil {
		return errors.New(
			"data employee tidak boleh kosong",
		)
	}

	employee.FullName =
		strings.TrimSpace(employee.FullName)

	if employee.FullName == "" {
		return errors.New(
			"nama employee wajib diisi",
		)
	}

	employee.Email =
		strings.ToLower(
			strings.TrimSpace(employee.Email),
		)

	if employee.Email == "" {
		return errors.New(
			"email wajib diisi",
		)
	}

	if employee.Password == "" {
		return errors.New(
			"password wajib diisi",
		)
	}

	if len(employee.Password) < 6 {
		return errors.New(
			"password minimal 6 karakter",
		)
	}

	if strings.TrimSpace(employee.Role) == "" {
		employee.Role = "employee"
	}

	employee.Status =
		strings.ToLower(
			strings.TrimSpace(employee.Status),
		)

	if employee.Status == "" {
		employee.Status = "active"
	}

	switch employee.Status {
	case "active":
	case "on-leave":
	case "inactive":
	default:
		return errors.New(
			"status employee tidak valid",
		)
	}

	hashedPassword, err :=
		helper.HashPassword(employee.Password)

	if err != nil {
		return err
	}

	employee.Password = hashedPassword

	return s.repository.Create(employee)
}

// ======================================================
// UPDATE EMPLOYEE
// ======================================================

func (s *employeeService) Update(
	employee *model.Employee,
) error {
	if employee == nil {
		return errors.New(
			"data employee tidak boleh kosong",
		)
	}

	if employee.ID <= 0 {
		return errors.New(
			"ID employee tidak valid",
		)
	}

	employee.FullName =
		strings.TrimSpace(employee.FullName)

	if employee.FullName == "" {
		return errors.New(
			"nama employee wajib diisi",
		)
	}

	employee.Email =
		strings.ToLower(
			strings.TrimSpace(employee.Email),
		)

	if employee.Email == "" {
		return errors.New(
			"email wajib diisi",
		)
	}

	employee.Status =
		strings.ToLower(
			strings.TrimSpace(employee.Status),
		)

	if employee.Status == "" {
		employee.Status = "active"
	}

	switch employee.Status {
	case "active":
	case "on-leave":
	case "inactive":
	default:
		return errors.New(
			"status employee tidak valid",
		)
	}

	return s.repository.Update(employee)
}

// ======================================================
// DELETE EMPLOYEE
// ======================================================

func (s *employeeService) Delete(
	id int64,
) error {
	if id <= 0 {
		return errors.New(
			"ID employee tidak valid",
		)
	}

	return s.repository.Delete(id)
}

// ======================================================
// LOGIN
// ======================================================

func (s *employeeService) Login(
	email string,
	password string,
) (*model.Employee, error) {
	email = strings.ToLower(
		strings.TrimSpace(email),
	)

	if email == "" {
		return nil, errors.New(
			"email wajib diisi",
		)
	}

	if password == "" {
		return nil, errors.New(
			"password wajib diisi",
		)
	}

	employee, err :=
		s.repository.Login(email)

	if err != nil {
		return nil, err
	}

	if employee == nil {
		return nil, nil
	}

	if strings.ToLower(employee.Status) != "active" {
		return nil, errors.New(
			"akun employee tidak aktif",
		)
	}

	if employee.Password == "" {
		return nil, errors.New(
			"password employee belum tersedia",
		)
	}

	if !helper.CheckPassword(
		password,
		employee.Password,
	) {
		return nil, nil
	}

	return employee, nil
}

// ======================================================
// REGISTER
// ======================================================

func (s *employeeService) Register(
	employee *model.Employee,
) error {
	return s.Create(employee)
}

// ======================================================
// UPDATE PROFILE
// ======================================================

func (s *employeeService) UpdateProfile(
	id int64,
	fullName string,
	email string,
) error {
	if id <= 0 {
		return errors.New(
			"ID employee tidak valid",
		)
	}

	fullName = strings.TrimSpace(fullName)

	if fullName == "" {
		return errors.New(
			"nama employee wajib diisi",
		)
	}

	email = strings.ToLower(
		strings.TrimSpace(email),
	)

	if email == "" {
		return errors.New(
			"email wajib diisi",
		)
	}

	return s.repository.UpdateProfile(
		id,
		fullName,
		email,
	)
}

// ======================================================
// CHANGE PASSWORD
// ======================================================

func (s *employeeService) ChangePassword(
	id int64,
	currentPassword string,
	newPassword string,
	confirmPassword string,
) error {
	if id <= 0 {
		return errors.New(
			"ID employee tidak valid",
		)
	}

	if currentPassword == "" {
		return errors.New(
			"password saat ini wajib diisi",
		)
	}

	if newPassword == "" {
		return errors.New(
			"password baru wajib diisi",
		)
	}

	if len(newPassword) < 6 {
		return errors.New(
			"password baru minimal 6 karakter",
		)
	}

	if confirmPassword == "" {
		return errors.New(
			"konfirmasi password wajib diisi",
		)
	}

	if newPassword != confirmPassword {
		return errors.New(
			"konfirmasi password tidak cocok",
		)
	}

	if currentPassword == newPassword {
		return errors.New(
			"password baru harus berbeda dengan password lama",
		)
	}

	// ==============================================
	// AMBIL PASSWORD HASH
	// ==============================================

	passwordHash, err :=
		s.repository.GetPasswordByID(id)

	if err != nil {
		return err
	}

	if passwordHash == "" {
		return errors.New(
			"password employee belum tersedia",
		)
	}

	// ==============================================
	// CEK PASSWORD LAMA
	// ==============================================

	if !helper.CheckPassword(
		currentPassword,
		passwordHash,
	) {
		return errors.New(
			"password saat ini salah",
		)
	}

	// ==============================================
	// HASH PASSWORD BARU
	// ==============================================

	newPasswordHash, err :=
		helper.HashPassword(newPassword)

	if err != nil {
		return err
	}

	// ==============================================
	// SIMPAN PASSWORD BARU
	// ==============================================

	return s.repository.UpdatePassword(
		id,
		newPasswordHash,
	)
}
