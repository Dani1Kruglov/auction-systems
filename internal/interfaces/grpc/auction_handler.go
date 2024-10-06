package grpc

import (
	"context"
	"time"

	"auction-systems/api/protobufs"
	"auction-systems/internal/application"
	"auction-systems/internal/domain"
	"auction-systems/traits"

	"go.uber.org/zap"
)

type AuctionHandler struct {
	logger   *zap.Logger
	services *application.Services
}

func NewAuctionHandler(logger *zap.Logger, services *application.Services) *AuctionHandler {
	return &AuctionHandler{
		logger:   logger,
		services: services,
	}
}

func (h *AuctionHandler) RegisterClient(ctx context.Context, request *protobufs.RegisterClientRequest) (*protobufs.RegisterClientResponse, error) {
	user := domain.User{
		Name:            request.Name,
		Role:            request.Role,
		Email:           request.Email,
		Password:        request.Password,
		NotificationUrl: request.NotificationUrl,
	}

	err := user.Validate()
	if err != nil {
		h.logger.Error("error user validation", zap.String("file", traits.GetFile()), zap.Int("line", traits.GetLine()), zap.Error(err))
		return nil, err
	}

	_, err = h.services.UserService.RegisterClient(&user)
	if err != nil {
		h.logger.Error("error register client", zap.String("file", traits.GetFile()), zap.Int("line", traits.GetLine()))
		return nil, err
	}

	return &protobufs.RegisterClientResponse{Success: true}, nil
}

func (h *AuctionHandler) CreateLot(ctx context.Context, request *protobufs.CreateLotRequest) (*protobufs.CreateLotResponse, error) {
	lot := &domain.Lot{
		SellerID:    int(request.SellerId),
		Title:       request.ItemName,
		Description: request.Description,
		StartingBid: request.StartingPrice,
		IsClosed:    false,
	}

	if err := lot.Validate(); err != nil {
		h.logger.Error("error validation", zap.String("file", traits.GetFile()), zap.Int("line", traits.GetLine()))
		return nil, err
	}

	lotID, err := h.services.LotService.CreateLot(lot)
	if err != nil {
		h.logger.Error("error creating lot", zap.String("file", traits.GetFile()), zap.Error(err))
		return nil, err
	}
	return &protobufs.CreateLotResponse{LotId: float64(lotID)}, nil
}

func (h *AuctionHandler) CreateAuction(ctx context.Context, request *protobufs.CreateAuctionRequest) (*protobufs.CreateAuctionResponse, error) {

	startTime, err := time.Parse(time.RFC3339, request.StartTime)
	if err != nil {
		h.logger.Error("error parsing start time", zap.String("file", traits.GetFile()), zap.Error(err))
		return nil, err
	}

	endTime, err := time.Parse(time.RFC3339, request.EndTime)
	if err != nil {
		h.logger.Error("error parsing time", zap.String("file", traits.GetFile()), zap.Error(err))
		return nil, err
	}

	auction := &domain.Auction{
		LotID:     int(request.LotId),
		Status:    "wait",
		MinStep:   request.MinStep,
		StartTime: startTime,
		EndTime:   endTime,
	}

	if err := auction.Validate(); err != nil {
		h.logger.Error("error validation", zap.String("file", traits.GetFile()), zap.Int("line", traits.GetLine()))
		return nil, err
	}

	auctionId, err := h.services.AuctionService.CreateAuction(auction)
	if err != nil {
		h.logger.Error("error create auction", zap.String("file", traits.GetFile()), zap.Error(err))
		return nil, err
	}
	return &protobufs.CreateAuctionResponse{AuctionId: float64(auctionId)}, nil
}

func (h *AuctionHandler) PlaceBid(ctx context.Context, request *protobufs.PlaceBidRequest) (*protobufs.PlaceBidResponse, error) {

	bid := &domain.Bid{
		AuctionId: int(request.AuctionId),
		BuyerId:   int(request.BuyerId),
		Amount:    request.BidAmount,
	}

	err := h.services.BidService.PlaceBid(bid)
	if err != nil {
		h.logger.Error("error placing bid", zap.String("file", traits.GetFile()), zap.Int("line", traits.GetLine()), zap.Error(err))
		return nil, err
	}

	return &protobufs.PlaceBidResponse{Success: true}, nil
}

func (h *AuctionHandler) AddBalance(ctx context.Context, request *protobufs.AddBalanceRequest) (*protobufs.AddBalanceResponse, error) {
	err := h.services.UserService.UpdateUserBalance(int(request.UserId), request.Amount)
	if err != nil {
		h.logger.Error("error updating user balance", zap.String("file", traits.GetFile()), zap.Int("line", traits.GetLine()), zap.Error(err))
	}
	return &protobufs.AddBalanceResponse{Success: true}, nil
}
