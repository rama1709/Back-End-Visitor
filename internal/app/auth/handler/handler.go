package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	firebaseConfig "visitor_management_backend/firebase"

	authPayload "visitor_management_backend/internal/app/auth/payload"
	"visitor_management_backend/internal/app/employee/model"
	"visitor_management_backend/internal/app/employee/port"
	"visitor_management_backend/internal/helper"
)

type AuthHandler struct {
	service  port.EmployeeService
	firebase *firebaseConfig.Firebase
}

func NewAuthHandler(
	service port.EmployeeService,
	firebaseClient *firebaseConfig.Firebase,
) *AuthHandler {
	return &AuthHandler{
		service:  service,
		firebase: firebaseClient,
	}
}

// ======================================================
// LOGIN
// ======================================================

func (h *AuthHandler) Login(c *gin.Context) {
	var request authPayload.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	employee, err := h.service.Login(
		request.Email,
		request.Password,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if employee == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Email atau password salah",
		})
		return
	}

	token, err := helper.GenerateToken(
		employee.ID,
		employee.Email,
		employee.Role,
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	firebaseToken := ""

	if h.firebase != nil {
		customToken, firebaseErr :=
			h.firebase.CreateCustomToken(
				employee.Email,
			)

		if firebaseErr == nil {
			firebaseToken = customToken
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Login berhasil",
		"token":         token,
		"firebaseToken": firebaseToken,
		"employee": gin.H{
			"id":         employee.ID,
			"full_name":  employee.FullName,
			"email":      employee.Email,
			"role":       employee.Role,
			"department": employee.Department,
			"position":   employee.Position,
			"phone":      employee.Phone,
			"status":     employee.Status,
		},
	})
}

// ======================================================
// REGISTER
// ======================================================

func (h *AuthHandler) Register(c *gin.Context) {
	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.service.Register(&employee)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Register berhasil",
	})
}

// ======================================================
// GET USER ID FROM JWT
// ======================================================

func getUserID(c *gin.Context) (int64, bool) {
	value, exists := c.Get("user_id")

	if !exists {
		return 0, false
	}

	userID, ok := value.(int64)

	if !ok || userID <= 0 {
		return 0, false
	}

	return userID, true
}

// ======================================================
// UPDATE PROFILE
// ======================================================
//
// PUT /api/profile
//
// Body:
//
// {
//   "full_name": "Test User Baru",
//   "email": "baru@gmail.com"
// }
//
// ID tidak dikirim dari frontend.
// ID diambil dari JWT.
//

func (h *AuthHandler) UpdateProfile(c *gin.Context) {
	userID, ok := getUserID(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User tidak terautentikasi",
		})
		return
	}

	var request struct {
		FullName string `json:"full_name"`
		Email    string `json:"email"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data profile tidak valid",
			"error":   err.Error(),
		})
		return
	}

	request.FullName = strings.TrimSpace(
		request.FullName,
	)

	request.Email = strings.ToLower(
		strings.TrimSpace(request.Email),
	)

	if request.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Nama wajib diisi",
		})
		return
	}

	if request.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email wajib diisi",
		})
		return
	}

	err := h.service.UpdateProfile(
		userID,
		request.FullName,
		request.Email,
	)

	if err != nil {
		message := err.Error()

		if strings.Contains(
			strings.ToLower(message),
			"duplicate",
		) ||
			strings.Contains(
				strings.ToLower(message),
				"unique",
			) {

			c.JSON(http.StatusConflict, gin.H{
				"message": "Email sudah digunakan oleh employee lain",
			})
			return
		}

		if strings.Contains(message, "wajib") ||
			strings.Contains(message, "tidak valid") {

			c.JSON(http.StatusBadRequest, gin.H{
				"message": message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate profile",
			"error":   message,
		})
		return
	}

	// Ambil data terbaru dari database.
	employee, err := h.service.GetByID(userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Profile berhasil disimpan tetapi gagal mengambil data terbaru",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Profile berhasil diperbarui",
		"employee": gin.H{
			"id":         employee.ID,
			"full_name":  employee.FullName,
			"email":      employee.Email,
			"role":       employee.Role,
			"department": employee.Department,
			"position":   employee.Position,
			"phone":      employee.Phone,
			"status":     employee.Status,
		},
	})
}

// ======================================================
// CHANGE PASSWORD
// ======================================================
//
// PUT /api/profile/password
//
// Body:
//
// {
//   "current_password": "...",
//   "new_password": "...",
//   "confirm_password": "..."
// }
//
// ID diambil dari JWT.
//

func (h *AuthHandler) ChangePassword(c *gin.Context) {
	userID, ok := getUserID(c)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User tidak terautentikasi",
		})
		return
	}

	var request struct {
		CurrentPassword string `json:"current_password"`
		NewPassword     string `json:"new_password"`
		ConfirmPassword string `json:"confirm_password"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data password tidak valid",
			"error":   err.Error(),
		})
		return
	}

	err := h.service.ChangePassword(
		userID,
		request.CurrentPassword,
		request.NewPassword,
		request.ConfirmPassword,
	)

	if err != nil {
		message := err.Error()

		if message == "password saat ini salah" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": message,
			})
			return
		}

		if strings.Contains(message, "wajib diisi") ||
			strings.Contains(message, "minimal") ||
			strings.Contains(message, "tidak cocok") ||
			strings.Contains(message, "harus berbeda") {

			c.JSON(http.StatusBadRequest, gin.H{
				"message": message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengubah password",
			"error":   message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Password berhasil diubah",
	})
}
