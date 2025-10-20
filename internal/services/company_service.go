package services

import (
	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/gabrielssssssssss/kestrel/internal/repository"
)

type CompanyService struct {
	repo *repository.CompanyRepository
}

func NewCompanyService() *CompanyService {
	return &CompanyService{
		repo: repository.NewCompanyRepository(),
	}
}

func (s *CompanyService) GetCompany(url string) (models.Company, error) {
	return s.repo.FetchRechercheEntreprise(url)
}
