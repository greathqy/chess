// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	auth.proto

It has these top-level messages:
	AuthArgs
	AuthRes
	TokenInfoArgs
	TokenInfoRes
	RefreshTokenArgs
	RefreshTokenRes
	DestroyTokenArgs
	DestroyTokenRes
*/
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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthArgs struct {
	UserId int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	Token  string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
}

func (m *AuthArgs) Reset()                    { *m = AuthArgs{} }
func (m *AuthArgs) String() string            { return proto1.CompactTextString(m) }
func (*AuthArgs) ProtoMessage()               {}
func (*AuthArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthArgs) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *AuthArgs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type AuthRes struct {
	Ret int32  `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`
}

func (m *AuthRes) Reset()                    { *m = AuthRes{} }
func (m *AuthRes) String() string            { return proto1.CompactTextString(m) }
func (*AuthRes) ProtoMessage()               {}
func (*AuthRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *AuthRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type TokenInfoArgs struct {
	UserId     int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	AppFrom    string `protobuf:"bytes,2,opt,name=AppFrom" json:"AppFrom,omitempty"`
	AppChannel string `protobuf:"bytes,3,opt,name=AppChannel" json:"AppChannel,omitempty"`
	AppVer     int32  `protobuf:"varint,4,opt,name=AppVer" json:"AppVer,omitempty"`
	UniqueId   string `protobuf:"bytes,5,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Token      string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
}

func (m *TokenInfoArgs) Reset()                    { *m = TokenInfoArgs{} }
func (m *TokenInfoArgs) String() string            { return proto1.CompactTextString(m) }
func (*TokenInfoArgs) ProtoMessage()               {}
func (*TokenInfoArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TokenInfoArgs) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *TokenInfoArgs) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *TokenInfoArgs) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

func (m *TokenInfoArgs) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *TokenInfoArgs) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *TokenInfoArgs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type TokenInfoRes struct {
	Ret    int32  `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`
	Expire int64  `protobuf:"varint,3,opt,name=Expire" json:"Expire,omitempty"`
}

func (m *TokenInfoRes) Reset()                    { *m = TokenInfoRes{} }
func (m *TokenInfoRes) String() string            { return proto1.CompactTextString(m) }
func (*TokenInfoRes) ProtoMessage()               {}
func (*TokenInfoRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TokenInfoRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *TokenInfoRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *TokenInfoRes) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

type RefreshTokenArgs struct {
	UserId       int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	AppFrom      string `protobuf:"bytes,2,opt,name=AppFrom" json:"AppFrom,omitempty"`
	AppChannel   string `protobuf:"bytes,3,opt,name=AppChannel" json:"AppChannel,omitempty"`
	AppVer       int32  `protobuf:"varint,4,opt,name=AppVer" json:"AppVer,omitempty"`
	UniqueId     string `protobuf:"bytes,5,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Token        string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
	RefreshToken string `protobuf:"bytes,7,opt,name=RefreshToken" json:"RefreshToken,omitempty"`
}

func (m *RefreshTokenArgs) Reset()                    { *m = RefreshTokenArgs{} }
func (m *RefreshTokenArgs) String() string            { return proto1.CompactTextString(m) }
func (*RefreshTokenArgs) ProtoMessage()               {}
func (*RefreshTokenArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RefreshTokenArgs) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RefreshTokenArgs) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *RefreshTokenArgs) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

