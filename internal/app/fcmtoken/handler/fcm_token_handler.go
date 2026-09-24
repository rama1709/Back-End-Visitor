package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"visitor_management_backend/internal/app/fcmtoken/model"
	"visitor_management_backend/internal/app/fcmtoken/port"
)

type FCMTokenHandler struct {
	service port.FCMTokenService
}

func NewFCMTokenHandler(service port.FCMTokenService) *FCMTokenHandler {
	return &FCMTokenHandler{
		service: service,
	}
}

func (h *FCMTokenHandler) Register(c *gin.Context) {

	var request model.FCMToken

	// Baca JSON
	if err := c.ShouldBindJSON(&request); err != nil {

		log.Println("❌ BIND JSON ERROR :", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Tampilkan data yang diterima
	log.Println("========== FCM REGISTER ==========")
	log.Println("Employee ID     :", request.EmployeeID)
	log.Println("Installation ID :", request.InstallationID)
	log.Println("Device Type     :", request.DeviceType)
	log.Println("Device Name     :", request.DeviceName)
	log.Println("Token Length    :", len(request.RegistrationToken))

	// Simpan
	if err := h.service.Register(&request); err != nil {

		log.Println("❌ SERVICE ERROR :", err)

		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	log.Println("✅ FCM TOKEN BERHASIL DISIMPAN")

	c.JSON(http.StatusCreated, gin.H{
		"message": "FCM token berhasil disimpan",
		"data":    request,
	})
}