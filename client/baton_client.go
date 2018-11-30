package client

import (
	"github.com/wenfengwang/iMQ/baton/pb"
	"google.golang.org/grpc"
)

func NewBatonClient() pb.BatonClient{
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:23456", opts...)
	if err != nil {
		panic("dial server err")
	}

	return pb.NewBatonClient(conn)
}