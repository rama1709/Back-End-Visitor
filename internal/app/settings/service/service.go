package service

import (
	"visitor_management_backend/internal/app/settings/model"
	"visitor_management_backend/internal/app/settings/port"
)

type settingsService struct {
	repository port.SettingsRepository
}

func NewSettingsService(
	repository port.SettingsRepository,
) port.SettingsService {
	return &settingsService{
		repository: repository,
	}
}

func (s *settingsService) Get() (*model.CompanySettings, error) {
	return s.repository.Get()
}

func (s *settingsService) Update(
	settings *model.CompanySettings,
) error {
	return s.repository.Update(settings)
}