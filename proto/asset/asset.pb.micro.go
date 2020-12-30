// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/asset/asset.proto

package asset

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AssetService service

type AssetService interface {
	AddOne(ctx context.Context, in *ReqAssetAdd, opts ...client.CallOption) (*ReplyAssetOne, error)
	GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAssetOne, error)
	RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error)
	GetList(ctx context.Context, in *ReqAssetList, opts ...client.CallOption) (*ReplyAssetList, error)
	GetByOwner(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAssetList, error)
	UpdateSnapshot(ctx context.Context, in *ReqAssetFlag, opts ...client.CallOption) (*ReplyInfo, error)
	UpdateSmall(ctx context.Context, in *ReqAssetFlag, opts ...client.CallOption) (*ReplyInfo, error)
}

type assetService struct {
	c    client.Client
	name string
}

func NewAssetService(name string, c client.Client) AssetService {
	return &assetService{
		c:    c,
		name: name,
	}
}

func (c *assetService) AddOne(ctx context.Context, in *ReqAssetAdd, opts ...client.CallOption) (*ReplyAssetOne, error) {
	req := c.c.NewRequest(c.name, "AssetService.AddOne", in)
	out := new(ReplyAssetOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAssetOne, error) {
	req := c.c.NewRequest(c.name, "AssetService.GetOne", in)
	out := new(ReplyAssetOne)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) RemoveOne(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "AssetService.RemoveOne", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetList(ctx context.Context, in *ReqAssetList, opts ...client.CallOption) (*ReplyAssetList, error) {
	req := c.c.NewRequest(c.name, "AssetService.GetList", in)
	out := new(ReplyAssetList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) GetByOwner(ctx context.Context, in *RequestInfo, opts ...client.CallOption) (*ReplyAssetList, error) {
	req := c.c.NewRequest(c.name, "AssetService.GetByOwner", in)
	out := new(ReplyAssetList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) UpdateSnapshot(ctx context.Context, in *ReqAssetFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "AssetService.UpdateSnapshot", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetService) UpdateSmall(ctx context.Context, in *ReqAssetFlag, opts ...client.CallOption) (*ReplyInfo, error) {
	req := c.c.NewRequest(c.name, "AssetService.UpdateSmall", in)
	out := new(ReplyInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AssetService service

type AssetServiceHandler interface {
	AddOne(context.Context, *ReqAssetAdd, *ReplyAssetOne) error
	GetOne(context.Context, *RequestInfo, *ReplyAssetOne) error
	RemoveOne(context.Context, *RequestInfo, *ReplyInfo) error
	GetList(context.Context, *ReqAssetList, *ReplyAssetList) error
	GetByOwner(context.Context, *RequestInfo, *ReplyAssetList) error
	UpdateSnapshot(context.Context, *ReqAssetFlag, *ReplyInfo) error
	UpdateSmall(context.Context, *ReqAssetFlag, *ReplyInfo) error
}

func RegisterAssetServiceHandler(s server.Server, hdlr AssetServiceHandler, opts ...server.HandlerOption) error {
	type assetService interface {
		AddOne(ctx context.Context, in *ReqAssetAdd, out *ReplyAssetOne) error
		GetOne(ctx context.Context, in *RequestInfo, out *ReplyAssetOne) error
		RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error
		GetList(ctx context.Context, in *ReqAssetList, out *ReplyAssetList) error
		GetByOwner(ctx context.Context, in *RequestInfo, out *ReplyAssetList) error
		UpdateSnapshot(ctx context.Context, in *ReqAssetFlag, out *ReplyInfo) error
		UpdateSmall(ctx context.Context, in *ReqAssetFlag, out *ReplyInfo) error
	}
	type AssetService struct {
		assetService
	}
	h := &assetServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AssetService{h}, opts...))
}

type assetServiceHandler struct {
	AssetServiceHandler
}

func (h *assetServiceHandler) AddOne(ctx context.Context, in *ReqAssetAdd, out *ReplyAssetOne) error {
	return h.AssetServiceHandler.AddOne(ctx, in, out)
}

func (h *assetServiceHandler) GetOne(ctx context.Context, in *RequestInfo, out *ReplyAssetOne) error {
	return h.AssetServiceHandler.GetOne(ctx, in, out)
}

func (h *assetServiceHandler) RemoveOne(ctx context.Context, in *RequestInfo, out *ReplyInfo) error {
	return h.AssetServiceHandler.RemoveOne(ctx, in, out)
}

func (h *assetServiceHandler) GetList(ctx context.Context, in *ReqAssetList, out *ReplyAssetList) error {
	return h.AssetServiceHandler.GetList(ctx, in, out)
}

func (h *assetServiceHandler) GetByOwner(ctx context.Context, in *RequestInfo, out *ReplyAssetList) error {
	return h.AssetServiceHandler.GetByOwner(ctx, in, out)
}

func (h *assetServiceHandler) UpdateSnapshot(ctx context.Context, in *ReqAssetFlag, out *ReplyInfo) error {
	return h.AssetServiceHandler.UpdateSnapshot(ctx, in, out)
}

func (h *assetServiceHandler) UpdateSmall(ctx context.Context, in *ReqAssetFlag, out *ReplyInfo) error {
	return h.AssetServiceHandler.UpdateSmall(ctx, in, out)
}
