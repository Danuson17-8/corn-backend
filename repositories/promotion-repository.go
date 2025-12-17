package repositories

import (
	"database/sql"

	"github.com/Danuson17-8/corn-backend/models"
)

type PromotionRepository struct {
	DB *sql.DB
}

func (r *PromotionRepository) GetActive() ([]models.Promotion, error) {
	rows, err := r.DB.Query(`
		SELECT 
			id, title, description, image,
			start_date, end_date, is_active, link
		FROM promotions
		WHERE is_active = 1
		ORDER BY start_date DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var promotions []models.Promotion

	for rows.Next() {
		var p models.Promotion
		var startDate, endDate sql.NullTime

		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.Image,
			&startDate,
			&endDate,
			&p.IsActive,
			&p.Link,
		); err != nil {
			return nil, err
		}

		if startDate.Valid {
			p.StartDate = startDate.Time.Format("2006-01-02")
		}
		if endDate.Valid {
			p.EndDate = endDate.Time.Format("2006-01-02")
		}

		promotions = append(promotions, p)
	}

	return promotions, nil
}
