package application

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"auction-systems/internal/domain"
	"auction-systems/internal/domain/repositories"
	"auction-systems/traits"

	"go.uber.org/zap"
)

type AuctionService struct {
	logger *zap.Logger
	repo   *repositories.Repositories
}

func NewAuctionService(logger *zap.Logger, repo *repositories.Repositories) (*AuctionService, error) {
	auctionService := &AuctionService{
		logger: logger,
		repo:   repo,
	}

	go auctionService.AuctionsListener()

	return auctionService, nil
}

func (as *AuctionService) AuctionsListener() {
	activeAuctions, err := as.repo.AuctionRepository.GetActiveAuctions()
	if err != nil {
		as.logger.Fatal("error get active auctions", zap.Error(err))
	}

	for _, auction := range activeAuctions {
		go func(auction *domain.Auction) { as.AuctionSchedule(auction) }(auction)
	}

}

func (as *AuctionService) CreateAuction(auction *domain.Auction) (int, error) {
	lot, err := as.repo.LotRepository.GetLotByID(auction.LotID)
	if err != nil {
		return 0, fmt.Errorf("failed to get lot by ID %d: %w", auction.LotID, err)
	}
	if lot == nil {
		return 0, errors.New("error: lot not found")
	}

	auctionId, err := as.repo.AuctionRepository.Create(auction)
	if err != nil {
		return 0, fmt.Errorf("failed to create auction: %w", err)
	}

	go func(auction *domain.Auction) { as.AuctionSchedule(auction) }(auction)

	return auctionId, nil
}

func (as *AuctionService) AuctionSchedule(auction *domain.Auction) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			currentTime := time.Now()

			if !currentTime.After(auction.StartTime) {
				continue
			}

			buyers := as.repo.UserRepository.GetBuyers()

			if auction.Status == "wait" {
				go func(buyers map[int]*domain.User) {
					for _, buyer := range buyers {
						err := traits.SendNotification(buyer.NotificationUrl, traits.Response{
							Status:      "notification",
							Code:        http.StatusOK,
							Description: fmt.Sprintf("auction with id %d started", auction.Id),
						})
						if err != nil {
							as.logger.Error("error sending notification", zap.Error(err))
						}
					}
				}(buyers)
			}

			err := as.repo.AuctionRepository.UpdateStatus(auction.Id, "open")
			if err != nil {
				as.logger.Fatal("error update auction status", zap.Error(err))
			}

			for {
				time.Sleep(1 * time.Second)
				currentTime = time.Now()
				if !currentTime.Before(auction.EndTime) {
					fmt.Println("STOP", auction.Id)
					err = as.repo.AuctionRepository.UpdateStatus(auction.Id, "closed")
					if err != nil {
						as.logger.Fatal("error update auction status", zap.Error(err))
					}

					bids, maxBid, err := as.repo.BidRepository.GetBidsByAuctionId(auction.Id)
					if err != nil {
						as.logger.Fatal("error get max bid from repository", zap.Error(err))
					}

					err = as.repo.AuctionRepository.UpdateWinnerId(auction.Id, maxBid.BuyerId)
					if err != nil {
						as.logger.Fatal("error update auction winner id", zap.Error(err))
					}

					if buyer, exists := buyers[maxBid.BuyerId]; exists {
						err = traits.SendNotification(buyer.NotificationUrl, traits.Response{
							Status:      "notification",
							Code:        http.StatusOK,
							Description: fmt.Sprintf("You have won the auction with ID %d. Your bid amount is %f.", auction.Id, maxBid.Amount),
						})
						if err != nil {
							as.logger.Error("error sending notification", zap.Error(err))
						}
					}

					for _, bid := range bids {
						_, err = as.repo.UserRepository.UpdateUserBalance(bid.BuyerId, bid.Amount, domain.PlusSymbol)
						if err != nil {
							as.logger.Fatal("error update user balance", zap.Error(err))
						}
					}
					return
				}
			}
		}
	}
}
