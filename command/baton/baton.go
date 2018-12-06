package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"net"
	"fmt"
	"strings"
	"google.golang.org/grpc"
	"github.com/wenfengwang/iMQ/pb"
	"github.com/wenfengwang/iMQ/baton"
)

var (
	pd        = flag.String("pd", "", "pd server address")
	localIP   = flag.String("local-ip", "0.0.0.0", "address register to baton")
	port      = flag.Int("port", 30000, "server port")
)

func main() {
	flag.Parse()

	log.SetOutput(os.Stdout)
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *localIP, *port))
	if err != nil {
		panic(fmt.Sprintf("listen: %s error: %v", fmt.Sprintf("%s:%d", *localIP, *port), err))
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	pb.RegisterBatonServer(server, baton.NewBatonServer(strings.Split(*pd, ";")))
	log.Info("baton server started.")
	server.Serve(listen)
}