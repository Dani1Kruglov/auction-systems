package application

import (
	"go.uber.org/zap"

	"auction-systems/internal/domain"
	"auction-systems/internal/domain/repositories"
)

type IUserService interface {
	RegisterClient(user *domain.User) (int, error)
	UpdateUserBalance(userId int, amount float64) error
}

type IAuctionService interface {
	CreateAuction(auction *domain.Auction) (int, error)
}

type ILotService interface {
	CreateLot(lot *domain.Lot) (int, error)
}

type IBidService interface {
	PlaceBid(bid *domain.Bid) error
}

type Services struct {
	UserService    IUserService
	AuctionService IAuctionService
	LotService     ILotService
	BidService     IBidService
}

func NewServices(logger *zap.Logger, repo *repositories.Repositories) (*Services, error) {
	auctionService, err := NewAuctionService(logger, repo)
	if err != nil {
		return nil, err
	}
	return &Services{
		UserService:    NewUserService(repo),
		AuctionService: auctionService,
		LotService:     NewLotService(repo),
		BidService:     NewBidService(repo),
	}, nil
}
