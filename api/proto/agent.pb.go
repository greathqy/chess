// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	agent.proto
	auth.proto
	centre.proto
	game.proto
	room.proto

It has these top-level messages:
	AutoId
	SeedInfo
	BaseAck
	BaseReq
	UserLoginReq
	UserLoginAck
	RoomSetTableReq
	RoomSetTableAck
	RoomGetTableReq
	RoomGetTableAck
	RoomGetPlayerReq
	RoomGetPlayerAck
	RoomPlayerJoinReq
	RoomPlayerJoinAck
	RoomPlayerGoneReq
	RoomPlayerGoneAck
	RoomPlayerBetReq
	RoomPlayerBetAck
	RoomButtonAck
	RoomDealAck
	RoomPotAck
	RoomActionAck
	RoomShowdownAck
	TableInfo
	CardInfo
	PlayerInfo
	TokenInfoArgs
	TokenInfoRes
	RefreshTokenArgs
	RefreshTokenRes
	DestroyTokenArgs
	DestroyTokenRes
	GameListArgs
	GameListRes
	GamePlaygroundArgs
	GamePlaygroundRes
	Room
	RoomInfoArgs
	RoomInfoRes
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

// 0,1 心跳包
type AutoId struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *AutoId) Reset()                    { *m = AutoId{} }
func (m *AutoId) String() string            { return proto1.CompactTextString(m) }
func (*AutoId) ProtoMessage()               {}
func (*AutoId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AutoId) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// 30,31 通信加密种子
type SeedInfo struct {
	ClientSendSeed    int32 `protobuf:"varint,1,opt,name=client_send_seed,json=clientSendSeed" json:"client_send_seed,omitempty"`
	ClientReceiveSeed int32 `protobuf:"varint,2,opt,name=client_receive_seed,json=clientReceiveSeed" json:"client_receive_seed,omitempty"`
}

func (m *SeedInfo) Reset()                    { *m = SeedInfo{} }
func (m *SeedInfo) String() string            { return proto1.CompactTextString(m) }
func (*SeedInfo) ProtoMessage()               {}
func (*SeedInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SeedInfo) GetClientSendSeed() int32 {
	if m != nil {
		return m.ClientSendSeed
	}
	return 0
}

func (m *SeedInfo) GetClientReceiveSeed() int32 {
	if m != nil {
		return m.ClientReceiveSeed
	}
	return 0
}

// 一般性回复payload
type BaseAck struct {
	Ret int32  `protobuf:"varint,1,opt,name=ret" json:"ret,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *BaseAck) Reset()                    { *m = BaseAck{} }
func (m *BaseAck) String() string            { return proto1.CompactTextString(m) }
func (*BaseAck) ProtoMessage()               {}
func (*BaseAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BaseAck) GetRet() int32 {
	if m != nil {
		return m.Ret
	}
	return 0
}

func (m *BaseAck) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

// 一般性请求payload
type BaseReq struct {
	AppFrom    string `protobuf:"bytes,1,opt,name=app_from,json=appFrom" json:"app_from,omitempty"`
	AppVer     int32  `protobuf:"varint,2,opt,name=app_ver,json=appVer" json:"app_ver,omitempty"`
	AppChannel string `protobuf:"bytes,3,opt,name=app_channel,json=appChannel" json:"app_channel,omitempty"`
}

func (m *BaseReq) Reset()                    { *m = BaseReq{} }
func (m *BaseReq) String() string            { return proto1.CompactTextString(m) }
func (*BaseReq) ProtoMessage()               {}
func (*BaseReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *BaseReq) GetAppFrom() string {
	if m != nil {
		return m.AppFrom
	}
	return ""
}

func (m *BaseReq) GetAppVer() int32 {
	if m != nil {
		return m.AppVer
	}
	return 0
}

func (m *BaseReq) GetAppChannel() string {
	if m != nil {
		return m.AppChannel
	}
	return ""
}

// 10, 用户登录
type UserLoginReq struct {
	BaseReq  *BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq" json:"base_req,omitempty"`
	UserId   int32    `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UniqueId string   `protobuf:"bytes,3,opt,name=unique_id,json=uniqueId" json:"unique_id,omitempty"`
	Token    string   `protobuf:"bytes,4,opt,name=token" json:"token,omitempty"`
}

func (m *UserLoginReq) Reset()                    { *m = UserLoginReq{} }
func (m *UserLoginReq) String() string            { return proto1.CompactTextString(m) }
func (*UserLoginReq) ProtoMessage()               {}
func (*UserLoginReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UserLoginReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *UserLoginReq) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserLoginReq) GetUniqueId() string {
	if m != nil {
		return m.UniqueId
	}
	return ""
}

func (m *UserLoginReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// 11, 用户登录回复
type UserLoginAck struct {
	BaseAck *BaseAck `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
}

