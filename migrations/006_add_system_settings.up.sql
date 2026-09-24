ALTER TABLE company_settings
    ADD COLUMN monday_friday_start TIME NULL DEFAULT '08:00:00',
    ADD COLUMN monday_friday_end TIME NULL DEFAULT '17:00:00',
    ADD COLUMN saturday_start TIME NULL DEFAULT '08:00:00',
    ADD COLUMN saturday_end TIME NULL DEFAULT '13:00:00',
    ADD COLUMN sunday_start TIME NULL DEFAULT NULL,
    ADD COLUMN sunday_end TIME NULL DEFAULT NULL,

    ADD COLUMN require_id_verification BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN require_host_approval BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN capture_visitor_photo BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN auto_checkout BOOLEAN NOT NULL DEFAULT FALSE,

    ADD COLUMN host_arrival_notification BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN pending_approval_alert BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN daily_summary_email BOOLEAN NOT NULL DEFAULT FALSE,

    ADD COLUMN badge_show_visitor_photo BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN badge_show_host_name BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN badge_show_qr_code BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN badge_show_company BOOLEAN NOT NULL DEFAULT TRUE,
    ADD COLUMN badge_footer_text TEXT NULL;