package model

import "time"

type CompanySettings struct {
	ID             int64  `json:"id"`
	CompanyName    string `json:"company_name"`
	CompanyEmail   string `json:"company_email"`
	CompanyPhone   string `json:"company_phone"`
	CompanyWebsite string `json:"company_website"`
	CompanyAddress string `json:"company_address"`

	MondayFridayStart string `json:"monday_friday_start"`
	MondayFridayEnd   string `json:"monday_friday_end"`
	SaturdayStart     string `json:"saturday_start"`
	SaturdayEnd       string `json:"saturday_end"`
	SundayStart       string `json:"sunday_start"`
	SundayEnd         string `json:"sunday_end"`

	RequireIDVerification bool `json:"require_id_verification"`
	RequireHostApproval   bool `json:"require_host_approval"`
	CaptureVisitorPhoto   bool `json:"capture_visitor_photo"`
	AutoCheckout          bool `json:"auto_checkout"`

	HostArrivalNotification bool `json:"host_arrival_notification"`
	PendingApprovalAlert    bool `json:"pending_approval_alert"`
	DailySummaryEmail       bool `json:"daily_summary_email"`

	BadgeShowVisitorPhoto bool   `json:"badge_show_visitor_photo"`
	BadgeShowHostName     bool   `json:"badge_show_host_name"`
	BadgeShowQRCode       bool   `json:"badge_show_qr_code"`
	BadgeShowCompany      bool   `json:"badge_show_company"`
	BadgeFooterText       string `json:"badge_footer_text"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}