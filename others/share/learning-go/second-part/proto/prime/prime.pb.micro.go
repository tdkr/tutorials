// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: prime/prime.proto

package prime

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

// Client API for Prime service

type PrimeService interface {
	GetPrime(ctx context.Context, in *PrimeRequest, opts ...client.CallOption) (*PrimeResponse, error)
}

type primeService struct {
	c    client.Client
	name string
}

func NewPrimeService(name string, c client.Client) PrimeService {
	return &primeService{
		c:    c,
		name: name,
	}
}

func (c *primeService) GetPrime(ctx context.Context, in *PrimeRequest, opts ...client.CallOption) (*PrimeResponse, error) {
	req := c.c.NewRequest(c.name, "Prime.GetPrime", in)
	out := new(PrimeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Prime service

type PrimeHandler interface {
	GetPrime(context.Context, *PrimeRequest, *PrimeResponse) error
}

func RegisterPrimeHandler(s server.Server, hdlr PrimeHandler, opts ...server.HandlerOption) error {
	type prime interface {
		GetPrime(ctx context.Context, in *PrimeRequest, out *PrimeResponse) error
	}
	type Prime struct {
		prime
	}
	h := &primeHandler{hdlr}
	return s.Handle(s.NewHandler(&Prime{h}, opts...))
}

type primeHandler struct {
	PrimeHandler
}

func (h *primeHandler) GetPrime(ctx context.Context, in *PrimeRequest, out *PrimeResponse) error {
	return h.PrimeHandler.GetPrime(ctx, in, out)
}
