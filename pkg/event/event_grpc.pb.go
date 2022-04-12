// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package event

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventServiceClient is the client API for EventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventServiceClient interface {
	PostEvent(ctx context.Context, opts ...grpc.CallOption) (EventService_PostEventClient, error)
	PostEventBatch(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error)
}

type eventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventServiceClient(cc grpc.ClientConnInterface) EventServiceClient {
	return &eventServiceClient{cc}
}

func (c *eventServiceClient) PostEvent(ctx context.Context, opts ...grpc.CallOption) (EventService_PostEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventService_ServiceDesc.Streams[0], "/proto.EventService/PostEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventServicePostEventClient{stream}
	return x, nil
}

type EventService_PostEventClient interface {
	Send(*Event) error
	CloseAndRecv() (*emptypb.Empty, error)
	grpc.ClientStream
}

type eventServicePostEventClient struct {
	grpc.ClientStream
}

func (x *eventServicePostEventClient) Send(m *Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventServicePostEventClient) CloseAndRecv() (*emptypb.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(emptypb.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventServiceClient) PostEventBatch(ctx context.Context, in *EventRequest, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/proto.EventService/PostEventBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventServiceServer is the server API for EventService service.
// All implementations must embed UnimplementedEventServiceServer
// for forward compatibility
type EventServiceServer interface {
	PostEvent(EventService_PostEventServer) error
	PostEventBatch(context.Context, *EventRequest) (*EventResponse, error)
	mustEmbedUnimplementedEventServiceServer()
}

// UnimplementedEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventServiceServer struct {
}

func (UnimplementedEventServiceServer) PostEvent(EventService_PostEventServer) error {
	return status.Errorf(codes.Unimplemented, "method PostEvent not implemented")
}
func (UnimplementedEventServiceServer) PostEventBatch(context.Context, *EventRequest) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostEventBatch not implemented")
}
func (UnimplementedEventServiceServer) mustEmbedUnimplementedEventServiceServer() {}

// UnsafeEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventServiceServer will
// result in compilation errors.
type UnsafeEventServiceServer interface {
	mustEmbedUnimplementedEventServiceServer()
}

func RegisterEventServiceServer(s grpc.ServiceRegistrar, srv EventServiceServer) {
	s.RegisterService(&EventService_ServiceDesc, srv)
}

func _EventService_PostEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventServiceServer).PostEvent(&eventServicePostEventServer{stream})
}

type EventService_PostEventServer interface {
	SendAndClose(*emptypb.Empty) error
	Recv() (*Event, error)
	grpc.ServerStream
}

type eventServicePostEventServer struct {
	grpc.ServerStream
}

func (x *eventServicePostEventServer) SendAndClose(m *emptypb.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventServicePostEventServer) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EventService_PostEventBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventServiceServer).PostEventBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.EventService/PostEventBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventServiceServer).PostEventBatch(ctx, req.(*EventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventService_ServiceDesc is the grpc.ServiceDesc for EventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.EventService",
	HandlerType: (*EventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostEventBatch",
			Handler:    _EventService_PostEventBatch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PostEvent",
			Handler:       _EventService_PostEvent_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/event.proto",
}