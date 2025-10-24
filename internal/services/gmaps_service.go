package services

import (
	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/gabrielssssssssss/kestrel/internal/repository"
)

type GoogleMapsService struct {
	repo *repository.GoogleMapsStruct
}

func NewGoogleMapsService() *GoogleMapsService {
	return &GoogleMapsService{
		repo: repository.NewMappyRepository(),
	}
}

func (s *GoogleMapsService) GetPlaceId(query string) (string, error) {
	return s.repo.FetchPlaceId(query)
}

func (s *GoogleMapsService) GetPlaceDetails(id string) (models.GoogleMaps, error) {
	return s.repo.FetchPlaceDetails(id)
}
