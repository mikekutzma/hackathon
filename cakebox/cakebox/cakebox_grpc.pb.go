// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package cakebox

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

// CakeBoxClient is the client API for CakeBox service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CakeBoxClient interface {
	// Sends a greeting
	UserFromBirthday(ctx context.Context, in *Birthday, opts ...grpc.CallOption) (*User, error)
	UsersFromBirthday(ctx context.Context, in *Birthday, opts ...grpc.CallOption) (CakeBox_UsersFromBirthdayClient, error)
	BirthdayFromUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Birthday, error)
}

type cakeBoxClient struct {
	cc grpc.ClientConnInterface
}

func NewCakeBoxClient(cc grpc.ClientConnInterface) CakeBoxClient {
	return &cakeBoxClient{cc}
}

func (c *cakeBoxClient) UserFromBirthday(ctx context.Context, in *Birthday, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/cakebox.CakeBox/UserFromBirthday", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cakeBoxClient) UsersFromBirthday(ctx context.Context, in *Birthday, opts ...grpc.CallOption) (CakeBox_UsersFromBirthdayClient, error) {
	stream, err := c.cc.NewStream(ctx, &CakeBox_ServiceDesc.Streams[0], "/cakebox.CakeBox/UsersFromBirthday", opts...)
	if err != nil {
		return nil, err
	}
	x := &cakeBoxUsersFromBirthdayClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CakeBox_UsersFromBirthdayClient interface {
	Recv() (*User, error)
	grpc.ClientStream
}

type cakeBoxUsersFromBirthdayClient struct {
	grpc.ClientStream
}

func (x *cakeBoxUsersFromBirthdayClient) Recv() (*User, error) {
	m := new(User)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *cakeBoxClient) BirthdayFromUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Birthday, error) {
	out := new(Birthday)
	err := c.cc.Invoke(ctx, "/cakebox.CakeBox/BirthdayFromUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CakeBoxServer is the server API for CakeBox service.
// All implementations must embed UnimplementedCakeBoxServer
// for forward compatibility
type CakeBoxServer interface {
	// Sends a greeting
	UserFromBirthday(context.Context, *Birthday) (*User, error)
	UsersFromBirthday(*Birthday, CakeBox_UsersFromBirthdayServer) error
	BirthdayFromUser(context.Context, *User) (*Birthday, error)
	mustEmbedUnimplementedCakeBoxServer()
}

// UnimplementedCakeBoxServer must be embedded to have forward compatible implementations.
type UnimplementedCakeBoxServer struct {
}

func (UnimplementedCakeBoxServer) UserFromBirthday(context.Context, *Birthday) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFromBirthday not implemented")
}
func (UnimplementedCakeBoxServer) UsersFromBirthday(*Birthday, CakeBox_UsersFromBirthdayServer) error {
	return status.Errorf(codes.Unimplemented, "method UsersFromBirthday not implemented")
}
func (UnimplementedCakeBoxServer) BirthdayFromUser(context.Context, *User) (*Birthday, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BirthdayFromUser not implemented")
}
func (UnimplementedCakeBoxServer) mustEmbedUnimplementedCakeBoxServer() {}

// UnsafeCakeBoxServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CakeBoxServer will
// result in compilation errors.
type UnsafeCakeBoxServer interface {
	mustEmbedUnimplementedCakeBoxServer()
}

func RegisterCakeBoxServer(s grpc.ServiceRegistrar, srv CakeBoxServer) {
	s.RegisterService(&CakeBox_ServiceDesc, srv)
}

func _CakeBox_UserFromBirthday_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Birthday)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CakeBoxServer).UserFromBirthday(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cakebox.CakeBox/UserFromBirthday",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CakeBoxServer).UserFromBirthday(ctx, req.(*Birthday))
	}
	return interceptor(ctx, in, info, handler)
}

func _CakeBox_UsersFromBirthday_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Birthday)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CakeBoxServer).UsersFromBirthday(m, &cakeBoxUsersFromBirthdayServer{stream})
}

type CakeBox_UsersFromBirthdayServer interface {
	Send(*User) error
	grpc.ServerStream
}

type cakeBoxUsersFromBirthdayServer struct {
	grpc.ServerStream
}

func (x *cakeBoxUsersFromBirthdayServer) Send(m *User) error {
	return x.ServerStream.SendMsg(m)
}

func _CakeBox_BirthdayFromUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CakeBoxServer).BirthdayFromUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cakebox.CakeBox/BirthdayFromUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CakeBoxServer).BirthdayFromUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// CakeBox_ServiceDesc is the grpc.ServiceDesc for CakeBox service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CakeBox_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cakebox.CakeBox",
	HandlerType: (*CakeBoxServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserFromBirthday",
			Handler:    _CakeBox_UserFromBirthday_Handler,
		},
		{
			MethodName: "BirthdayFromUser",
			Handler:    _CakeBox_BirthdayFromUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UsersFromBirthday",
			Handler:       _CakeBox_UsersFromBirthday_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cakebox/cakebox.proto",
}