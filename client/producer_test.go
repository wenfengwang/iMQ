package client

import (
	"testing"
	"time"
	"fmt"
	"github.com/wenfengwang/iMQ/pb"
)

func TestProducer_Publish(t *testing.T) {
	rHub = &RouteHub{batonAddress: "localhost:30000", bh: &BrokerHub{}}
	rHub.start()
	p := &producer{topicName:"testTopic", quitCh:make(chan interface{})}
	p.start()

	//time.Sleep(time.Hour)

	count := 0

	go func() {
		ticker := time.NewTicker(time.Second)
		last := 0
		for  {
			select {
			case <- ticker.C:
				fmt.Println("TPS: ", count - last)
				last = count
			}
		}
	}()

	for i:=0; i < 16 ; i++  {
		go func() {
			for {
				msgs := make([]*pb.Message, 16)
				for i := 0; i < 16; i++ {
					msgs[i] = &pb.Message{Body:"testsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestsssssstestssssss"}
				}
				res := p.PublishBatch(msgs)
				if res == pb.PublishResult_SUCCESS {
					count += 16
				}
			}
		}()
	}
	time.Sleep(time.Hour)
}