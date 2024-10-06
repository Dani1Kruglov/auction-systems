// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: api/protobufs/auction.proto

package protobufs

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId float64 `protobuf:"fixed64,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *AddBalanceRequest) Reset() {
	*x = AddBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBalanceRequest) ProtoMessage() {}

func (x *AddBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBalanceRequest.ProtoReflect.Descriptor instead.
func (*AddBalanceRequest) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{0}
}

func (x *AddBalanceRequest) GetUserId() float64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AddBalanceRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AddBalanceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AddBalanceResponse) Reset() {
	*x = AddBalanceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBalanceResponse) ProtoMessage() {}

func (x *AddBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBalanceResponse.ProtoReflect.Descriptor instead.
func (*AddBalanceResponse) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{1}
}

func (x *AddBalanceResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CreateLotRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SellerId      float64 `protobuf:"fixed64,1,opt,name=seller_id,json=sellerId,proto3" json:"seller_id,omitempty"`
	ItemName      string  `protobuf:"bytes,2,opt,name=item_name,json=itemName,proto3" json:"item_name,omitempty"`
	Description   string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	StartingPrice float64 `protobuf:"fixed64,4,opt,name=starting_price,json=startingPrice,proto3" json:"starting_price,omitempty"`
}

func (x *CreateLotRequest) Reset() {
	*x = CreateLotRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLotRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLotRequest) ProtoMessage() {}

func (x *CreateLotRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLotRequest.ProtoReflect.Descriptor instead.
func (*CreateLotRequest) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLotRequest) GetSellerId() float64 {
	if x != nil {
		return x.SellerId
	}
	return 0
}

func (x *CreateLotRequest) GetItemName() string {
	if x != nil {
		return x.ItemName
	}
	return ""
}

func (x *CreateLotRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateLotRequest) GetStartingPrice() float64 {
	if x != nil {
		return x.StartingPrice
	}
	return 0
}

type CreateLotResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LotId float64 `protobuf:"fixed64,1,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
}

func (x *CreateLotResponse) Reset() {
	*x = CreateLotResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLotResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLotResponse) ProtoMessage() {}

func (x *CreateLotResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLotResponse.ProtoReflect.Descriptor instead.
func (*CreateLotResponse) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLotResponse) GetLotId() float64 {
	if x != nil {
		return x.LotId
	}
	return 0
}

type CreateAuctionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LotId     float64 `protobuf:"fixed64,1,opt,name=lot_id,json=lotId,proto3" json:"lot_id,omitempty"`
	MinStep   float64 `protobuf:"fixed64,2,opt,name=min_step,json=minStep,proto3" json:"min_step,omitempty"`
	StartTime string  `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime   string  `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
}

func (x *CreateAuctionRequest) Reset() {
	*x = CreateAuctionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuctionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuctionRequest) ProtoMessage() {}

func (x *CreateAuctionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuctionRequest.ProtoReflect.Descriptor instead.
func (*CreateAuctionRequest) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAuctionRequest) GetLotId() float64 {
	if x != nil {
		return x.LotId
	}
	return 0
}

func (x *CreateAuctionRequest) GetMinStep() float64 {
	if x != nil {
		return x.MinStep
	}
	return 0
}

func (x *CreateAuctionRequest) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *CreateAuctionRequest) GetEndTime() string {
	if x != nil {
		return x.EndTime
	}
	return ""
}

type CreateAuctionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId float64 `protobuf:"fixed64,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
}

func (x *CreateAuctionResponse) Reset() {
	*x = CreateAuctionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAuctionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAuctionResponse) ProtoMessage() {}

func (x *CreateAuctionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAuctionResponse.ProtoReflect.Descriptor instead.
func (*CreateAuctionResponse) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAuctionResponse) GetAuctionId() float64 {
	if x != nil {
		return x.AuctionId
	}
	return 0
}

type PlaceBidRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuctionId float64 `protobuf:"fixed64,1,opt,name=auction_id,json=auctionId,proto3" json:"auction_id,omitempty"`
	BuyerId   float64 `protobuf:"fixed64,2,opt,name=buyer_id,json=buyerId,proto3" json:"buyer_id,omitempty"`
	BidAmount float64 `protobuf:"fixed64,3,opt,name=bid_amount,json=bidAmount,proto3" json:"bid_amount,omitempty"`
}

func (x *PlaceBidRequest) Reset() {
	*x = PlaceBidRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceBidRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceBidRequest) ProtoMessage() {}

func (x *PlaceBidRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceBidRequest.ProtoReflect.Descriptor instead.
func (*PlaceBidRequest) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{6}
}

func (x *PlaceBidRequest) GetAuctionId() float64 {
	if x != nil {
		return x.AuctionId
	}
	return 0
}

func (x *PlaceBidRequest) GetBuyerId() float64 {
	if x != nil {
		return x.BuyerId
	}
	return 0
}

func (x *PlaceBidRequest) GetBidAmount() float64 {
	if x != nil {
		return x.BidAmount
	}
	return 0
}

type PlaceBidResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PlaceBidResponse) Reset() {
	*x = PlaceBidResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlaceBidResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlaceBidResponse) ProtoMessage() {}

func (x *PlaceBidResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlaceBidResponse.ProtoReflect.Descriptor instead.
func (*PlaceBidResponse) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{7}
}

func (x *PlaceBidResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type RegisterClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name            string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email           string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Password        string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Role            string `protobuf:"bytes,4,opt,name=role,proto3" json:"role,omitempty"`
	NotificationUrl string `protobuf:"bytes,5,opt,name=notification_url,json=notificationUrl,proto3" json:"notification_url,omitempty"`
}

func (x *RegisterClientRequest) Reset() {
	*x = RegisterClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientRequest) ProtoMessage() {}

func (x *RegisterClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientRequest.ProtoReflect.Descriptor instead.
func (*RegisterClientRequest) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{8}
}

func (x *RegisterClientRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterClientRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegisterClientRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterClientRequest) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *RegisterClientRequest) GetNotificationUrl() string {
	if x != nil {
		return x.NotificationUrl
	}
	return ""
}

type RegisterClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RegisterClientResponse) Reset() {
	*x = RegisterClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_protobufs_auction_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterClientResponse) ProtoMessage() {}

func (x *RegisterClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_protobufs_auction_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterClientResponse.ProtoReflect.Descriptor instead.
func (*RegisterClientResponse) Descriptor() ([]byte, []int) {
	return file_api_protobufs_auction_proto_rawDescGZIP(), []int{9}
}

func (x *RegisterClientResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_api_protobufs_auction_proto protoreflect.FileDescriptor

var file_api_protobufs_auction_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x2f,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x44, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x42,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2e,
	0x0a, 0x12, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x95,
	0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x74, 0x65, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0x2a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6c,
	0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x6f, 0x74,
	0x49, 0x64, 0x22, 0x82, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x6c,
	0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x6c, 0x6f, 0x74,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74, 0x65, 0x70, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x36, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22,
	0x6a, 0x0a, 0x0f, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x75, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x07, 0x62, 0x75, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x62, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x09, 0x62, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2c, 0x0a, 0x10, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x15, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x29, 0x0a,
	0x10, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x32, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xad, 0x04, 0x0a,
	0x0e, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5e, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x74, 0x12, 0x19, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f,
	0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x74, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12,
	0x6d, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x1a, 0x2e,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01,
	0x2a, 0x22, 0x1b, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x6e,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x6a,
	0x0a, 0x08, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x12, 0x18, 0x2e, 0x61, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50,
	0x6c, 0x61, 0x63, 0x65, 0x42, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x3a, 0x01, 0x2a, 0x22, 0x1e, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x7b, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x62, 0x69, 0x64, 0x73, 0x12, 0x70, 0x0a, 0x0e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x42, 0x19, 0x5a, 0x17,
	0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_protobufs_auction_proto_rawDescOnce sync.Once
	file_api_protobufs_auction_proto_rawDescData = file_api_protobufs_auction_proto_rawDesc
)

func file_api_protobufs_auction_proto_rawDescGZIP() []byte {
	file_api_protobufs_auction_proto_rawDescOnce.Do(func() {
		file_api_protobufs_auction_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_protobufs_auction_proto_rawDescData)
	})
	return file_api_protobufs_auction_proto_rawDescData
}

var file_api_protobufs_auction_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_protobufs_auction_proto_goTypes = []any{
	(*AddBalanceRequest)(nil),      // 0: auction.AddBalanceRequest
	(*AddBalanceResponse)(nil),     // 1: auction.AddBalanceResponse
	(*CreateLotRequest)(nil),       // 2: auction.CreateLotRequest
	(*CreateLotResponse)(nil),      // 3: auction.CreateLotResponse
	(*CreateAuctionRequest)(nil),   // 4: auction.CreateAuctionRequest
	(*CreateAuctionResponse)(nil),  // 5: auction.CreateAuctionResponse
	(*PlaceBidRequest)(nil),        // 6: auction.PlaceBidRequest
	(*PlaceBidResponse)(nil),       // 7: auction.PlaceBidResponse
	(*RegisterClientRequest)(nil),  // 8: auction.RegisterClientRequest
	(*RegisterClientResponse)(nil), // 9: auction.RegisterClientResponse
}
var file_api_protobufs_auction_proto_depIdxs = []int32{
	2, // 0: auction.AuctionService.CreateLot:input_type -> auction.CreateLotRequest
	0, // 1: auction.AuctionService.AddBalance:input_type -> auction.AddBalanceRequest
	4, // 2: auction.AuctionService.CreateAuction:input_type -> auction.CreateAuctionRequest
	6, // 3: auction.AuctionService.PlaceBid:input_type -> auction.PlaceBidRequest
	8, // 4: auction.AuctionService.RegisterClient:input_type -> auction.RegisterClientRequest
	3, // 5: auction.AuctionService.CreateLot:output_type -> auction.CreateLotResponse
	1, // 6: auction.AuctionService.AddBalance:output_type -> auction.AddBalanceResponse
	5, // 7: auction.AuctionService.CreateAuction:output_type -> auction.CreateAuctionResponse
	7, // 8: auction.AuctionService.PlaceBid:output_type -> auction.PlaceBidResponse
	9, // 9: auction.AuctionService.RegisterClient:output_type -> auction.RegisterClientResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_protobufs_auction_proto_init() }
func file_api_protobufs_auction_proto_init() {
	if File_api_protobufs_auction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_protobufs_auction_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddBalanceRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protobufs_auction_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddBalanceResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protobufs_auction_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateLotRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protobufs_auction_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateLotResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protobufs_auction_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAuctionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protobufs_auction_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateAuctionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protobufs_auction_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*PlaceBidRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protobufs_auction_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PlaceBidResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protobufs_auction_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterClientRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_protobufs_auction_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*RegisterClientResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_protobufs_auction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_protobufs_auction_proto_goTypes,
		DependencyIndexes: file_api_protobufs_auction_proto_depIdxs,
		MessageInfos:      file_api_protobufs_auction_proto_msgTypes,
	}.Build()
	File_api_protobufs_auction_proto = out.File
	file_api_protobufs_auction_proto_rawDesc = nil
	file_api_protobufs_auction_proto_goTypes = nil
	file_api_protobufs_auction_proto_depIdxs = nil
}
