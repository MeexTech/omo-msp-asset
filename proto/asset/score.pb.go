// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/asset/score.proto

package asset

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ScoreInfo struct {
	Uid                  string       `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Created              int64        `protobuf:"varint,2,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64        `protobuf:"varint,3,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string       `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string       `protobuf:"bytes,5,opt,name=operator,proto3" json:"operator,omitempty"`
	Name                 string       `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Scene                string       `protobuf:"bytes,7,opt,name=scene,proto3" json:"scene,omitempty"`
	Entity               string       `protobuf:"bytes,8,opt,name=entity,proto3" json:"entity,omitempty"`
	Type                 uint32       `protobuf:"varint,9,opt,name=type,proto3" json:"type,omitempty"`
	Total                uint64       `protobuf:"varint,10,opt,name=total,proto3" json:"total,omitempty"`
	List                 []*ScorePair `protobuf:"bytes,21,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ScoreInfo) Reset()         { *m = ScoreInfo{} }
func (m *ScoreInfo) String() string { return proto.CompactTextString(m) }
func (*ScoreInfo) ProtoMessage()    {}
func (*ScoreInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00efef9d65092e04, []int{0}
}

func (m *ScoreInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScoreInfo.Unmarshal(m, b)
}
func (m *ScoreInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScoreInfo.Marshal(b, m, deterministic)
}
func (m *ScoreInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScoreInfo.Merge(m, src)
}
func (m *ScoreInfo) XXX_Size() int {
	return xxx_messageInfo_ScoreInfo.Size(m)
}
func (m *ScoreInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ScoreInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ScoreInfo proto.InternalMessageInfo

func (m *ScoreInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ScoreInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *ScoreInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *ScoreInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ScoreInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ScoreInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ScoreInfo) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ScoreInfo) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ScoreInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ScoreInfo) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *ScoreInfo) GetList() []*ScorePair {
	if m != nil {
		return m.List
	}
	return nil
}

type ScorePair struct {
	Type                 uint32   `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	Count                uint32   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScorePair) Reset()         { *m = ScorePair{} }
func (m *ScorePair) String() string { return proto.CompactTextString(m) }
func (*ScorePair) ProtoMessage()    {}
func (*ScorePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_00efef9d65092e04, []int{1}
}

func (m *ScorePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScorePair.Unmarshal(m, b)
}
func (m *ScorePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScorePair.Marshal(b, m, deterministic)
}
func (m *ScorePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScorePair.Merge(m, src)
}
func (m *ScorePair) XXX_Size() int {
	return xxx_messageInfo_ScorePair.Size(m)
}
func (m *ScorePair) XXX_DiscardUnknown() {
	xxx_messageInfo_ScorePair.DiscardUnknown(m)
}

var xxx_messageInfo_ScorePair proto.InternalMessageInfo

func (m *ScorePair) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ScorePair) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReqScoreAdd struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Scene                string   `protobuf:"bytes,2,opt,name=scene,proto3" json:"scene,omitempty"`
	Entity               string   `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	Kind                 uint32   `protobuf:"varint,5,opt,name=kind,proto3" json:"kind,omitempty"`
	Count                uint32   `protobuf:"varint,6,opt,name=count,proto3" json:"count,omitempty"`
	Type                 uint32   `protobuf:"varint,7,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqScoreAdd) Reset()         { *m = ReqScoreAdd{} }
func (m *ReqScoreAdd) String() string { return proto.CompactTextString(m) }
func (*ReqScoreAdd) ProtoMessage()    {}
func (*ReqScoreAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_00efef9d65092e04, []int{2}
}

func (m *ReqScoreAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqScoreAdd.Unmarshal(m, b)
}
func (m *ReqScoreAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqScoreAdd.Marshal(b, m, deterministic)
}
func (m *ReqScoreAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqScoreAdd.Merge(m, src)
}
func (m *ReqScoreAdd) XXX_Size() int {
	return xxx_messageInfo_ReqScoreAdd.Size(m)
}
func (m *ReqScoreAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqScoreAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqScoreAdd proto.InternalMessageInfo

func (m *ReqScoreAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqScoreAdd) GetScene() string {
	if m != nil {
		return m.Scene
	}
	return ""
}

func (m *ReqScoreAdd) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *ReqScoreAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqScoreAdd) GetKind() uint32 {
	if m != nil {
		return m.Kind
	}
	return 0
}

func (m *ReqScoreAdd) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *ReqScoreAdd) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

type ReplyScoreInfo struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Info                 *ScoreInfo   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyScoreInfo) Reset()         { *m = ReplyScoreInfo{} }
func (m *ReplyScoreInfo) String() string { return proto.CompactTextString(m) }
func (*ReplyScoreInfo) ProtoMessage()    {}
func (*ReplyScoreInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_00efef9d65092e04, []int{3}
}

