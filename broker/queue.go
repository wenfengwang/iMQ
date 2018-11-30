package broker

import (
	"encoding/binary"
	"fmt"
	"github.com/pingcap/tidb/store/tikv"
	"time"
)

const (
	MaxBatchSize          = 64
	queueIDShift          = 40
	queueMetadataPrefix   = "/imq/metadata/queue"
	flushMetadataInterval = 10 * time.Second
)

type Queue interface {
	QueueId() uint64
	Put([][]byte) error
	Get(num int) ([][]byte, error)
	Destroy()
}

func NewQueue(qId uint64, c *tikv.RawKVClient) Queue {
	q := &queue{id: qId, tikvClient: c}
	q.init()
	go q.flushMetadata()
	return q
}

type queue struct {
	tikvClient    *tikv.RawKVClient
	id            uint64
	writePosition uint64
	readPosition  uint64
	batchPutKeys  [][]byte
	getStartKey   []byte
	quit          chan interface{}
}

func (q *queue) QueueId() uint64 {
	return q.id
}

func (q *queue) Put(values [][]byte) error {
	for i := 0; i < len(values); i++ {
		binary.BigEndian.PutUint64(q.batchPutKeys[i], (q.id<<queueIDShift)|q.writePosition)
		q.writePosition++
	}
	return q.tikvClient.BatchPut(q.batchPutKeys[:len(values)], values)
}

func (q *queue) Get(num int) ([][]byte, error) {
	binary.BigEndian.PutUint64(q.getStartKey, (q.id<<queueIDShift)|q.readPosition)
	_, values, err := q.tikvClient.Scan(q.getStartKey, num)
	q.readPosition += uint64(len(values))
	return values, err
}

func (q *queue) Destroy() {
	q.quit <- "exit"
}

func (q *queue) init() {
	q.batchPutKeys = make([][]byte, MaxBatchSize)
	for i := 0; i < MaxBatchSize; i++ {
		q.batchPutKeys[i] = make([]byte, 8)
	}
	q.getStartKey = make([]byte, 8)

}

func (q *queue) flushMetadata() {
	ticker := time.NewTicker(flushMetadataInterval)
	wPosKey := []byte(fmt.Sprint("%s/%d/writePosition", queueMetadataPrefix, q.id))
	rPosKey := []byte(fmt.Sprint("%s/%d/readPosition", queueMetadataPrefix, q.id))
	value := make([]byte, 8)
	flush := func() {
		binary.BigEndian.PutUint64(wPosKey, q.writePosition)
		q.tikvClient.Put(wPosKey, value)
		binary.BigEndian.PutUint64(rPosKey, q.readPosition)
		q.tikvClient.Put(rPosKey, value)
	}
	for {
		select {
		case <-ticker.C:
			flush()
		case <-q.quit:
			flush()
			break
		}
	}
}
