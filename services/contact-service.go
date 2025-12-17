package services

import (
	"errors"

	"github.com/Danuson17-8/corn-backend/models"
	"github.com/Danuson17-8/corn-backend/repositories"
)

var ErrContactInvalid = errors.New("contact data invalid")

type ContactService struct {
	Repo *repositories.ContactRepository
}

func (s *ContactService) Create(contact *models.ContactUser) error {
	if contact.Topic == "" ||
		contact.Name == "" ||
		contact.Email == "" ||
		contact.Message == "" {
		return ErrContactInvalid
	}

	return s.Repo.Create(contact)
}
