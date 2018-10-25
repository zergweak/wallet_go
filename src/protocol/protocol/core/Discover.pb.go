// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/Discover.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Endpoint struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	NodeId               []byte   `protobuf:"bytes,3,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{0}
}

func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (m *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(m, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

type PingMessage struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Endpoint `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Version              int32     `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp            int64     `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{1}
}

func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (m *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(m, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PingMessage) GetTo() *Endpoint {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PingMessage) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PingMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PongMessage struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Echo                 int32     `protobuf:"varint,2,opt,name=echo,proto3" json:"echo,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PongMessage) Reset()         { *m = PongMessage{} }
func (m *PongMessage) String() string { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()    {}
func (*PongMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{2}
}

func (m *PongMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongMessage.Unmarshal(m, b)
}
func (m *PongMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongMessage.Marshal(b, m, deterministic)
}
func (m *PongMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongMessage.Merge(m, src)
}
func (m *PongMessage) XXX_Size() int {
	return xxx_messageInfo_PongMessage.Size(m)
}
func (m *PongMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PongMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PongMessage proto.InternalMessageInfo

func (m *PongMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PongMessage) GetEcho() int32 {
	if m != nil {
		return m.Echo
	}
	return 0
}

func (m *PongMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type FindNeighbours struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	TargetId             []byte    `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FindNeighbours) Reset()         { *m = FindNeighbours{} }
func (m *FindNeighbours) String() string { return proto.CompactTextString(m) }
func (*FindNeighbours) ProtoMessage()    {}
func (*FindNeighbours) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{3}
}

func (m *FindNeighbours) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNeighbours.Unmarshal(m, b)
}
func (m *FindNeighbours) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNeighbours.Marshal(b, m, deterministic)
}
func (m *FindNeighbours) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNeighbours.Merge(m, src)
}
func (m *FindNeighbours) XXX_Size() int {
	return xxx_messageInfo_FindNeighbours.Size(m)
}
func (m *FindNeighbours) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNeighbours.DiscardUnknown(m)
}

var xxx_messageInfo_FindNeighbours proto.InternalMessageInfo

func (m *FindNeighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *FindNeighbours) GetTargetId() []byte {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *FindNeighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Neighbours struct {
	From                 *Endpoint   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Neighbours           []*Endpoint `protobuf:"bytes,2,rep,name=neighbours,proto3" json:"neighbours,omitempty"`
	Timestamp            int64       `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Neighbours) Reset()         { *m = Neighbours{} }
func (m *Neighbours) String() string { return proto.CompactTextString(m) }
func (*Neighbours) ProtoMessage()    {}
func (*Neighbours) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{4}
}

func (m *Neighbours) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbours.Unmarshal(m, b)
}
func (m *Neighbours) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbours.Marshal(b, m, deterministic)
}
func (m *Neighbours) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbours.Merge(m, src)
}
func (m *Neighbours) XXX_Size() int {
	return xxx_messageInfo_Neighbours.Size(m)
}
func (m *Neighbours) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbours.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbours proto.InternalMessageInfo

func (m *Neighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Neighbours) GetNeighbours() []*Endpoint {
	if m != nil {
		return m.Neighbours
	}
	return nil
}

func (m *Neighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type BackupMessage struct {
	Flag                 bool     `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Priority             int32    `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupMessage) Reset()         { *m = BackupMessage{} }
func (m *BackupMessage) String() string { return proto.CompactTextString(m) }
func (*BackupMessage) ProtoMessage()    {}
func (*BackupMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0471e8a4adb61f9e, []int{5}
}

func (m *BackupMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupMessage.Unmarshal(m, b)
}
func (m *BackupMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupMessage.Marshal(b, m, deterministic)
}
func (m *BackupMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupMessage.Merge(m, src)
}
func (m *BackupMessage) XXX_Size() int {
	return xxx_messageInfo_BackupMessage.Size(m)
}
func (m *BackupMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BackupMessage proto.InternalMessageInfo

func (m *BackupMessage) GetFlag() bool {
	if m != nil {
		return m.Flag
	}
	return false
}

func (m *BackupMessage) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "protocol.Endpoint")
	proto.RegisterType((*PingMessage)(nil), "protocol.PingMessage")
	proto.RegisterType((*PongMessage)(nil), "protocol.PongMessage")
	proto.RegisterType((*FindNeighbours)(nil), "protocol.FindNeighbours")
	proto.RegisterType((*Neighbours)(nil), "protocol.Neighbours")
	proto.RegisterType((*BackupMessage)(nil), "protocol.BackupMessage")
}

func init() { proto.RegisterFile("core/Discover.proto", fileDescriptor_0471e8a4adb61f9e) }

