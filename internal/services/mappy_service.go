package services

import (
	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/gabrielssssssssss/kestrel/internal/repository"
)

type MappyService struct {
	repo *repository.MappyStruct
}

func NewMappyService() *MappyService {
	return &MappyService{
		repo: repository.NewMappyRepository(),
	}
}

func (s *MappyService) GetMappySearch(query string) (models.MappySearch, error) {
	return s.repo.FetchMappySearch(query)
}

func (s *MappyService) GetMappyGeo(id string) (models.MappyGeo, error) {
	return s.repo.FetchMappyGeo(id)
}
