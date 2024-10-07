package domain

import (
	"errors"
	"time"
)

type Auction struct {
	Id        int       `json:"id"`
	LotID     int       `json:"lot_id"`
	WinnerID  int       `json:"winner_id"`
	Status    string    `json:"status"`
	MaxBid    float64   `json:"max_bid"`
	MinStep   float64   `json:"min_step"`
	Lot       Lot       `json:"lot"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (a *Auction) Validate() error {
	if a.LotID == 0 {
		return errors.New("lot_id is required")
	}
	if a.StartTime.IsZero() {
		return errors.New("start_time is required")
	}
	if a.EndTime.IsZero() {
		return errors.New("end_time is required")
	}
	if a.MinStep == 0 {
		return errors.New("min_step is required")
	}
	return nil
}
