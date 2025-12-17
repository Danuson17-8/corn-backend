package repositories

import (
	"database/sql"

	"github.com/Danuson17-8/corn-backend/models"
)

type ContactRepository struct {
	DB *sql.DB
}

func (r *ContactRepository) Create(contact *models.ContactUser) error {
	_, err := r.DB.Exec(`
		INSERT INTO contacts (topic, name, email, message)
		VALUES (?, ?, ?, ?)
	`, contact.Topic, contact.Name, contact.Email, contact.Message)
	return err
}
