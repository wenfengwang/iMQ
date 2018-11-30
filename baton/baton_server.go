package baton

import (
	"context"
	"fmt"
	. "github.com/wenfengwang/iMQ/baton/pb"
	"sync/atomic"
	"time"
)

type batonServer struct {

}

func NewBatonServer() BatonServer {
	go func() {
		ticker := time.NewTicker(time.Second)
		var last int64 = 0;
		for {
			select {
			case  <- ticker.C:
				fmt.Println("tps: ", count - last)
				last = count
			}
		}
	}()
	return & batonServer{}
}
var count int64 = 0
func (s * batonServer) CreateTopic(ctx context.Context, request *CreateTopicRequest) (*Response, error) {
	atomic.AddInt64(&count, 1)
	return &Response{ResponseCode: Code_SUCCESS}, nil
}

func (s * batonServer) DeleteTopic(ctx context.Context, request *DeleteTopicRequest) (*Response, error) {
	return nil, nil
}

func (s * batonServer) CreateSubscription(ctx context.Context, request *CreateSubscriptionRequest) (*Response, error) {
	return nil, nil
}

func (s * batonServer) DeleteSubscription(ctx context.Context, request *DeleteSubscriptionRequest) (*Response, error) {
	return nil, nil
}

func (s * batonServer) BrokerHeartBeat(ctx context.Context, request *BrokerHBRequest) (*BrokerHBResponse, error) {
	return nil, nil
}

func (s * batonServer) RegisterBroker(ctx context.Context, request *RegisterBrokerRequest) (*Response, error) {
	return nil, nil
}

func (s * batonServer) UpdateRoute(ctx context.Context, request *UpdateRouteRequest) (*UpdateRouteResponse, error) {
	return nil, nil
}
