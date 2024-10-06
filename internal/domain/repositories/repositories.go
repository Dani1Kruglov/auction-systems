package repositories

import (
	"auction-systems/internal/domain"
	"auction-systems/internal/infrastructure"
	"database/sql"
)

type IUserRepository interface {
	GetBuyers() map[int]*domain.User
	Create(user *domain.User) (int, error)
	GetUserById(userId int, role string) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	AddUserBalance(userID int, amount float64) error
	UpdateUserBalance(userID int, amount float64, symbol string) (bool, error)
	GetBalanceByUserId(userID int) (float64, error)
}

type ILotRepository interface {
	Create(lot *domain.Lot) (int, error)
	GetLotByID(id int) (*domain.Lot, error)
	Update(lot *domain.Lot) error
}

type IBidRepository interface {
	GetBidsByAuctionId(auctionId int) ([]*domain.Bid, *domain.Bid, error)
	GetBidByUserIdAndAuctionId(userId, auctionId int, amount float64) (bool, error)
	AddBid(bid *domain.Bid) error
}

type IAuctionRepository interface {
	GetActiveAuctions() ([]*domain.Auction, error)
	GetActiveAuctionById(auctionId int) (*domain.Auction, error)
	Create(auction *domain.Auction) (int, error)
	UpdateStatus(auctionId int, newStatus string) error
	UpdateWinnerId(auctionId int, winnerId int) error
}

type Repositories struct {
	UserRepository    IUserRepository
	LotRepository     ILotRepository
	AuctionRepository IAuctionRepository
	BidRepository     IBidRepository
}

func NewRepositories(pgConn *sql.DB) (*Repositories, error) {
	userRepo, err := infrastructure.NewUserRepository(pgConn)
	if err != nil {
		return nil, err
	}
	return &Repositories{
		UserRepository:    userRepo,
		LotRepository:     infrastructure.NewLotRepository(pgConn),
		AuctionRepository: infrastructure.NewAuctionRepository(pgConn),
		BidRepository:     infrastructure.NewBidRepository(pgConn),
	}, nil
}
