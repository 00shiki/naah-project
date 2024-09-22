// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: proto/user.proto

package proto

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
	UserService_CreateUser_FullMethodName    = "/user.UserService/CreateUser"
	UserService_ValidateUser_FullMethodName  = "/user.UserService/ValidateUser"
	UserService_GetUserDetail_FullMethodName = "/user.UserService/GetUserDetail"
	UserService_VerifyEmail_FullMethodName   = "/user.UserService/VerifyEmail"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	ValidateUser(ctx context.Context, in *ValidateUserRequest, opts ...grpc.CallOption) (*ValidateUserResponse, error)
	GetUserDetail(ctx context.Context, in *GetUserDetailRequest, opts ...grpc.CallOption) (*GetUserDetailResponse, error)
	VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, UserService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ValidateUser(ctx context.Context, in *ValidateUserRequest, opts ...grpc.CallOption) (*ValidateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ValidateUserResponse)
	err := c.cc.Invoke(ctx, UserService_ValidateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserDetail(ctx context.Context, in *GetUserDetailRequest, opts ...grpc.CallOption) (*GetUserDetailResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUserDetailResponse)
	err := c.cc.Invoke(ctx, UserService_GetUserDetail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) VerifyEmail(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_VerifyEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations should embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	ValidateUser(context.Context, *ValidateUserRequest) (*ValidateUserResponse, error)
	GetUserDetail(context.Context, *GetUserDetailRequest) (*GetUserDetailResponse, error)
	VerifyEmail(context.Context, *VerifyEmailRequest) (*emptypb.Empty, error)
}

// UnimplementedUserServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) ValidateUser(context.Context, *ValidateUserRequest) (*ValidateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserDetail(context.Context, *GetUserDetailRequest) (*GetUserDetailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserDetail not implemented")
}
func (UnimplementedUserServiceServer) VerifyEmail(context.Context, *VerifyEmailRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyEmail not implemented")
}
func (UnimplementedUserServiceServer) testEmbeddedByValue() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ValidateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ValidateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_ValidateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ValidateUser(ctx, req.(*ValidateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserDetail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserDetail(ctx, req.(*GetUserDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_VerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).VerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_VerifyEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).VerifyEmail(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "ValidateUser",
			Handler:    _UserService_ValidateUser_Handler,
		},
		{
			MethodName: "GetUserDetail",
			Handler:    _UserService_GetUserDetail_Handler,
		},
		{
			MethodName: "VerifyEmail",
			Handler:    _UserService_VerifyEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
