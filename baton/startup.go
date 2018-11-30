package baton

import (
	"github.com/wenfengwang/iMQ/baton/pb"
	"google.golang.org/grpc"
	"net"
)

func StartBatonServer() {
	listen, err := net.Listen("tcp", "localhost:23456")
	if err != nil {
		panic("listen err")
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBatonServer(grpcServer, &batonServer{})
	grpcServer.Serve(listen)
}
