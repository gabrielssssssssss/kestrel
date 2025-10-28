package services

import (
	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
	repository "github.com/gabrielssssssssss/kestrel/internal/repository/companies"
)

type WebsiteService struct {
	repo repository.WebsiteRepository
}

func NewWebsiteService() *WebsiteService {
	return &WebsiteService{
		repo: *repository.NewWebsiteRepository(),
	}
}

func (s *WebsiteService) GetWebsite(domain string) (models.Website, error) {
	return s.repo.FetchWebsite(domain)
}
