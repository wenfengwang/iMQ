package client

import "github.com/wenfengwang/iMQ/broker/pb"
import (
	"github.com/wenfengwang/iMQ/baton/pb"
	pb "github.com/wenfengwang/iMQ/client/pb"
	"sync"
	"sync/atomic"
	"time"
)

type producer struct {
	pId       uint64
	topicName string
	mutex     sync.RWMutex
	routes    []*QueueRoute
	count     int64
	quitCh    chan interface{}
}

func (p *producer) Publish(*brokerpb.Message) pb.PublishResult {
	return pb.PublishResult_SUCCESS
}

func (p *producer) PublishBatch(msgs []*brokerpb.Message) pb.PublishResult {
	q := p.getQueue()
	if q == nil {
		return pb.PublishResult_NO_QUEUE_ROUTE
	}

	for _, msg := range msgs {
		msg.QueueId = q.queueId
	}
	_, err := q.broker.publish(&brokerpb.PublishRequest{Msg: msgs})
	if err != nil {
		return pb.PublishResult_FAILED
	}

	return pb.PublishResult_SUCCESS
}

func (p *producer) PublishAsync(*brokerpb.Message, func(pb.PublishResult)) {

}

func (p *producer) PublishBatchAsync([]*brokerpb.Message, func(pb.PublishResult)) {

}

func (p *producer) start() {
	go func() {
		ticker := time.NewTicker(20 * time.Second)
		for {
			select {
			case <-ticker.C:
				newRoutes := rHub.getQueueRoute(batonpb.Action_PUB, p.pId, p.topicName)
				p.mutex.Lock()
				p.routes = newRoutes
				p.mutex.RUnlock()
			case <-p.quitCh:
				return
			}
		}
	}()
}

func (p *producer) shutdown() {
	p.quitCh <- "exit"
}

func (p *producer) getQueue() *QueueRoute {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	if p.routes == nil || len(p.routes) == 0 {
		return nil
	}
	return p.routes[atomic.AddInt64(&p.count, 1)%int64(len(p.routes))]
}
