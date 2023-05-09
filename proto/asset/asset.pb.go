// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/asset/asset.proto

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

type OwnerType int32

const (
	OwnerType_Single OwnerType = 0
	OwnerType_Group  OwnerType = 1
)

var OwnerType_name = map[int32]string{
	0: "Single",
	1: "Group",
}

var OwnerType_value = map[string]int32{
	"Single": 0,
	"Group":  1,
}

func (x OwnerType) String() string {
	return proto.EnumName(OwnerType_name, int32(x))
}

func (OwnerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{0}
}

type AssetInfo struct {
	Uid                  string        `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Md5                  string        `protobuf:"bytes,3,opt,name=md5,proto3" json:"md5,omitempty"`
	Format               string        `protobuf:"bytes,4,opt,name=format,proto3" json:"format,omitempty"`
	Version              string        `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Size                 uint64        `protobuf:"varint,6,opt,name=size,proto3" json:"size,omitempty"`
	Language             string        `protobuf:"bytes,7,opt,name=language,proto3" json:"language,omitempty"`
	Type                 OwnerType     `protobuf:"varint,8,opt,name=type,proto3,enum=omo.msp.asset.OwnerType" json:"type,omitempty"`
	Owner                string        `protobuf:"bytes,9,opt,name=owner,proto3" json:"owner,omitempty"`
	Url                  string        `protobuf:"bytes,10,opt,name=url,proto3" json:"url,omitempty"`
	Created              int64         `protobuf:"varint,11,opt,name=created,proto3" json:"created,omitempty"`
	Updated              int64         `protobuf:"varint,12,opt,name=updated,proto3" json:"updated,omitempty"`
	Creator              string        `protobuf:"bytes,13,opt,name=creator,proto3" json:"creator,omitempty"`
	Operator             string        `protobuf:"bytes,14,opt,name=operator,proto3" json:"operator,omitempty"`
	Width                uint32        `protobuf:"varint,15,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32        `protobuf:"varint,16,opt,name=height,proto3" json:"height,omitempty"`
	Snapshot             string        `protobuf:"bytes,17,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	Small                string        `protobuf:"bytes,18,opt,name=small,proto3" json:"small,omitempty"`
	Remark               string        `protobuf:"bytes,20,opt,name=remark,proto3" json:"remark,omitempty"`
	Meta                 string        `protobuf:"bytes,21,opt,name=meta,proto3" json:"meta,omitempty"`
	Weight               uint32        `protobuf:"varint,22,opt,name=weight,proto3" json:"weight,omitempty"`
	Status               uint32        `protobuf:"varint,23,opt,name=status,proto3" json:"status,omitempty"`
	Thumbs               []*ThumbBrief `protobuf:"bytes,19,rep,name=thumbs,proto3" json:"thumbs,omitempty"`
	Links                []string      `protobuf:"bytes,31,rep,name=links,proto3" json:"links,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AssetInfo) Reset()         { *m = AssetInfo{} }
func (m *AssetInfo) String() string { return proto.CompactTextString(m) }
func (*AssetInfo) ProtoMessage()    {}
func (*AssetInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{0}
}

func (m *AssetInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetInfo.Unmarshal(m, b)
}
func (m *AssetInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetInfo.Marshal(b, m, deterministic)
}
func (m *AssetInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetInfo.Merge(m, src)
}
func (m *AssetInfo) XXX_Size() int {
	return xxx_messageInfo_AssetInfo.Size(m)
}
func (m *AssetInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AssetInfo proto.InternalMessageInfo

func (m *AssetInfo) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *AssetInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AssetInfo) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *AssetInfo) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *AssetInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AssetInfo) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *AssetInfo) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *AssetInfo) GetType() OwnerType {
	if m != nil {
		return m.Type
	}
	return OwnerType_Single
}

func (m *AssetInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *AssetInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AssetInfo) GetCreated() int64 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *AssetInfo) GetUpdated() int64 {
	if m != nil {
		return m.Updated
	}
	return 0
}

func (m *AssetInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *AssetInfo) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *AssetInfo) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *AssetInfo) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *AssetInfo) GetSnapshot() string {
	if m != nil {
		return m.Snapshot
	}
	return ""
}

func (m *AssetInfo) GetSmall() string {
	if m != nil {
		return m.Small
	}
	return ""
}

func (m *AssetInfo) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *AssetInfo) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *AssetInfo) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *AssetInfo) GetStatus() uint32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *AssetInfo) GetThumbs() []*ThumbBrief {
	if m != nil {
		return m.Thumbs
	}
	return nil
}

func (m *AssetInfo) GetLinks() []string {
	if m != nil {
		return m.Links
	}
	return nil
}

type ThumbBrief struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Owner                string   `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Face                 string   `protobuf:"bytes,3,opt,name=face,proto3" json:"face,omitempty"`
	Probably             float32  `protobuf:"fixed32,4,opt,name=probably,proto3" json:"probably,omitempty"`
	Similar              float32  `protobuf:"fixed32,5,opt,name=similar,proto3" json:"similar,omitempty"`
	Blur                 float32  `protobuf:"fixed32,6,opt,name=blur,proto3" json:"blur,omitempty"`
	Url                  string   `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ThumbBrief) Reset()         { *m = ThumbBrief{} }
func (m *ThumbBrief) String() string { return proto.CompactTextString(m) }
func (*ThumbBrief) ProtoMessage()    {}
func (*ThumbBrief) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{1}
}

func (m *ThumbBrief) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ThumbBrief.Unmarshal(m, b)
}
func (m *ThumbBrief) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ThumbBrief.Marshal(b, m, deterministic)
}
func (m *ThumbBrief) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ThumbBrief.Merge(m, src)
}
func (m *ThumbBrief) XXX_Size() int {
	return xxx_messageInfo_ThumbBrief.Size(m)
}
func (m *ThumbBrief) XXX_DiscardUnknown() {
	xxx_messageInfo_ThumbBrief.DiscardUnknown(m)
}

var xxx_messageInfo_ThumbBrief proto.InternalMessageInfo

func (m *ThumbBrief) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ThumbBrief) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ThumbBrief) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

func (m *ThumbBrief) GetProbably() float32 {
	if m != nil {
		return m.Probably
	}
	return 0
}

func (m *ThumbBrief) GetSimilar() float32 {
	if m != nil {
		return m.Similar
	}
	return 0
}

func (m *ThumbBrief) GetBlur() float32 {
	if m != nil {
		return m.Blur
	}
	return 0
}

func (m *ThumbBrief) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type ReqAssetAdd struct {
	Name                 string    `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Md5                  string    `protobuf:"bytes,2,opt,name=md5,proto3" json:"md5,omitempty"`
	Format               string    `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	Version              string    `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Size                 uint64    `protobuf:"varint,5,opt,name=size,proto3" json:"size,omitempty"`
	Language             string    `protobuf:"bytes,6,opt,name=language,proto3" json:"language,omitempty"`
	Uuid                 string    `protobuf:"bytes,7,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Owner                string    `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Type                 OwnerType `protobuf:"varint,9,opt,name=type,proto3,enum=omo.msp.asset.OwnerType" json:"type,omitempty"`
	Operator             string    `protobuf:"bytes,10,opt,name=operator,proto3" json:"operator,omitempty"`
	Width                uint32    `protobuf:"varint,11,opt,name=width,proto3" json:"width,omitempty"`
	Height               uint32    `protobuf:"varint,12,opt,name=height,proto3" json:"height,omitempty"`
	Snapshot             string    `protobuf:"bytes,13,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	Small                string    `protobuf:"bytes,14,opt,name=small,proto3" json:"small,omitempty"`
	Remark               string    `protobuf:"bytes,16,opt,name=remark,proto3" json:"remark,omitempty"`
	Meta                 string    `protobuf:"bytes,17,opt,name=meta,proto3" json:"meta,omitempty"`
	Thumbs               []string  `protobuf:"bytes,15,rep,name=thumbs,proto3" json:"thumbs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReqAssetAdd) Reset()         { *m = ReqAssetAdd{} }
func (m *ReqAssetAdd) String() string { return proto.CompactTextString(m) }
func (*ReqAssetAdd) ProtoMessage()    {}
func (*ReqAssetAdd) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{2}
}

