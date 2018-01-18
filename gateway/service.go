package main

import (
	"github.com/liangpengcheng/qserver/protocol"
	"google.golang.org/grpc"
)

type service struct {
	serv    *protocol.Service
	rpc     protocol.GatewayCallServiceClient
	rpcConn *grpc.ClientConn
}

func (s *service) onDestroy() {
	if s.rpcConn != nil {
		s.rpcConn.Close()
	}
}
