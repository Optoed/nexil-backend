package repositories

import (
	"example.com/mod/internal/models"
	"github.com/jmoiron/sqlx"
)

type PurchaseRepository struct {
	db *sqlx.DB
}

func NewPurchaseRepository(db *sqlx.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

func (r *PurchaseRepository) GetAllPurchases() ([]models.Purchase, error) {
	var purchases []models.Purchase
	query := `SELECT * FROM purchase`
	err := r.db.Select(&purchases, query)
	if err != nil {
		return nil, err
	}
	return purchases, nil
}

func (r *PurchaseRepository) GetPurchasesByDateRange(startDate, endDate string) ([]models.Purchase, error) {
	var purchases []models.Purchase
	query := `SELECT * FROM purchase WHERE date BETWEEN $1 and $2`
	err := r.db.Select(&purchases, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	return purchases, nil
}