func (m *ReqAssetAdd) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAssetAdd.Unmarshal(m, b)
}
func (m *ReqAssetAdd) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAssetAdd.Marshal(b, m, deterministic)
}
func (m *ReqAssetAdd) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAssetAdd.Merge(m, src)
}
func (m *ReqAssetAdd) XXX_Size() int {
	return xxx_messageInfo_ReqAssetAdd.Size(m)
}
func (m *ReqAssetAdd) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAssetAdd.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAssetAdd proto.InternalMessageInfo

func (m *ReqAssetAdd) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqAssetAdd) GetMd5() string {
	if m != nil {
		return m.Md5
	}
	return ""
}

func (m *ReqAssetAdd) GetFormat() string {
	if m != nil {
		return m.Format
	}
	return ""
}

func (m *ReqAssetAdd) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ReqAssetAdd) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *ReqAssetAdd) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *ReqAssetAdd) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *ReqAssetAdd) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReqAssetAdd) GetType() OwnerType {
	if m != nil {
		return m.Type
	}
	return OwnerType_Single
}

func (m *ReqAssetAdd) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *ReqAssetAdd) GetWidth() uint32 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *ReqAssetAdd) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ReqAssetAdd) GetSnapshot() string {
	if m != nil {
		return m.Snapshot
	}
	return ""
}

