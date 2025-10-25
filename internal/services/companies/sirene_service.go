package services

import repository "github.com/gabrielssssssssss/kestrel/internal/repository/companies"

type SireneService struct {
	repo *repository.SireneRepository
}

func NewSireneService() *SireneService {
	return &SireneService{
		repo: repository.NewSireneRepository(),
	}
}

func (s *SireneService) GetSirene(company, sector string) (string, error) {
	return s.repo.FetchSirene(company, sector)
}
