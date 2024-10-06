// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.20.3
// source: api/protobufs/auction.proto

package protobufs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	AuctionService_CreateLot_FullMethodName      = "/auction.AuctionService/CreateLot"
	AuctionService_AddBalance_FullMethodName     = "/auction.AuctionService/AddBalance"
	AuctionService_CreateAuction_FullMethodName  = "/auction.AuctionService/CreateAuction"
	AuctionService_PlaceBid_FullMethodName       = "/auction.AuctionService/PlaceBid"
	AuctionService_RegisterClient_FullMethodName = "/auction.AuctionService/RegisterClient"
)

// AuctionServiceClient is the client API for AuctionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuctionServiceClient interface {
	CreateLot(ctx context.Context, in *CreateLotRequest, opts ...grpc.CallOption) (*CreateLotResponse, error)
	AddBalance(ctx context.Context, in *AddBalanceRequest, opts ...grpc.CallOption) (*AddBalanceResponse, error)
	CreateAuction(ctx context.Context, in *CreateAuctionRequest, opts ...grpc.CallOption) (*CreateAuctionResponse, error)
	PlaceBid(ctx context.Context, in *PlaceBidRequest, opts ...grpc.CallOption) (*PlaceBidResponse, error)
	RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*RegisterClientResponse, error)
}

type auctionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuctionServiceClient(cc grpc.ClientConnInterface) AuctionServiceClient {
	return &auctionServiceClient{cc}
}

func (c *auctionServiceClient) CreateLot(ctx context.Context, in *CreateLotRequest, opts ...grpc.CallOption) (*CreateLotResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateLotResponse)
	err := c.cc.Invoke(ctx, AuctionService_CreateLot_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) AddBalance(ctx context.Context, in *AddBalanceRequest, opts ...grpc.CallOption) (*AddBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddBalanceResponse)
	err := c.cc.Invoke(ctx, AuctionService_AddBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) CreateAuction(ctx context.Context, in *CreateAuctionRequest, opts ...grpc.CallOption) (*CreateAuctionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateAuctionResponse)
	err := c.cc.Invoke(ctx, AuctionService_CreateAuction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) PlaceBid(ctx context.Context, in *PlaceBidRequest, opts ...grpc.CallOption) (*PlaceBidResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceBidResponse)
	err := c.cc.Invoke(ctx, AuctionService_PlaceBid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auctionServiceClient) RegisterClient(ctx context.Context, in *RegisterClientRequest, opts ...grpc.CallOption) (*RegisterClientResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterClientResponse)
	err := c.cc.Invoke(ctx, AuctionService_RegisterClient_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuctionServiceServer is the server API for AuctionService service.
// All implementations must embed UnimplementedAuctionServiceServer
// for forward compatibility.
type AuctionServiceServer interface {
	CreateLot(context.Context, *CreateLotRequest) (*CreateLotResponse, error)
	AddBalance(context.Context, *AddBalanceRequest) (*AddBalanceResponse, error)
	CreateAuction(context.Context, *CreateAuctionRequest) (*CreateAuctionResponse, error)
	PlaceBid(context.Context, *PlaceBidRequest) (*PlaceBidResponse, error)
	RegisterClient(context.Context, *RegisterClientRequest) (*RegisterClientResponse, error)
	//mustEmbedUnimplementedAuctionServiceServer()
}

// UnimplementedAuctionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuctionServiceServer struct{}

func (UnimplementedAuctionServiceServer) CreateLot(context.Context, *CreateLotRequest) (*CreateLotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLot not implemented")
}
func (UnimplementedAuctionServiceServer) AddBalance(context.Context, *AddBalanceRequest) (*AddBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBalance not implemented")
}
func (UnimplementedAuctionServiceServer) CreateAuction(context.Context, *CreateAuctionRequest) (*CreateAuctionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAuction not implemented")
}
func (UnimplementedAuctionServiceServer) PlaceBid(context.Context, *PlaceBidRequest) (*PlaceBidResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceBid not implemented")
}
func (UnimplementedAuctionServiceServer) RegisterClient(context.Context, *RegisterClientRequest) (*RegisterClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterClient not implemented")
}
func (UnimplementedAuctionServiceServer) mustEmbedUnimplementedAuctionServiceServer() {}
func (UnimplementedAuctionServiceServer) testEmbeddedByValue()                        {}

// UnsafeAuctionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuctionServiceServer will
// result in compilation errors.
type UnsafeAuctionServiceServer interface {
	mustEmbedUnimplementedAuctionServiceServer()
}

func RegisterAuctionServiceServer(s grpc.ServiceRegistrar, srv AuctionServiceServer) {
	// If the following call pancis, it indicates UnimplementedAuctionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AuctionService_ServiceDesc, srv)
}

func _AuctionService_CreateLot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateLotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).CreateLot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_CreateLot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).CreateLot(ctx, req.(*CreateLotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_AddBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).AddBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_AddBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).AddBalance(ctx, req.(*AddBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_CreateAuction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAuctionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).CreateAuction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_CreateAuction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).CreateAuction(ctx, req.(*CreateAuctionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_PlaceBid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceBidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).PlaceBid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_PlaceBid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).PlaceBid(ctx, req.(*PlaceBidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuctionService_RegisterClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterClientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuctionServiceServer).RegisterClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuctionService_RegisterClient_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuctionServiceServer).RegisterClient(ctx, req.(*RegisterClientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuctionService_ServiceDesc is the grpc.ServiceDesc for AuctionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuctionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auction.AuctionService",
	HandlerType: (*AuctionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLot",
			Handler:    _AuctionService_CreateLot_Handler,
		},
		{
			MethodName: "AddBalance",
			Handler:    _AuctionService_AddBalance_Handler,
		},
		{
			MethodName: "CreateAuction",
			Handler:    _AuctionService_CreateAuction_Handler,
		},
		{
			MethodName: "PlaceBid",
			Handler:    _AuctionService_PlaceBid_Handler,
		},
		{
			MethodName: "RegisterClient",
			Handler:    _AuctionService_RegisterClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/protobufs/auction.proto",
}
