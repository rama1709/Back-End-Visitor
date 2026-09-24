package handler	

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"visitor_management_backend/internal/app/settings/model"
	"visitor_management_backend/internal/app/settings/port"
)

type SettingsHandler struct {
	service port.SettingsService
}

func NewSettingsHandler(
	service port.SettingsService,
) *SettingsHandler {
	return &SettingsHandler{
		service: service,
	}
}

func (h *SettingsHandler) Get(c *gin.Context) {
	settings, err := h.service.Get()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if settings == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Settings tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, settings)
}

func (h *SettingsHandler) Update(c *gin.Context) {
	var settings model.CompanySettings

	if err := c.ShouldBindJSON(&settings); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "ID settings tidak valid",
		})
		return
	}

	settings.ID = id

	if err := h.service.Update(&settings); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Settings berhasil diupdate",
	})
}