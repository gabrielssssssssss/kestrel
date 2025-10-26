package handlers

import (
	"net/http"

	models "github.com/gabrielssssssssss/kestrel/internal/models/companies"
	services "github.com/gabrielssssssssss/kestrel/internal/services/companies"
	"github.com/gin-gonic/gin"
)

type CompaniesHandler struct {
	serviceSirene   *services.SireneService
	serviceCompany  *services.OrganizationService
	serviceGmaps    *services.GoogleMapsService
	serviceEmployee *services.EmployeeService
}

func NewCompaniesHandler() *CompaniesHandler {
	return &CompaniesHandler{
		serviceSirene:   services.NewSireneService(),
		serviceCompany:  services.NewOrganizationService(),
		serviceGmaps:    services.NewGoogleMapsService(),
		serviceEmployee: services.NewEmployeeService(),
	}
}

func (h *CompaniesHandler) GetSireneHandler(c *gin.Context) {
	company, sector := c.Query("company"), c.Query("sector")
	sirene, err := h.serviceSirene.GetSirene(company, sector)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, sirene)
}

func (h *CompaniesHandler) GetCompanyHandler(c *gin.Context) {
	var (
		payload    models.Company
		googleMaps models.GoogleMaps
		sirene     = c.Query("sirene")
	)

	organization, err := h.serviceCompany.GetOrganization(sirene)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	placeId, err := h.serviceGmaps.GetPlaceId(organization.Result[0].NomComplet + " " + organization.Result[0].Siege.Adresse)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	if placeId != "" {
		googleMaps, err = h.serviceGmaps.GetPlaceDetails(placeId)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
	}

	payload = models.Company{
		Organization: organization,
		GoogleMaps:   googleMaps,
	}
	c.JSON(http.StatusOK, payload)
}

func (h *CompaniesHandler) GetEmployeeHandler(c *gin.Context) {
	query := c.Query("query")
	employees, err := h.serviceEmployee.GetEmployees(query)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, employees)
}
