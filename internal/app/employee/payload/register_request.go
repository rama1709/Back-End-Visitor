package payload

type RegisterRequest struct {
	FullName   string `json:"full_name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Role       string `json:"role"`
	Department string `json:"department"`
	Position   string `json:"position"`
	Phone      string `json:"phone"`
}
