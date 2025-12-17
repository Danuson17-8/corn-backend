package repositories

import (
	"database/sql"
	"time"
)

type OTPRepository struct {
	DB *sql.DB
}

func (r *OTPRepository) Upsert(email string, code int, session string, expired time.Time) error {
	_, err := r.DB.Exec(`
		INSERT INTO email_verification (email, code, otp_session, expired_at)
		VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
			code=?, otp_session=?, expired_at=?
	`, email, code, session, expired,
		code, session, expired)
	return err
}

func (r *OTPRepository) GetByEmail(email string) (string, string, time.Time, error) {
	var code, session string
	var expired time.Time

	err := r.DB.QueryRow(`
		SELECT code, otp_session, expired_at
		FROM email_verification
		WHERE email = ?
	`, email).Scan(&code, &session, &expired)

	return code, session, expired, err
}