func (m *ReplyScoreInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyScoreInfo.Unmarshal(m, b)
}
func (m *ReplyScoreInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyScoreInfo.Marshal(b, m, deterministic)
}
func (m *ReplyScoreInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyScoreInfo.Merge(m, src)
}
func (m *ReplyScoreInfo) XXX_Size() int {
	return xxx_messageInfo_ReplyScoreInfo.Size(m)
}
func (m *ReplyScoreInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyScoreInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyScoreInfo proto.InternalMessageInfo

func (m *ReplyScoreInfo) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyScoreInfo) GetInfo() *ScoreInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

type ReplyScoreList struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	List                 []*ScoreInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyScoreList) Reset()         { *m = ReplyScoreList{} }
func (m *ReplyScoreList) String() string { return proto.CompactTextString(m) }
func (*ReplyScoreList) ProtoMessage()    {}
func (*ReplyScoreList) Descriptor() ([]byte, []int) {
	return fileDescriptor_00efef9d65092e04, []int{4}
}

func (m *ReplyScoreList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyScoreList.Unmarshal(m, b)
}
func (m *ReplyScoreList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyScoreList.Marshal(b, m, deterministic)
}
func (m *ReplyScoreList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyScoreList.Merge(m, src)
}
func (m *ReplyScoreList) XXX_Size() int {
	return xxx_messageInfo_ReplyScoreList.Size(m)
}
func (m *ReplyScoreList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyScoreList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyScoreList proto.InternalMessageInfo

func (m *ReplyScoreList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyScoreList) GetList() []*ScoreInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*ScoreInfo)(nil), "omo.msp.asset.ScoreInfo")
	proto.RegisterType((*ScorePair)(nil), "omo.msp.asset.ScorePair")
	proto.RegisterType((*ReqScoreAdd)(nil), "omo.msp.asset.ReqScoreAdd")
	proto.RegisterType((*ReplyScoreInfo)(nil), "omo.msp.asset.ReplyScoreInfo")
	proto.RegisterType((*ReplyScoreList)(nil), "omo.msp.asset.ReplyScoreList")
}

func init() {
	proto.RegisterFile("proto/asset/score.proto", fileDescriptor_00efef9d65092e04)
}

