package examples

import (
	"fmt"
	"github.com/wenfengwang/iMQ/baton/pb"
	"google.golang.org/grpc"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:23456")
	if err != nil {
		panic("listen err")
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBatonServer(grpcServer, NewBatonServer())
	fmt.Println("baton server started...")
	grpcServer.Serve(listen)
	fmt.Println("baton server started...")
}
