package repositories

import (
	"database/sql"

	"github.com/Danuson17-8/corn-backend/models"
)

type MenuRepository struct {
	DB *sql.DB
}

func (r *MenuRepository) GetAll() ([]models.CornMenu, error) {
	rows, err := r.DB.Query(`
		SELECT 
			id, name_th, name_en, price, image, description, stock
		FROM corn_menu
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	menus := []models.CornMenu{}

	for rows.Next() {
		var menu models.CornMenu
		err := rows.Scan(
			&menu.ID,
			&menu.Name.Th,
			&menu.Name.En,
			&menu.Price,
			&menu.Image,
			&menu.Description,
			&menu.Stock,
		)
		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}
