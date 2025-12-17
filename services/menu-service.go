package services

import (
	"github.com/Danuson17-8/corn-backend/models"
	"github.com/Danuson17-8/corn-backend/repositories"
)

type MenuService struct {
	Repo *repositories.MenuRepository
}

func (s *MenuService) GetMenu() ([]models.CornMenu, error) {
	return s.Repo.GetAll()
}
