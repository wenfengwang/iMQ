package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/wenfengwang/iMQ/baton/pb"
	"google.golang.org/grpc"
	"sync"
)

type RouteHub struct {
	batonAddress  string
	client        batonpb.BatonClient
	queueRouteMap sync.Map
	topicRouteMap sync.Map
	bh            *BrokerHub
	count uint64
}

type topicRoute struct {
	queueToInstance sync.Map
	instanceToQueue sync.Map
}

func (rh *RouteHub) start() error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(rh.batonAddress, opts...)
	if err != nil {
		return errors.New(fmt.Sprintf("dial baton: %s error %v", rh.batonAddress, err))
	}

	log.Info("Dial to %s success\n", rh.batonAddress)
	rh.client = batonpb.NewBatonClient(conn)
	return nil
}

func (rh *RouteHub) helper() *batonpb.BrokerInfo{
	//switch atomic.AddUint64(&rh.count,1) % 3 {
	//case 0:
	//	return &batonpb.BrokerInfo{BrokerId: 0, Address: "localhost:23456"}
	//case 1:
	//	return &batonpb.BrokerInfo{BrokerId: 1, Address: "localhost:23456"}
	//case 2:
	//	return &batonpb.BrokerInfo{BrokerId: 2, Address: "localhost:23456"}
	//}
	return &batonpb.BrokerInfo{BrokerId: 0, Address: "localhost:23456"}
}

func (rh *RouteHub) getQueueRoute(action batonpb.Action, id uint64, topicName string) []*QueueRoute {
	newRoutes := make([]*QueueRoute, 16)
	for i := 0; i < 16 ; i++  {
		newRoutes[i] = &QueueRoute{
			queueId: 2*10e8 + uint64(i),
			broker: rHub.bh.getBroker(rh.helper()),
		}
	}

	return newRoutes
	//tr, exist := rh.topicRouteMap.Load(topicName)
	//if !exist {
	//	tr, _ = rh.topicRouteMap.LoadOrStore(topicName, &topicRoute{})
	//	rh.updateRouteInfo(action, id, topicName)
	//}
	//
	//qrs, exist := tr.(*topicRoute).instanceToQueue.Load(id)
	//if !exist {
	//	return nil
	//}
	//return qrs.([]*QueueRoute)
}

func (rh *RouteHub) updateRouteInfo(action batonpb.Action, id uint64, name string) {
	log.Infof("prepare updateRouteInfo for [Action: %s, instanceId: %d, topicName: %s]",
		batonpb.Action_name[int32(action)], id, name)

	res, _ := rh.client.UpdateRoute(context.Background(), &batonpb.UpdateRouteRequest{Id: id, Name: name, Action: action})
	tr, _ := rh.topicRouteMap.Load(name)
	if res != nil {
		for _, le := range res.Leases {
			queue, exist := rh.queueRouteMap.Load(le.QueueId)
			if !exist {
				nq := &QueueRoute{assignedId: id,
					broker:  rh.bh.getBroker(&batonpb.BrokerInfo{BrokerId: le.BrokerId, Address: le.BrokerAddr}),
					queueId: le.QueueId,
					perm:    le.Perm}

				queue, _ = rh.queueRouteMap.LoadOrStore(nq.queueId, nq)
			}
			queue.(*QueueRoute).expiredTime = le.ExpiredTime
			_, exist = tr.(*topicRoute).queueToInstance.Load(le.QueueId)
			if !exist {
				tr.(*topicRoute).queueToInstance.Store(le.QueueId, queue)
				qs, _ := tr.(*topicRoute).instanceToQueue.Load(id)
				var qqs []*QueueRoute
				if qs == nil {
					qqs = make([]*QueueRoute, 0)
				} else {
					qqs = qs.([]*QueueRoute)
				}
				tr.(*topicRoute).instanceToQueue.Store(id, append(qqs, queue.(*QueueRoute)))
			}
		}
	}
}

type QueueRoute struct {
	queueId     uint64
	assignedId  uint64
	broker      *pubsub
	expiredTime uint64
	perm        batonpb.Permission
}