var fileDescriptor_00efef9d65092e04 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x25, 0x90, 0x86, 0x72, 0x29, 0xd5, 0x64, 0xed, 0xc3, 0x42, 0x9b, 0x84, 0xf2, 0xc4, 0xc3,
	0x44, 0x25, 0xa6, 0xfd, 0x80, 0x32, 0x6d, 0xa8, 0x53, 0xa7, 0x4e, 0xae, 0xf6, 0xb2, 0xb7, 0x2c,
	0xb9, 0x95, 0xac, 0x11, 0x3b, 0xb3, 0x4d, 0x27, 0x7e, 0xc8, 0x7e, 0xc0, 0x1e, 0xf7, 0x2f, 0x27,
	0xdf, 0x84, 0x00, 0x6d, 0x28, 0x12, 0x6f, 0x3e, 0xf7, 0x1e, 0x9f, 0x7b, 0xee, 0x71, 0x00, 0x5e,
	0x15, 0x46, 0x3b, 0x7d, 0x91, 0x58, 0x8b, 0xee, 0xc2, 0xa6, 0xda, 0xe0, 0x84, 0x2a, 0x6c, 0xa0,
	0x73, 0x3d, 0xc9, 0x6d, 0x31, 0xa1, 0xd6, 0x90, 0x6f, 0xf3, 0x52, 0x9d, 0xe7, 0x5a, 0x95, 0xc4,
	0xf8, 0x6f, 0x1b, 0x7a, 0xb7, 0xfe, 0xe2, 0x95, 0xba, 0xd3, 0xec, 0x19, 0x74, 0x96, 0x32, 0xe3,
	0xc1, 0x28, 0x18, 0xf7, 0x84, 0x3f, 0x32, 0x0e, 0xdd, 0xd4, 0x60, 0xe2, 0x30, 0xe3, 0xed, 0x51,
	0x30, 0xee, 0x88, 0x35, 0xf4, 0x9d, 0x65, 0x91, 0x51, 0xa7, 0x53, 0x76, 0x2a, 0x58, 0xdf, 0xd1,
	0x86, 0x87, 0xa4, 0xb4, 0x86, 0x6c, 0x08, 0xa7, 0xba, 0x40, 0x43, 0xad, 0x13, 0x6a, 0xd5, 0x98,
	0x31, 0x08, 0x55, 0x92, 0x23, 0x8f, 0xa8, 0x4e, 0x67, 0xf6, 0x1c, 0x4e, 0x6c, 0x8a, 0x0a, 0x79,
	0x97, 0x8a, 0x25, 0x60, 0x2f, 0x21, 0x42, 0xe5, 0xa4, 0x5b, 0xf1, 0x53, 0x2a, 0x57, 0xc8, 0x2b,
	0xb8, 0x55, 0x81, 0xbc, 0x37, 0x0a, 0xc6, 0x03, 0x41, 0x67, 0xaf, 0xe0, 0xb4, 0x4b, 0x16, 0x1c,
	0x46, 0xc1, 0x38, 0x14, 0x25, 0x60, 0x6f, 0x21, 0x5c, 0x48, 0xeb, 0xf8, 0x8b, 0x51, 0x67, 0xdc,
	0x9f, 0xf2, 0xc9, 0x4e, 0x5a, 0x13, 0xca, 0xe3, 0x6b, 0x22, 0x8d, 0x20, 0x56, 0xfc, 0xbe, 0x8a,
	0xc8, 0x97, 0xea, 0x21, 0xc1, 0xee, 0x90, 0x54, 0x2f, 0x95, 0xa3, 0x88, 0x06, 0xa2, 0x04, 0xf1,
	0xbf, 0x00, 0xfa, 0x02, 0x7f, 0xd1, 0xd5, 0xcb, 0x2c, 0xab, 0x17, 0x0c, 0x9a, 0x16, 0x6c, 0x37,
	0x2f, 0xd8, 0xd9, 0x59, 0x70, 0x3b, 0xbe, 0xf0, 0x71, 0x7c, 0x3f, 0xa5, 0xca, 0x28, 0xd6, 0x81,
	0xa0, 0xf3, 0xc6, 0x57, 0xb4, 0xe5, 0xab, 0xde, 0xa0, 0xbb, 0xd9, 0x20, 0x36, 0x70, 0x2e, 0xb0,
	0x58, 0xac, 0x36, 0x9f, 0xc2, 0x14, 0x22, 0xeb, 0x12, 0xb7, 0xb4, 0xe4, 0xb7, 0x3f, 0x1d, 0x3e,
	0x08, 0xa9, 0xa4, 0x13, 0x43, 0x54, 0x4c, 0x1f, 0xab, 0x54, 0x77, 0x9a, 0x96, 0xd9, 0x13, 0xab,
	0xd7, 0x16, 0xc4, 0xda, 0x9d, 0x79, 0x2d, 0xad, 0x3b, 0x76, 0x26, 0x3d, 0x65, 0x7b, 0xff, 0x53,
	0x96, 0x33, 0x3d, 0x6b, 0xfa, 0x27, 0x84, 0x33, 0xaa, 0xdd, 0xa2, 0xb9, 0x97, 0x29, 0xb2, 0x8f,
	0x10, 0x5d, 0x66, 0xd9, 0x8d, 0x42, 0xf6, 0x78, 0x58, 0xfd, 0x74, 0xc3, 0x37, 0x8d, 0x46, 0xd6,
	0xda, 0x71, 0xcb, 0xcb, 0xcc, 0xd1, 0xed, 0x91, 0x59, 0xa2, 0x75, 0x9e, 0x77, 0x58, 0xe6, 0x03,
	0xf4, 0x04, 0xe6, 0xfa, 0x1e, 0x0f, 0x29, 0xf1, 0x26, 0xa5, 0x4a, 0xe4, 0x0a, 0x60, 0x8e, 0x6e,
	0xb6, 0xba, 0xf9, 0xad, 0xd0, 0x1c, 0xe9, 0xc7, 0x3f, 0x47, 0xdc, 0x62, 0x9f, 0xe1, 0xfc, 0x1b,
	0xfd, 0xa8, 0x67, 0xab, 0x4f, 0x72, 0xe1, 0xd0, 0xb0, 0xd7, 0xcd, 0x72, 0x25, 0xeb, 0x49, 0x5b,
	0xd7, 0xd0, 0x27, 0x5b, 0x4f, 0x0b, 0x95, 0xdd, 0xc3, 0xce, 0xbe, 0xc0, 0xd9, 0x1c, 0x9d, 0xff,
	0x16, 0xa4, 0x75, 0x32, 0x3d, 0x4a, 0x6e, 0x7d, 0x39, 0x6e, 0xcd, 0x06, 0xdf, 0xfb, 0x5b, 0x7f,
	0x91, 0x3f, 0x22, 0x02, 0xef, 0xfe, 0x07, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x4e, 0x05, 0xc1, 0x60,
	0x05, 0x00, 0x00,
}
