package client

import (
	"context"
	"github.com/wenfengwang/iMQ/baton/pb"
	"github.com/wenfengwang/iMQ/broker/pb"
)

type Consumer interface {
	Subscribe(uint64, func(*brokerpb.Message)) error
	Pull(uint64) ([]*brokerpb.Message, error)
}

type consumer struct {
	batonClient  batonpb.BatonClient
	brokerClient brokerpb.PubSubClient
	pullStream   brokerpb.PubSub_PullMessageClient
	topicName    string
}

func (c *consumer) Subscribe(qId uint64, cfunc func(*brokerpb.Message)) error {

	return nil
}

func (c *consumer) Pull(qId uint64, number int32) ([]*brokerpb.Message, error) {
	if c.pullStream == nil {
		stream, _ := c.brokerClient.PullMessage(context.Background())
		c.pullStream = stream
	}
	// TODO error
	c.pullStream.Send(&brokerpb.PullMessageRequest{QueueId: qId, Numbers: number})

	response, _ := c.pullStream.Recv()
	return response.Msg, nil
}
