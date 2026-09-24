package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"visitor_management_backend/internal/app/visitor/model"
	"visitor_management_backend/internal/app/visitor/port"
)

type VisitorHandler struct {
	service port.VisitorService
}

func NewVisitorHandler(service port.VisitorService) *VisitorHandler {
	return &VisitorHandler{
		service: service,
	}
}

// ========================================
// GET ALL
// ========================================

func (h *VisitorHandler) GetAll(c *gin.Context) {

	visitors, err := h.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, visitors)
}

// ========================================
// GET BY ID
// ========================================

func (h *VisitorHandler) GetByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid visitor id",
		})
		return
	}

	visitor, err := h.service.GetByID(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if visitor == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Visitor not found",
		})
		return
	}

	c.JSON(http.StatusOK, visitor)
}

// ========================================
// CREATE
// ========================================

func (h *VisitorHandler) Create(c *gin.Context) {

	var visitor model.Visitor

	if err := c.ShouldBindJSON(&visitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	visitor.Status = "pending"

	if err := h.service.Create(&visitor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	created, _ := h.service.GetByID(visitor.ID)

	c.JSON(http.StatusCreated, created)
}

// ========================================
// UPDATE
// ========================================

func (h *VisitorHandler) Update(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid visitor id",
		})
		return
	}

	var visitor model.Visitor

	if err := c.ShouldBindJSON(&visitor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	visitor.ID = id

	if err := h.service.Update(&visitor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	updated, _ := h.service.GetByID(id)

	c.JSON(http.StatusOK, updated)
}

// ========================================
// UPDATE STATUS
// ========================================

func (h *VisitorHandler) UpdateStatus(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid visitor id",
		})
		return
	}

	var payload struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	switch payload.Status {
	case "approved",
		"checked-in",
		"checked-out",
		"rejected":
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid status",
		})
		return
	}

	if err := h.service.UpdateStatus(id, payload.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	visitor, _ := h.service.GetByID(id)

	c.JSON(http.StatusOK, visitor)
}

// ========================================
// DELETE
// ========================================

func (h *VisitorHandler) Delete(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid visitor id",
		})
		return
	}

	if err := h.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Visitor deleted successfully",
	})
}