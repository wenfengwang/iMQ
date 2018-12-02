package client

import (
	"github.com/wenfengwang/iMQ/broker/pb"
	pb "github.com/wenfengwang/iMQ/client/pb"
	"sync"
	"sync/atomic"
)

type consumer struct {
	cId       uint64
	topicName string
	mutex     sync.RWMutex
	routes    []*QueueRoute
	count     int64
	quitCh    chan interface{}
	pullMsgCh chan *brokerpb.PullMessageResponse
}

func (c *consumer) Pull(numbers int32) []*brokerpb.Message {
	queue := c.getQueue()
	if queue == nil {
		return nil
	}
	queue.broker.pullMessage(&brokerpb.PullMessageRequest{QueueId: queue.queueId, Numbers: numbers}, c.pullMsgCh)
	res := <-c.pullMsgCh
	return res.Msg
}

func (c *consumer) Push(func([]*brokerpb.Message) pb.ConsumeResult) {

}

func (c *consumer) shutdown() {
	c.quitCh <- "exit"
}

func (c *consumer) getQueue() *QueueRoute {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if c.routes == nil || len(c.routes) == 0 {
		return nil
	}
	return c.routes[atomic.AddInt64(&c.count, 1)%int64(len(c.routes))]
}
