package handlers

import (
	"example.com/mod/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ProcurementHandler struct {
	service *services.ProcurementService
}

func NewProcurementHandler(service *services.ProcurementService) *ProcurementHandler {
	return &ProcurementHandler{service: service}
}

func (h *ProcurementHandler) GetProcurements(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		procurements, err := h.service.GetAllProcurements()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, gin.H{"procurements": procurements})
		return
	}

	layout := "2006-01-02"
	startDateDate, err := time.Parse(layout, startDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "StartDate is not in correct format (YYYY-MM-DD)"})
		return
	}
	EndDateDate, err := time.Parse(layout, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "EndDate is not in correct format (YYYY-MM-DD)"})
		return
	}

	// Преобразуем `time.Time` в формат "YYYY-MM-DD"
	startDateStr := startDateDate.Format("2006-01-02")
	endDateStr := EndDateDate.Format("2006-01-02")

	procurements, err := h.service.GetProcurementsByDateRange(startDateStr, endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"procurements": procurements})
	return
}
