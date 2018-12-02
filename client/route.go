package client

import (
	"github.com/wenfengwang/iMQ/baton/pb"
	"sync"
	"context"
)

type RouteHub struct {
	client batonpb.BatonClient
	queueRouteMap sync.Map
	topicRouteMap sync.Map
	bh *BrokerHub
}

type topicRoute struct {
	queueToInstance sync.Map
	instanceToQueue sync.Map
}

func (rh *RouteHub) getQueueRoute(action batonpb.Action, id uint64, topicName string) []*QueueRoute {
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

func (rh *RouteHub) updateRouteInfo(action batonpb.Action, id uint64, name string) {
	res, _ := rh.client.UpdateRoute(context.Background(), &batonpb.UpdateRouteRequest{Id: id, Name: name, Action: action})
	tr, _ := rh.topicRouteMap.Load(name)
	if res != nil {
		for _, le := range res.Leases  {
			queue, exist := rh.queueRouteMap.Load(le.QueueId)
			if !exist {
				nq := &QueueRoute{assignedId: id,
					broker: rh.bh.getBroker(&batonpb.BrokerInfo{BrokerId: le.BrokerId, Address: le.BrokerAddr}),
					queueId: le.QueueId,
					perm: le.Perm}

				queue, _ = rh.queueRouteMap.LoadOrStore(nq.queueId, nq)
			}
			queue.(*QueueRoute).expiredTime = le.ExpiredTime
			_, exist = tr.(*topicRoute).queueToInstance.Load(le.QueueId)
			if !exist {
				tr.(*topicRoute).queueToInstance.Store(le, queue)
				qs, _ := tr.(*topicRoute).instanceToQueue.Load(id)
				qqs := qs.([]*QueueRoute)
				qqqs := append(qqs, queue.(*QueueRoute))
				tr.(*topicRoute).instanceToQueue.Store(id, qqqs)
			}
		}
	}
}

type QueueRoute struct {
	queueId uint64
	assignedId uint64
	broker *pubsub
	expiredTime uint64
	perm batonpb.Permission
}
