package client

import (
	"context"
	"github.com/wenfengwang/iMQ/baton/pb"
	"github.com/wenfengwang/iMQ/broker/pb"
)

type Producer interface {
	Start() error
	Shutdown() error
	Publish(msg *brokerpb.Message) error
}

type producer struct {
	batonClient  batonpb.BatonClient
	brokerClient brokerpb.PubSubClient
	topicName    string
}

func (p *producer) Publish(msgs []*brokerpb.Message) error {
	_, err := p.brokerClient.Publish(context.Background(), &brokerpb.PublishRequest{Msg: msgs})
	if err != nil {
		return err
	}
	return nil
}
