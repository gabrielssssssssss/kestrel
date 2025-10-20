package handlers

import (
	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/gabrielssssssssss/kestrel/internal/services"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler() *CompanyHandler {
	return &CompanyHandler{
		service: services.NewCompanyService(),
	}
}

func (h *CompanyHandler) HandleCompanyRequest(url string) (models.Company, error) {
	return h.service.GetCompany(url)
}
