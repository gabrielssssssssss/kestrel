package handlers

import (
	"net/http"

	"github.com/gabrielssssssssss/kestrel/internal/services"
	"github.com/gin-gonic/gin"
)

type MappyHandler struct {
	service *services.MappyService
}

func NewMappyHandler() *MappyHandler {
	return &MappyHandler{
		service: services.NewMappyService(),
	}
}

func (h *MappyHandler) HandleMappySearchRequest(c *gin.Context) {
	payload, err := h.service.GetMappySearch(c.Query("q"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, payload.POIS[0])
}

func (h *MappyHandler) HandleMappyGeoRequest(c *gin.Context) {
	payload, err := h.service.GetMappyGeo(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, payload)
}