func (m *UserLoginAck) Reset()                    { *m = UserLoginAck{} }
func (m *UserLoginAck) String() string            { return proto1.CompactTextString(m) }
func (*UserLoginAck) ProtoMessage()               {}
func (*UserLoginAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *UserLoginAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

// 2003, 创建牌桌
type RoomSetTableReq struct {
	BaseReq    *BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq" json:"base_req,omitempty"`
	SmallBlind int32    `protobuf:"varint,2,opt,name=small_blind,json=smallBlind" json:"small_blind,omitempty"`
	BigBlind   int32    `protobuf:"varint,3,opt,name=big_blind,json=bigBlind" json:"big_blind,omitempty"`
	Timeout    int32    `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
	Max        int32    `protobuf:"varint,5,opt,name=max" json:"max,omitempty"`
}

func (m *RoomSetTableReq) Reset()                    { *m = RoomSetTableReq{} }
func (m *RoomSetTableReq) String() string            { return proto1.CompactTextString(m) }
func (*RoomSetTableReq) ProtoMessage()               {}
func (*RoomSetTableReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RoomSetTableReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *RoomSetTableReq) GetSmallBlind() int32 {
	if m != nil {
		return m.SmallBlind
	}
	return 0
}

func (m *RoomSetTableReq) GetBigBlind() int32 {
	if m != nil {
		return m.BigBlind
	}
	return 0
}

func (m *RoomSetTableReq) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *RoomSetTableReq) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

// 2004, 创建牌桌回复
type RoomSetTableAck struct {
	BaseAck *BaseAck   `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	Table   *TableInfo `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
}

func (m *RoomSetTableAck) Reset()                    { *m = RoomSetTableAck{} }
func (m *RoomSetTableAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomSetTableAck) ProtoMessage()               {}
func (*RoomSetTableAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RoomSetTableAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomSetTableAck) GetTable() *TableInfo {
	if m != nil {
		return m.Table
	}
	return nil
}

// 2005, 查询牌桌信息
type RoomGetTableReq struct {
	BaseReq *BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq" json:"base_req,omitempty"`
	RoomId  int32    `protobuf:"varint,2,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
	TableId string   `protobuf:"bytes,3,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
}

func (m *RoomGetTableReq) Reset()                    { *m = RoomGetTableReq{} }
func (m *RoomGetTableReq) String() string            { return proto1.CompactTextString(m) }
func (*RoomGetTableReq) ProtoMessage()               {}
func (*RoomGetTableReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RoomGetTableReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *RoomGetTableReq) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RoomGetTableReq) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

// 2006, 查询牌桌信息回复 (当玩家加入牌桌后，服务器会向此用户推送牌桌信息)
type RoomGetTableAck struct {
	BaseAck *BaseAck   `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	Table   *TableInfo `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
}

func (m *RoomGetTableAck) Reset()                    { *m = RoomGetTableAck{} }
func (m *RoomGetTableAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomGetTableAck) ProtoMessage()               {}
func (*RoomGetTableAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RoomGetTableAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomGetTableAck) GetTable() *TableInfo {
	if m != nil {
		return m.Table
	}
	return nil
}

// 2007, 查询玩家信息
type RoomGetPlayerReq struct {
	BaseReq  *BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq" json:"base_req,omitempty"`
	PlayerId int32    `protobuf:"varint,2,opt,name=player_id,json=playerId" json:"player_id,omitempty"`
}

func (m *RoomGetPlayerReq) Reset()                    { *m = RoomGetPlayerReq{} }
func (m *RoomGetPlayerReq) String() string            { return proto1.CompactTextString(m) }
func (*RoomGetPlayerReq) ProtoMessage()               {}
func (*RoomGetPlayerReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RoomGetPlayerReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *RoomGetPlayerReq) GetPlayerId() int32 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

// 2008, 查询玩家信息回复
type RoomGetPlayerAck struct {
	BaseAck *BaseAck    `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	Player  *PlayerInfo `protobuf:"bytes,2,opt,name=player" json:"player,omitempty"`
}

func (m *RoomGetPlayerAck) Reset()                    { *m = RoomGetPlayerAck{} }
func (m *RoomGetPlayerAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomGetPlayerAck) ProtoMessage()               {}
func (*RoomGetPlayerAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *RoomGetPlayerAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomGetPlayerAck) GetPlayer() *PlayerInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

// 2101, 玩家加入游戏
type RoomPlayerJoinReq struct {
	BaseReq *BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq" json:"base_req,omitempty"`
	RoomId  int32    `protobuf:"varint,2,opt,name=room_id,json=roomId" json:"room_id,omitempty"`
	TableId string   `protobuf:"bytes,3,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
}

func (m *RoomPlayerJoinReq) Reset()                    { *m = RoomPlayerJoinReq{} }
func (m *RoomPlayerJoinReq) String() string            { return proto1.CompactTextString(m) }
func (*RoomPlayerJoinReq) ProtoMessage()               {}
func (*RoomPlayerJoinReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *RoomPlayerJoinReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *RoomPlayerJoinReq) GetRoomId() int32 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RoomPlayerJoinReq) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

// 2102, 通报加入游戏的玩家
type RoomPlayerJoinAck struct {
	BaseAck *BaseAck    `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	Player  *PlayerInfo `protobuf:"bytes,2,opt,name=player" json:"player,omitempty"`
}

