package handlers

import (
	"net/http"

	"github.com/gabrielssssssssss/kestrel/internal/services"
	"github.com/gin-gonic/gin"
)

type WebsiteHandler struct {
	service *services.WebsiteService
}

func NewWebsiteHandler() *WebsiteHandler {
	return &WebsiteHandler{
		service: services.NewWebsiteService(),
	}
}

func (h *WebsiteHandler) HandleWebsite(c *gin.Context) {
	payload, err := h.service.GetWebsite(c.Query("d"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, payload)
}
