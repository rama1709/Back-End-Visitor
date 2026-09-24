package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"visitor_management_backend/config"
	firebaseConfig "visitor_management_backend/firebase"

	authHandler "visitor_management_backend/internal/app/auth/handler"

	employeeHandler "visitor_management_backend/internal/app/employee/handler"
	employeeRepository "visitor_management_backend/internal/app/employee/repository"
	employeeService "visitor_management_backend/internal/app/employee/service"

	fcmHandler "visitor_management_backend/internal/app/fcmtoken/handler"
	fcmRepository "visitor_management_backend/internal/app/fcmtoken/repository"
	fcmService "visitor_management_backend/internal/app/fcmtoken/service"

	settingsHandler "visitor_management_backend/internal/app/settings/handler"
	settingsRepository "visitor_management_backend/internal/app/settings/repository"
	settingsService "visitor_management_backend/internal/app/settings/service"

	visitorHandler "visitor_management_backend/internal/app/visitor/handler"
	visitorRepository "visitor_management_backend/internal/app/visitor/repository"
	visitorService "visitor_management_backend/internal/app/visitor/service"

	visitorRequestHandler "visitor_management_backend/internal/app/visitorrequest/handler"
	visitorRequestRepository "visitor_management_backend/internal/app/visitorrequest/repository"
	visitorRequestService "visitor_management_backend/internal/app/visitorrequest/service"

	"visitor_management_backend/internal/middleware"
)

