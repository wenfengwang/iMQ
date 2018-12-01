package baton

import (
	"context"
	"fmt"
	"github.com/pingcap/tidb/config"
	"github.com/pingcap/tidb/store/tikv"
	"github.com/pkg/errors"
	. "github.com/wenfengwang/iMQ/baton/pb"
	"sync/atomic"
)

var (
	ErrTopicNotFound        = errors.New("topic not found.")
	ErrSubscriptionNotFound = errors.New("subscription not found")
)

type batonServer struct {
	mdm               *metadataManager
	rm                *routeManager
	brokerIdGenerator uint64
}

func NewBatonServer(pdAddrs []string) BatonServer {
	tikvClient, err := tikv.NewRawKVClient(pdAddrs, config.Security{})
	if err != nil {
		panic(fmt.Sprintf("new TiKV client error: %v", err))
	}
	mdm := &metadataManager{tikvClient: tikvClient}
	mdm.init()
	return &batonServer{mdm, &routeManager{}, 0}
}

var count int64 = 0

func (s *batonServer) CreateTopic(ctx context.Context, request *CreateTopicRequest) (*Response, error) {
	atomic.AddInt64(&count, 1)
	return &Response{ResponseCode: Code_SUCCESS}, nil
}

func (s *batonServer) DeleteTopic(ctx context.Context, request *DeleteTopicRequest) (*Response, error) {
	return nil, nil
}

func (s *batonServer) CreateSubscription(ctx context.Context, request *CreateSubscriptionRequest) (*Response, error) {
	return nil, nil
}

func (s *batonServer) DeleteSubscription(ctx context.Context, request *DeleteSubscriptionRequest) (*Response, error) {
	return nil, nil
}

func (s *batonServer) BrokerHeartBeat(ctx context.Context, request *BrokerHBRequest) (*BrokerHBResponse, error) {
	return nil, nil
}

func (s *batonServer) RegisterBroker(ctx context.Context, request *RegisterBrokerRequest) (*Response, error) {
	return nil, nil
}

func (s *batonServer) UpdateRoute(ctx context.Context, request *UpdateRouteRequest) (*UpdateRouteResponse, error) {
	return nil, nil
}
