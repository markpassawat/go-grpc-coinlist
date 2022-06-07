// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: proto/coinlist/coin_list.proto

package go_grpc_coinlist

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoinListClient is the client API for CoinList service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoinListClient interface {
	GetCoins(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CoinList_GetCoinsClient, error)
	GetCoin(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CoinInfo, error)
	CreateCoins(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error)
	UpdateCoins(ctx context.Context, in *CoinInfo, opts ...grpc.CallOption) (*Status, error)
	DeleteCoin(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error)
	SearchCoins(ctx context.Context, in *InputText, opts ...grpc.CallOption) (CoinList_SearchCoinsClient, error)
}

type coinListClient struct {
	cc grpc.ClientConnInterface
}

func NewCoinListClient(cc grpc.ClientConnInterface) CoinListClient {
	return &coinListClient{cc}
}

func (c *coinListClient) GetCoins(ctx context.Context, in *Empty, opts ...grpc.CallOption) (CoinList_GetCoinsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoinList_ServiceDesc.Streams[0], "/coinlist.CoinList/GetCoins", opts...)
	if err != nil {
		return nil, err
	}
	x := &coinListGetCoinsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CoinList_GetCoinsClient interface {
	Recv() (*CoinInfo, error)
	grpc.ClientStream
}

type coinListGetCoinsClient struct {
	grpc.ClientStream
}

func (x *coinListGetCoinsClient) Recv() (*CoinInfo, error) {
	m := new(CoinInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *coinListClient) GetCoin(ctx context.Context, in *Id, opts ...grpc.CallOption) (*CoinInfo, error) {
	out := new(CoinInfo)
	err := c.cc.Invoke(ctx, "/coinlist.CoinList/GetCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinListClient) CreateCoins(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/coinlist.CoinList/CreateCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinListClient) UpdateCoins(ctx context.Context, in *CoinInfo, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/coinlist.CoinList/UpdateCoins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinListClient) DeleteCoin(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/coinlist.CoinList/DeleteCoin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coinListClient) SearchCoins(ctx context.Context, in *InputText, opts ...grpc.CallOption) (CoinList_SearchCoinsClient, error) {
	stream, err := c.cc.NewStream(ctx, &CoinList_ServiceDesc.Streams[1], "/coinlist.CoinList/SearchCoins", opts...)
	if err != nil {
		return nil, err
	}
	x := &coinListSearchCoinsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CoinList_SearchCoinsClient interface {
	Recv() (*CoinInfo, error)
	grpc.ClientStream
}

type coinListSearchCoinsClient struct {
	grpc.ClientStream
}

func (x *coinListSearchCoinsClient) Recv() (*CoinInfo, error) {
	m := new(CoinInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CoinListServer is the server API for CoinList service.
// All implementations should embed UnimplementedCoinListServer
// for forward compatibility
type CoinListServer interface {
	GetCoins(*Empty, CoinList_GetCoinsServer) error
	GetCoin(context.Context, *Id) (*CoinInfo, error)
	CreateCoins(context.Context, *Id) (*Status, error)
	UpdateCoins(context.Context, *CoinInfo) (*Status, error)
	DeleteCoin(context.Context, *Id) (*Status, error)
	SearchCoins(*InputText, CoinList_SearchCoinsServer) error
}

// UnimplementedCoinListServer should be embedded to have forward compatible implementations.
type UnimplementedCoinListServer struct {
}

func (UnimplementedCoinListServer) GetCoins(*Empty, CoinList_GetCoinsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetCoins not implemented")
}
func (UnimplementedCoinListServer) GetCoin(context.Context, *Id) (*CoinInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCoin not implemented")
}
func (UnimplementedCoinListServer) CreateCoins(context.Context, *Id) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoins not implemented")
}
func (UnimplementedCoinListServer) UpdateCoins(context.Context, *CoinInfo) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCoins not implemented")
}
func (UnimplementedCoinListServer) DeleteCoin(context.Context, *Id) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoin not implemented")
}
func (UnimplementedCoinListServer) SearchCoins(*InputText, CoinList_SearchCoinsServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchCoins not implemented")
}

// UnsafeCoinListServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoinListServer will
// result in compilation errors.
type UnsafeCoinListServer interface {
	mustEmbedUnimplementedCoinListServer()
}

func RegisterCoinListServer(s grpc.ServiceRegistrar, srv CoinListServer) {
	s.RegisterService(&CoinList_ServiceDesc, srv)
}

func _CoinList_GetCoins_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoinListServer).GetCoins(m, &coinListGetCoinsServer{stream})
}

type CoinList_GetCoinsServer interface {
	Send(*CoinInfo) error
	grpc.ServerStream
}

type coinListGetCoinsServer struct {
	grpc.ServerStream
}

func (x *coinListGetCoinsServer) Send(m *CoinInfo) error {
	return x.ServerStream.SendMsg(m)
}

func _CoinList_GetCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinListServer).GetCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinlist.CoinList/GetCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinListServer).GetCoin(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinList_CreateCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinListServer).CreateCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinlist.CoinList/CreateCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinListServer).CreateCoins(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinList_UpdateCoins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CoinInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinListServer).UpdateCoins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinlist.CoinList/UpdateCoins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinListServer).UpdateCoins(ctx, req.(*CoinInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinList_DeleteCoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoinListServer).DeleteCoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/coinlist.CoinList/DeleteCoin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoinListServer).DeleteCoin(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoinList_SearchCoins_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InputText)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CoinListServer).SearchCoins(m, &coinListSearchCoinsServer{stream})
}

type CoinList_SearchCoinsServer interface {
	Send(*CoinInfo) error
	grpc.ServerStream
}

type coinListSearchCoinsServer struct {
	grpc.ServerStream
}

func (x *coinListSearchCoinsServer) Send(m *CoinInfo) error {
	return x.ServerStream.SendMsg(m)
}

// CoinList_ServiceDesc is the grpc.ServiceDesc for CoinList service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoinList_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "coinlist.CoinList",
	HandlerType: (*CoinListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCoin",
			Handler:    _CoinList_GetCoin_Handler,
		},
		{
			MethodName: "CreateCoins",
			Handler:    _CoinList_CreateCoins_Handler,
		},
		{
			MethodName: "UpdateCoins",
			Handler:    _CoinList_UpdateCoins_Handler,
		},
		{
			MethodName: "DeleteCoin",
			Handler:    _CoinList_DeleteCoin_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetCoins",
			Handler:       _CoinList_GetCoins_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SearchCoins",
			Handler:       _CoinList_SearchCoins_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/coinlist/coin_list.proto",
}
