package domain

import "time"

type Bid struct {
	Id        int       `json:"id"`
	AuctionId int       `json:"auction_id"`
	BuyerId   int       `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
