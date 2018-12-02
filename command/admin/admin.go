package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/wenfengwang/iMQ/baton/pb"
	"google.golang.org/grpc"
	"flag"
	"context"
)

var (
	baton        = flag.String("baton-addr", "localhost:30000", "baton server address")
	topicName   = flag.String("topicName", "", "topicName")
	queueNumbers      = flag.Int("queue-numbers", 4, "queue numbers")
)

func main()  {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(*baton, opts...)
	if err != nil {
		panic( fmt.Sprintf("dial baton: %s error %v", *baton, err))
	}
	if *topicName == "" {
		log.Error("topic name can not be nil")
		return
	}

	log.Info("Dial to %s success\n", *baton)
	client := batonpb.NewBatonClient(conn)
	_, err = client.CreateTopic(context.Background(),
		&batonpb.CreateTopicRequest{
			Name: *topicName,
			QueueNumbers: int32(*queueNumbers),
			AutoScaling: false})

	if err != nil {
		fmt.Printf("crate topic: %s to %s failed. error: %v", *topicName, *baton, err)
	} else {
		fmt.Printf("crate topic: %s to %s successed.", *topicName, *baton)
	}
}