func (m *ReqAssetAdd) GetSmall() string {
	if m != nil {
		return m.Small
	}
	return ""
}

func (m *ReqAssetAdd) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqAssetAdd) GetMeta() string {
	if m != nil {
		return m.Meta
	}
	return ""
}

func (m *ReqAssetAdd) GetThumbs() []string {
	if m != nil {
		return m.Thumbs
	}
	return nil
}

type ReqAssetFlag struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Flag                 string   `protobuf:"bytes,2,opt,name=flag,proto3" json:"flag,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqAssetFlag) Reset()         { *m = ReqAssetFlag{} }
func (m *ReqAssetFlag) String() string { return proto.CompactTextString(m) }
func (*ReqAssetFlag) ProtoMessage()    {}
func (*ReqAssetFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{3}
}

func (m *ReqAssetFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAssetFlag.Unmarshal(m, b)
}
func (m *ReqAssetFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAssetFlag.Marshal(b, m, deterministic)
}
func (m *ReqAssetFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAssetFlag.Merge(m, src)
}
func (m *ReqAssetFlag) XXX_Size() int {
	return xxx_messageInfo_ReqAssetFlag.Size(m)
}
func (m *ReqAssetFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAssetFlag.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAssetFlag proto.InternalMessageInfo

func (m *ReqAssetFlag) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqAssetFlag) GetFlag() string {
	if m != nil {
		return m.Flag
	}
	return ""
}

func (m *ReqAssetFlag) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReqAssetWeight struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=weight,proto3" json:"weight,omitempty"`
	Operator             string   `protobuf:"bytes,3,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqAssetWeight) Reset()         { *m = ReqAssetWeight{} }
func (m *ReqAssetWeight) String() string { return proto.CompactTextString(m) }
func (*ReqAssetWeight) ProtoMessage()    {}
func (*ReqAssetWeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{4}
}

func (m *ReqAssetWeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAssetWeight.Unmarshal(m, b)
}
func (m *ReqAssetWeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAssetWeight.Marshal(b, m, deterministic)
}
func (m *ReqAssetWeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAssetWeight.Merge(m, src)
}
func (m *ReqAssetWeight) XXX_Size() int {
	return xxx_messageInfo_ReqAssetWeight.Size(m)
}
func (m *ReqAssetWeight) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAssetWeight.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAssetWeight proto.InternalMessageInfo

func (m *ReqAssetWeight) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqAssetWeight) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *ReqAssetWeight) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyAssetOne struct {
	Info                 *AssetInfo   `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyAssetOne) Reset()         { *m = ReplyAssetOne{} }
func (m *ReplyAssetOne) String() string { return proto.CompactTextString(m) }
func (*ReplyAssetOne) ProtoMessage()    {}
func (*ReplyAssetOne) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{5}
}

