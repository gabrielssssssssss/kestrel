package services

import (
	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/gabrielssssssssss/kestrel/internal/repository"
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
