package services

import (
	"github.com/gabrielssssssssss/kestrel/internal/repository"
	"github.com/gabrielssssssssss/kestrel/pkg/duckduckgo"
)

type LinkedinService struct {
	repo *repository.LinkedinRepository
}

func NewLinkedinService() *LinkedinService {
	return &LinkedinService{
		repo: repository.NewLinkedinRepository(),
	}
}

func (s *LinkedinService) GetLinkedinProfiles(company string) ([]duckduckgo.SearchResult, error) {
	return s.repo.FetchLinkedinCompany(company)
}