func (m *ReplyAssetOne) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyAssetOne.Unmarshal(m, b)
}
func (m *ReplyAssetOne) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyAssetOne.Marshal(b, m, deterministic)
}
func (m *ReplyAssetOne) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyAssetOne.Merge(m, src)
}
func (m *ReplyAssetOne) XXX_Size() int {
	return xxx_messageInfo_ReplyAssetOne.Size(m)
}
func (m *ReplyAssetOne) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyAssetOne.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyAssetOne proto.InternalMessageInfo

func (m *ReplyAssetOne) GetInfo() *AssetInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func (m *ReplyAssetOne) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReqAssetList struct {
	List                 []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqAssetList) Reset()         { *m = ReqAssetList{} }
func (m *ReqAssetList) String() string { return proto.CompactTextString(m) }
func (*ReqAssetList) ProtoMessage()    {}
func (*ReqAssetList) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{6}
}

func (m *ReqAssetList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAssetList.Unmarshal(m, b)
}
func (m *ReqAssetList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAssetList.Marshal(b, m, deterministic)
}
func (m *ReqAssetList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAssetList.Merge(m, src)
}
func (m *ReqAssetList) XXX_Size() int {
	return xxx_messageInfo_ReqAssetList.Size(m)
}
func (m *ReqAssetList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAssetList.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAssetList proto.InternalMessageInfo

func (m *ReqAssetList) GetList() []string {
	if m != nil {
		return m.List
	}
	return nil
}

type ReqAssetUpdate struct {
	Uid                  string   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Remark               string   `protobuf:"bytes,3,opt,name=remark,proto3" json:"remark,omitempty"`
	Operator             string   `protobuf:"bytes,4,opt,name=operator,proto3" json:"operator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqAssetUpdate) Reset()         { *m = ReqAssetUpdate{} }
func (m *ReqAssetUpdate) String() string { return proto.CompactTextString(m) }
func (*ReqAssetUpdate) ProtoMessage()    {}
func (*ReqAssetUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{7}
}

func (m *ReqAssetUpdate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqAssetUpdate.Unmarshal(m, b)
}
func (m *ReqAssetUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqAssetUpdate.Marshal(b, m, deterministic)
}
func (m *ReqAssetUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqAssetUpdate.Merge(m, src)
}
func (m *ReqAssetUpdate) XXX_Size() int {
	return xxx_messageInfo_ReqAssetUpdate.Size(m)
}
func (m *ReqAssetUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqAssetUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReqAssetUpdate proto.InternalMessageInfo

func (m *ReqAssetUpdate) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *ReqAssetUpdate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqAssetUpdate) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *ReqAssetUpdate) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

type ReplyAssetList struct {
	Owner                string       `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	List                 []*AssetInfo `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
	Status               *ReplyStatus `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyAssetList) Reset()         { *m = ReplyAssetList{} }
func (m *ReplyAssetList) String() string { return proto.CompactTextString(m) }
func (*ReplyAssetList) ProtoMessage()    {}
func (*ReplyAssetList) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{8}
}

func (m *ReplyAssetList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyAssetList.Unmarshal(m, b)
}
func (m *ReplyAssetList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyAssetList.Marshal(b, m, deterministic)
}
func (m *ReplyAssetList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyAssetList.Merge(m, src)
}
func (m *ReplyAssetList) XXX_Size() int {
	return xxx_messageInfo_ReplyAssetList.Size(m)
}
func (m *ReplyAssetList) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyAssetList.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyAssetList proto.InternalMessageInfo

func (m *ReplyAssetList) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *ReplyAssetList) GetList() []*AssetInfo {
	if m != nil {
		return m.List
	}
	return nil
}

func (m *ReplyAssetList) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type ReplyAssetToken struct {
	Status               *ReplyStatus `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Domain               string       `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	Limit                uint32       `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Expire               uint32       `protobuf:"varint,4,opt,name=expire,proto3" json:"expire,omitempty"`
	Bucket               string       `protobuf:"bytes,5,opt,name=bucket,proto3" json:"bucket,omitempty"`
	Token                string       `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ReplyAssetToken) Reset()         { *m = ReplyAssetToken{} }
