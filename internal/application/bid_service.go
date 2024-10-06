package application

import (
	"errors"
	"fmt"

	"auction-systems/internal/domain"
	"auction-systems/internal/domain/repositories"
)

type BidService struct {
	repo *repositories.Repositories
}

func NewBidService(repo *repositories.Repositories) *BidService {
	return &BidService{
		repo: repo,
	}
}

func (bs *BidService) PlaceBid(bid *domain.Bid) error {
	_, err := bs.repo.UserRepository.GetUserById(bid.BuyerId, "buyer")
	if err != nil {
		return fmt.Errorf("failed to get user by ID %d: %w", bid.BuyerId, err)
	}

	userBalance, err := bs.repo.UserRepository.GetBalanceByUserId(bid.BuyerId)
	if err != nil {
		return fmt.Errorf("failed to get balance for user ID %d: %w", bid.BuyerId, err)
	}
	if userBalance == 0 || userBalance < bid.Amount {
		return errors.New("there are not enough funds on the balance sheet")
	}

	auction, err := bs.repo.AuctionRepository.GetActiveAuctionById(bid.AuctionId)
	if err != nil {
		return fmt.Errorf("failed to get active auction by ID %d: %w", bid.AuctionId, err)
	}

	if auction.MinStep > bid.Amount {
		return errors.New("the minimum step is too low")
	}

	exists, err := bs.repo.BidRepository.GetBidByUserIdAndAuctionId(bid.BuyerId, bid.AuctionId, bid.Amount)
	if err != nil {
		return fmt.Errorf("failed to check if bid exists for user ID %d in auction ID %d: %w", bid.BuyerId, bid.AuctionId, err)
	}

	if !exists {
		err = bs.repo.BidRepository.AddBid(bid)
		if err != nil {
			return fmt.Errorf("failed to add bid for user ID %d in auction ID %d: %w", bid.BuyerId, bid.AuctionId, err)
		}
	}

	_, err = bs.repo.UserRepository.UpdateUserBalance(bid.BuyerId, bid.Amount, domain.MinusSymbol)
	if err != nil {
		return fmt.Errorf("failed to update balance for user ID %d: %w", bid.BuyerId, err)
	}

	return nil
}
