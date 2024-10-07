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

	existsBid, err := bs.repo.BidRepository.GetBidByUserIdAndAuctionId(bid.BuyerId, bid.AuctionId)
	if err != nil {
		return fmt.Errorf("failed to check if bid exists for user ID %d in auction ID %d: %w", bid.BuyerId, bid.AuctionId, err)
	}

	var newBid float64

	if existsBid != nil {
		existsBid.Amount = existsBid.Amount + bid.Amount
		if existsBid.Amount-auction.MaxBid < auction.MinStep {
			return errors.New("bid is too low")
		}
		newBid = existsBid.Amount
		err = bs.repo.BidRepository.UpdateBid(existsBid)

	} else {
		if auction.MaxBid != 0 {
			if bid.Amount < auction.MaxBid+auction.MinStep {
				return errors.New("bid is too low")
			}
		} else {
			if bid.Amount < auction.Lot.StartingBid {
				return errors.New("bid is too low")
			}
		}
		newBid = bid.Amount
		err = bs.repo.BidRepository.AddBid(bid)
		if err != nil {
			return fmt.Errorf("failed to add bid for user ID %d in auction ID %d: %w", bid.BuyerId, bid.AuctionId, err)
		}
	}

	err = bs.repo.AuctionRepository.UpdateWinnerIdAndMaxBid(auction.Id, bid.BuyerId, newBid)
	if err != nil {
		return fmt.Errorf("failed to update auction max bid on user Id %d: %w", bid.BuyerId, err)
	}

	_, err = bs.repo.UserRepository.UpdateUserBalance(bid.BuyerId, bid.Amount, domain.MinusSymbol)
	if err != nil {
		return fmt.Errorf("failed to update balance for user ID %d: %w", bid.BuyerId, err)
	}

	return nil
}