func (m *ReplyAssetToken) String() string { return proto.CompactTextString(m) }
func (*ReplyAssetToken) ProtoMessage()    {}
func (*ReplyAssetToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_78e4568551c90fa4, []int{9}
}

func (m *ReplyAssetToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReplyAssetToken.Unmarshal(m, b)
}
func (m *ReplyAssetToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReplyAssetToken.Marshal(b, m, deterministic)
}
func (m *ReplyAssetToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplyAssetToken.Merge(m, src)
}
func (m *ReplyAssetToken) XXX_Size() int {
	return xxx_messageInfo_ReplyAssetToken.Size(m)
}
func (m *ReplyAssetToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplyAssetToken.DiscardUnknown(m)
}

var xxx_messageInfo_ReplyAssetToken proto.InternalMessageInfo

func (m *ReplyAssetToken) GetStatus() *ReplyStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ReplyAssetToken) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ReplyAssetToken) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ReplyAssetToken) GetExpire() uint32 {
	if m != nil {
		return m.Expire
	}
	return 0
}

func (m *ReplyAssetToken) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ReplyAssetToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterEnum("omo.msp.asset.OwnerType", OwnerType_name, OwnerType_value)
	proto.RegisterType((*AssetInfo)(nil), "omo.msp.asset.AssetInfo")
	proto.RegisterType((*ThumbBrief)(nil), "omo.msp.asset.ThumbBrief")
	proto.RegisterType((*ReqAssetAdd)(nil), "omo.msp.asset.ReqAssetAdd")
	proto.RegisterType((*ReqAssetFlag)(nil), "omo.msp.asset.ReqAssetFlag")
	proto.RegisterType((*ReqAssetWeight)(nil), "omo.msp.asset.ReqAssetWeight")
	proto.RegisterType((*ReplyAssetOne)(nil), "omo.msp.asset.ReplyAssetOne")
	proto.RegisterType((*ReqAssetList)(nil), "omo.msp.asset.ReqAssetList")
	proto.RegisterType((*ReqAssetUpdate)(nil), "omo.msp.asset.ReqAssetUpdate")
	proto.RegisterType((*ReplyAssetList)(nil), "omo.msp.asset.ReplyAssetList")
	proto.RegisterType((*ReplyAssetToken)(nil), "omo.msp.asset.ReplyAssetToken")
}

func init() {
	proto.RegisterFile("proto/asset/asset.proto", fileDescriptor_78e4568551c90fa4)
}

