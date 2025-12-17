package services

import (
	"database/sql"
	"errors"

	"github.com/Danuson17-8/corn-backend/models"
	"github.com/Danuson17-8/corn-backend/repositories"
)

var ErrUserNotFound = errors.New("user not found")

type ProfileService struct {
	UserRepo *repositories.UserRepository
}

func (s *ProfileService) GetProfileByEmail(email string) (*models.User, error) {
	user, err := s.UserRepo.GetByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}
