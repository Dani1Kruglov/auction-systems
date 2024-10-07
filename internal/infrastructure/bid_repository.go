package infrastructure

import (
	"database/sql"
	"errors"
	"time"

	"auction-systems/internal/domain"
)

type BidRepository struct {
	db *sql.DB
}

func NewBidRepository(db *sql.DB) *BidRepository {
	return &BidRepository{db: db}
}

func (br *BidRepository) GetBidByUserIdAndAuctionId(userId, auctionId int) (*domain.Bid, error) {
	var bid domain.Bid

	err := br.db.QueryRow(`SELECT id, user_id, auction_id, bid_amount FROM bids WHERE user_id = $1 AND auction_id = $2;`, userId, auctionId).
		Scan(&bid.Id, &bid.BuyerId, &bid.AuctionId, &bid.Amount)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if bid.Id != 0 {
		return &bid, nil
	}

	return nil, nil
}

func (br *BidRepository) UpdateBid(bid *domain.Bid) error {
	_, err := br.db.Exec(`UPDATE bids SET bid_amount = $2, updated_at = $3 WHERE id = $1`,
		bid.Id, bid.Amount, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (br *BidRepository) GetBidsByAuctionId(auctionId int) ([]*domain.Bid, *domain.Bid, error) {
	rows, err := br.db.Query(`SELECT id, user_id, auction_id, bid_amount FROM bids WHERE auction_id = $1 ORDER BY bid_amount DESC;`, auctionId)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	var (
		bids        []*domain.Bid
		maxBid      *domain.Bid
		maxBidIndex int
	)

	for rows.Next() {
		var bid domain.Bid
		err := rows.Scan(&bid.Id, &bid.BuyerId, &bid.AuctionId, &bid.Amount)
		if err != nil {
			return nil, nil, err
		}

		bids = append(bids, &bid)

		if maxBid == nil || bid.Amount > maxBid.Amount {
			maxBid = &bid
			maxBidIndex = len(bids) - 1
		}

	}

	if maxBid != nil {
		bids = append(bids[:maxBidIndex], bids[maxBidIndex+1:]...)
	}

	return bids, maxBid, nil
}

func (br *BidRepository) AddBid(bid *domain.Bid) error {
	_, err := br.db.Exec(`INSERT INTO bids (user_id, auction_id, bid_amount) VALUES ($1, $2, $3)`,
		bid.BuyerId, bid.AuctionId, bid.Amount)
	if err != nil {
		return err
	}

	return nil
}
