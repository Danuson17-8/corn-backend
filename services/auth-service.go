package services

import (
	"database/sql"
	"errors"

	"github.com/Danuson17-8/corn-backend/repositories"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmailExists     = errors.New("email already exists")
	ErrEmailNotFound   = errors.New("email not found")
	ErrInvalidPassword = errors.New("invalid password")
)

type AuthService struct {
	AccountRepo *repositories.AccountRepository
}

func (s *AuthService) Register(email, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	err = s.AccountRepo.Create(email, string(hash))
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
			return ErrEmailExists
		}
		return err
	}

	return nil
}

func (s *AuthService) Login(email, password string) error {
	hash, err := s.AccountRepo.GetPasswordHash(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrEmailNotFound
		}
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return ErrInvalidPassword
	}

	return nil
}
