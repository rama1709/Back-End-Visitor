package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"visitor_management_backend/internal/app/visitorrequest/model"
	"visitor_management_backend/internal/app/visitorrequest/port"
)

type VisitorRequestHandler struct {
	service port.VisitorRequestService
}

func NewVisitorRequestHandler(
	service port.VisitorRequestService,
) *VisitorRequestHandler {
	return &VisitorRequestHandler{
		service: service,
	}
}

// ======================================================
// GET ALL
// ======================================================

func (h *VisitorRequestHandler) GetAll(c *gin.Context) {

	requests, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Gagal mengambil appointment",
			"error":   err.Error(),
		})
		return
	}

	if requests == nil {
		requests = []model.VisitorRequest{}
	}

	c.JSON(http.StatusOK, requests)
}

// ======================================================
// GET BY ID
// ======================================================

func (h *VisitorRequestHandler) GetByID(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID tidak valid",
		})
		return
	}

	request, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if request == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Appointment tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, request)
}

// ======================================================
// CREATE
// ======================================================

func (h *VisitorRequestHandler) Create(c *gin.Context) {

	var request model.VisitorRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Data tidak valid",
			"error":   err.Error(),
		})
		return
	}

	request.Status = "requested"

	if err := h.service.Create(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	created, _ := h.service.GetByID(request.ID)

	c.JSON(http.StatusCreated, created)
}

// ======================================================
// UPDATE
// ======================================================

func (h *VisitorRequestHandler) Update(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID tidak valid",
		})
		return
	}

	var request model.VisitorRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	request.ID = id

	if err := h.service.Update(&request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	updated, _ := h.service.GetByID(id)

	c.JSON(http.StatusOK, updated)
}

// ======================================================
// APPROVE
// ======================================================

func (h *VisitorRequestHandler) Approve(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID tidak valid",
		})
		return
	}

	if err := h.service.Approve(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, _ := h.service.GetByID(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment berhasil disetujui",
		"data":    data,
	})
}

// ======================================================
// REJECT
// ======================================================

func (h *VisitorRequestHandler) Reject(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID tidak valid",
		})
		return
	}

	if err := h.service.Reject(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, _ := h.service.GetByID(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment berhasil ditolak",
		"data":    data,
	})
}

// ======================================================
// CHECK IN
// ======================================================

func (h *VisitorRequestHandler) CheckIn(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID tidak valid",
		})
		return
	}

	if err := h.service.CheckIn(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, _ := h.service.GetByID(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Visitor berhasil check in",
		"data":    data,
	})
}

// ======================================================
// CHECK OUT
// ======================================================

func (h *VisitorRequestHandler) CheckOut(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID tidak valid",
		})
		return
	}

	if err := h.service.CheckOut(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	data, _ := h.service.GetByID(id)

	c.JSON(http.StatusOK, gin.H{
		"message": "Visitor berhasil check out",
		"data":    data,
	})
}

// ======================================================
// DELETE
// ======================================================

func (h *VisitorRequestHandler) Delete(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID tidak valid",
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
		"message": "Appointment berhasil dihapus",
	})
}