package domain

import "errors"

type Lot struct {
	Id          int
	SellerID    int
	Title       string
	Description string
	StartingBid float64
	IsClosed    bool
}

func (l *Lot) Validate() error {
	if l.SellerID == 0 {
		return errors.New("seller_id is required")
	}
	if l.Title == "" {
		return errors.New("title is required")
	}
	if l.Description == "" {
		return errors.New("description is required")
	}
	if l.StartingBid <= 0 {
		return errors.New("starting bid must be greater than zero")
	}
	return nil
}