var fileDescriptor_0471e8a4adb61f9e = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6a, 0xe3, 0x30,
	0x10, 0xc6, 0xf1, 0x9f, 0x64, 0xbd, 0x93, 0x64, 0x17, 0xb4, 0xb0, 0x98, 0xd2, 0x83, 0xf1, 0xa1,
	0x84, 0x1e, 0x1c, 0x48, 0x1f, 0xa0, 0x10, 0xda, 0x42, 0x0e, 0x2d, 0xc1, 0xc7, 0xde, 0x1c, 0x4b,
	0x51, 0x44, 0x13, 0x8d, 0x19, 0x29, 0x81, 0xbe, 0x40, 0xef, 0x7d, 0xe3, 0x62, 0x35, 0x4a, 0xff,
	0xd0, 0x92, 0xb6, 0x27, 0xcf, 0x67, 0x7d, 0xd2, 0xef, 0x9b, 0x61, 0xe0, 0x5f, 0x8d, 0x24, 0x46,
	0x17, 0xca, 0xd4, 0xb8, 0x15, 0x54, 0x34, 0x84, 0x16, 0x59, 0xe2, 0x3e, 0x35, 0xae, 0xf2, 0x19,
	0x24, 0x97, 0x9a, 0x37, 0xa8, 0xb4, 0x65, 0x29, 0xfc, 0xaa, 0x38, 0x27, 0x61, 0x4c, 0x1a, 0x64,
	0xc1, 0xb0, 0x5f, 0x7a, 0xc9, 0x18, 0xc4, 0x0d, 0x92, 0x4d, 0xc3, 0x2c, 0x18, 0x76, 0x4a, 0x57,
	0xb3, 0xff, 0xd0, 0xd5, 0xc8, 0xc5, 0x94, 0xa7, 0x91, 0x33, 0xef, 0x54, 0xfe, 0x18, 0x40, 0x6f,
	0xa6, 0xb4, 0xbc, 0x16, 0xc6, 0x54, 0x52, 0xb0, 0x13, 0x88, 0x17, 0x84, 0x6b, 0xf7, 0x64, 0x6f,
	0xcc, 0x0a, 0x8f, 0x2e, 0x3c, 0xb7, 0x74, 0xe7, 0x2c, 0x87, 0xd0, 0xa2, 0x23, 0x7c, 0xec, 0x0a,
	0x2d, 0xb6, 0x09, 0xb7, 0x82, 0x8c, 0x42, 0xed, 0xa0, 0x9d, 0xd2, 0x4b, 0x76, 0x0c, 0xbf, 0xad,
	0x5a, 0x0b, 0x63, 0xab, 0x75, 0x93, 0xc6, 0x59, 0x30, 0x8c, 0xca, 0x97, 0x1f, 0xb9, 0x84, 0xde,
	0x0c, 0xbf, 0x1f, 0x89, 0x41, 0x2c, 0xea, 0x25, 0xfa, 0xb6, 0xdb, 0xfa, 0x2d, 0x28, 0x7a, 0x0f,
	0x22, 0xf8, 0x73, 0xa5, 0x34, 0xbf, 0x11, 0x4a, 0x2e, 0xe7, 0xb8, 0x21, 0xf3, 0x65, 0xd6, 0x11,
	0x24, 0xb6, 0x22, 0x29, 0xec, 0x94, 0x3b, 0x5e, 0xbf, 0xdc, 0xeb, 0x03, 0xcc, 0x87, 0x00, 0xe0,
	0x07, 0xc0, 0x31, 0x80, 0xde, 0xdf, 0x4a, 0xc3, 0x2c, 0xfa, 0xc4, 0xfd, 0xca, 0x75, 0x20, 0xc8,
	0x39, 0x0c, 0x26, 0x55, 0x7d, 0xb7, 0x69, 0xfc, 0x9c, 0x19, 0xc4, 0x8b, 0x55, 0x25, 0x5d, 0x94,
	0xa4, 0x74, 0x75, 0xdb, 0x67, 0x43, 0x0a, 0x49, 0xd9, 0xfb, 0xdd, 0x5c, 0xf7, 0x7a, 0x72, 0x0a,
	0x7f, 0x91, 0x64, 0x61, 0x09, 0xf5, 0x73, 0x10, 0x33, 0x49, 0xfc, 0xe6, 0xde, 0x0e, 0x7c, 0xb4,
	0x51, 0xbb, 0xd1, 0xf3, 0xae, 0x93, 0x67, 0x4f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x35, 0xf3, 0xb9,
	0xea, 0xe0, 0x02, 0x00, 0x00,
}