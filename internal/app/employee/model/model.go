package model

import "time"

type Employee struct {
	ID         int64     `json:"id"`
	FullName   string    `json:"full_name"`
	Email      string    `json:"email"`
	Password   string    `json:"password,omitempty"`
	Role       string    `json:"role"`
	Department string    `json:"department"`
	Position   string    `json:"position"`
	Phone      string    `json:"phone"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}