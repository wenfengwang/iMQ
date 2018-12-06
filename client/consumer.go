package client

import (
	"sync"
	"sync/atomic"
	"github.com/wenfengwang/iMQ/pb"
	"github.com/prometheus/common/log"
)

type consumer struct {
	cId       uint64
	topicName string
	mutex     sync.RWMutex
	routes    []*QueueRoute
	count     int64
	quitCh    chan interface{}
	pullMsgCh chan *pb.PullMessageResponse
}

func (c *consumer) Pull(numbers int32) []*pb.Message {
	queue := c.getQueue()
	if queue == nil {
		return nil
	}
	queue.broker.pullMessage(&pb.PullMessageRequest{QueueId: queue.queueId, Numbers: numbers}, c.pullMsgCh)
	res := <-c.pullMsgCh
	return res.Msg
}

func (c *consumer) Push(func([]*pb.Message) pb.ConsumeResult) {

}

func (c *consumer) start() {
	getRoute := func() {
		newRoutes := rHub.getQueueRoute(pb.Action_SUB, c.cId, c.topicName)
		if newRoutes != nil{
			for _, r := range newRoutes {
				log.Infof("got Route[queueId: %d, brokerAddress: %s, brokerId: %d]",
					r.queueId, r.broker.address, r.broker.brokerId)
			}
		} else {
			log.Info("none route got")
		}
		c.mutex.Lock()
		defer c.mutex.Unlock()
		c.routes = newRoutes
	}
	getRoute()
	//go func() {
	//	ticker := time.NewTicker(20 * time.Second)
	//	for {
	//		select {
	//		case <-ticker.C:
	//			getRoute()
	//		case <-p.quitCh:
	//			return
	//		}
	//	}
	//}()
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
