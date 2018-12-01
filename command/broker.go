package main

import (
	"flag"
	"fmt"
	"github.com/wenfengwang/iMQ/broker"
	"github.com/wenfengwang/iMQ/broker/pb"
	"google.golang.org/grpc"
	"net"
	"strings"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	pd        = flag.String("pd", "", "pd server address")
	localIP   = flag.String("local-ip", "localhost", "address register to baton")
	port      = flag.Int("port", 23456, "server port")
	batonAddr = flag.String("baton-addr", "", "baton ip:port address")
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
	brokerpb.RegisterPubSubServer(server, broker.NewPubSubServer(broker.BrokerConfig{
		PD:           strings.Split(*pd, ";"),
		Address:      fmt.Sprintf("%s:%d", *localIP, *port),
		BatonAddress: *batonAddr}))
	log.Info("broker start serving...")
	server.Serve(listen)
}