func (m *RoomPlayerJoinAck) Reset()                    { *m = RoomPlayerJoinAck{} }
func (m *RoomPlayerJoinAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomPlayerJoinAck) ProtoMessage()               {}
func (*RoomPlayerJoinAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *RoomPlayerJoinAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomPlayerJoinAck) GetPlayer() *PlayerInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

// 2103, 玩家离开牌桌
type RoomPlayerGoneReq struct {
	BaseReq *BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq" json:"base_req,omitempty"`
	TableId string   `protobuf:"bytes,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
}

func (m *RoomPlayerGoneReq) Reset()                    { *m = RoomPlayerGoneReq{} }
func (m *RoomPlayerGoneReq) String() string            { return proto1.CompactTextString(m) }
func (*RoomPlayerGoneReq) ProtoMessage()               {}
func (*RoomPlayerGoneReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *RoomPlayerGoneReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *RoomPlayerGoneReq) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

// 2104, 服务器广播离开房间的玩家
type RoomPlayerGoneAck struct {
	BaseAck *BaseAck    `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	Player  *PlayerInfo `protobuf:"bytes,2,opt,name=player" json:"player,omitempty"`
}

func (m *RoomPlayerGoneAck) Reset()                    { *m = RoomPlayerGoneAck{} }
func (m *RoomPlayerGoneAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomPlayerGoneAck) ProtoMessage()               {}
func (*RoomPlayerGoneAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *RoomPlayerGoneAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomPlayerGoneAck) GetPlayer() *PlayerInfo {
	if m != nil {
		return m.Player
	}
	return nil
}

// 2105, 玩家下注
// 玩家有四种下注方式，下注数分别对应为：
// 弃牌：< 0 或空 表示弃牌 (fold)
// 看注：= 0 表示看注 (check)
// 跟注：小于等于单注额 (call)
// 加注：大于单注额 (raise)
// 全押：等于玩家手中所有筹码 (allin)
type RoomPlayerBetReq struct {
	BaseReq *BaseReq `protobuf:"bytes,1,opt,name=base_req,json=baseReq" json:"base_req,omitempty"`
	TableId string   `protobuf:"bytes,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	Bet     int32    `protobuf:"varint,3,opt,name=bet" json:"bet,omitempty"`
}

func (m *RoomPlayerBetReq) Reset()                    { *m = RoomPlayerBetReq{} }
func (m *RoomPlayerBetReq) String() string            { return proto1.CompactTextString(m) }
func (*RoomPlayerBetReq) ProtoMessage()               {}
func (*RoomPlayerBetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *RoomPlayerBetReq) GetBaseReq() *BaseReq {
	if m != nil {
		return m.BaseReq
	}
	return nil
}

func (m *RoomPlayerBetReq) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

func (m *RoomPlayerBetReq) GetBet() int32 {
	if m != nil {
		return m.Bet
	}
	return 0
}

// 2106, 玩家下注回复
type RoomPlayerBetAck struct {
	BaseAck *BaseAck `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	TableId string   `protobuf:"bytes,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	Action  string   `protobuf:"bytes,3,opt,name=action" json:"action,omitempty"`
	Bet     int32    `protobuf:"varint,4,opt,name=bet" json:"bet,omitempty"`
	Chips   int32    `protobuf:"varint,5,opt,name=chips" json:"chips,omitempty"`
}

func (m *RoomPlayerBetAck) Reset()                    { *m = RoomPlayerBetAck{} }
func (m *RoomPlayerBetAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomPlayerBetAck) ProtoMessage()               {}
func (*RoomPlayerBetAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *RoomPlayerBetAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomPlayerBetAck) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

func (m *RoomPlayerBetAck) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RoomPlayerBetAck) GetBet() int32 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func (m *RoomPlayerBetAck) GetChips() int32 {
	if m != nil {
		return m.Chips
	}
	return 0
}

// 2107, 通报本局庄家 (服务器广播此消息，代表游戏开始并确定本局庄家)
type RoomButtonAck struct {
	BaseAck   *BaseAck `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	TableId   string   `protobuf:"bytes,2,opt,name=table_id,json=tableId" json:"table_id,omitempty"`
	ButtonPos int32    `protobuf:"varint,3,opt,name=button_pos,json=buttonPos" json:"button_pos,omitempty"`
}

