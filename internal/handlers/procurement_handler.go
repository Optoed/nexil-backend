package handlers

import (
	"example.com/mod/internal/models"
	"example.com/mod/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProcurementHandler struct {
	service *services.ProcurementService
}

func NewProcurementHandler(service *services.ProcurementService) *ProcurementHandler {
	return &ProcurementHandler{service: service}
}

func (h *ProcurementHandler) GetProcurements(c *gin.Context) {
	var procurements []models.Procurement
	procurements, err := h.service.GetAllProcurements()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"procurements": procurements})
}
