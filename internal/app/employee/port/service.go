package port

import "visitor_management_backend/internal/app/employee/model"

type EmployeeService interface {
	GetAll() ([]model.Employee, error)
	GetByID(id int64) (*model.Employee, error)

	Create(employee *model.Employee) error
	Update(employee *model.Employee) error
	Delete(id int64) error

	Login(
		email string,
		password string,
	) (*model.Employee, error)

	Register(employee *model.Employee) error

	// Profile
	UpdateProfile(
		id int64,
		fullName string,
		email string,
	) error

	// Password
	ChangePassword(
		id int64,
		currentPassword string,
		newPassword string,
		confirmPassword string,
	) error
}
