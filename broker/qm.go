package broker

import (
	"github.com/pingcap/tidb/store/tikv"
	"sync"
)

type QueueManager struct {
	queueMap   sync.Map
	tikvClient *tikv.RawKVClient
}

func (qm *QueueManager) getQueue(qId uint64) Queue {
	v, exist := qm.queueMap.Load(qId)
	if exist {
		return v.(Queue)
	}
	q := NewQueue(qId, qm.tikvClient)
	qm.queueMap.Store(qId, q)
	return q
}
