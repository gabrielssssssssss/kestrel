package services

import (
	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
	repository "github.com/gabrielssssssssss/kestrel/internal/repository/companies"
)

type MapsService struct {
	repo *repository.MapsStruct
}

func NewMapsService() *MapsService {
	return &MapsService{
		repo: repository.NewMappyRepository(),
	}
}

func (s *MapsService) GetPlaceId(query string) (string, error) {
	return s.repo.FetchPlaceId(query)
}

func (s *MapsService) GetPlaceDetails(id string) (models.Maps, error) {
	return s.repo.FetchPlaceDetails(id)
}
