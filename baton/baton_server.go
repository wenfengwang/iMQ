package baton

import (
	"context"
	"fmt"
	"github.com/pingcap/tidb/config"
	"github.com/pingcap/tidb/store/tikv"
	"github.com/pkg/errors"
	. "github.com/wenfengwang/iMQ/baton/pb"
	"sync/atomic"
	log "github.com/sirupsen/logrus"
)

var (
	ErrTopicNotFound        = errors.New("topic not found.")
	ErrSubscriptionNotFound = errors.New("subscription not found")
)

var (
	producerIdGenerator uint64 = 0
	consumerIdGenerator uint64 = 0
	brokerIdGenerator   uint64 = 0
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
	return &batonServer{mdm, &routeManager{mdm: mdm}, 0}
}

var count int64 = 0

func (s *batonServer) CreateTopic(ctx context.Context, request *CreateTopicRequest) (*Response, error) {
	fmt.Println(fmt.Sprintf("CreateTopic: [name: %s, queueNumbers: %d, autoScaling:%v]",
		request.Name, request.QueueNumbers, request.AutoScaling))
	log.Info(fmt.Sprintf("CreateTopic: [name: %s, queueNumbers: %d, autoScaling:%v]",
		request.Name, request.QueueNumbers, request.AutoScaling))
	t := s.mdm.getTopic(request.Name)
	if t != nil {
		return &Response{ResponseCode: Code_TOPIC_ALREADY_EXIST}, nil
	}

	t = &topic{}
	t.topicId, _ = s.mdm.generateTopicId()
	t.topicName = request.Name
	t.autoScaling = request.AutoScaling
	t.queues = make([]*Queue, request.QueueNumbers)
	t.path = TopicMetaPathPrefix + fmt.Sprint(t.topicId)
	t.status = &topicStatus{producers: make(map[uint64]*ProducerInfo), consumers: make(map[uint64]*ConsumerInfo)}
	var i int32
	for i = 0; i < request.QueueNumbers; i++ {
		qId, _ := s.mdm.generateQueueId()
		q := &Queue{queueId: qId, path: QueueMetaPathPrefix + fmt.Sprint(qId), perm: Permission_READ_WRITE, status: &QueueStatus{}}
		s.mdm.putQueue(q) // TODO error
		t.queues[i] = q
	}

	s.mdm.putTopic(t)
	log.Infof("create topic: %s success, queue numbers: %d, topicId: %d", t.topicName, len(t.queues), t.topicId )
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

func (s *batonServer) RegisterBroker(ctx context.Context, request *RegisterBrokerRequest) (*RegisterBrokerResponse, error) {
	info := &BrokerInfo{BrokerId: atomic.AddUint64(&brokerIdGenerator, 1), Address: request.Addr}
	s.rm.brokerOnline(info)
	log.Infof("broker %s register success, ID: %d", info.Address, info.BrokerId)
	return &RegisterBrokerResponse{Id: info.BrokerId}, nil
}

func (s *batonServer) UpdateRoute(ctx context.Context, request *UpdateRouteRequest) (*UpdateRouteResponse, error) {
	resp := &UpdateRouteResponse{}
	t := s.mdm.getTopic(request.Name)
	if t == nil {
		// TODO error code
		return nil, ErrTopicNotFound
	}
	id := request.Id
	switch request.Action {
	case Action_PUB:
		if id == 0 {
			id = atomic.AddUint64(&producerIdGenerator, 1)
		}
		resp.Leases = s.rm.updateRouteLeaseForProducer(id, t)
	case Action_SUB:
		if id == 0 {
			id = atomic.AddUint64(&consumerIdGenerator, 1)
		}
		resp.Leases = s.rm.updateRouteLeaseForConsumer(id, t)
	}
	resp.Id = id
	return resp, nil
}
