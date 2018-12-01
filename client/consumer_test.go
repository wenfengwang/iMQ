package client

import (
	"testing"
	"fmt"
	"github.com/wenfengwang/iMQ/broker/pb"
	"google.golang.org/grpc"
	"time"
)

func TestConsumer_Pull(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:23456", opts...)
	if err != nil {
		panic("dial server err")
	}

	fmt.Printf("Dial to %s success\n", "localhost:23456")
	c := consumer{topicName: "testTopic", brokerClient: brokerpb.NewPubSubClient(conn)}

	for  {
		msgs, _ := c.Pull(10e9+1, 16)
		for _, v := range msgs {
			fmt.Println("msg:", v.Body)
		}
		//fmt.Println("get message numbers:", len(msgs))
		time.Sleep(time.Second)
	}

}
