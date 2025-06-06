// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0--rc1
// source: push/v1/push.proto

package v1

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
	Push_PushMsg_FullMethodName         = "/api.push.v1.Push/PushMsg"
	Push_GetOnlineStatus_FullMethodName = "/api.push.v1.Push/GetOnlineStatus"
)

// PushClient is the client API for Push service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PushClient interface {
	PushMsg(ctx context.Context, in *PushMsgRequest, opts ...grpc.CallOption) (*PushMsgReply, error)
	GetOnlineStatus(ctx context.Context, in *GetOnlineStatusRequest, opts ...grpc.CallOption) (*GetOnlineStatusReply, error)
}

type pushClient struct {
	cc grpc.ClientConnInterface
}

func NewPushClient(cc grpc.ClientConnInterface) PushClient {
	return &pushClient{cc}
}

func (c *pushClient) PushMsg(ctx context.Context, in *PushMsgRequest, opts ...grpc.CallOption) (*PushMsgReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushMsgReply)
	err := c.cc.Invoke(ctx, Push_PushMsg_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) GetOnlineStatus(ctx context.Context, in *GetOnlineStatusRequest, opts ...grpc.CallOption) (*GetOnlineStatusReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOnlineStatusReply)
	err := c.cc.Invoke(ctx, Push_GetOnlineStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServer is the server API for Push service.
// All implementations must embed UnimplementedPushServer
// for forward compatibility.
type PushServer interface {
	PushMsg(context.Context, *PushMsgRequest) (*PushMsgReply, error)
	GetOnlineStatus(context.Context, *GetOnlineStatusRequest) (*GetOnlineStatusReply, error)
	mustEmbedUnimplementedPushServer()
}

// UnimplementedPushServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPushServer struct{}

func (UnimplementedPushServer) PushMsg(context.Context, *PushMsgRequest) (*PushMsgReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushMsg not implemented")
}
func (UnimplementedPushServer) GetOnlineStatus(context.Context, *GetOnlineStatusRequest) (*GetOnlineStatusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineStatus not implemented")
}
func (UnimplementedPushServer) mustEmbedUnimplementedPushServer() {}
func (UnimplementedPushServer) testEmbeddedByValue()              {}

// UnsafePushServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PushServer will
// result in compilation errors.
type UnsafePushServer interface {
	mustEmbedUnimplementedPushServer()
}

func RegisterPushServer(s grpc.ServiceRegistrar, srv PushServer) {
	// If the following call pancis, it indicates UnimplementedPushServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Push_ServiceDesc, srv)
}

func _Push_PushMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushMsgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).PushMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Push_PushMsg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).PushMsg(ctx, req.(*PushMsgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_GetOnlineStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOnlineStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).GetOnlineStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Push_GetOnlineStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).GetOnlineStatus(ctx, req.(*GetOnlineStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Push_ServiceDesc is the grpc.ServiceDesc for Push service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Push_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.push.v1.Push",
	HandlerType: (*PushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushMsg",
			Handler:    _Push_PushMsg_Handler,
		},
		{
			MethodName: "GetOnlineStatus",
			Handler:    _Push_GetOnlineStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push/v1/push.proto",
}
