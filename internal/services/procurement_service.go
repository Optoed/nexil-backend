package services

import (
	"example.com/mod/internal/models"
	"example.com/mod/internal/repositories"
)

type ProcurementService struct {
	repo *repositories.ProcurementRepository
}

func NewProcurementService(repo *repositories.ProcurementRepository) *ProcurementService {
	return &ProcurementService{repo: repo}
}

func (s *ProcurementService) GetAllProcurements() ([]models.Procurement, error) {
	return s.repo.GetAllProcurements()
}
