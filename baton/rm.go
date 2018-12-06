package baton

import (
	"github.com/wenfengwang/iMQ/pb"
	"sync"
	"time"
	log "github.com/sirupsen/logrus"
	"fmt"
)

// TODO 后期实现Subscription

const (
	// unit: Second
	RouteLeaseLength int64 = 30
)

type routeManager struct {
	mdm     *metadataManager
	brokers []*pb.BrokerInfo
	count   int
}

func (rm *routeManager) brokerOnline(info *pb.BrokerInfo) {
	rm.brokers = append(rm.brokers, info)
}

func (rm *routeManager) brokerOffline(info *pb.BrokerInfo) {

}

func (rm *routeManager) updateRouteLeaseForProducer(producerId uint64, t *topic) []*pb.Lease {
	if t == nil {
		return nil
	}

	qs := t.getQueueForProducer(producerId)

	if qs == nil {
		if !t.addProducer(producerId) {
			if t.autoScaling {
				rm.mdm.scale(t)
			} else {
				return nil
			}
			// TODO scale
		}
	}

	qs = t.getQueueForProducer(producerId)

	if len(qs) == 0 {
		return nil
	}

	leases := make([]*pb.Lease, len(qs))

	for i := 0; i < len(qs); i++ {
		qStats := qs[i].status
		if qStats.brokerForPub == nil {
			qStats.brokerForPub = rm.selectBroker()
		}
		leases[i] = &pb.Lease{
			BrokerId:   qStats.brokerForPub.BrokerId,
			BrokerAddr: qStats.brokerForPub.Address,
		}
	}

	return leases
}

func (rm *routeManager) updateRouteLeaseForConsumer(consumerId uint64, t *topic) []*pb.Lease {
	if t == nil {
		return nil
	}

	qs := t.getQueueForConsumer(consumerId)

	// Don't support scale for consumer
	if qs == nil || len(qs) == 0 {
		return nil
	}

	leases := make([]*pb.Lease, len(qs))
	for i := 0; i < len(qs); i++ {
		qStats := qs[i].status
		if qStats.brokerForSub == nil {
			qStats.brokerForSub = rm.selectBroker()
		}
		leases[i] = &pb.Lease{
			BrokerId:   qStats.brokerForSub.BrokerId,
			BrokerAddr: qStats.brokerForSub.Address,
		}
	}

	return leases
}

func (rm *routeManager) selectBroker() *pb.BrokerInfo {
	count++
	return rm.brokers[rm.count % len(rm.brokers)]
}

type topic struct {
	topicId     uint64
	path        string
	topicName   string
	queues      []*Queue
	autoScaling bool
	status      *topicStatus
}

func (t *topic) getQueueForProducer(pId uint64) []*Queue {
	t.status.pMutex.RLock()
	defer t.status.pMutex.RUnlock()

	p, exist := t.status.producers[pId]
	if !exist {
		return nil
	}
	qs := make([]*Queue, 0)
	for i := 0; i < len(t.queues); i++ {
		qStatus := t.queues[i].status
		if qStatus.curPID == p.pId || (qStatus.nextPID == p.pId && qStatus.pExpired <= time.Now().Unix()) {
			qStatus.pExpired = time.Now().Unix() + RouteLeaseLength
			qs = append(qs, t.queues[i])
		}
	}
	return qs
}

func (t *topic) getQueueForConsumer(cId uint64) []*Queue {
	t.status.cMutex.RLock()
	defer t.status.cMutex.RUnlock()

	c, exist := t.status.consumers[cId]
	if !exist {
		return nil
	}

	qs := make([]*Queue, 0)
	for i := 0; i < len(t.queues); i++ {
		qStatus := t.queues[i].status
		if qStatus.curCID == c.cId || (qStatus.nextCID == c.cId && qStatus.pExpired <= time.Now().Unix()) {
			qStatus.pExpired = time.Now().Unix() + RouteLeaseLength
			qs = append(qs, t.queues[i])
		}
	}
	return qs
}

func (t *topic) addProducer(pId uint64) bool {
	log.Info(fmt.Sprintf("add producer: %d", pId))
	_, exist := t.status.producers[pId]
	if exist {
		return true
	}
	status := t.status
	if len(t.queues) <= len(status.producers) {
		return false
	}

	t.status.pMutex.Lock()
	defer t.status.pMutex.Unlock()

	status.producers[pId] = &ProducerInfo{pId: pId}
	queuePerProducer := len(t.queues) / len(status.producers)
	remainingQueue := len(t.queues) - queuePerProducer*len(status.producers)
	queueIndex := 0
	// TODO guarantee order
	for _, p := range status.producers {
		for j := 0; j < queuePerProducer; j++ {
			t.queues[queueIndex].status.nextPID = p.pId
			queueIndex++
		}
		if remainingQueue > 0 {
			t.queues[queueIndex].status.nextPID = p.pId
			queueIndex++
			remainingQueue--
		}
	}

	return true
}

func (t *topic) removeProducer(pId uint64) {

}

func (t *topic) addConsumer(cId uint64) bool {
	_, exist := t.status.consumers[cId]
	if exist {
		return true
	}
	status := t.status
	if len(t.queues) >= len(status.consumers) {
		return false
	}

	t.status.cMutex.Lock()
	defer t.status.cMutex.Unlock()

	status.consumers[cId] = &ConsumerInfo{cId: cId}
	queuePerConsumer := len(t.queues) / len(status.consumers)
	remainingQueue := len(t.queues) - queuePerConsumer*len(status.consumers)
	queueIndex := 0
	// TODO guarantee order
	for _, c := range status.consumers {
		for j := 0; j < queuePerConsumer; j++ {
			t.queues[queueIndex].status.nextCID = c.cId
			queueIndex++
		}
		if remainingQueue > 0 {
			t.queues[queueIndex].status.nextCID = c.cId
			queueIndex++
			remainingQueue--
		}
	}

	return true
}

func (t *topic) removeConsumer(cId uint64) {

}

// TODO move to pb
type ProducerInfo struct {
	pId uint64
}

type ConsumerInfo struct {
	cId uint64
}

type topicStatus struct {
	pMutex    sync.RWMutex
	producers map[uint64]*ProducerInfo

	cMutex    sync.RWMutex
	consumers map[uint64]*ConsumerInfo
}

type Queue struct {
	queueId uint64

	// TopicMetaPathPrefix + /topicId
	path   string
	perm   pb.Permission
	mutex  sync.Mutex
	status *QueueStatus
}

func (q *Queue) QueueId() uint64 {
	return q.queueId
}

func (q *Queue) Path() string {
	return q.path
}

func (q *Queue) Permission() pb.Permission {
	return q.perm
}

func (q *Queue) MinimumOffset() uint64 {
	return 0
}

func (q *Queue) MaximumOffset() uint64 {
	return 0
}

type QueueStatus struct {
	curPID       uint64
	nextPID      uint64
	pExpired     int64
	pMutex       sync.RWMutex
	brokerForPub *pb.BrokerInfo

	curCID       uint64
	nextCID      uint64
	cExpired     int64
	cMutex       sync.RWMutex
	brokerForSub *pb.BrokerInfo
}
