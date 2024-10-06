package application

import (
	"fmt"

	"auction-systems/internal/domain"
	"auction-systems/internal/domain/repositories"
)

type LotService struct {
	repo *repositories.Repositories
}

func NewLotService(repo *repositories.Repositories) *LotService {
	return &LotService{
		repo: repo,
	}
}

func (ls *LotService) CreateLot(lot *domain.Lot) (int, error) {
	_, err := ls.repo.UserRepository.GetUserById(lot.SellerID, domain.SellerRole)
	if err != nil {
		return 0, fmt.Errorf("failed to get user by ID %d with role %s: %w", lot.SellerID, domain.SellerRole, err)
	}

	lotId, err := ls.repo.LotRepository.Create(lot)
	if err != nil {
		return 0, fmt.Errorf("failed to create lot: %w", err)
	}

	return lotId, nil
}
