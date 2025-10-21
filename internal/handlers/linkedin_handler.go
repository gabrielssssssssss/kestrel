package handlers

import (
	"fmt"
	"net/http"

	"github.com/gabrielssssssssss/kestrel/internal/services"
	"github.com/gin-gonic/gin"
)

type LinkedinHandlers struct {
	service *services.LinkedinService
}

func NewLinkedinHandler() *LinkedinHandlers {
	return &LinkedinHandlers{
		service: services.NewLinkedinService(),
	}
}

func (h *LinkedinHandlers) HandleLinkedin(c *gin.Context) {
	fmt.Println(c.Query("company"))
	payload, err := h.service.GetLinkedinProfiles(c.Query("company"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, payload)
}
