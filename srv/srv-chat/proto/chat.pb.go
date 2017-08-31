// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Chat struct {
}

func (m *Chat) Reset()                    { *m = Chat{} }
func (m *Chat) String() string            { return proto1.CompactTextString(m) }
func (*Chat) ProtoMessage()               {}
func (*Chat) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type Chat_Nil struct {
}

func (m *Chat_Nil) Reset()                    { *m = Chat_Nil{} }
func (m *Chat_Nil) String() string            { return proto1.CompactTextString(m) }
func (*Chat_Nil) ProtoMessage()               {}
func (*Chat_Nil) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 0} }

type Chat_List struct {
	Messages []*Chat_Message `protobuf:"bytes,1,rep,name=Messages" json:"Messages,omitempty"`
}

func (m *Chat_List) Reset()                    { *m = Chat_List{} }
func (m *Chat_List) String() string            { return proto1.CompactTextString(m) }
func (*Chat_List) ProtoMessage()               {}
func (*Chat_List) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 1} }

func (m *Chat_List) GetMessages() []*Chat_Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

type Chat_Id struct {
	Id string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
}

func (m *Chat_Id) Reset()                    { *m = Chat_Id{} }
func (m *Chat_Id) String() string            { return proto1.CompactTextString(m) }
func (*Chat_Id) ProtoMessage()               {}
func (*Chat_Id) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 2} }

func (m *Chat_Id) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Chat_Message struct {
	Id     string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Body   []byte `protobuf:"bytes,2,opt,name=Body,proto3" json:"Body,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=Offset" json:"Offset,omitempty"`
}

func (m *Chat_Message) Reset()                    { *m = Chat_Message{} }
func (m *Chat_Message) String() string            { return proto1.CompactTextString(m) }
func (*Chat_Message) ProtoMessage()               {}
func (*Chat_Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 3} }

func (m *Chat_Message) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Chat_Message) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *Chat_Message) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type Chat_Consumer struct {
	Id   string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	From int64  `protobuf:"varint,2,opt,name=From" json:"From,omitempty"`
}

func (m *Chat_Consumer) Reset()                    { *m = Chat_Consumer{} }
func (m *Chat_Consumer) String() string            { return proto1.CompactTextString(m) }
func (*Chat_Consumer) ProtoMessage()               {}
func (*Chat_Consumer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 4} }

func (m *Chat_Consumer) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Chat_Consumer) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

type Chat_ConsumeRange struct {
	Id   string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	From int64  `protobuf:"varint,2,opt,name=From" json:"From,omitempty"`
	To   int64  `protobuf:"varint,3,opt,name=To" json:"To,omitempty"`
}

func (m *Chat_ConsumeRange) Reset()                    { *m = Chat_ConsumeRange{} }
func (m *Chat_ConsumeRange) String() string            { return proto1.CompactTextString(m) }
func (*Chat_ConsumeRange) ProtoMessage()               {}
func (*Chat_ConsumeRange) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 5} }

func (m *Chat_ConsumeRange) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Chat_ConsumeRange) GetFrom() int64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *Chat_ConsumeRange) GetTo() int64 {
	if m != nil {
		return m.To
	}
	return 0
}

type Chat_ConsumeLatest struct {
	Id     string `protobuf:"bytes,1,opt,name=Id" json:"Id,omitempty"`
	Length int64  `protobuf:"varint,2,opt,name=Length" json:"Length,omitempty"`
}

func (m *Chat_ConsumeLatest) Reset()                    { *m = Chat_ConsumeLatest{} }
func (m *Chat_ConsumeLatest) String() string            { return proto1.CompactTextString(m) }
func (*Chat_ConsumeLatest) ProtoMessage()               {}
func (*Chat_ConsumeLatest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0, 6} }

