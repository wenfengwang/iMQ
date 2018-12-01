package baton

import (
	"github.com/wenfengwang/iMQ/baton/pb"
	"sync"
	"time"
)

// TODO 后期实现Subscription

const (
	// unit: Second
	RouteLeaseLength int64 = 30
)

type routeManager struct {
}

func (rm *routeManager) brokerOnline(info *BrokerInfo) {

}

func (rm *routeManager) brokerOffline(info *BrokerInfo) {

}

func (rm *routeManager) allocateRouteLeaseForProducer(producerId uint64, t *topic) []*pb.Lease {
	if t == nil {
		return nil
	}

	return nil
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
	qs := make([]*Queue, 0)
	for i := 0; i < len(t.queues); i++ {
		qStatus := t.queues[i].status
		if qStatus.curPID == pId || (qStatus.nextPID == pId && qStatus.pExpired <= time.Now().Unix()) {
			qStatus.pExpired = time.Now().Unix() + RouteLeaseLength
			qs = append(qs, t.queues[i])
		}
	}
	return qs
}

func (t *topic) getQueueForConsumer(cId uint64) []*Queue {
	qs := make([]*Queue, 0)
	for i := 0; i < len(t.queues); i++ {
		qStatus := t.queues[i].status
		if qStatus.curCID == cId || (qStatus.nextCID == cId && qStatus.pExpired <= time.Now().Unix()) {
			qStatus.pExpired = time.Now().Unix() + RouteLeaseLength
			qs = append(qs, t.queues[i])
		}
	}
	return qs
}

func (t *topic) addProducer(pId uint64) bool {
	_, exist := t.status.producers[pId]
	if exist {
		return true
	}
	status := t.status
	if len(t.queues) >= len(status.producers) {
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

func (t *topic) scale() {
	// TODO
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
	path string
	perm pb.Permission

	mutex  sync.Mutex
	status *QueueStatus
	//assignedConsumer uint64
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

//
//func (q *Queue) assignBroker(info *BrokerInfo) bool {
//	q.mutex.Lock()
//	defer q.mutex.Unlock()
//	if q.assignedBroker != nil {
//		return false
//	}
//
//	q.assignedBroker = info
//	return true
//}

func (q *Queue) MinimumOffset() uint64 {
	return 0
}

func (q *Queue) MaximumOffset() uint64 {
	return 0
}

type BrokerInfo struct {
	brokerId uint64
	address  string
}

type QueueStatus struct {
	curPID       uint64
	nextPID      uint64
	pExpired     int64
	pMutex       sync.RWMutex
	brokerForPub *BrokerInfo

	curCID       uint64
	nextCID      uint64
	cExpired     int64
	cMutex       sync.RWMutex
	brokerForSub *BrokerInfo
}
