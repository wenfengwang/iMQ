package client

import (
	"github.com/wenfengwang/iMQ/pb"
	"sync"
)

var BatonAddress string

var mutex sync.Mutex
var rHub *RouteHub

type Topic interface {
	Publish(*pb.Message) pb.PublishResult
	PublishBatch([]*pb.Message) pb.PublishResult
	PublishAsync(*pb.Message, func(pb.PublishResult))
	PublishBatchAsync([]*pb.Message, func(pb.PublishResult))
	Subscribe(string) Subscription
}

type Subscription interface {
	Pull(int32) []*pb.Message
	Push(func([]*pb.Message) pb.ConsumeResult)
}

func NewTopic(topicName string) Topic {
	if rHub == nil {
		mutex.Lock()
		rHub = &RouteHub{batonAddress: BatonAddress}
		rHub.start()
		mutex.Unlock()
	}
	pro := &producer{pId: 0, topicName: topicName}
	pro.start()
	return &entry{p: pro}
}

type entry struct {
	p    *producer
	cMap sync.Map
}

func (e *entry) Publish(msg *pb.Message) pb.PublishResult {
	return e.p.Publish(msg)
}

func (e *entry) PublishBatch(msgs []*pb.Message) pb.PublishResult {
	return e.p.PublishBatch(msgs)
}

func (e *entry) PublishAsync(msg *pb.Message, f func(pb.PublishResult)) {
	e.p.PublishAsync(msg, f)
}

func (e *entry) PublishBatchAsync(msgs []*pb.Message, f func(pb.PublishResult)) {
	e.p.PublishBatchAsync(msgs, f)
}

func (e *entry) Subscribe(name string) Subscription {
	con, exist := e.cMap.Load(name)

	if !exist {
		con, _ = e.cMap.LoadOrStore(name, &consumer{cId: 0, topicName: name})
		con.(*consumer).start()
	}
	return con.(*consumer)
}
