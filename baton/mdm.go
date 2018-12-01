package baton

import (
	"github.com/pingcap/tidb/store/tikv"
	"sync"
	"github.com/wenfengwang/iMQ/baton/pb"
	"github.com/prometheus/common/log"
	"github.com/pkg/errors"
	"fmt"
	"strings"
	"strconv"
	"encoding/binary"
)

var (

)

const (
	TopicMetaPathPrefix = "/imq/metadata/topic/"
	QueueMetaPathPrefix = "/imq/metadata/queue/"
	topicIdGeneratorPath = "/imq/metadata/topic_generator"
	queueIdGeneratorPath = "/imq/metadata/queue_generator"
)

var (
	ErrTopicAlreadyExist = errors.New("topic already exist.")
	ErrQueueAlreadyExist = errors.New("queue already exist.")
)

type metadataManager struct {
	tikvClient *tikv.RawKVClient

	topicMapMutex sync.RWMutex
	topicMap map[string]*topic

	queueMapMutex sync.RWMutex
	queueMap map[uint64]*Queue

	tgMutex sync.Mutex
	topicIdGenerator uint64

	qGMutex sync.Mutex
	queueIdGenerator uint64
}

func (mdm *metadataManager) init() {
	value, err := mdm.tikvClient.Get([]byte(topicIdGeneratorPath))
	if err != nil {
		panic(fmt.Sprintf("error: %v, get topicIdGenerator FAILED !!!", err))
	}
	mdm.topicIdGenerator = binary.BigEndian.Uint64(value)

	value, err = mdm.tikvClient.Get([]byte(queueIdGeneratorPath))
	if err != nil {
		panic(fmt.Sprintf("error: %v, get queueIdGenerator FAILED !!!", err))
	}
	mdm.queueIdGenerator = binary.BigEndian.Uint64(value)
	mdm.loadTopicsAndQueues()
}

func (mdm *metadataManager) loadTopicsAndQueues() {
	mdm.topicMap = make(map[string]*topic)
	value, err := mdm.tikvClient.Get([]byte(TopicMetaPathPrefix + "/allTopics"))
	if err != nil {
		panic(fmt.Sprintf("error: %v, load topics FAILED !!!", err))
	}

	for i := 0; i< len(value) ; i += 8  {
		topicId := binary.BigEndian.Uint64(value[i:i+8])
		t := mdm.loadTopic(topicId)
		if t == nil {
			log.Errorf("load topicId: %v but it doesn't exist.", topicId)
			continue
		}
		for j := 0; j < len(t.queues); j++ {
			mdm.queueMap[t.queues[i].queueId] = t.queues[i]
		}
		mdm.topicMap[t.topicName] = t
	}
}

func (mdm *metadataManager) loadTopic(topicId uint64) *topic {
	value, err := mdm.tikvClient.Get([]byte(TopicMetaPathPrefix + fmt.Sprint(topicId)))
	if err != nil {
		log.Errorf("load topic with Id error: %v", err)
		return nil
	}

	return mdm.decodeTopic(value)
}

func (mdm *metadataManager) loadQueue(queueId uint64) *Queue {
	value, err := mdm.tikvClient.Get([]byte(QueueMetaPathPrefix + fmt.Sprint(queueId)))
	if err != nil {
		log.Errorf("load queue with Id error: %v", err)
		return nil
	}

	return mdm.decodeQueue(value)
}

func (mdm *metadataManager) generateTopicId() (uint64, error) {
	mdm.tgMutex.Lock()
	defer mdm.tgMutex.Unlock()

	topicId := mdm.queueIdGenerator
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, topicId)
	err := mdm.tikvClient.Put([]byte(topicIdGeneratorPath), bs)
	if err != nil {
		return 0, err
	}
	mdm.topicIdGenerator++
	return topicId, nil
}

func (mdm *metadataManager) generateQueueId() (uint64, error) {
	mdm.qGMutex.Lock()
	defer mdm.qGMutex.Unlock()

	queueId := mdm.queueIdGenerator
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, queueId)
	err := mdm.tikvClient.Put([]byte(queueIdGeneratorPath), bs)
	if err != nil {
		return 0, err
	}

	mdm.queueIdGenerator++
	return queueId, nil
}

func (mdm *metadataManager) getTopic(topicName string) *topic {
	mdm.topicMapMutex.RLock()
	defer mdm.topicMapMutex.RUnlock()
	
	t, exist := mdm.topicMap[topicName]
	if !exist {
		log.Infof("Topic: %s doesn't exist.", topicName)
		return nil
	}
	return t
}

