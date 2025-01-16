package services

import (
	"example.com/mod/internal/models"
	"example.com/mod/internal/repositories"
)

type PurchaseService struct {
	repo *repositories.PurchaseRepository
}

func NewPurchaseService(repo *repositories.PurchaseRepository) *PurchaseService {
	return &PurchaseService{repo: repo}
}

func (s *PurchaseService) GetAllPurchases() ([]models.Purchase, error) {
	return s.repo.GetAllPurchases()
}

func (s *PurchaseService) GetPurchasesByDateRange(startDateStr, endDateStr string) ([]models.Purchase, error) {
	return s.repo.GetPurchasesByDateRange(startDateStr, endDateStr)
}
