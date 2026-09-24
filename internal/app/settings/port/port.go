package port

import "visitor_management_backend/internal/app/settings/model"

type SettingsRepository interface {
	Get() (*model.CompanySettings, error)
	Update(settings *model.CompanySettings) error
}

type SettingsService interface {
	Get() (*model.CompanySettings, error)
	Update(settings *model.CompanySettings) error
}