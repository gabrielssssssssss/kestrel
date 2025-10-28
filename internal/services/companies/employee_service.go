package services

import (
	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
	repository "github.com/gabrielssssssssss/kestrel/internal/repository/companies"
)

type EmployeeService struct {
	repo *repository.EmployeeStruct
}

func NewEmployeeService() *EmployeeService {
	return &EmployeeService{
		repo: repository.NewEmployeeRepository(),
	}
}

func (s *EmployeeService) GetEmployees(query string, naf string) (models.Employee, error) {
	return s.repo.FetchEmployee(query, naf)
}