var fileDescriptor_78e4568551c90fa4 = []byte{
	// 1017 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x57, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0xe3, 0x34, 0x6d, 0x4e, 0x7e, 0x9a, 0x1d, 0x96, 0xdd, 0x21, 0xec, 0x42, 0xe4, 0xab,
	0x08, 0xa1, 0x20, 0x8a, 0xf6, 0x01, 0x5a, 0xa0, 0xa1, 0x68, 0xd1, 0x22, 0xb7, 0x80, 0xc4, 0x9d,
	0x13, 0x4f, 0x12, 0x53, 0xdb, 0xe3, 0x1d, 0xdb, 0x2d, 0xe1, 0x09, 0xb8, 0xe4, 0x21, 0xb8, 0xe5,
	0x01, 0x78, 0x25, 0x9e, 0x02, 0x9d, 0x33, 0xb6, 0xf3, 0x53, 0x3b, 0x89, 0x76, 0x6f, 0xaa, 0xf9,
	0x66, 0x4e, 0xbe, 0x9c, 0xf3, 0x9d, 0xef, 0x9c, 0xa8, 0xf0, 0x3c, 0x52, 0x32, 0x91, 0x5f, 0x38,
	0x71, 0x2c, 0x12, 0xfd, 0x77, 0x44, 0x37, 0xac, 0x23, 0x03, 0x39, 0x0a, 0xe2, 0x68, 0x44, 0x97,
	0x7d, 0xbe, 0x1e, 0x37, 0x95, 0x41, 0x20, 0x43, 0x1d, 0x68, 0xfd, 0x57, 0x87, 0xe6, 0x05, 0x5e,
	0x5f, 0x87, 0x33, 0xc9, 0x7a, 0x60, 0xa6, 0x9e, 0xcb, 0x8d, 0x81, 0x31, 0x6c, 0xda, 0x78, 0x64,
	0x0c, 0xea, 0xa1, 0x13, 0x08, 0x5e, 0xa3, 0x2b, 0x3a, 0x63, 0x54, 0xe0, 0xbe, 0xe2, 0xa6, 0x8e,
	0x0a, 0xdc, 0x57, 0xec, 0x19, 0x34, 0x66, 0x52, 0x05, 0x4e, 0xc2, 0xeb, 0x74, 0x99, 0x21, 0xc6,
	0xe1, 0xe4, 0x5e, 0xa8, 0xd8, 0x93, 0x21, 0x3f, 0xa6, 0x87, 0x1c, 0x22, 0x6f, 0xec, 0xfd, 0x21,
	0x78, 0x63, 0x60, 0x0c, 0xeb, 0x36, 0x9d, 0x59, 0x1f, 0x4e, 0x7d, 0x27, 0x9c, 0xa7, 0xce, 0x5c,
	0xf0, 0x13, 0x0a, 0x2f, 0x30, 0xfb, 0x1c, 0xea, 0xc9, 0x32, 0x12, 0xfc, 0x74, 0x60, 0x0c, 0xbb,
	0xe7, 0x7c, 0xb4, 0x51, 0xdf, 0xe8, 0xcd, 0x43, 0x28, 0xd4, 0xed, 0x32, 0x12, 0x36, 0x45, 0xb1,
	0xa7, 0x70, 0x2c, 0xf1, 0x8a, 0x37, 0x89, 0x46, 0x03, 0xaa, 0x4e, 0xf9, 0x1c, 0xb2, 0xea, 0x94,
	0x8f, 0xf9, 0x4d, 0x95, 0x70, 0x12, 0xe1, 0xf2, 0xd6, 0xc0, 0x18, 0x9a, 0x76, 0x0e, 0xf1, 0x25,
	0x8d, 0x5c, 0x7a, 0x69, 0xeb, 0x97, 0x0c, 0x16, 0x9f, 0x91, 0x8a, 0x77, 0x74, 0x4d, 0x19, 0xc4,
	0xfc, 0x65, 0x24, 0x14, 0x3d, 0x75, 0x75, 0xfe, 0x39, 0xc6, 0x8c, 0x1e, 0x3c, 0x37, 0x59, 0xf0,
	0xb3, 0x81, 0x31, 0xec, 0xd8, 0x1a, 0xa0, 0x6e, 0x0b, 0xe1, 0xcd, 0x17, 0x09, 0xef, 0xd1, 0x75,
	0x86, 0x90, 0x29, 0x0e, 0x9d, 0x28, 0x5e, 0xc8, 0x84, 0x3f, 0xd1, 0x4c, 0x39, 0x46, 0xa6, 0x38,
	0x70, 0x7c, 0x9f, 0x33, 0x5d, 0x1b, 0x01, 0x64, 0x52, 0x22, 0x70, 0xd4, 0x1d, 0x7f, 0xaa, 0x3b,
	0xa0, 0x11, 0xea, 0x1c, 0x88, 0xc4, 0xe1, 0x1f, 0xea, 0xfe, 0xe1, 0x19, 0x63, 0x1f, 0xf4, 0xb7,
	0x3e, 0xd3, 0xdf, 0xaa, 0x11, 0xde, 0xc7, 0x89, 0x93, 0xa4, 0x31, 0x7f, 0xae, 0xef, 0x35, 0x62,
	0x5f, 0x42, 0x23, 0x59, 0xa4, 0xc1, 0x24, 0xe6, 0x1f, 0x0c, 0xcc, 0x61, 0xeb, 0xfc, 0xa3, 0x2d,
	0xf5, 0x6f, 0xf1, 0xf1, 0x52, 0x79, 0x62, 0x66, 0x67, 0x81, 0x98, 0xa4, 0xef, 0x85, 0x77, 0x31,
	0xff, 0x74, 0x60, 0x62, 0x92, 0x04, 0xac, 0xbf, 0x0d, 0x80, 0x55, 0x70, 0x89, 0xdb, 0x8a, 0xbe,
	0xd5, 0xd6, 0xfb, 0xc6, 0xa0, 0x3e, 0x73, 0xa6, 0x22, 0x33, 0x1c, 0x9d, 0x51, 0xa1, 0x48, 0xc9,
	0x89, 0x33, 0xf1, 0x97, 0xe4, 0xb9, 0x9a, 0x5d, 0x60, 0xec, 0x50, 0xec, 0x05, 0x9e, 0xef, 0x28,
	0x72, 0x5d, 0xcd, 0xce, 0x21, 0x32, 0x4d, 0xfc, 0x54, 0x91, 0xeb, 0x6a, 0x36, 0x9d, 0x73, 0x57,
	0x9c, 0x14, 0xae, 0xb0, 0xfe, 0x31, 0xa1, 0x65, 0x8b, 0xb7, 0x34, 0x16, 0x17, 0xee, 0x6a, 0x06,
	0x8c, 0xc7, 0x33, 0x50, 0x2b, 0x9b, 0x01, 0xb3, 0x6a, 0x06, 0xea, 0xe5, 0x33, 0x70, 0x5c, 0x31,
	0x03, 0x8d, 0xad, 0x19, 0x60, 0x50, 0x4f, 0x51, 0x30, 0x9d, 0x2a, 0x9d, 0x57, 0x8a, 0x9d, 0xae,
	0x2b, 0x96, 0x4f, 0x4b, 0xf3, 0xa0, 0x69, 0x59, 0xf7, 0x2d, 0x54, 0xf9, 0xb6, 0x55, 0xee, 0xdb,
	0x76, 0xa5, 0x6f, 0x3b, 0x55, 0xbe, 0xed, 0x96, 0xfb, 0xb6, 0x57, 0xea, 0xdb, 0x27, 0x9b, 0xbe,
	0xcd, 0x7c, 0x78, 0x46, 0xae, 0xca, 0x90, 0xf5, 0x23, 0xb4, 0xf3, 0x76, 0x5d, 0xf9, 0xce, 0xbc,
	0x7c, 0x8b, 0xcd, 0x7c, 0x67, 0x9e, 0x6f, 0x31, 0x3c, 0x6f, 0x54, 0x6d, 0x6e, 0x56, 0x6d, 0xfd,
	0x0c, 0xdd, 0x9c, 0xf1, 0x17, 0x5d, 0xd9, 0x63, 0xce, 0xd5, 0x14, 0xd5, 0x36, 0xa6, 0x68, 0x17,
	0xef, 0x5b, 0xe8, 0xd8, 0x22, 0xf2, 0x97, 0xc4, 0xfc, 0x26, 0xa4, 0xb5, 0xe6, 0x85, 0x33, 0x49,
	0xbc, 0xad, 0x47, 0x8d, 0x2a, 0x16, 0xb3, 0x4d, 0x51, 0xec, 0xbc, 0x18, 0xd0, 0x1a, 0xc5, 0xf7,
	0xb7, 0xe2, 0x89, 0xfb, 0x86, 0x22, 0xf2, 0xe1, 0xb5, 0xac, 0x95, 0x38, 0xaf, 0xbd, 0x38, 0x41,
	0x29, 0x7c, 0x2f, 0x4e, 0xb8, 0x41, 0x12, 0xd2, 0xd9, 0xfa, 0x6d, 0x55, 0xee, 0x4f, 0xb4, 0xe5,
	0x0e, 0xfc, 0x21, 0x58, 0x35, 0xcf, 0xdc, 0x68, 0xde, 0xba, 0x04, 0xf5, 0x2d, 0x09, 0xfe, 0x34,
	0xf0, 0xcb, 0x72, 0x0d, 0x28, 0xa5, 0xc2, 0xc3, 0xc6, 0x96, 0x87, 0x29, 0xd1, 0x1a, 0xed, 0x9c,
	0x1d, 0xd2, 0x60, 0xd4, 0x9a, 0x34, 0xe6, 0xc1, 0xd2, 0xfc, 0x6b, 0xc0, 0xd9, 0x2a, 0x95, 0x5b,
	0x79, 0x27, 0xc2, 0x35, 0x1e, 0xe3, 0x50, 0x1e, 0x94, 0xc1, 0x95, 0x81, 0xe3, 0x85, 0x99, 0x38,
	0x19, 0xd2, 0x4b, 0x30, 0xf0, 0xf4, 0x42, 0xe8, 0xd8, 0x1a, 0x60, 0xb4, 0xf8, 0x3d, 0xf2, 0x94,
	0x20, 0x69, 0x3a, 0x76, 0x86, 0xf0, 0x7e, 0x92, 0x4e, 0xef, 0x44, 0x92, 0xfd, 0x54, 0x66, 0x08,
	0x59, 0x12, 0x4c, 0x2d, 0x5b, 0x07, 0x1a, 0x7c, 0x66, 0x41, 0xb3, 0x18, 0x63, 0x06, 0xd0, 0xb8,
	0xf1, 0xc2, 0xb9, 0x2f, 0x7a, 0x47, 0xac, 0x09, 0xc7, 0x63, 0x25, 0xd3, 0xa8, 0x67, 0x9c, 0xff,
	0x75, 0x0a, 0x6d, 0x2a, 0xed, 0x46, 0xa8, 0x7b, 0x6f, 0x2a, 0xd8, 0x37, 0xd0, 0xb8, 0x70, 0x5d,
	0xf4, 0xdd, 0xe3, 0xb2, 0x8a, 0x75, 0xd7, 0x7f, 0x51, 0x56, 0x72, 0xee, 0x58, 0xeb, 0x08, 0x59,
	0xc6, 0xda, 0xbd, 0x25, 0x2c, 0xa9, 0x88, 0xa9, 0x2d, 0x7b, 0x59, 0xbe, 0x86, 0xa6, 0x2d, 0x02,
	0x79, 0x2f, 0xf6, 0x11, 0xf1, 0x32, 0x22, 0x7c, 0xb1, 0x8e, 0xd8, 0x18, 0x4e, 0xc6, 0x99, 0x89,
	0x3e, 0xae, 0xa8, 0x08, 0x1f, 0xfb, 0x2f, 0x2b, 0x93, 0xc1, 0x67, 0xeb, 0x88, 0x5d, 0x03, 0x8c,
	0x45, 0x72, 0xb9, 0x24, 0x4d, 0x77, 0xa6, 0xb3, 0x97, 0xea, 0x35, 0xb4, 0x88, 0xea, 0xca, 0xf3,
	0x13, 0xa1, 0xd8, 0x8b, 0x72, 0x2e, 0xfd, 0x7a, 0x48, 0x62, 0x5d, 0x3d, 0x92, 0x37, 0xf9, 0x1e,
	0xad, 0x2a, 0x14, 0x57, 0xdf, 0x4e, 0xb1, 0xae, 0xa0, 0x95, 0x51, 0xd1, 0xe6, 0x7d, 0x67, 0x9e,
	0xef, 0xe0, 0x74, 0x9c, 0x8f, 0xcb, 0x2e, 0xa5, 0x3e, 0xa9, 0xac, 0x8d, 0x3e, 0x4b, 0xed, 0x03,
	0x9d, 0xd1, 0xa5, 0x13, 0x0b, 0xf6, 0xb2, 0x22, 0x21, 0x1d, 0xb2, 0x33, 0xa5, 0x6f, 0x73, 0xa2,
	0x1f, 0xf0, 0x77, 0xe2, 0x9d, 0x2b, 0xbb, 0x86, 0xb6, 0xa6, 0xc9, 0x96, 0x7e, 0x55, 0x46, 0xfa,
	0xf9, 0x30, 0x2a, 0xbd, 0x2b, 0xde, 0x87, 0xea, 0xfb, 0xdc, 0x02, 0xfb, 0x3c, 0xb5, 0x5f, 0xa8,
	0xcb, 0xce, 0xaf, 0xad, 0xb5, 0x7f, 0x05, 0x26, 0x0d, 0x02, 0x5f, 0xfd, 0x1f, 0x00, 0x00, 0xff,
	0xff, 0xc8, 0xeb, 0xa8, 0x8c, 0x48, 0x0c, 0x00, 0x00,
}