func (m *Chat_ConsumeLatest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Chat_ConsumeLatest) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func init() {
	proto1.RegisterType((*Chat)(nil), "proto.Chat")
	proto1.RegisterType((*Chat_Nil)(nil), "proto.Chat.Nil")
	proto1.RegisterType((*Chat_List)(nil), "proto.Chat.List")
	proto1.RegisterType((*Chat_Id)(nil), "proto.Chat.Id")
	proto1.RegisterType((*Chat_Message)(nil), "proto.Chat.Message")
	proto1.RegisterType((*Chat_Consumer)(nil), "proto.Chat.Consumer")
	proto1.RegisterType((*Chat_ConsumeRange)(nil), "proto.Chat.ConsumeRange")
	proto1.RegisterType((*Chat_ConsumeLatest)(nil), "proto.Chat.ConsumeLatest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ChatService service

type ChatServiceClient interface {
	Subscribe(ctx context.Context, in *Chat_Consumer, opts ...grpc.CallOption) (ChatService_SubscribeClient, error)
	Reg(ctx context.Context, in *Chat_Id, opts ...grpc.CallOption) (*Chat_Nil, error)
	Unreg(ctx context.Context, in *Chat_Id, opts ...grpc.CallOption) (*Chat_Nil, error)
	Query(ctx context.Context, in *Chat_ConsumeRange, opts ...grpc.CallOption) (*Chat_List, error)
	Latest(ctx context.Context, in *Chat_ConsumeLatest, opts ...grpc.CallOption) (*Chat_List, error)
	Send(ctx context.Context, in *Chat_Message, opts ...grpc.CallOption) (*Chat_Nil, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Subscribe(ctx context.Context, in *Chat_Consumer, opts ...grpc.CallOption) (ChatService_SubscribeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChatService_serviceDesc.Streams[0], c.cc, "/proto.ChatService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_SubscribeClient interface {
	Recv() (*Chat_Message, error)
	grpc.ClientStream
}

type chatServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *chatServiceSubscribeClient) Recv() (*Chat_Message, error) {
	m := new(Chat_Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) Reg(ctx context.Context, in *Chat_Id, opts ...grpc.CallOption) (*Chat_Nil, error) {
	out := new(Chat_Nil)
	err := grpc.Invoke(ctx, "/proto.ChatService/Reg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Unreg(ctx context.Context, in *Chat_Id, opts ...grpc.CallOption) (*Chat_Nil, error) {
	out := new(Chat_Nil)
	err := grpc.Invoke(ctx, "/proto.ChatService/Unreg", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Query(ctx context.Context, in *Chat_ConsumeRange, opts ...grpc.CallOption) (*Chat_List, error) {
	out := new(Chat_List)
	err := grpc.Invoke(ctx, "/proto.ChatService/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Latest(ctx context.Context, in *Chat_ConsumeLatest, opts ...grpc.CallOption) (*Chat_List, error) {
	out := new(Chat_List)
	err := grpc.Invoke(ctx, "/proto.ChatService/Latest", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Send(ctx context.Context, in *Chat_Message, opts ...grpc.CallOption) (*Chat_Nil, error) {
	out := new(Chat_Nil)
	err := grpc.Invoke(ctx, "/proto.ChatService/Send", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ChatService service

type ChatServiceServer interface {
	Subscribe(*Chat_Consumer, ChatService_SubscribeServer) error
	Reg(context.Context, *Chat_Id) (*Chat_Nil, error)
	Unreg(context.Context, *Chat_Id) (*Chat_Nil, error)
	Query(context.Context, *Chat_ConsumeRange) (*Chat_List, error)
	Latest(context.Context, *Chat_ConsumeLatest) (*Chat_List, error)
	Send(context.Context, *Chat_Message) (*Chat_Nil, error)
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Chat_Consumer)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).Subscribe(m, &chatServiceSubscribeServer{stream})
}

type ChatService_SubscribeServer interface {
	Send(*Chat_Message) error
	grpc.ServerStream
}

type chatServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *chatServiceSubscribeServer) Send(m *Chat_Message) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_Reg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat_Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Reg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChatService/Reg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Reg(ctx, req.(*Chat_Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Unreg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat_Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Unreg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChatService/Unreg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Unreg(ctx, req.(*Chat_Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat_ConsumeRange)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChatService/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Query(ctx, req.(*Chat_ConsumeRange))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Latest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat_ConsumeLatest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Latest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChatService/Latest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Latest(ctx, req.(*Chat_ConsumeLatest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Chat_Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ChatService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Send(ctx, req.(*Chat_Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Reg",
			Handler:    _ChatService_Reg_Handler,
		},
		{
			MethodName: "Unreg",
			Handler:    _ChatService_Unreg_Handler,
		},
		{
			MethodName: "Query",
			Handler:    _ChatService_Query_Handler,
		},
		{
			MethodName: "Latest",
			Handler:    _ChatService_Latest_Handler,
		},
		{
			MethodName: "Send",
			Handler:    _ChatService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ChatService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat.proto",
}

func init() { proto1.RegisterFile("chat.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcd, 0x6a, 0xf2, 0x40,
	0x14, 0x25, 0xbf, 0x9f, 0x5e, 0xfd, 0x6c, 0xb9, 0x15, 0x49, 0x67, 0x15, 0xba, 0x28, 0x59, 0x94,
	0xb4, 0x28, 0xc5, 0xae, 0x95, 0x16, 0x04, 0x6b, 0x69, 0xb4, 0x0f, 0x10, 0xcd, 0xa8, 0x01, 0xcd,
	0x94, 0x99, 0xb1, 0xe0, 0x73, 0xf5, 0x45, 0xfa, 0x48, 0x25, 0xe3, 0xb4, 0x04, 0x9d, 0x45, 0x37,
	0xc9, 0x1c, 0xce, 0xcf, 0xdc, 0x7b, 0x18, 0x80, 0xc5, 0x3a, 0x95, 0xf1, 0x3b, 0x67, 0x92, 0xa1,
	0xa7, 0x7e, 0x57, 0x5f, 0x36, 0xb8, 0xc3, 0x75, 0x2a, 0x89, 0x07, 0xce, 0x24, 0xdf, 0x90, 0x3e,
	0xb8, 0xe3, 0x5c, 0x48, 0xbc, 0x85, 0xda, 0x33, 0x15, 0x22, 0x5d, 0x51, 0x11, 0x58, 0xa1, 0x13,
	0x35, 0xba, 0x17, 0x07, 0x63, 0x5c, 0xaa, 0x63, 0xcd, 0x25, 0xbf, 0x22, 0xd2, 0x06, 0x7b, 0x94,
	0x61, 0xab, 0xfc, 0x06, 0x56, 0x68, 0x45, 0xf5, 0xc4, 0x1e, 0x65, 0xe4, 0x11, 0xfe, 0x69, 0xc5,
	0x31, 0x85, 0x08, 0xee, 0x80, 0x65, 0xfb, 0xc0, 0x0e, 0xad, 0xa8, 0x99, 0xa8, 0x33, 0x76, 0xc0,
	0x7f, 0x59, 0x2e, 0x05, 0x95, 0x81, 0x13, 0x5a, 0x91, 0x93, 0x68, 0x44, 0x62, 0xa8, 0x0d, 0x59,
	0x21, 0x76, 0x5b, 0xca, 0x4d, 0x39, 0x4f, 0x9c, 0x6d, 0x55, 0x8e, 0x93, 0xa8, 0x33, 0x19, 0x40,
	0x53, 0xeb, 0x93, 0xb4, 0x30, 0xdf, 0x7d, 0xec, 0x29, 0x35, 0x33, 0xa6, 0xef, 0xb5, 0x67, 0x8c,
	0xf4, 0xe1, 0xbf, 0xce, 0x18, 0xa7, 0x92, 0x0a, 0x79, 0x12, 0xd2, 0x01, 0x7f, 0x4c, 0x8b, 0x95,
	0x5c, 0xeb, 0x18, 0x8d, 0xba, 0x9f, 0x36, 0x34, 0xca, 0x92, 0xa6, 0x94, 0x7f, 0xe4, 0x0b, 0x8a,
	0x0f, 0x50, 0x9f, 0xee, 0xe6, 0x62, 0xc1, 0xf3, 0x39, 0xc5, 0x76, 0xb5, 0xc5, 0x9f, 0x9d, 0x88,
	0xa9, 0xdb, 0x3b, 0x0b, 0xaf, 0xc1, 0x49, 0xe8, 0x0a, 0x5b, 0x55, 0x76, 0x94, 0x91, 0xb3, 0x2a,
	0x9e, 0xe4, 0x1b, 0x8c, 0xc0, 0x7b, 0x2b, 0xf8, 0x5f, 0x94, 0x3d, 0xf0, 0x5e, 0x77, 0x94, 0xef,
	0x31, 0x30, 0xcc, 0xa1, 0xba, 0x22, 0xe7, 0x55, 0x46, 0xbd, 0x85, 0x7b, 0xf0, 0x75, 0x05, 0x97,
	0x06, 0xd7, 0x81, 0x32, 0xd8, 0x6e, 0xc0, 0x9d, 0xd2, 0x22, 0x43, 0xd3, 0x72, 0x27, 0x93, 0xcd,
	0x7d, 0x85, 0x7b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x29, 0x8e, 0xfa, 0xeb, 0xa4, 0x02, 0x00,
	0x00,
}
