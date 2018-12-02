package client

import (
	"testing"
	"time"
	"fmt"
	"github.com/wenfengwang/iMQ/broker/pb"
)

func TestConsumer_Pull(t *testing.T) {
	rHub = &RouteHub{batonAddress: "localhost:30000", bh: &BrokerHub{}}
	rHub.start()
	c := &consumer{topicName:"testTopic", quitCh:make(chan interface{}),pullMsgCh: make(chan *brokerpb.PullMessageResponse, 4096) }
	c.start()

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
				res := c.Pull(16)
				count += len(res)
			}
		}()
	}
	time.Sleep(time.Hour)
}