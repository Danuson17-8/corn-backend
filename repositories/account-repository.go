package repositories

import "database/sql"

type AccountRepository struct {
	DB *sql.DB
}

func (r *AccountRepository) Create(email, password string) error {
	_, err := r.DB.Exec(`
		INSERT INTO accounts (email, password)
		VALUES (?, ?)
	`, email, password)
	return err
}

func (r *AccountRepository) GetPasswordHash(email string) (string, error) {
	var hash string
	err := r.DB.QueryRow(`
		SELECT password FROM accounts WHERE email = ?
	`, email).Scan(&hash)
	return hash, err
}
