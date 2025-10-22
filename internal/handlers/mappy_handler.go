package handlers

import (
	"fmt"
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

func (h *MappyHandler) HandleMappy(c *gin.Context) {
	mappyId, err := h.service.GetMappySearch(c.Query("q"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	fmt.Print(mappyId)
	payload, err := h.service.GetMappyGeo(mappyId.POIS[0].ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, payload)
}