func (m *RoomButtonAck) Reset()                    { *m = RoomButtonAck{} }
func (m *RoomButtonAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomButtonAck) ProtoMessage()               {}
func (*RoomButtonAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *RoomButtonAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomButtonAck) GetTableId() string {
	if m != nil {
		return m.TableId
	}
	return ""
}

func (m *RoomButtonAck) GetButtonPos() int32 {
	if m != nil {
		return m.ButtonPos
	}
	return 0
}

// 2108, 发牌 - 共有四轮发牌，按顺序分别为：preflop (底牌), flop (翻牌), turn (转牌), river(河牌)
type RoomDealAck struct {
	BaseAck *BaseAck    `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	Action  string      `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Cards   []*CardInfo `protobuf:"bytes,3,rep,name=cards" json:"cards,omitempty"`
	//
	// 这套手牌的权重等级，一共有10个等级
	// 皇家同花顺：10
	// 同花顺    ：9
	// 四条      ：8
	// 葫芦      ：7
	// 同花      ：6
	// 顺子      ：5
	// 三条      ：4
	// 两对      ：3
	// 一对      ：2
	// 高牌      ：1
	HandLevel int32 `protobuf:"varint,4,opt,name=hand_level,json=handLevel" json:"hand_level,omitempty"`
}

func (m *RoomDealAck) Reset()                    { *m = RoomDealAck{} }
func (m *RoomDealAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomDealAck) ProtoMessage()               {}
func (*RoomDealAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *RoomDealAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomDealAck) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *RoomDealAck) GetCards() []*CardInfo {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *RoomDealAck) GetHandLevel() int32 {
	if m != nil {
		return m.HandLevel
	}
	return 0
}

// 2109, 通报奖池
type RoomPotAck struct {
	BaseAck *BaseAck `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	Pot     []int32  `protobuf:"varint,2,rep,packed,name=pot" json:"pot,omitempty"`
}

func (m *RoomPotAck) Reset()                    { *m = RoomPotAck{} }
func (m *RoomPotAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomPotAck) ProtoMessage()               {}
func (*RoomPotAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *RoomPotAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomPotAck) GetPot() []int32 {
	if m != nil {
		return m.Pot
	}
	return nil
}

// 2110, 通报当前下注玩家
type RoomActionAck struct {
	BaseAck *BaseAck `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	Pos     int32    `protobuf:"varint,2,opt,name=pos" json:"pos,omitempty"`
	BaseBet int32    `protobuf:"varint,3,opt,name=base_bet,json=baseBet" json:"base_bet,omitempty"`
}

func (m *RoomActionAck) Reset()                    { *m = RoomActionAck{} }
func (m *RoomActionAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomActionAck) ProtoMessage()               {}
func (*RoomActionAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *RoomActionAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomActionAck) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *RoomActionAck) GetBaseBet() int32 {
	if m != nil {
		return m.BaseBet
	}
	return 0
}

// 2111, 摊牌和比牌
type RoomShowdownAck struct {
	BaseAck *BaseAck   `protobuf:"bytes,1,opt,name=base_ack,json=baseAck" json:"base_ack,omitempty"`
	Table   *TableInfo `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"`
}

func (m *RoomShowdownAck) Reset()                    { *m = RoomShowdownAck{} }
func (m *RoomShowdownAck) String() string            { return proto1.CompactTextString(m) }
func (*RoomShowdownAck) ProtoMessage()               {}
func (*RoomShowdownAck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *RoomShowdownAck) GetBaseAck() *BaseAck {
	if m != nil {
		return m.BaseAck
	}
	return nil
}

func (m *RoomShowdownAck) GetTable() *TableInfo {
	if m != nil {
		return m.Table
	}
	return nil
}

// 牌桌信息
type TableInfo struct {
	Id         string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	SmallBlind int32         `protobuf:"varint,2,opt,name=small_blind,json=smallBlind" json:"small_blind,omitempty"`
	BigBlind   int32         `protobuf:"varint,3,opt,name=big_blind,json=bigBlind" json:"big_blind,omitempty"`
	Bet        int32         `protobuf:"varint,4,opt,name=bet" json:"bet,omitempty"`
	Timeout    int32         `protobuf:"varint,5,opt,name=timeout" json:"timeout,omitempty"`
	Cards      []*CardInfo   `protobuf:"bytes,6,rep,name=cards" json:"cards,omitempty"`
	Pot        []int32       `protobuf:"varint,7,rep,packed,name=pot" json:"pot,omitempty"`
	Chips      []int32       `protobuf:"varint,8,rep,packed,name=chips" json:"chips,omitempty"`
	Button     int32         `protobuf:"varint,9,opt,name=button" json:"button,omitempty"`
	N          int32         `protobuf:"varint,10,opt,name=n" json:"n,omitempty"`
	Max        int32         `protobuf:"varint,11,opt,name=max" json:"max,omitempty"`
	Players    []*PlayerInfo `protobuf:"bytes,12,rep,name=players" json:"players,omitempty"`
}

