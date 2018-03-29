package mz

import (
	"net"

	"github.com/mzmico/toolkit/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type RPCServer = grpc.Server

type AddRpcServerHandler func(server *RPCServer)

var (
	rpcAddServerHandlers = make([]AddRpcServerHandler, 0)
)

type RpcService struct {
	Service
	s *RPCServer

	options []grpc.ServerOption
}

func (m *RpcService) AddRpcOptions(opts ...grpc.ServerOption) {
	m.options = append(m.options, opts...)
}

func (m *RpcService) Run() error {

	rpcInitLogger(m)

	m.s = grpc.NewServer(
		m.options...,
	)

	for _, handler := range rpcAddServerHandlers {
		handler(m.s)
	}

	l, err := net.Listen("tcp", m.Address())

	if err != nil {
		return errors.New("run service %s fail. %s", m.Address(), err)
	}

	grpclog.Infof("service %s listen on %s.", m.Name(), m.Address())

	err = m.s.Serve(l)

	if err != nil {
		return errors.New("run grpc service %s fail. %s", m.Name(), err)
	}

	return nil
}

func AddRpcServer(handler AddRpcServerHandler) {
	rpcAddServerHandlers = append(rpcAddServerHandlers, handler)
}

func NewRpcService(opts ...ServiceOption) (*RpcService, error) {

	service := &RpcService{}

	service.init()

	for _, opt := range opts {
		opt(service)
	}

	return service, nil
}
