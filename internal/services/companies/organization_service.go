package services

import (
	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
	repository "github.com/gabrielssssssssss/kestrel/internal/repository/companies"
)

type OrganizationService struct {
	repo *repository.OrganizationRepository
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{
		repo: repository.NewCompanyRepository(),
	}
}

func (s *OrganizationService) GetOrganization(sirene string) (models.OrganizationResult, error) {
	return s.repo.FetchOrganization(sirene)
}