func (m *TableInfo) Reset()                    { *m = TableInfo{} }
func (m *TableInfo) String() string            { return proto1.CompactTextString(m) }
func (*TableInfo) ProtoMessage()               {}
func (*TableInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *TableInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TableInfo) GetSmallBlind() int32 {
	if m != nil {
		return m.SmallBlind
	}
	return 0
}

func (m *TableInfo) GetBigBlind() int32 {
	if m != nil {
		return m.BigBlind
	}
	return 0
}

func (m *TableInfo) GetBet() int32 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func (m *TableInfo) GetTimeout() int32 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *TableInfo) GetCards() []*CardInfo {
	if m != nil {
		return m.Cards
	}
	return nil
}

func (m *TableInfo) GetPot() []int32 {
	if m != nil {
		return m.Pot
	}
	return nil
}

func (m *TableInfo) GetChips() []int32 {
	if m != nil {
		return m.Chips
	}
	return nil
}

func (m *TableInfo) GetButton() int32 {
	if m != nil {
		return m.Button
	}
	return 0
}

func (m *TableInfo) GetN() int32 {
	if m != nil {
		return m.N
	}
	return 0
}

func (m *TableInfo) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *TableInfo) GetPlayers() []*PlayerInfo {
	if m != nil {
		return m.Players
	}
	return nil
}

// 牌
type CardInfo struct {
	Suit int32 `protobuf:"varint,1,opt,name=suit" json:"suit,omitempty"`
	Val  int32 `protobuf:"varint,2,opt,name=val" json:"val,omitempty"`
}

func (m *CardInfo) Reset()                    { *m = CardInfo{} }
func (m *CardInfo) String() string            { return proto1.CompactTextString(m) }
func (*CardInfo) ProtoMessage()               {}
func (*CardInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{24} }

func (m *CardInfo) GetSuit() int32 {
	if m != nil {
		return m.Suit
	}
	return 0
}

func (m *CardInfo) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

