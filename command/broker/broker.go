package main

import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/wenfengwang/iMQ/broker"
	"github.com/wenfengwang/iMQ/pb"
	"google.golang.org/grpc"
	"net"
	"os"
	"strings"
)

var (
	pd        = flag.String("pd", "", "pd server address")
	localIP   = flag.String("local-ip", "localhost", "address register to baton")
	port      = flag.Int("port", 23456, "server port")
	batonAddr = flag.String("baton-addr", "localhost:30000", "baton ip:port address")
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
	pb.RegisterPubSubServer(server, broker.NewPubSubServer(broker.BrokerConfig{
		PD:           strings.Split(*pd, ";"),
		Address:      fmt.Sprintf("%s:%d", *localIP, *port),
		BatonAddress: *batonAddr}))
	log.Info("broker start serving...")
	server.Serve(listen)
}
