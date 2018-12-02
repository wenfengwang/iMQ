package client

import (
	"github.com/wenfengwang/iMQ/broker/pb"
	pb "github.com/wenfengwang/iMQ/client/pb"
	"sync"
)

var BatonAddress string

var mutex sync.Mutex
var rHub *RouteHub

type Topic interface {
	Publish(*brokerpb.Message) pb.PublishResult
	PublishBatch([]*brokerpb.Message) pb.PublishResult
	PublishAsync(*brokerpb.Message, func(pb.PublishResult))
	PublishBatchAsync([]*brokerpb.Message, func(pb.PublishResult))
	Subscribe(string) Subscription
}

type Subscription interface {
	Pull(int32) []*brokerpb.Message
	Push(func([]*brokerpb.Message) pb.ConsumeResult)
}

func NewTopic(topicName string) Topic {
	if rHub == nil {
		mutex.Lock()
		rHub = &RouteHub{batonAddress: BatonAddress}
		rHub.start()
		mutex.Unlock()
	}
	return &entry{p: &producer{pId: 0, topicName: topicName}}
}

type entry struct {
	p    *producer
	cMap sync.Map
}

func (e *entry) Publish(msg *brokerpb.Message) pb.PublishResult {
	return e.p.Publish(msg)
}

func (e *entry) PublishBatch(msgs []*brokerpb.Message) pb.PublishResult {
	return e.p.PublishBatch(msgs)
}

func (e *entry) PublishAsync(msg *brokerpb.Message, f func(pb.PublishResult)) {
	e.p.PublishAsync(msg, f)
}

func (e *entry) PublishBatchAsync(msgs []*brokerpb.Message, f func(pb.PublishResult)) {
	e.p.PublishBatchAsync(msgs, f)
}

func (e *entry) Subscribe(name string) Subscription {
	con, exist := e.cMap.Load(name)

	if !exist {
		con, _ = e.cMap.LoadOrStore(name, &consumer{cId: 0, topicName: name})
	}
	return con.(*consumer)
}
