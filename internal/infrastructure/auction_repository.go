package infrastructure

import (
	"database/sql"
	"fmt"

	"auction-systems/internal/domain"
)

type AuctionRepository struct {
	db *sql.DB
}

func NewAuctionRepository(db *sql.DB) *AuctionRepository {
	return &AuctionRepository{db: db}
}

func (ar *AuctionRepository) GetActiveAuctionById(auctionId int) (*domain.Auction, error) {
	var auction domain.Auction
	err := ar.db.QueryRow(`SELECT a.id, a.lot_id, a.status, a.max_bid, a.min_step, l.starting_price FROM auctions a JOIN lots l ON a.lot_id = l.id WHERE a.id = $1 AND a.status = 'open';`, auctionId).
		Scan(&auction.Id, &auction.LotID, &auction.Status, &auction.MaxBid, &auction.MinStep, &auction.Lot.StartingBid)
	if err != nil {
		return nil, err
	}
	return &auction, nil
}

func (ar *AuctionRepository) GetActiveAuctions() ([]*domain.Auction, error) {
	rows, err := ar.db.Query(`SELECT id, lot_id, status, min_step, start_time, end_time FROM auctions WHERE (status = 'wait' OR status = 'open')`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var auctions []*domain.Auction

	for rows.Next() {
		var auction domain.Auction
		err := rows.Scan(&auction.Id, &auction.LotID, &auction.Status, &auction.MinStep, &auction.StartTime, &auction.EndTime)
		if err != nil {
			return nil, err
		}
		auctions = append(auctions, &auction)
	}
	return auctions, nil
}

func (ar *AuctionRepository) Create(auction *domain.Auction) (int, error) {
	query := `INSERT INTO auctions (lot_id, status, min_step, start_time, end_time) 
        VALUES ($1, $2, $3, $4, $5) 
        RETURNING id`

	err := ar.db.QueryRow(query, auction.LotID, auction.Status, auction.MinStep, auction.StartTime, auction.EndTime).Scan(&auction.Id)
	if err != nil {
		return 0, err
	}
	return auction.Id, nil
}

func (ar *AuctionRepository) UpdateStatus(auctionId int, newStatus string) error {

	_, err := ar.db.Exec(`UPDATE auctions SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`, newStatus, auctionId)
	if err != nil {
		return fmt.Errorf("failed to update auction status: %w", err)
	}
	return nil
}

func (ar *AuctionRepository) UpdateWinnerIdAndMaxBid(auctionId int, winnerId int, maxBid float64) error {

	_, err := ar.db.Exec(`UPDATE auctions SET winner_id = $1, max_bid = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3 AND $2 > max_bid;`, winnerId, maxBid, auctionId)
	if err != nil {
		return fmt.Errorf("failed to update auction winner id: %w", err)
	}
	return nil
}
