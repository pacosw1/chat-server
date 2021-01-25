// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ChatClient is the client API for Chat service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatClient interface {
	CreateSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*UserData, error)
	JoinRoom(ctx context.Context, in *RoomID, opts ...grpc.CallOption) (Chat_JoinRoomClient, error)
	SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error)
}

type chatClient struct {
	cc grpc.ClientConnInterface
}

func NewChatClient(cc grpc.ClientConnInterface) ChatClient {
	return &chatClient{cc}
}

func (c *chatClient) CreateSession(ctx context.Context, in *SessionRequest, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/proto.Chat/CreateSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatClient) JoinRoom(ctx context.Context, in *RoomID, opts ...grpc.CallOption) (Chat_JoinRoomClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[0], "/proto.Chat/JoinRoom", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatJoinRoomClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Chat_JoinRoomClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type chatJoinRoomClient struct {
	grpc.ClientStream
}

func (x *chatJoinRoomClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatClient) SendMessage(ctx context.Context, opts ...grpc.CallOption) (Chat_SendMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Chat_serviceDesc.Streams[1], "/proto.Chat/SendMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatSendMessageClient{stream}
	return x, nil
}

type Chat_SendMessageClient interface {
	Send(*Message) error
	CloseAndRecv() (*MessageSent, error)
	grpc.ClientStream
}

type chatSendMessageClient struct {
	grpc.ClientStream
}

func (x *chatSendMessageClient) Send(m *Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatSendMessageClient) CloseAndRecv() (*MessageSent, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(MessageSent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServer is the server API for Chat service.
// All implementations must embed UnimplementedChatServer
// for forward compatibility
type ChatServer interface {
	CreateSession(context.Context, *SessionRequest) (*UserData, error)
	JoinRoom(*RoomID, Chat_JoinRoomServer) error
	SendMessage(Chat_SendMessageServer) error
	mustEmbedUnimplementedChatServer()
}

// UnimplementedChatServer must be embedded to have forward compatible implementations.
type UnimplementedChatServer struct {
}

func (UnimplementedChatServer) CreateSession(context.Context, *SessionRequest) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSession not implemented")
}
func (UnimplementedChatServer) JoinRoom(*RoomID, Chat_JoinRoomServer) error {
	return status.Errorf(codes.Unimplemented, "method JoinRoom not implemented")
}
func (UnimplementedChatServer) SendMessage(Chat_SendMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServer) mustEmbedUnimplementedChatServer() {}

// UnsafeChatServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServer will
// result in compilation errors.
type UnsafeChatServer interface {
	mustEmbedUnimplementedChatServer()
}

func RegisterChatServer(s grpc.ServiceRegistrar, srv ChatServer) {
	s.RegisterService(&_Chat_serviceDesc, srv)
}

func _Chat_CreateSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServer).CreateSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Chat/CreateSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServer).CreateSession(ctx, req.(*SessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chat_JoinRoom_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RoomID)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServer).JoinRoom(m, &chatJoinRoomServer{stream})
}

type Chat_JoinRoomServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type chatJoinRoomServer struct {
	grpc.ServerStream
}

func (x *chatJoinRoomServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Chat_SendMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServer).SendMessage(&chatSendMessageServer{stream})
}

type Chat_SendMessageServer interface {
	SendAndClose(*MessageSent) error
	Recv() (*Message, error)
	grpc.ServerStream
}

type chatSendMessageServer struct {
	grpc.ServerStream
}

func (x *chatSendMessageServer) SendAndClose(m *MessageSent) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatSendMessageServer) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Chat_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Chat",
	HandlerType: (*ChatServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSession",
			Handler:    _Chat_CreateSession_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "JoinRoom",
			Handler:       _Chat_JoinRoom_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendMessage",
			Handler:       _Chat_SendMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "proto/contract.proto",
}
