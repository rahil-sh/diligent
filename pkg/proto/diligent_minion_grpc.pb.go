// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: diligent_minion.proto

package proto

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

// MinionClient is the client API for Minion service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MinionClient interface {
	// Ping
	Ping(ctx context.Context, in *MinionPingRequest, opts ...grpc.CallOption) (*MinionPingResponse, error)
	// Job Control
	PrepareJob(ctx context.Context, in *MinionPrepareJobRequest, opts ...grpc.CallOption) (*MinionPrepareJobResponse, error)
	RunJob(ctx context.Context, in *MinionRunJobRequest, opts ...grpc.CallOption) (*MinionRunJobResponse, error)
	AbortJob(ctx context.Context, in *MinionAbortJobRequest, opts ...grpc.CallOption) (*MinionAbortJobResponse, error)
	QueryJob(ctx context.Context, in *MinionQueryJobRequest, opts ...grpc.CallOption) (*MinionQueryJobResponse, error)
}

type minionClient struct {
	cc grpc.ClientConnInterface
}

func NewMinionClient(cc grpc.ClientConnInterface) MinionClient {
	return &minionClient{cc}
}

func (c *minionClient) Ping(ctx context.Context, in *MinionPingRequest, opts ...grpc.CallOption) (*MinionPingResponse, error) {
	out := new(MinionPingResponse)
	err := c.cc.Invoke(ctx, "/proto.Minion/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minionClient) PrepareJob(ctx context.Context, in *MinionPrepareJobRequest, opts ...grpc.CallOption) (*MinionPrepareJobResponse, error) {
	out := new(MinionPrepareJobResponse)
	err := c.cc.Invoke(ctx, "/proto.Minion/PrepareJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minionClient) RunJob(ctx context.Context, in *MinionRunJobRequest, opts ...grpc.CallOption) (*MinionRunJobResponse, error) {
	out := new(MinionRunJobResponse)
	err := c.cc.Invoke(ctx, "/proto.Minion/RunJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minionClient) AbortJob(ctx context.Context, in *MinionAbortJobRequest, opts ...grpc.CallOption) (*MinionAbortJobResponse, error) {
	out := new(MinionAbortJobResponse)
	err := c.cc.Invoke(ctx, "/proto.Minion/AbortJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minionClient) QueryJob(ctx context.Context, in *MinionQueryJobRequest, opts ...grpc.CallOption) (*MinionQueryJobResponse, error) {
	out := new(MinionQueryJobResponse)
	err := c.cc.Invoke(ctx, "/proto.Minion/QueryJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MinionServer is the server API for Minion service.
// All implementations must embed UnimplementedMinionServer
// for forward compatibility
type MinionServer interface {
	// Ping
	Ping(context.Context, *MinionPingRequest) (*MinionPingResponse, error)
	// Job Control
	PrepareJob(context.Context, *MinionPrepareJobRequest) (*MinionPrepareJobResponse, error)
	RunJob(context.Context, *MinionRunJobRequest) (*MinionRunJobResponse, error)
	AbortJob(context.Context, *MinionAbortJobRequest) (*MinionAbortJobResponse, error)
	QueryJob(context.Context, *MinionQueryJobRequest) (*MinionQueryJobResponse, error)
	mustEmbedUnimplementedMinionServer()
}

// UnimplementedMinionServer must be embedded to have forward compatible implementations.
type UnimplementedMinionServer struct {
}

func (UnimplementedMinionServer) Ping(context.Context, *MinionPingRequest) (*MinionPingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedMinionServer) PrepareJob(context.Context, *MinionPrepareJobRequest) (*MinionPrepareJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareJob not implemented")
}
func (UnimplementedMinionServer) RunJob(context.Context, *MinionRunJobRequest) (*MinionRunJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunJob not implemented")
}
func (UnimplementedMinionServer) AbortJob(context.Context, *MinionAbortJobRequest) (*MinionAbortJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortJob not implemented")
}
func (UnimplementedMinionServer) QueryJob(context.Context, *MinionQueryJobRequest) (*MinionQueryJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryJob not implemented")
}
func (UnimplementedMinionServer) mustEmbedUnimplementedMinionServer() {}

// UnsafeMinionServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MinionServer will
// result in compilation errors.
type UnsafeMinionServer interface {
	mustEmbedUnimplementedMinionServer()
}

func RegisterMinionServer(s grpc.ServiceRegistrar, srv MinionServer) {
	s.RegisterService(&Minion_ServiceDesc, srv)
}

func _Minion_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinionPingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinionServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Minion/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinionServer).Ping(ctx, req.(*MinionPingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Minion_PrepareJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinionPrepareJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinionServer).PrepareJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Minion/PrepareJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinionServer).PrepareJob(ctx, req.(*MinionPrepareJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Minion_RunJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinionRunJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinionServer).RunJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Minion/RunJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinionServer).RunJob(ctx, req.(*MinionRunJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Minion_AbortJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinionAbortJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinionServer).AbortJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Minion/AbortJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinionServer).AbortJob(ctx, req.(*MinionAbortJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Minion_QueryJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MinionQueryJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinionServer).QueryJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Minion/QueryJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinionServer).QueryJob(ctx, req.(*MinionQueryJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Minion_ServiceDesc is the grpc.ServiceDesc for Minion service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Minion_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Minion",
	HandlerType: (*MinionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _Minion_Ping_Handler,
		},
		{
			MethodName: "PrepareJob",
			Handler:    _Minion_PrepareJob_Handler,
		},
		{
			MethodName: "RunJob",
			Handler:    _Minion_RunJob_Handler,
		},
		{
			MethodName: "AbortJob",
			Handler:    _Minion_AbortJob_Handler,
		},
		{
			MethodName: "QueryJob",
			Handler:    _Minion_QueryJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "diligent_minion.proto",
}