func main() {

	// ========================================
	// DATABASE
	// ========================================

	config.ConnectDatabase()

	if config.DB == nil {
		log.Fatal("Database connection failed")
	}

	// ========================================
	// FIREBASE
	// ========================================

	firebaseClient, err := firebaseConfig.Initialize()

	if err != nil {
		log.Fatal(err)
	}

	defer firebaseClient.Firestore.Close()

	fmt.Println("===================================")
	fmt.Println("Visitor Management Backend")
	fmt.Println("===================================")

	// ========================================
	// ROUTER
	// ========================================

	router := gin.Default()

	router.Use(
		middleware.CORSMiddleware(),
	)

	// ========================================
	// EMPLOYEE
	// ========================================

	employeeRepo :=
		employeeRepository.NewEmployeeRepository(
			config.DB,
		)

	employeeSvc :=
		employeeService.NewEmployeeService(
			employeeRepo,
		)

	employeeHdl :=
		employeeHandler.NewEmployeeHandler(
			employeeSvc,
		)

	// ========================================
	// AUTH
	// ========================================

	authHdl :=
		authHandler.NewAuthHandler(
			employeeSvc,
			firebaseClient,
		)

	// ========================================
	// VISITOR
	// ========================================

	visitorRepo :=
		visitorRepository.NewVisitorRepository(
			config.DB,
		)

	visitorSvc :=
		visitorService.NewVisitorService(
			visitorRepo,
		)

	visitorHdl :=
		visitorHandler.NewVisitorHandler(
			visitorSvc,
		)

	// ========================================
	// VISITOR REQUEST
	// ========================================

	visitorRequestRepo :=
		visitorRequestRepository.NewVisitorRequestRepository(
			config.DB,
		)

	visitorRequestSvc :=
		visitorRequestService.NewVisitorRequestService(
			visitorRequestRepo,
			visitorRepo,
			firebaseClient,
		)

	visitorRequestHdl :=
		visitorRequestHandler.NewVisitorRequestHandler(
			visitorRequestSvc,
		)

	// ========================================
	// FCM TOKEN
	// ========================================

	fcmRepo :=
		fcmRepository.NewFCMTokenRepository(
			config.DB,
		)

	fcmSvc :=
		fcmService.NewFCMTokenService(
			fcmRepo,
		)

	fcmHdl :=
		fcmHandler.NewFCMTokenHandler(
			fcmSvc,
		)

	// ========================================
	// SETTINGS
	// ========================================

	settingsRepo :=
		settingsRepository.NewSettingsRepository(
			config.DB,
		)

	settingsSvc :=
		settingsService.NewSettingsService(
			settingsRepo,
		)

	settingsHdl :=
		settingsHandler.NewSettingsHandler(
			settingsSvc,
		)

	// ========================================
	// API
	// ========================================

	api := router.Group("/api")

	// ========================================
	// PUBLIC ROUTES
	// ========================================

	api.POST(
		"/register",
		authHdl.Register,
	)

	api.POST(
		"/login",
		authHdl.Login,
	)

	api.POST(
		"/visitors",
		visitorHdl.Create,
	)

	api.GET(
		"/hosts",
		employeeHdl.GetHosts,
	)

	api.POST(
		"/visitor-requests",
		visitorRequestHdl.Create,
	)

	// ========================================
	// PROTECTED ROUTES
	// ========================================

	auth := api.Group("/")

	// IMPORTANT:
	// Gunakan AuthMiddleware karena middleware ini
	// menyimpan user_id, user_email, user_role
	// ke Gin Context.

	auth.Use(
		middleware.AuthMiddleware(),
	)

	// ========================================
	// PROFILE
	// ========================================

	auth.PUT(
		"/profile",
		authHdl.UpdateProfile,
	)

	auth.PUT(
		"/profile/password",
		authHdl.ChangePassword,
	)

	// ========================================
	// EMPLOYEE
	// ========================================

	auth.GET(
		"/employees",
		employeeHdl.GetAll,
	)

	auth.GET(
		"/employees/:id",
		employeeHdl.GetByID,
	)

	auth.POST(
		"/employees",
		employeeHdl.Create,
	)

	auth.PUT(
		"/employees/:id",
		employeeHdl.Update,
	)

	auth.DELETE(
		"/employees/:id",
		employeeHdl.Delete,
	)

	// ========================================
	// VISITOR
	// ========================================

	auth.GET(
		"/visitors",
		visitorHdl.GetAll,
	)

	auth.GET(
		"/visitors/:id",
		visitorHdl.GetByID,
	)

	auth.PUT(
		"/visitors/:id",
		visitorHdl.Update,
	)

	auth.PUT(
		"/visitors/:id/status",
		visitorHdl.UpdateStatus,
	)

	auth.DELETE(
		"/visitors/:id",
		visitorHdl.Delete,
	)

	// ========================================
	// FCM TOKEN
	// ========================================

	auth.POST(
		"/fcm/register",
		fcmHdl.Register,
	)

	// ========================================
	// VISITOR REQUEST / APPOINTMENT
	// ========================================

	auth.GET(
		"/visitor-requests",
		visitorRequestHdl.GetAll,
	)

	auth.GET(
		"/visitor-requests/:id",
		visitorRequestHdl.GetByID,
	)

	auth.PUT(
		"/visitor-requests/:id",
		visitorRequestHdl.Update,
	)

	auth.DELETE(
		"/visitor-requests/:id",
		visitorRequestHdl.Delete,
	)

	auth.PUT(
		"/visitor-requests/:id/approve",
		visitorRequestHdl.Approve,
	)

	auth.PUT(
		"/visitor-requests/:id/reject",
		visitorRequestHdl.Reject,
	)

	auth.PUT(
		"/visitor-requests/:id/checkin",
		visitorRequestHdl.CheckIn,
	)

	auth.PUT(
		"/visitor-requests/:id/checkout",
		visitorRequestHdl.CheckOut,
	)

	// ========================================
	// SETTINGS
	// ========================================

	auth.GET(
		"/settings",
		settingsHdl.Get,
	)

	auth.PUT(
		"/settings/:id",
		settingsHdl.Update,
	)

	// ========================================
	// START SERVER
	// ========================================

	fmt.Println(
		"Server running : http://localhost:8080",
	)

	fmt.Println(
		"API Base URL   : http://localhost:8080/api",
	)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