func (m *RefreshTokenArgs) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *RefreshTokenArgs) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *RefreshTokenArgs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RefreshTokenArgs) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type RefreshTokenRes struct {
	Ret          int32  `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
	Msg          string `protobuf:"bytes,2,opt,name=Msg" json:"Msg,omitempty"`
	UserId       int32  `protobuf:"varint,3,opt,name=UserId" json:"UserId,omitempty"`
	Token        string `protobuf:"bytes,4,opt,name=Token" json:"Token,omitempty"`
	RefreshToken string `protobuf:"bytes,5,opt,name=RefreshToken" json:"RefreshToken,omitempty"`
	Expire       int64  `protobuf:"varint,6,opt,name=Expire" json:"Expire,omitempty"`
}

func (m *RefreshTokenRes) Reset()                    { *m = RefreshTokenRes{} }
func (m *RefreshTokenRes) String() string            { return proto1.CompactTextString(m) }
func (*RefreshTokenRes) ProtoMessage()               {}
func (*RefreshTokenRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RefreshTokenRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *RefreshTokenRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *RefreshTokenRes) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *RefreshTokenRes) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RefreshTokenRes) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *RefreshTokenRes) GetExpire() int64 {
	if m != nil {
		return m.Expire
	}
	return 0
}

type DestroyTokenArgs struct {
	UserId     int32  `protobuf:"varint,1,opt,name=UserId" json:"UserId,omitempty"`
	AppFrom    string `protobuf:"bytes,2,opt,name=AppFrom" json:"AppFrom,omitempty"`
	AppChannel string `protobuf:"bytes,3,opt,name=AppChannel" json:"AppChannel,omitempty"`
	AppVer     int32  `protobuf:"varint,4,opt,name=AppVer" json:"AppVer,omitempty"`
	UniqueId   string `protobuf:"bytes,5,opt,name=UniqueId" json:"UniqueId,omitempty"`
	Token      string `protobuf:"bytes,6,opt,name=Token" json:"Token,omitempty"`
}

func (m *DestroyTokenArgs) Reset()                    { *m = DestroyTokenArgs{} }
func (m *DestroyTokenArgs) String() string            { return proto1.CompactTextString(m) }
func (*DestroyTokenArgs) ProtoMessage()               {}
func (*DestroyTokenArgs) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DestroyTokenArgs) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *DestroyTokenArgs) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *DestroyTokenArgs) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

func (m *DestroyTokenArgs) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *DestroyTokenArgs) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *DestroyTokenArgs) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type DestroyTokenRes struct {
	Ret int32 `protobuf:"varint,1,opt,name=Ret" json:"Ret,omitempty"`
}

func (m *DestroyTokenRes) Reset()                    { *m = DestroyTokenRes{} }
func (m *DestroyTokenRes) String() string            { return proto1.CompactTextString(m) }
func (*DestroyTokenRes) ProtoMessage()               {}
func (*DestroyTokenRes) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DestroyTokenRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func init() {
	proto1.RegisterType((*AuthArgs)(nil), "proto.AuthArgs")
	proto1.RegisterType((*AuthRes)(nil), "proto.AuthRes")
	proto1.RegisterType((*TokenInfoArgs)(nil), "proto.TokenInfoArgs")
	proto1.RegisterType((*TokenInfoRes)(nil), "proto.TokenInfoRes")
	proto1.RegisterType((*RefreshTokenArgs)(nil), "proto.RefreshTokenArgs")
	proto1.RegisterType((*RefreshTokenRes)(nil), "proto.RefreshTokenRes")
	proto1.RegisterType((*DestroyTokenArgs)(nil), "proto.DestroyTokenArgs")
	proto1.RegisterType((*DestroyTokenRes)(nil), "proto.DestroyTokenRes")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthService service

type AuthServiceClient interface {
	Auth(ctx context.Context, in *AuthArgs, opts ...grpc.CallOption) (*AuthRes, error)
	TokenInfo(ctx context.Context, in *TokenInfoArgs, opts ...grpc.CallOption) (*TokenInfoRes, error)
	RefreshToken(ctx context.Context, in *RefreshTokenArgs, opts ...grpc.CallOption) (*RefreshTokenRes, error)
	DestroyToken(ctx context.Context, in *DestroyTokenArgs, opts ...grpc.CallOption) (*DestroyTokenRes, error)
}

type authServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthServiceClient(cc *grpc.ClientConn) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Auth(ctx context.Context, in *AuthArgs, opts ...grpc.CallOption) (*AuthRes, error) {
	out := new(AuthRes)
	err := grpc.Invoke(ctx, "/proto.AuthService/Auth", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) TokenInfo(ctx context.Context, in *TokenInfoArgs, opts ...grpc.CallOption) (*TokenInfoRes, error) {
	out := new(TokenInfoRes)
	err := grpc.Invoke(ctx, "/proto.AuthService/TokenInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenArgs, opts ...grpc.CallOption) (*RefreshTokenRes, error) {
	out := new(RefreshTokenRes)
	err := grpc.Invoke(ctx, "/proto.AuthService/RefreshToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authServiceClient) DestroyToken(ctx context.Context, in *DestroyTokenArgs, opts ...grpc.CallOption) (*DestroyTokenRes, error) {
	out := new(DestroyTokenRes)
	err := grpc.Invoke(ctx, "/proto.AuthService/DestroyToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceServer interface {
	Auth(context.Context, *AuthArgs) (*AuthRes, error)
	TokenInfo(context.Context, *TokenInfoArgs) (*TokenInfoRes, error)
	RefreshToken(context.Context, *RefreshTokenArgs) (*RefreshTokenRes, error)
	DestroyToken(context.Context, *DestroyTokenArgs) (*DestroyTokenRes, error)
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Auth(ctx, req.(*AuthArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_TokenInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInfoArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).TokenInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/TokenInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).TokenInfo(ctx, req.(*TokenInfoArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).RefreshToken(ctx, req.(*RefreshTokenArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthService_DestroyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyTokenArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).DestroyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AuthService/DestroyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).DestroyToken(ctx, req.(*DestroyTokenArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _AuthService_Auth_Handler,
		},
		{
			MethodName: "TokenInfo",
			Handler:    _AuthService_TokenInfo_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthService_RefreshToken_Handler,
		},
		{
			MethodName: "DestroyToken",
			Handler:    _AuthService_DestroyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}

func init() { proto1.RegisterFile("auth.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 384 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0x66, 0x2d, 0x2d, 0x30, 0xa2, 0x90, 0x95, 0xe0, 0xa6, 0x07, 0x43, 0xd6, 0x0b, 0x1e, 0xe4,
	0xa0, 0x17, 0xae, 0x8d, 0x3f, 0x09, 0x26, 0x5e, 0xaa, 0x78, 0x47, 0x19, 0x28, 0x51, 0xdb, 0x75,
	0x5b, 0x8c, 0x3e, 0x8d, 0x0f, 0x60, 0xe2, 0xc3, 0xf8, 0x34, 0x1e, 0xcd, 0x6e, 0x4b, 0xd3, 0x56,
	0x24, 0x7a, 0xd3, 0x53, 0xf7, 0x9b, 0xe9, 0x37, 0xfb, 0xcd, 0xb7, 0x33, 0x00, 0xa3, 0x79, 0xe4,
	0xf5, 0x84, 0x0c, 0xa2, 0x80, 0x9a, 0xfa, 0xc3, 0xfb, 0x50, 0x75, 0xe6, 0x91, 0xe7, 0xc8, 0x69,
	0x48, 0xdb, 0x60, 0x0d, 0x43, 0x94, 0x83, 0x31, 0x23, 0x1d, 0xd2, 0x35, 0xdd, 0x04, 0xd1, 0x16,
	0x98, 0x97, 0xc1, 0x2d, 0xfa, 0xcc, 0xea, 0x90, 0x6e, 0xcd, 0x8d, 0x01, 0xdf, 0x87, 0x8a, 0x62,
	0xba, 0x18, 0xd2, 0x26, 0x18, 0x2e, 0x46, 0x09, 0x4b, 0x1d, 0x55, 0xe4, 0x3c, 0x9c, 0xb2, 0x35,
	0x4d, 0x50, 0x47, 0xfe, 0x4a, 0x60, 0x43, 0x13, 0x07, 0xfe, 0x24, 0x58, 0x79, 0x1d, 0x83, 0x8a,
	0x23, 0xc4, 0xa9, 0x0c, 0xee, 0x13, 0xfe, 0x02, 0xd2, 0x1d, 0x00, 0x47, 0x88, 0x23, 0x6f, 0xe4,
	0xfb, 0x78, 0xc7, 0x0c, 0x9d, 0xcc, 0x44, 0x54, 0x45, 0x47, 0x88, 0x2b, 0x94, 0xac, 0x1c, 0x57,
	0x8c, 0x11, 0xb5, 0xa1, 0x3a, 0xf4, 0x67, 0x0f, 0x73, 0x1c, 0x8c, 0x99, 0xa9, 0x59, 0x29, 0xfe,
	0xa6, 0xb9, 0x33, 0xa8, 0xa7, 0x62, 0x7f, 0xd8, 0xa1, 0xba, 0xfd, 0xe4, 0x49, 0xcc, 0x24, 0x6a,
	0x65, 0x86, 0x9b, 0x20, 0xfe, 0x4e, 0xa0, 0xe9, 0xe2, 0x44, 0x62, 0xe8, 0xe9, 0x9a, 0x7f, 0xbd,
	0x79, 0xca, 0xa1, 0x9e, 0xd5, 0xcb, 0x2a, 0x3a, 0x99, 0x8b, 0xf1, 0x17, 0x02, 0x8d, 0x6c, 0xe0,
	0x17, 0x26, 0x25, 0x7d, 0x1b, 0xcb, 0x67, 0xac, 0xbc, 0x4a, 0x89, 0xf9, 0x55, 0x49, 0xc6, 0x76,
	0x2b, 0x67, 0xfb, 0x1b, 0x81, 0xe6, 0x31, 0x86, 0x91, 0x0c, 0x9e, 0xff, 0x85, 0xed, 0x7c, 0x17,
	0x1a, 0x59, 0xbd, 0x4b, 0x1d, 0x3d, 0xf8, 0x20, 0xb0, 0xae, 0xd6, 0xee, 0x02, 0xe5, 0xe3, 0xec,
	0x06, 0xe9, 0x1e, 0x94, 0x15, 0xa4, 0x8d, 0x78, 0xad, 0x7b, 0x8b, 0x65, 0xb6, 0x37, 0x33, 0x01,
	0x17, 0x43, 0x5e, 0xa2, 0x7d, 0xa8, 0xa5, 0x33, 0x4d, 0x5b, 0x49, 0x3a, 0xb7, 0x92, 0xf6, 0x56,
	0x31, 0x1a, 0x33, 0x9d, 0xfc, 0x33, 0xd0, 0xed, 0xe4, 0xb7, 0xe2, 0x54, 0xdb, 0xed, 0x25, 0x89,
	0xb4, 0x44, 0xb6, 0xb9, 0xb4, 0x44, 0xf1, 0x85, 0xd2, 0x12, 0x05, 0x2b, 0x78, 0xe9, 0xda, 0xd2,
	0x89, 0xc3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdd, 0xcc, 0xec, 0x71, 0xc6, 0x04, 0x00, 0x00,
}