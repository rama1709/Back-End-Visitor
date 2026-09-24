ALTER TABLE company_settings
    DROP COLUMN monday_friday_start,
    DROP COLUMN monday_friday_end,
    DROP COLUMN saturday_start,
    DROP COLUMN saturday_end,
    DROP COLUMN sunday_start,
    DROP COLUMN sunday_end,

    DROP COLUMN require_id_verification,
    DROP COLUMN require_host_approval,
    DROP COLUMN capture_visitor_photo,
    DROP COLUMN auto_checkout,

    DROP COLUMN host_arrival_notification,
    DROP COLUMN pending_approval_alert,
    DROP COLUMN daily_summary_email,

    DROP COLUMN badge_show_visitor_photo,
    DROP COLUMN badge_show_host_name,
    DROP COLUMN badge_show_qr_code,
    DROP COLUMN badge_show_company,
    DROP COLUMN badge_footer_text;