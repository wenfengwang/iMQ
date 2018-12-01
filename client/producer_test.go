package client

import (
	"fmt"
	"github.com/wenfengwang/iMQ/broker/pb"
	"google.golang.org/grpc"
	"testing"
	"sync/atomic"
)

func TestProducer_Publish(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:23456", opts...)
	if err != nil {
		panic("dial server err")
	}

	fmt.Printf("Dial to %s success\n", "localhost:23456")
	p := producer{topicName: "testTopic", brokerClient: brokerpb.NewPubSubClient(conn)}
	var count uint64 = 0
	for i := 0; i < 100; i++ {
		batch := 1

		//fmt.Println("prepare to send message...")
		msgs := make([]*brokerpb.Message, batch)
		for j := 0; j < batch ; j++  {
			msgs[j] = &brokerpb.Message{MessageId: 10e12, QueueId: 10e9 + 1, Body: fmt.Sprint(i)}
		}
		err := p.Publish(msgs)
		if err != nil {
			//fmt.Printf("send mssage error: %v\n", err)
			//return
		} else {
			atomic.AddUint64(&count, uint64(batch))
			//fmt.Println("send message success.")
		}
		//time.Sleep(time.Second)
	}

	//ticker := time.NewTicker(time.Second)
	//var last uint64 = 0
	//for  {
	//	select {
	//	case <- ticker.C:
	//		fmt.Println("TPS:", count - last)
	//		last = count
	//	}
	//}
}
