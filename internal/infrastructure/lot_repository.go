package infrastructure

import (
	"database/sql"
	"errors"

	"auction-systems/internal/domain"
)

type LotRepository struct {
	db *sql.DB
}

func NewLotRepository(db *sql.DB) *LotRepository {
	return &LotRepository{db: db}
}

func (lr *LotRepository) Create(lot *domain.Lot) (int, error) {
	query := `INSERT INTO lots (seller_id, title, description, starting_price) 
        VALUES ($1, $2, $3, $4) 
        RETURNING id`

	err := lr.db.QueryRow(query, lot.SellerID, lot.Title, lot.Description, lot.StartingBid).Scan(&lot.Id)
	if err != nil {
		return 0, err
	}
	return lot.Id, nil
}

func (lr *LotRepository) GetLotByID(id int) (*domain.Lot, error) {
	var lot domain.Lot
	query := `SELECT id, seller_id, title, description, starting_price FROM lots WHERE id = $1`

	err := lr.db.QueryRow(query, id).Scan(&lot.Id, &lot.SellerID, &lot.Title, &lot.Description, &lot.StartingBid)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &lot, nil
}

func (lr *LotRepository) Update(lot *domain.Lot) error {
	query := `
        UPDATE lots 
        SET title = $1, description = $2, starting_bid = $3, is_closed = $4 
        WHERE id = $5`

	_, err := lr.db.Exec(query, lot.Title, lot.Description, lot.StartingBid, lot.IsClosed, lot.Id)
	return err
}
