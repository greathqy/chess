// Code generated by protoc-gen-go. DO NOT EDIT.
// source: centre.proto

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

type BaseRes struct {
	Ret int32  `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *BaseRes) Reset()                    { *m = BaseRes{} }
func (m *BaseRes) String() string            { return proto1.CompactTextString(m) }
func (*BaseRes) ProtoMessage()               {}
func (*BaseRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *BaseRes) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *BaseRes) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type UpdateRoomInfoArgs struct {
	ServiceId string    `protobuf:"bytes,1,opt,name=service_id,json=serviceId" json:"service_id,omitempty"`
	RoomId    int32     `protobuf:"varint,2,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
	RoomInfo  *RoomInfo `protobuf:"bytes,3,opt,name=room_info,json=roomInfo" json:"room_info,omitempty"`
}

func (m *UpdateRoomInfoArgs) Reset()                    { *m = UpdateRoomInfoArgs{} }
func (m *UpdateRoomInfoArgs) String() string            { return proto1.CompactTextString(m) }
func (*UpdateRoomInfoArgs) ProtoMessage()               {}
func (*UpdateRoomInfoArgs) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *UpdateRoomInfoArgs) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *UpdateRoomInfoArgs) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *UpdateRoomInfoArgs) GetRoomInfo() *RoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

type RoomListArgs struct {
}

func (m *RoomListArgs) Reset()                    { *m = RoomListArgs{} }
func (m *RoomListArgs) String() string            { return proto1.CompactTextString(m) }
func (*RoomListArgs) ProtoMessage()               {}
func (*RoomListArgs) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