func (mdm *metadataManager) putTopic(t *topic) error {
	mdm.topicMapMutex.Lock()
	defer mdm.topicMapMutex.Unlock()

	_, exist := mdm.topicMap[t.topicName]
	if exist {
		return ErrTopicAlreadyExist
	}

	keys := make([][]byte, 2)
	values := make([][]byte, 2)

	keys[0] = []byte(TopicMetaPathPrefix +"allTopics")
	value, err := mdm.tikvClient.Get(keys[0])

	if err != nil {
		log.Warnf("get topics from TiKV failed, error: %v", err)
		return err
	}

	bytesOfUint64 := make([]byte, 8)
	binary.BigEndian.PutUint64(bytesOfUint64, t.topicId)
	values[0] = make([]byte, len(value) + 8)
	values[0][:len(value)] = value
	values[0][len(value):] = bytesOfUint64

	keys[1] = []byte(t.path)
	values[1] = mdm.encodeTopic(t)

	// TODO: TiKV guarantee request all success or not?
	err = mdm.tikvClient.BatchPut(keys, values)
	if err != nil {
		log.Warnf("put topic to TiKV failed, error: %v", err)
		return err
	}
	mdm.topicMap[t.topicName] = t
	return nil
}

func (mdm *metadataManager) getQueue(queueId uint64) *Queue {
	mdm.queueMapMutex.RLock()
	defer mdm.queueMapMutex.RUnlock()

	q, exist := mdm.queueMap[queueId]
	if !exist {
		log.Infof("Queue: %s doesn't exist.", queueId)
		return nil
	}
	return q
}

func (mdm *metadataManager) putQueue(q *Queue) error {
	mdm.queueMapMutex.Lock()
	defer mdm.queueMapMutex.Unlock()

	_, exist := mdm.queueMap[q.queueId]
	if exist {
		return ErrQueueAlreadyExist
	}
	return mdm.tikvClient.Put([]byte(q.path), mdm.encodeQueue(q))
}

func (mdm *metadataManager) encodeTopic(t *topic) []byte {
	qStr := ""
	for i := 0; i < len(t.queues) -1; i++  {
		qStr = fmt.Sprint("%s,%d", qStr, t.queues[i].queueId)
	}
	qStr = fmt.Sprint("%s,%d", qStr, t.queues[len(t.queues) - 1].queueId)
	str := fmt.Sprint("%d;%s;%s;%s,%v", t.topicId, t.path, t.topicName, qStr, t.autoScaling)

	return []byte(str)
}

func (mdm *metadataManager) decodeTopic(data []byte) *topic {
	t := &topic{}
	values := strings.Split(string(data), ";")
	if len(values) != 5 {
		log.Errorf("decode topic failed, origin values: %v", values)
		return nil
	}
	v, err := strconv.Atoi(values[0])
	if err != nil {
		log.Errorf("convert topicId error")
		return nil
	}
	t.topicId = uint64(v)
	t.path = values[1]
	t.topicName = values[2]
	b, err := strconv.ParseBool(values[4])
	if err != nil {
		log.Errorf("convert autoScaling error")
		return nil
	}
	t.autoScaling = b
	t.queues = make([]*Queue, 0)
	queues := strings.Split(values[3], ",")

	for i := 0; i < len(queues); i++ {
		v, err := strconv.Atoi(values[0])
		if err != nil {
			log.Errorf("convert queueId error")
			return nil
		}
		q := mdm.loadQueue(uint64(v))
		if q == nil {
			log.Errorf("queueId: %d doesn't exist.", v)
			return nil
		}
		t.queues = append(t.queues, q)
	}
	return t;
}

func (mdm *metadataManager) encodeQueue(q *Queue) []byte {
	return []byte(fmt.Sprintf("%d,%s,%v", q.queueId, q.path, pb.Permission_name[int32(q.perm)]))
}

func (mdm *metadataManager) decodeQueue(data []byte) *Queue {
	q := &Queue{}
	values := strings.Split(string(data), ",")
	if len(values) != 3 {
		log.Errorf("decode Queue error, numbers of values: %v != 3", values)
		return nil
	}
	v, err := strconv.Atoi(values[0])
	if err != nil {
		 log.Errorf("convert queueId error: %v, value: %s", err, values[0])
		 return nil
	}
	q.queueId = uint64(v)
	q.path = values[1]
	p, exist := pb.Permission_value[values[3]]
	if !exist {
		log.Errorf("perm string of %s error", values[3])
		return nil
	}
	q.perm = pb.Permission(p)
	return q;
}

type topic struct {
	topicId uint64
	path string
	topicName string
	queues []*Queue
	autoScaling bool
}

type Queue struct {
	queueId uint64

	// TopicMetaPathPrefix + /topicId
	path string
	perm pb.Permission

	mutex sync.Mutex
	// -1 is none
	assignedBroker *BrokerInfo
	//assignedConsumer uint64
}

func (q *Queue) QueueId() uint64 {
	return q.queueId;
}

func (q *Queue) Path() string {
	return q.path;
}

func (q *Queue) Permission() pb.Permission {
	return q.perm;
}

func (q *Queue) assignBroker(info *BrokerInfo) bool {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if q.assignedBroker != nil {
		return false
	}

	q.assignedBroker = info
	return true
}

func (q *Queue) MinimumOffset() uint64 {
	return 0;
}

func (q *Queue) MaximumOffset() uint64 {
	return 0
}

type BrokerInfo struct {
	brokerId uint64;
	address string;
}
