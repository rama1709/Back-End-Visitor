package handler

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"visitor_management_backend/internal/app/employee/model"
	"visitor_management_backend/internal/app/employee/port"
)

type EmployeeHandler struct {
	service port.EmployeeService
}

func NewEmployeeHandler(
	service port.EmployeeService,
) *EmployeeHandler {
	return &EmployeeHandler{
		service: service,
	}
}

// ======================================================
// GET ALL EMPLOYEES
// ======================================================

func (h *EmployeeHandler) GetAll(c *gin.Context) {
	employees, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data employee",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, employees)
}

// ======================================================
// GET HOSTS
// ======================================================

func (h *EmployeeHandler) GetHosts(c *gin.Context) {
	employees, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data host",
			"error":   err.Error(),
		})
		return
	}

	hosts := make([]gin.H, 0)

	for _, employee := range employees {
		if strings.ToLower(employee.Status) != "active" {
			continue
		}

		hosts = append(hosts, gin.H{
			"id":         employee.ID,
			"name":       employee.FullName,
			"department": employee.Department,
			"position":   employee.Position,
			"email":      employee.Email,
			"phone":      employee.Phone,
		})
	}

	c.JSON(http.StatusOK, hosts)
}

// ======================================================
// GET EMPLOYEE BY ID
// ======================================================

func (h *EmployeeHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID employee tidak valid",
		})
		return
	}

	employee, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil data employee",
			"error":   err.Error(),
		})
		return
	}

	if employee == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Employee tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, employee)
}

// ======================================================
// CREATE EMPLOYEE
// ======================================================

func (h *EmployeeHandler) Create(c *gin.Context) {
	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data employee tidak valid",
			"error":   err.Error(),
		})
		return
	}

	if err := h.service.Create(&employee); err != nil {
		message := err.Error()

		if strings.Contains(message, "wajib diisi") ||
			strings.Contains(message, "minimal") ||
			strings.Contains(message, "tidak valid") {

			c.JSON(http.StatusBadRequest, gin.H{
				"message": message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal membuat employee",
			"error":   message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Employee berhasil dibuat",
	})
}

// ======================================================
// UPDATE EMPLOYEE
// ======================================================

func (h *EmployeeHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID employee tidak valid",
		})
		return
	}

	var employee model.Employee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data employee tidak valid",
			"error":   err.Error(),
		})
		return
	}

	employee.ID = id

	if err := h.service.Update(&employee); err != nil {
		message := err.Error()

		if strings.Contains(message, "wajib diisi") ||
			strings.Contains(message, "tidak valid") {

			c.JSON(http.StatusBadRequest, gin.H{
				"message": message,
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengupdate employee",
			"error":   message,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee berhasil diupdate",
	})
}

// ======================================================
// DELETE EMPLOYEE
// ======================================================

func (h *EmployeeHandler) Delete(c *gin.Context) {
	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)

	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID employee tidak valid",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal menghapus employee",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Employee berhasil dihapus",
	})
}
