package main

import (
	pb "at_modem_gateway/pb"
)

// Rpc server

type RpcServer struct {
	pb.UnimplementedSmsGatewayServer
}
