package services

import (
	"github.com/Danuson17-8/corn-backend/models"
	"github.com/Danuson17-8/corn-backend/repositories"
)

type PromotionService struct {
	Repo *repositories.PromotionRepository
}

func (s *PromotionService) GetActivePromotions() ([]models.Promotion, error) {
	return s.Repo.GetActive()
}
