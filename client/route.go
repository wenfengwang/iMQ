package client

import (
	"context"
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/wenfengwang/iMQ/pb"
	"google.golang.org/grpc"
	"sync"
)

type RouteHub struct {
	batonAddress  string
	client        pb.BatonClient
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
	rh.client = pb.NewBatonClient(conn)
	return nil
}

func (rh *RouteHub) getQueueRoute(action pb.Action, id uint64, topicName string) []*QueueRoute {
	tr, exist := rh.topicRouteMap.Load(topicName)
	if !exist {
		tr, _ = rh.topicRouteMap.LoadOrStore(topicName, &topicRoute{})
		rh.updateRouteInfo(action, id, topicName)
	}

	qrs, exist := tr.(*topicRoute).instanceToQueue.Load(id)
	if !exist {
		return nil
	}
	return qrs.([]*QueueRoute)
}

func (rh *RouteHub) updateRouteInfo(action pb.Action, id uint64, name string) {
	log.Infof("prepare updateRouteInfo for [Action: %s, instanceId: %d, topicName: %s]",
		pb.Action_name[int32(action)], id, name)

	res, _ := rh.client.UpdateRoute(context.Background(), &pb.UpdateRouteRequest{Id: id, Name: name, Action: action})
	tr, _ := rh.topicRouteMap.Load(name)
	if res != nil {
		for _, le := range res.Leases {
			queue, exist := rh.queueRouteMap.Load(le.QueueId)
			if !exist {
				nq := &QueueRoute{assignedId: id,
					broker:  rh.bh.getBroker(&pb.BrokerInfo{BrokerId: le.BrokerId, Address: le.BrokerAddr}),
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
	perm        pb.Permission
}
