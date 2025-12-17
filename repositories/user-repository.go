package repositories

import (
	"database/sql"

	"github.com/Danuson17-8/corn-backend/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	var user models.User

	err := r.DB.QueryRow(`
		SELECT email, name, role, address, mobile_phone
		FROM users
		WHERE email = ?
	`, email).Scan(
		&user.Account,
		&user.Name,
		&user.Role,
		&user.Address,
		&user.Contact.MobilePhone,
	)

	user.Contact.Email = user.Account

	return &user, err
}