// 玩家信息
type PlayerInfo struct {
	Pos      int32       `protobuf:"varint,1,opt,name=pos" json:"pos,omitempty"`
	Id       int32       `protobuf:"varint,2,opt,name=id" json:"id,omitempty"`
	Nickname string      `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	Avatar   string      `protobuf:"bytes,4,opt,name=avatar" json:"avatar,omitempty"`
	Level    string      `protobuf:"bytes,5,opt,name=level" json:"level,omitempty"`
	Chips    int32       `protobuf:"varint,6,opt,name=chips" json:"chips,omitempty"`
	Bet      int32       `protobuf:"varint,7,opt,name=bet" json:"bet,omitempty"`
	Action   string      `protobuf:"bytes,8,opt,name=action" json:"action,omitempty"`
	Cards    []*CardInfo `protobuf:"bytes,9,rep,name=cards" json:"cards,omitempty"`
}

func (m *PlayerInfo) Reset()                    { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string            { return proto1.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()               {}
func (*PlayerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{25} }

func (m *PlayerInfo) GetPos() int32 {
	if m != nil {
		return m.Pos
	}
	return 0
}

func (m *PlayerInfo) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PlayerInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *PlayerInfo) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *PlayerInfo) GetLevel() string {
	if m != nil {
		return m.Level
	}
	return ""
}

func (m *PlayerInfo) GetChips() int32 {
	if m != nil {
		return m.Chips
	}
	return 0
}

func (m *PlayerInfo) GetBet() int32 {
	if m != nil {
		return m.Bet
	}
	return 0
}

func (m *PlayerInfo) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *PlayerInfo) GetCards() []*CardInfo {
	if m != nil {
		return m.Cards
	}
	return nil
}

func init() {
	proto1.RegisterType((*AutoId)(nil), "proto.AutoId")
	proto1.RegisterType((*SeedInfo)(nil), "proto.SeedInfo")
	proto1.RegisterType((*BaseAck)(nil), "proto.BaseAck")
	proto1.RegisterType((*BaseReq)(nil), "proto.BaseReq")
	proto1.RegisterType((*UserLoginReq)(nil), "proto.UserLoginReq")
	proto1.RegisterType((*UserLoginAck)(nil), "proto.UserLoginAck")
	proto1.RegisterType((*RoomSetTableReq)(nil), "proto.RoomSetTableReq")
	proto1.RegisterType((*RoomSetTableAck)(nil), "proto.RoomSetTableAck")
	proto1.RegisterType((*RoomGetTableReq)(nil), "proto.RoomGetTableReq")
	proto1.RegisterType((*RoomGetTableAck)(nil), "proto.RoomGetTableAck")
	proto1.RegisterType((*RoomGetPlayerReq)(nil), "proto.RoomGetPlayerReq")
	proto1.RegisterType((*RoomGetPlayerAck)(nil), "proto.RoomGetPlayerAck")
	proto1.RegisterType((*RoomPlayerJoinReq)(nil), "proto.RoomPlayerJoinReq")
	proto1.RegisterType((*RoomPlayerJoinAck)(nil), "proto.RoomPlayerJoinAck")
	proto1.RegisterType((*RoomPlayerGoneReq)(nil), "proto.RoomPlayerGoneReq")
	proto1.RegisterType((*RoomPlayerGoneAck)(nil), "proto.RoomPlayerGoneAck")
	proto1.RegisterType((*RoomPlayerBetReq)(nil), "proto.RoomPlayerBetReq")
	proto1.RegisterType((*RoomPlayerBetAck)(nil), "proto.RoomPlayerBetAck")
	proto1.RegisterType((*RoomButtonAck)(nil), "proto.RoomButtonAck")
	proto1.RegisterType((*RoomDealAck)(nil), "proto.RoomDealAck")
	proto1.RegisterType((*RoomPotAck)(nil), "proto.RoomPotAck")
	proto1.RegisterType((*RoomActionAck)(nil), "proto.RoomActionAck")
	proto1.RegisterType((*RoomShowdownAck)(nil), "proto.RoomShowdownAck")
	proto1.RegisterType((*TableInfo)(nil), "proto.TableInfo")
	proto1.RegisterType((*CardInfo)(nil), "proto.CardInfo")
	proto1.RegisterType((*PlayerInfo)(nil), "proto.PlayerInfo")
}

func init() { proto1.RegisterFile("agent.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 916 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xef, 0x6e, 0xe3, 0x44,
	0x10, 0x97, 0x93, 0x26, 0xb6, 0x27, 0xa5, 0x97, 0x1a, 0x44, 0x0d, 0x27, 0x74, 0xa7, 0x95, 0x40,
	0x3d, 0x21, 0x2a, 0x54, 0x3e, 0xf1, 0x31, 0x39, 0x44, 0x15, 0x74, 0x1f, 0x4e, 0x2e, 0x20, 0xc1,
	0x07, 0xac, 0xb5, 0x3d, 0x97, 0x98, 0xd8, 0xbb, 0xae, 0xbd, 0xc9, 0x1d, 0x0f, 0xc0, 0x2b, 0x20,
	0x9e, 0x80, 0x77, 0xe0, 0x69, 0x78, 0x15, 0x34, 0xbb, 0xeb, 0x38, 0x2d, 0x55, 0x25, 0x97, 0xf6,
	0x93, 0x3d, 0x7f, 0x76, 0x66, 0x7e, 0x33, 0xb3, 0xb3, 0x03, 0x13, 0xbe, 0x44, 0xa1, 0xce, 0xaa,
	0x5a, 0x2a, 0x19, 0x8c, 0xf4, 0x87, 0x85, 0x30, 0x9e, 0x6d, 0x94, 0x5c, 0x64, 0xc1, 0x11, 0x0c,
	0xf2, 0x2c, 0x74, 0x9e, 0x3b, 0xa7, 0xa3, 0x68, 0x90, 0x67, 0x2c, 0x03, 0xef, 0x12, 0x31, 0x5b,
	0x88, 0x37, 0x32, 0x38, 0x85, 0x69, 0x5a, 0xe4, 0x28, 0x54, 0xdc, 0xa0, 0xc8, 0xe2, 0x06, 0xb1,
	0xd5, 0x3c, 0x32, 0xfc, 0x4b, 0x14, 0x19, 0x69, 0x07, 0x67, 0xf0, 0xbe, 0xd5, 0xac, 0x31, 0xc5,
	0x7c, 0x8b, 0x46, 0x79, 0xa0, 0x95, 0x8f, 0x8d, 0x28, 0x32, 0x12, 0xd2, 0x67, 0x5f, 0x80, 0x3b,
	0xe7, 0x0d, 0xce, 0xd2, 0x75, 0x30, 0x85, 0x61, 0x8d, 0xca, 0xda, 0xa5, 0x5f, 0xe2, 0x94, 0xcd,
	0x52, 0x1f, 0xf6, 0x23, 0xfa, 0x65, 0xbf, 0x18, 0xf5, 0x08, 0xaf, 0x82, 0x8f, 0xc0, 0xe3, 0x55,
	0x15, 0xbf, 0xa9, 0x65, 0xa9, 0xcf, 0xf8, 0x91, 0xcb, 0xab, 0xea, 0xdb, 0x5a, 0x96, 0xc1, 0x09,
	0xd0, 0x6f, 0xbc, 0xc5, 0xda, 0x3a, 0x1e, 0xf3, 0xaa, 0xfa, 0x11, 0xeb, 0xe0, 0x19, 0x4c, 0x48,
	0x90, 0xae, 0xb8, 0x10, 0x58, 0x84, 0x43, 0x7d, 0x0c, 0x78, 0x55, 0xbd, 0x34, 0x1c, 0xf6, 0xbb,
	0x03, 0x87, 0x3f, 0x34, 0x58, 0xbf, 0x92, 0xcb, 0x5c, 0x90, 0x97, 0x17, 0xe0, 0x25, 0xbc, 0xc1,
	0xb8, 0xc6, 0x2b, 0xed, 0x65, 0x72, 0x7e, 0x64, 0x12, 0x78, 0x66, 0xe3, 0x88, 0xdc, 0xc4, 0x06,
	0x74, 0x02, 0xee, 0xa6, 0xc1, 0x3a, 0xce, 0x5b, 0xb8, 0x63, 0x22, 0x17, 0x59, 0xf0, 0x14, 0xfc,
	0x8d, 0xc8, 0xaf, 0x36, 0x48, 0x22, 0xe3, 0xd3, 0x33, 0x8c, 0x45, 0x16, 0x7c, 0x00, 0x23, 0x25,
	0xd7, 0x28, 0xc2, 0x03, 0x2d, 0x30, 0x04, 0xfb, 0x7a, 0x2f, 0x0c, 0xca, 0x4d, 0x1b, 0x06, 0x4f,
	0xd7, 0xb7, 0x84, 0x31, 0x4b, 0xd7, 0x26, 0x8c, 0x59, 0xba, 0x66, 0x7f, 0x39, 0xf0, 0x24, 0x92,
	0xb2, 0xbc, 0x44, 0xf5, 0x3d, 0x4f, 0x0a, 0xec, 0x89, 0xe2, 0x19, 0x4c, 0x9a, 0x92, 0x17, 0x45,
	0x9c, 0x14, 0xb9, 0x68, 0x91, 0x80, 0x66, 0xcd, 0x89, 0x43, 0x68, 0x92, 0x7c, 0x69, 0xc5, 0x43,
	0x2d, 0xf6, 0x92, 0x7c, 0x69, 0x84, 0x21, 0xb8, 0x2a, 0x2f, 0x51, 0x6e, 0x94, 0xc6, 0x33, 0x8a,
	0x5a, 0x52, 0xd7, 0x92, 0xbf, 0x0b, 0x47, 0xa6, 0xba, 0x25, 0x7f, 0xc7, 0xb2, 0xeb, 0x71, 0xf6,
	0x83, 0x19, 0x7c, 0x06, 0x23, 0x45, 0xc7, 0x74, 0x84, 0x93, 0xf3, 0xa9, 0xd5, 0xd3, 0xa6, 0xa8,
	0x67, 0x23, 0x23, 0x66, 0x95, 0xf1, 0x72, 0x71, 0xbf, 0x6c, 0x9c, 0x80, 0x5b, 0x4b, 0x59, 0xee,
	0xd5, 0x94, 0xc8, 0x45, 0x46, 0xdd, 0xa7, 0xed, 0x77, 0x25, 0x75, 0x35, 0xbd, 0xc8, 0x5a, 0x5c,
	0x17, 0x8f, 0x8b, 0xeb, 0x67, 0x98, 0x5a, 0x2f, 0xaf, 0x0b, 0xfe, 0x1b, 0xd6, 0x3d, 0x81, 0x3d,
	0x05, 0xbf, 0xd2, 0xe7, 0x3a, 0x68, 0x9e, 0x61, 0x2c, 0x32, 0xb6, 0xba, 0x61, 0xbb, 0x27, 0x84,
	0x17, 0x30, 0x36, 0xa6, 0x2c, 0x86, 0x63, 0xab, 0x68, 0x8c, 0x69, 0x10, 0x56, 0x81, 0xd5, 0x70,
	0x4c, 0x9e, 0x8c, 0xe4, 0x3b, 0x79, 0x9f, 0x3b, 0xd7, 0xbb, 0x3e, 0xf9, 0x4d, 0x9f, 0x8f, 0x07,
	0xef, 0xa7, 0x7d, 0x57, 0x17, 0x52, 0xf4, 0x6d, 0xbf, 0x7d, 0x14, 0x83, 0x3b, 0x50, 0x90, 0xe9,
	0xc7, 0x43, 0xf1, 0xab, 0x69, 0x07, 0x23, 0x99, 0xa3, 0x7a, 0x30, 0x10, 0x34, 0x14, 0x12, 0x54,
	0x76, 0x8a, 0xd0, 0x2f, 0xfb, 0xd3, 0xb9, 0xe1, 0xac, 0x27, 0xac, 0x3b, 0x9c, 0x7d, 0x08, 0x63,
	0x9e, 0xaa, 0x5c, 0x0a, 0xdb, 0x10, 0x96, 0x6a, 0x83, 0x38, 0xd8, 0x05, 0x41, 0x33, 0x39, 0x5d,
	0xe5, 0x55, 0x63, 0xa7, 0x95, 0x21, 0x98, 0x82, 0xf7, 0x28, 0xb2, 0xf9, 0x46, 0x29, 0x29, 0x1e,
	0x2e, 0xac, 0x4f, 0x00, 0x12, 0x6d, 0x32, 0xae, 0x64, 0x63, 0x53, 0xe1, 0x1b, 0xce, 0x6b, 0xd9,
	0xb0, 0x3f, 0x1c, 0x98, 0x90, 0xdb, 0x6f, 0x90, 0x17, 0x3d, 0x9d, 0x76, 0x80, 0x07, 0xd7, 0x00,
	0x7f, 0x0a, 0xa3, 0x94, 0xd7, 0x19, 0x39, 0x1b, 0x9e, 0x4e, 0xce, 0x9f, 0xd8, 0xf3, 0x2f, 0x79,
	0x9d, 0x99, 0x09, 0xa3, 0xa5, 0x14, 0xd8, 0x8a, 0x8b, 0x2c, 0x2e, 0x70, 0x8b, 0x85, 0x4d, 0x8f,
	0x4f, 0x9c, 0x57, 0xc4, 0x60, 0x0b, 0x00, 0x5d, 0x28, 0xd9, 0xb7, 0x44, 0x53, 0x18, 0x56, 0x52,
	0x85, 0x83, 0xe7, 0x43, 0xca, 0x77, 0x25, 0x15, 0x5b, 0x9a, 0xcc, 0xce, 0x74, 0x78, 0xf7, 0xb1,
	0xd6, 0xd8, 0xdb, 0x4f, 0xbf, 0x94, 0x6b, 0x7d, 0xb8, 0xeb, 0x2c, 0xad, 0x3c, 0x47, 0xb5, 0x7b,
	0x72, 0x56, 0xf2, 0x6d, 0x26, 0xdf, 0x8a, 0x47, 0x1a, 0xcd, 0x7f, 0x0f, 0xc0, 0xdf, 0x31, 0xf7,
	0xf6, 0x2a, 0x9f, 0xf6, 0xaa, 0xff, 0xf9, 0xc0, 0xfe, 0xb7, 0x59, 0xf7, 0x9e, 0xdc, 0xd1, 0xf5,
	0x27, 0x77, 0x57, 0xe7, 0xf1, 0x9d, 0x75, 0xb6, 0xf5, 0x70, 0x77, 0xf5, 0xe8, 0xfa, 0xdf, 0xd3,
	0x3c, 0x43, 0x50, 0x3b, 0x99, 0xb6, 0x0c, 0x7d, 0x33, 0x6a, 0x0d, 0x15, 0x1c, 0x82, 0x23, 0x42,
	0xd0, 0x2c, 0x47, 0xb4, 0xef, 0xfc, 0x64, 0xf7, 0xce, 0x07, 0x9f, 0x83, 0x6b, 0x06, 0x49, 0x13,
	0x1e, 0xea, 0x40, 0x6e, 0x19, 0x35, 0xad, 0x06, 0xfb, 0x12, 0xbc, 0x36, 0xbe, 0x20, 0x80, 0x83,
	0x66, 0x93, 0xb7, 0x1b, 0xa1, 0xfe, 0x27, 0xf3, 0x5b, 0x5e, 0xb4, 0xe5, 0xde, 0xf2, 0x82, 0xfd,
	0xe3, 0x00, 0x74, 0x96, 0xda, 0x7e, 0x70, 0xba, 0x7e, 0x30, 0x05, 0x18, 0xb4, 0x8b, 0x6d, 0xf0,
	0x31, 0x78, 0x22, 0x4f, 0xd7, 0x82, 0x97, 0xd8, 0x6e, 0x63, 0x2d, 0xad, 0xaf, 0xcc, 0x96, 0x2b,
	0x5e, 0xdb, 0x75, 0xcc, 0x52, 0x94, 0x11, 0x73, 0x0d, 0x46, 0x66, 0x4b, 0xd3, 0x44, 0x97, 0xa7,
	0xf1, 0xde, 0x9c, 0x68, 0x4b, 0xe4, 0x76, 0x25, 0xea, 0x2e, 0xa2, 0x77, 0xfb, 0x45, 0xf4, 0xef,
	0x2a, 0x50, 0x32, 0xd6, 0xec, 0xaf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x7e, 0x0c, 0xe0,
	0xc0, 0x0b, 0x00, 0x00,
}