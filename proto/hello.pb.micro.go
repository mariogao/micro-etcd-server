// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: hello.proto

package mymicro

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Hi service

func NewHiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Hi service

type HiService interface {
	MyHello(ctx context.Context, in *HelloReq, opts ...client.CallOption) (*HelloRes, error)
}

type hiService struct {
	c    client.Client
	name string
}

func NewHiService(name string, c client.Client) HiService {
	return &hiService{
		c:    c,
		name: name,
	}
}

func (c *hiService) MyHello(ctx context.Context, in *HelloReq, opts ...client.CallOption) (*HelloRes, error) {
	req := c.c.NewRequest(c.name, "Hi.MyHello", in)
	out := new(HelloRes)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hi service

type HiHandler interface {
	MyHello(context.Context, *HelloReq, *HelloRes) error
}

func RegisterHiHandler(s server.Server, hdlr HiHandler, opts ...server.HandlerOption) error {
	type hi interface {
		MyHello(ctx context.Context, in *HelloReq, out *HelloRes) error
	}
	type Hi struct {
		hi
	}
	h := &hiHandler{hdlr}
	return s.Handle(s.NewHandler(&Hi{h}, opts...))
}

type hiHandler struct {
	HiHandler
}

func (h *hiHandler) MyHello(ctx context.Context, in *HelloReq, out *HelloRes) error {
	return h.HiHandler.MyHello(ctx, in, out)
}
