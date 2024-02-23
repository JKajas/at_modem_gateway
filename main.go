package main

import (
	"context"
	"log"
	"net"

	pb "github.com/JKajas/at_modem_gateway/pb"
	"github.com/tarm/serial"
	"google.golang.org/grpc"
)

var Modem *serial.Port

func init() {
	Modem = ConnectModem()
	InitModem(Modem)
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:8000")
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSmsGatewayServer(grpcServer, &RpcServer{})
	defer Modem.Close()
	grpcServer.Serve(lis)
}

func (s *RpcServer) SendSms(ctx context.Context, in *pb.GtwRequest) (*pb.GtwResponse, error) {
	_, err := WriteSMS(Modem, in.Message, in.PhoneNumber)
	if err != nil {
		return &pb.GtwResponse{StatusCode: 404}, err
	}
	return &pb.GtwResponse{StatusCode: 200}, nil
}

func (s *RpcServer) mustEmbedUnimplementedSmsGatewayServer() {}
