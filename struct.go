package main

import (
	pb "github.com/JKajas/at_modem_gateway/pb"
)

// Rpc server

type RpcServer struct {
	pb.UnimplementedSmsGatewayServer
}
