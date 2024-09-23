// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: voucher.proto

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
	VoucherService_AddVoucher_FullMethodName     = "/VoucherService/AddVoucher"
	VoucherService_GetVoucher_FullMethodName     = "/VoucherService/GetVoucher"
	VoucherService_GetVoucherList_FullMethodName = "/VoucherService/GetVoucherList"
)

// VoucherServiceClient is the client API for VoucherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VoucherServiceClient interface {
	AddVoucher(ctx context.Context, in *AddVoucherRequest, opts ...grpc.CallOption) (*AddVoucherResponse, error)
	GetVoucher(ctx context.Context, in *GetVoucherRequest, opts ...grpc.CallOption) (*GetVoucherResponse, error)
	GetVoucherList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetVoucherListResponse, error)
}

type voucherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVoucherServiceClient(cc grpc.ClientConnInterface) VoucherServiceClient {
	return &voucherServiceClient{cc}
}

func (c *voucherServiceClient) AddVoucher(ctx context.Context, in *AddVoucherRequest, opts ...grpc.CallOption) (*AddVoucherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddVoucherResponse)
	err := c.cc.Invoke(ctx, VoucherService_AddVoucher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voucherServiceClient) GetVoucher(ctx context.Context, in *GetVoucherRequest, opts ...grpc.CallOption) (*GetVoucherResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVoucherResponse)
	err := c.cc.Invoke(ctx, VoucherService_GetVoucher_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *voucherServiceClient) GetVoucherList(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetVoucherListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetVoucherListResponse)
	err := c.cc.Invoke(ctx, VoucherService_GetVoucherList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VoucherServiceServer is the server API for VoucherService service.
// All implementations should embed UnimplementedVoucherServiceServer
// for forward compatibility.
type VoucherServiceServer interface {
	AddVoucher(context.Context, *AddVoucherRequest) (*AddVoucherResponse, error)
	GetVoucher(context.Context, *GetVoucherRequest) (*GetVoucherResponse, error)
	GetVoucherList(context.Context, *emptypb.Empty) (*GetVoucherListResponse, error)
}

// UnimplementedVoucherServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedVoucherServiceServer struct{}

func (UnimplementedVoucherServiceServer) AddVoucher(context.Context, *AddVoucherRequest) (*AddVoucherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVoucher not implemented")
}
func (UnimplementedVoucherServiceServer) GetVoucher(context.Context, *GetVoucherRequest) (*GetVoucherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoucher not implemented")
}
func (UnimplementedVoucherServiceServer) GetVoucherList(context.Context, *emptypb.Empty) (*GetVoucherListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVoucherList not implemented")
}
func (UnimplementedVoucherServiceServer) testEmbeddedByValue() {}

// UnsafeVoucherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VoucherServiceServer will
// result in compilation errors.
type UnsafeVoucherServiceServer interface {
	mustEmbedUnimplementedVoucherServiceServer()
}

func RegisterVoucherServiceServer(s grpc.ServiceRegistrar, srv VoucherServiceServer) {
	// If the following call pancis, it indicates UnimplementedVoucherServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&VoucherService_ServiceDesc, srv)
}

func _VoucherService_AddVoucher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVoucherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServiceServer).AddVoucher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoucherService_AddVoucher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServiceServer).AddVoucher(ctx, req.(*AddVoucherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoucherService_GetVoucher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVoucherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServiceServer).GetVoucher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoucherService_GetVoucher_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServiceServer).GetVoucher(ctx, req.(*GetVoucherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VoucherService_GetVoucherList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VoucherServiceServer).GetVoucherList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VoucherService_GetVoucherList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VoucherServiceServer).GetVoucherList(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// VoucherService_ServiceDesc is the grpc.ServiceDesc for VoucherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VoucherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VoucherService",
	HandlerType: (*VoucherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddVoucher",
			Handler:    _VoucherService_AddVoucher_Handler,
		},
		{
			MethodName: "GetVoucher",
			Handler:    _VoucherService_GetVoucher_Handler,
		},
		{
			MethodName: "GetVoucherList",
			Handler:    _VoucherService_GetVoucherList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voucher.proto",
}
