package main

import (
	"net"

	"google.golang.org/grpc"
)

type server struct {
	orders.UnimplementedOrderServiceServer
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic("Unable to start tcp server")
	}

	grpcServer := grpc.NewServer()
	orders.RegisterOrderServiceServer(&server{})
	grpcServer.Serve(listener)

}
