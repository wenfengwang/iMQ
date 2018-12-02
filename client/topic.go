package client

import (
	"github.com/wenfengwang/iMQ/broker/pb"
	pb "github.com/wenfengwang/iMQ/client/pb"
)

type Topic interface {
	Publish(*brokerpb.Message) pb.PublishResult
	PublishBatch([]*brokerpb.Message) pb.PublishResult
	PublishAsync(*brokerpb.Message, func(pb.PublishResult))
	PublishBatchAsync([]*brokerpb.Message, func(pb.PublishResult))
	Subscribe(string) Subscription
}

type Subscription interface {
	Pull(int32) PullResult
	Push(func([]*brokerpb.Message) pb.ConsumeResult)
}

func NewTopic(topicName string) Topic {
	return nil
}
