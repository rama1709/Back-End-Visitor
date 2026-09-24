package repository

import (
	"database/sql"

	"visitor_management_backend/internal/app/settings/model"
	"visitor_management_backend/internal/app/settings/port"
)

type settingsRepository struct {
	db *sql.DB
}

func NewSettingsRepository(db *sql.DB) port.SettingsRepository {
	return &settingsRepository{
		db: db,
	}
}

func (r *settingsRepository) Get() (*model.CompanySettings, error) {
	query := `
		SELECT
			id,
			company_name,
			company_email,
			company_phone,
			company_website,
			company_address,

			TIME_FORMAT(monday_friday_start, '%H:%i'),
			TIME_FORMAT(monday_friday_end, '%H:%i'),
			TIME_FORMAT(saturday_start, '%H:%i'),
			TIME_FORMAT(saturday_end, '%H:%i'),
			COALESCE(TIME_FORMAT(sunday_start, '%H:%i'), ''),
			COALESCE(TIME_FORMAT(sunday_end, '%H:%i'), ''),

			require_id_verification,
			require_host_approval,
			capture_visitor_photo,
			auto_checkout,

			host_arrival_notification,
			pending_approval_alert,
			daily_summary_email,

			badge_show_visitor_photo,
			badge_show_host_name,
			badge_show_qr_code,
			badge_show_company,
			COALESCE(badge_footer_text, ''),

			created_at,
			updated_at
		FROM company_settings
		ORDER BY id ASC
		LIMIT 1
	`

	var settings model.CompanySettings

	err := r.db.QueryRow(query).Scan(
		&settings.ID,
		&settings.CompanyName,
		&settings.CompanyEmail,
		&settings.CompanyPhone,
		&settings.CompanyWebsite,
		&settings.CompanyAddress,

		&settings.MondayFridayStart,
		&settings.MondayFridayEnd,
		&settings.SaturdayStart,
		&settings.SaturdayEnd,
		&settings.SundayStart,
		&settings.SundayEnd,

		&settings.RequireIDVerification,
		&settings.RequireHostApproval,
		&settings.CaptureVisitorPhoto,
		&settings.AutoCheckout,

		&settings.HostArrivalNotification,
		&settings.PendingApprovalAlert,
		&settings.DailySummaryEmail,

		&settings.BadgeShowVisitorPhoto,
		&settings.BadgeShowHostName,
		&settings.BadgeShowQRCode,
		&settings.BadgeShowCompany,
		&settings.BadgeFooterText,

		&settings.CreatedAt,
		&settings.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &settings, nil
}

func (r *settingsRepository) Update(
	settings *model.CompanySettings,
) error {
	query := `
		UPDATE company_settings
		SET
			company_name = ?,
			company_email = ?,
			company_phone = ?,
			company_website = ?,
			company_address = ?,

			monday_friday_start = NULLIF(?, ''),
			monday_friday_end = NULLIF(?, ''),
			saturday_start = NULLIF(?, ''),
			saturday_end = NULLIF(?, ''),
			sunday_start = NULLIF(?, ''),
			sunday_end = NULLIF(?, ''),

			require_id_verification = ?,
			require_host_approval = ?,
			capture_visitor_photo = ?,
			auto_checkout = ?,

			host_arrival_notification = ?,
			pending_approval_alert = ?,
			daily_summary_email = ?,

			badge_show_visitor_photo = ?,
			badge_show_host_name = ?,
			badge_show_qr_code = ?,
			badge_show_company = ?,
			badge_footer_text = ?,

			updated_at = CURRENT_TIMESTAMP
		WHERE id = ?
	`

	_, err := r.db.Exec(
		query,
		settings.CompanyName,
		settings.CompanyEmail,
		settings.CompanyPhone,
		settings.CompanyWebsite,
		settings.CompanyAddress,

		settings.MondayFridayStart,
		settings.MondayFridayEnd,
		settings.SaturdayStart,
		settings.SaturdayEnd,
		settings.SundayStart,
		settings.SundayEnd,

		settings.RequireIDVerification,
		settings.RequireHostApproval,
		settings.CaptureVisitorPhoto,
		settings.AutoCheckout,

		settings.HostArrivalNotification,
		settings.PendingApprovalAlert,
		settings.DailySummaryEmail,

		settings.BadgeShowVisitorPhoto,
		settings.BadgeShowHostName,
		settings.BadgeShowQRCode,
		settings.BadgeShowCompany,
		settings.BadgeFooterText,

		settings.ID,
	)

	return err
}