// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: order.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OrderService_AddOrder_FullMethodName             = "/OrderService/AddOrder"
	OrderService_CallbackNotification_FullMethodName = "/OrderService/CallbackNotification"
	OrderService_GetOrderList_FullMethodName         = "/OrderService/GetOrderList"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	AddOrder(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderResponse, error)
	CallbackNotification(ctx context.Context, in *CallbackNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*GetOrderListResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) AddOrder(ctx context.Context, in *AddOrderRequest, opts ...grpc.CallOption) (*AddOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_AddOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CallbackNotification(ctx context.Context, in *CallbackNotificationRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, OrderService_CallbackNotification_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) GetOrderList(ctx context.Context, in *GetOrderListRequest, opts ...grpc.CallOption) (*GetOrderListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderListResponse)
	err := c.cc.Invoke(ctx, OrderService_GetOrderList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations should embed UnimplementedOrderServiceServer
// for forward compatibility.
type OrderServiceServer interface {
	AddOrder(context.Context, *AddOrderRequest) (*AddOrderResponse, error)
	CallbackNotification(context.Context, *CallbackNotificationRequest) (*emptypb.Empty, error)
	GetOrderList(context.Context, *GetOrderListRequest) (*GetOrderListResponse, error)
}

// UnimplementedOrderServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderServiceServer struct{}

func (UnimplementedOrderServiceServer) AddOrder(context.Context, *AddOrderRequest) (*AddOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrder not implemented")
}
func (UnimplementedOrderServiceServer) CallbackNotification(context.Context, *CallbackNotificationRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallbackNotification not implemented")
}
func (UnimplementedOrderServiceServer) GetOrderList(context.Context, *GetOrderListRequest) (*GetOrderListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderList not implemented")
}
func (UnimplementedOrderServiceServer) testEmbeddedByValue() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_AddOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).AddOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_AddOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).AddOrder(ctx, req.(*AddOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CallbackNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CallbackNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CallbackNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CallbackNotification_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CallbackNotification(ctx, req.(*CallbackNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_GetOrderList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).GetOrderList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_GetOrderList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).GetOrderList(ctx, req.(*GetOrderListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddOrder",
			Handler:    _OrderService_AddOrder_Handler,
		},
		{
			MethodName: "CallbackNotification",
			Handler:    _OrderService_CallbackNotification_Handler,
		},
		{
			MethodName: "GetOrderList",
			Handler:    _OrderService_GetOrderList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
