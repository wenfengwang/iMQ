package baton

import (
	"context"
	. "github.com/wenfengwang/iMQ/baton/pb"
)

type batonServer struct {

}

func (s * batonServer) CreateTopic(context.Context, *CreateTopicRequest) (*Response, error) {
	return nil, nil
}

func (s * batonServer) DeleteTopic(context.Context, *DeleteTopicRequest) (*Response, error) {
	return nil, nil
}

func (s * batonServer) CreateSubscription(context.Context, *CreateSubscriptionRequest) (*Response, error) {
	return nil, nil
}

func (s * batonServer) DeleteSubscription(context.Context, *DeleteSubscriptionRequest) (*Response, error) {
	return nil, nil
}

func (s * batonServer) BrokerHeartBeat(context.Context, *BrokerHBRequest) (*BrokerHBResponse, error) {
	return nil, nil
}

func (s * batonServer) RegisterBroker(context.Context, *RegisterBrokerRequest) (*Response, error) {
	return nil, nil
}

func (s * batonServer) UpdateRoute(context.Context, *UpdateRouteRequest) (*UpdateRouteResponse, error) {
	return nil, nil
}
