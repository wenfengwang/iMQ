package client

import (
	"github.com/wenfengwang/iMQ/pb"
	"sync"
	"sync/atomic"
	"github.com/prometheus/common/log"
)

type producer struct {
	pId       uint64
	topicName string
	mutex     sync.RWMutex
	routes    []*QueueRoute
	count     int64
	quitCh    chan interface{}
}

func (p *producer) Publish(*pb.Message) pb.PublishResult {
	return pb.PublishResult_SUCCESS
}

func (p *producer) PublishBatch(msgs []*pb.Message) pb.PublishResult {
	q := p.getQueue()
	if q == nil {
		return pb.PublishResult_NO_QUEUE_ROUTE
	}

	for _, msg := range msgs {
		msg.QueueId = q.queueId
	}
	_, err := q.broker.publish(&pb.PublishRequest{Msg: msgs})
	if err != nil {
		return pb.PublishResult_FAILED
	}

	return pb.PublishResult_SUCCESS
}

func (p *producer) PublishAsync(*pb.Message, func(pb.PublishResult)) {

}

func (p *producer) PublishBatchAsync([]*pb.Message, func(pb.PublishResult)) {

}

func (p *producer) start() {
	getRoute := func() {
		newRoutes := rHub.getQueueRoute(pb.Action_PUB, p.pId, p.topicName)
		if newRoutes != nil{
			for _, r := range newRoutes {
				log.Infof("got Route[queueId: %d, brokerAddress: %s, brokerId: %d]",
					r.queueId, r.broker.address, r.broker.brokerId)
			}
		} else {
			log.Info("none route got")
		}
		p.mutex.Lock()
		defer p.mutex.Unlock()
		p.routes = newRoutes
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
