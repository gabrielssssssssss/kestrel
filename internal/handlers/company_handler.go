package handlers

import (
	"net/http"
	"net/url"

	"github.com/gabrielssssssssss/kestrel/internal/services"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	service *services.CompanyService
}

func NewCompanyHandler() *CompanyHandler {
	return &CompanyHandler{
		service: services.NewCompanyService(),
	}
}

func (h *CompanyHandler) HandleCompanyRequest(c *gin.Context) {
	params := url.Values{}
	if q := c.Query("q"); q != "" {
		params.Add("q", q)
	}
	if cp := c.Query("code_postal"); cp != "" {
		params.Add("code_postal", cp)
	}
	if dep := c.Query("departement"); dep != "" {
		params.Add("departement", dep)
	}

	payload, err := h.service.GetCompany("https://recherche-entreprises.api.gouv.fr/search?" + params.Encode())
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, payload)
}