type RoomListRes struct {
	List map[int32]*RoomInfo `protobuf:"bytes,1,rep,name=list" json:"list,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *RoomListRes) Reset()                    { *m = RoomListRes{} }
func (m *RoomListRes) String() string            { return proto1.CompactTextString(m) }
func (*RoomListRes) ProtoMessage()               {}
func (*RoomListRes) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *RoomListRes) GetList() map[int32]*RoomInfo {
	if m != nil {
		return m.List
	}
	return nil
}

type RoomInfo struct {
	TableNumber  int32 `protobuf:"varint,1,opt,name=table_number,json=tableNumber" json:"table_number,omitempty"`
	PlayerNumber int32 `protobuf:"varint,2,opt,name=player_number,json=playerNumber" json:"player_number,omitempty"`
}

func (m *RoomInfo) Reset()                    { *m = RoomInfo{} }
func (m *RoomInfo) String() string            { return proto1.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()               {}
func (*RoomInfo) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *RoomInfo) GetTableNumber() int32 {
	if m != nil {
		return m.TableNumber
	}
	return 0
}

func (m *RoomInfo) GetPlayerNumber() int32 {
	if m != nil {
		return m.PlayerNumber
	}
	return 0
}

func init() {
	proto1.RegisterType((*BaseRes)(nil), "proto.BaseRes")
	proto1.RegisterType((*UpdateRoomInfoArgs)(nil), "proto.UpdateRoomInfoArgs")
	proto1.RegisterType((*RoomListArgs)(nil), "proto.RoomListArgs")
	proto1.RegisterType((*RoomListRes)(nil), "proto.RoomListRes")
	proto1.RegisterType((*RoomInfo)(nil), "proto.RoomInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CentreService service

type CentreServiceClient interface {
	// 获取所有房间信息
	RoomList(ctx context.Context, in *RoomListArgs, opts ...grpc.CallOption) (*RoomListRes, error)
	// 更新房间信息
	UpdateRoomInfo(ctx context.Context, in *UpdateRoomInfoArgs, opts ...grpc.CallOption) (*BaseRes, error)
}

type centreServiceClient struct {
	cc *grpc.ClientConn
}

func NewCentreServiceClient(cc *grpc.ClientConn) CentreServiceClient {
	return &centreServiceClient{cc}
}

func (c *centreServiceClient) RoomList(ctx context.Context, in *RoomListArgs, opts ...grpc.CallOption) (*RoomListRes, error) {
	out := new(RoomListRes)
	err := grpc.Invoke(ctx, "/proto.CentreService/RoomList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centreServiceClient) UpdateRoomInfo(ctx context.Context, in *UpdateRoomInfoArgs, opts ...grpc.CallOption) (*BaseRes, error) {
	out := new(BaseRes)
	err := grpc.Invoke(ctx, "/proto.CentreService/UpdateRoomInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CentreService service

type CentreServiceServer interface {
	// 获取所有房间信息
	RoomList(context.Context, *RoomListArgs) (*RoomListRes, error)
	// 更新房间信息
	UpdateRoomInfo(context.Context, *UpdateRoomInfoArgs) (*BaseRes, error)
}

func RegisterCentreServiceServer(s *grpc.Server, srv CentreServiceServer) {
	s.RegisterService(&_CentreService_serviceDesc, srv)
}

func _CentreService_RoomList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoomListArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentreServiceServer).RoomList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CentreService/RoomList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentreServiceServer).RoomList(ctx, req.(*RoomListArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentreService_UpdateRoomInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRoomInfoArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentreServiceServer).UpdateRoomInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.CentreService/UpdateRoomInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentreServiceServer).UpdateRoomInfo(ctx, req.(*UpdateRoomInfoArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _CentreService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CentreService",
	HandlerType: (*CentreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RoomList",
			Handler:    _CentreService_RoomList_Handler,
		},
		{
			MethodName: "UpdateRoomInfo",
			Handler:    _CentreService_UpdateRoomInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "centre.proto",
}

func init() { proto1.RegisterFile("centre.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xcd, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x9b, 0xf6, 0xf5, 0x23, 0x37, 0x69, 0xdf, 0xe3, 0xbe, 0xc5, 0xeb, 0x2b, 0x0a, 0x35,
	0x22, 0x64, 0xa1, 0x45, 0x2a, 0x82, 0x08, 0x2e, 0x54, 0x04, 0x0b, 0xe2, 0x62, 0xc4, 0x75, 0x49,
	0x9b, 0xdb, 0x12, 0x9a, 0x64, 0xca, 0x64, 0x5a, 0x88, 0x7b, 0x17, 0xfe, 0xd7, 0x32, 0x1f, 0x11,
	0xb5, 0xae, 0x66, 0xf2, 0xbb, 0x97, 0x33, 0xe7, 0xdc, 0x1b, 0xf0, 0xe7, 0x94, 0x4b, 0x41, 0xa3,
	0xb5, 0xe0, 0x92, 0x63, 0x53, 0x1f, 0xc1, 0x09, 0xb4, 0x6f, 0xa2, 0x82, 0x18, 0x15, 0xf8, 0x07,
	0x1a, 0x82, 0x64, 0xdf, 0x19, 0x3a, 0x61, 0x93, 0xa9, 0xab, 0x22, 0x59, 0xb1, 0xec, 0xd7, 0x87,
	0x4e, 0xe8, 0x32, 0x75, 0x0d, 0x5e, 0x00, 0x9f, 0xd7, 0x71, 0x24, 0x89, 0x71, 0x9e, 0x4d, 0xf2,
	0x05, 0xbf, 0x16, 0xcb, 0x02, 0xf7, 0x01, 0x0a, 0x12, 0xdb, 0x64, 0x4e, 0xd3, 0x24, 0xd6, 0x02,
	0x2e, 0x73, 0x2d, 0x99, 0xc4, 0xf8, 0x0f, 0xda, 0x82, 0xf3, 0x4c, 0xd5, 0xea, 0x5a, 0xbc, 0xa5,
	0x3e, 0x27, 0x31, 0x1e, 0x83, 0x6b, 0x0a, 0xf9, 0x82, 0xf7, 0x1b, 0x43, 0x27, 0xf4, 0xc6, 0xbf,
	0x8d, 0xbd, 0x51, 0xa5, 0xcf, 0x3a, 0xc2, 0xde, 0x82, 0x1e, 0xf8, 0x8a, 0x3e, 0x24, 0x85, 0x54,
	0xaf, 0x06, 0x6f, 0x0e, 0x78, 0x15, 0x50, 0xfe, 0x4f, 0xe1, 0x57, 0x9a, 0x14, 0x2a, 0x40, 0x23,
	0xf4, 0xc6, 0x7b, 0x9f, 0x84, 0x6c, 0xc7, 0x48, 0x9d, 0x77, 0xb9, 0x14, 0x25, 0xd3, 0x9d, 0x83,
	0x7b, 0x70, 0x3f, 0x90, 0x0a, 0xbb, 0xa2, 0xb2, 0x8a, 0xbf, 0xa2, 0x12, 0x8f, 0xa0, 0xb9, 0x8d,
	0xd2, 0x0d, 0x69, 0xd7, 0x3f, 0x58, 0x33, 0xd5, 0xcb, 0xfa, 0x85, 0x13, 0x30, 0xe8, 0x54, 0x18,
	0x0f, 0xc0, 0x97, 0xd1, 0x2c, 0xa5, 0x69, 0xbe, 0xc9, 0x66, 0x24, 0xac, 0xa2, 0xa7, 0xd9, 0xa3,
	0x46, 0x78, 0x08, 0xdd, 0x75, 0x1a, 0x95, 0x24, 0xaa, 0x1e, 0x33, 0x17, 0xdf, 0x40, 0xd3, 0x34,
	0x7e, 0x75, 0xa0, 0x7b, 0xab, 0x57, 0xf6, 0x64, 0x46, 0x89, 0xe7, 0xe6, 0x15, 0xe5, 0x19, 0xff,
	0x7e, 0xcb, 0xa7, 0x46, 0x32, 0xc0, 0xdd, 0xd0, 0x41, 0x0d, 0xaf, 0xa0, 0xf7, 0x75, 0x69, 0xf8,
	0xdf, 0xf6, 0xed, 0xee, 0x72, 0xd0, 0xb3, 0x25, 0xfb, 0x57, 0x04, 0xb5, 0x59, 0x4b, 0x83, 0xb3,
	0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0x9a, 0xc3, 0x38, 0x40, 0x02, 0x00, 0x00,
}