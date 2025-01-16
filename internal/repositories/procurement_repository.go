package repositories

import (
	"example.com/mod/internal/models"
	"github.com/jmoiron/sqlx"
)

type ProcurementRepository struct {
	db *sqlx.DB
}

func NewProcurementRepository(db *sqlx.DB) *ProcurementRepository {
	return &ProcurementRepository{db: db}
}

func (r *ProcurementRepository) GetAllProcurements() ([]models.Procurement, error) {
	var procurements []models.Procurement
	query := `SELECT * FROM procurement`
	err := r.db.Select(&procurements, query)
	if err != nil {
		return nil, err
	}
	return procurements, nil
}

func (r *ProcurementRepository) GetProcurementsByDateRange(startDate, endDate string) ([]models.Procurement, error) {
	var procurements []models.Procurement
	query := `SELECT * FROM procurement WHERE date BETWEEN $1 and $2`
	err := r.db.Select(&procurements, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return procurements, nil
}
