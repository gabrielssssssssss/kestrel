package handlers

import (
	"net/http"

	"github.com/gabrielssssssssss/kestrel/internal/models"
	"github.com/gabrielssssssssss/kestrel/internal/services"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	serviceCompany *services.CompanyService
	serviceGmaps   *services.GoogleMapsService
}

func NewCompanyHandler() *CompanyHandler {
	return &CompanyHandler{
		serviceCompany: services.NewCompanyService(),
		serviceGmaps:   services.NewGoogleMapsService(),
	}
}

func (h *CompanyHandler) GetCompanyHandler(c *gin.Context) {
	var payload models.Company
	sirene := c.Query("sirene")

	company, err := h.serviceCompany.GetCompany(sirene)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	placeId, err := h.serviceGmaps.GetPlaceId(company.Result[0].NomComplet + " " + company.Result[0].Siege.Adresse)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	gmapsResult, err := h.serviceGmaps.GetPlaceDetails(placeId)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	payload = models.Company{
		Result:     company.Result,
		GoogleMaps: gmapsResult,
	}

	c.JSON(http.StatusOK, payload)
}
