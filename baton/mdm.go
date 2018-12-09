package baton

import (
	"encoding/binary"
	"fmt"
	"github.com/pingcap/tidb/store/tikv"
	"github.com/pkg/errors"
	"github.com/prometheus/common/log"
	"github.com/wenfengwang/iMQ/pb"
	"strconv"
	"strings"
	"sync"
)

const (
	TopicMetaPathPrefix  = "/imq/metadata/topic/"
	QueueMetaPathPrefix  = "/imq/metadata/queue/"

	allTopicsPath = TopicMetaPathPrefix + "allTopics"
	topicIdGeneratorPath = "/imq/metadata/topic_generator"
	queueIdGeneratorPath = "/imq/metadata/queue_generator"
)

var (
	allTopicsPathKey = []byte(allTopicsPath)
	topicIdGeneratorPathKey = []byte(topicIdGeneratorPath)
	queueIdGeneratorPathKey = []byte(queueIdGeneratorPath)
)

var (
	ErrTopicAlreadyExist = errors.New("topic already exist.")
	ErrQueueAlreadyExist = errors.New("queue already exist.")
)

type metadataManager struct {
	tikvClient *tikv.RawKVClient

	topicMapMutex sync.RWMutex
	topicMap      map[uint64]*topic

	queueMapMutex sync.RWMutex
	queueMap      map[uint64]*Queue

	tgMutex          sync.Mutex
	topicIdGenerator uint64

	qGMutex          sync.Mutex
	queueIdGenerator uint64
}

func (mdm *metadataManager) init() {
	value, err := mdm.tikvClient.Get(topicIdGeneratorPathKey)
	if err != nil {
		panic(fmt.Sprintf("error: %v, get topicIdGenerator FAILED !!!", err))
	}

	if len(value) != 8 {
		panic(fmt.Sprintf("topicIdGenerator value invalid!!! it's length should be 8, but is: %v", value))
	}
	mdm.topicIdGenerator = binary.BigEndian.Uint64(value)

	value, err = mdm.tikvClient.Get(queueIdGeneratorPathKey)
	if err != nil {
		panic(fmt.Sprintf("error: %v, get queueIdGenerator FAILED !!!", err))
	}

	if len(value) != 8 {
		panic(fmt.Sprintf("queueIdGeneratorPathKey value invalid!!! it's length should be 8, but is: %v", value))
	}
	mdm.queueIdGenerator = binary.BigEndian.Uint64(value)
	mdm.loadTopicsAndQueues()
}

func (mdm *metadataManager) loadTopicsAndQueues() {
	mdm.topicMap = make(map[uint64]*topic)
	mdm.queueMap = make(map[uint64]*Queue)
	value, err := mdm.tikvClient.Get(allTopicsPathKey)
	if err != nil {
		panic(fmt.Sprintf("error: %v, load topics FAILED !!!", err))
	}

	for i := 0; i < len(value); i += 8 {
		// load topic
		topicId := binary.BigEndian.Uint64(value[i:i+8])
		t := mdm.loadTopic(topicId)
		if t == nil {
			log.Errorf("load topicId: %v but it doesn't exist.", topicId)
			continue
		}

		// load queue
		qs := make([]*Queue, 0)
		for j := 0; j < len(t.queues); j++ {
			q := mdm.loadQueue(t.queues[i].queueId)
			if q == nil {
				log.Errorf("load queue:[%d] but got nil", t.queues[i].queueId)
				continue
			}
			qs = append(qs, q)
			mdm.queueMap[q.queueId] = q
		}
		t.queues = qs
		mdm.topicMap[t.topicId] = t
	}
}

func (mdm *metadataManager) loadTopic(topicId uint64) *topic {
	value, err := mdm.tikvClient.Get(keyJoin(TopicMetaPathPrefix, fmt.Sprint(topicId)))
	if err != nil {
		log.Errorf("load topic with Id: %v error: %v", topicId, err)
		return nil
	}

	return decodeTopic(value)
}

func (mdm *metadataManager) loadQueue(queueId uint64) *Queue {
	value, err := mdm.tikvClient.Get([]byte(QueueMetaPathPrefix + fmt.Sprint(queueId)))
	if err != nil {
		log.Errorf("load queue with Id error: %v", err)
		return nil
	}

	return decodeQueue(value)
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

	t, _ := mdm.topicMap[1] // TODO how to transform topicName to topicId
	if t == nil {
		log.Info(fmt.Sprintf("Topic: %s doesn't exist.", topicName))
		return nil
	}
	return t
}

func (mdm *metadataManager) putTopic(t *topic) error {
	mdm.topicMapMutex.Lock()
	defer mdm.topicMapMutex.Unlock()

	_, exist := mdm.topicMap[t.topicId]
	if exist {
		return ErrTopicAlreadyExist
	}

	keys := make([][]byte, 2)
	values := make([][]byte, 2)
	keys[0] = allTopicsPathKey
	value, err := mdm.tikvClient.Get(keys[0])

	if err != nil {
		log.Warnf("get topics from TiKV failed, error: %v", err)
		return err
	}

	bytesOfUint64 := make([]byte, 8)
	binary.BigEndian.PutUint64(bytesOfUint64, t.topicId)
	values[0] = make([]byte, len(value)+8)
	copy(values[0][:len(value)], value)
	copy(values[0][len(value):], bytesOfUint64)

	keys[1] = []byte(t.path)
	values[1] = encodeTopic(t)

	// TODO: TiKV guarantee request all success or not?
	err = mdm.tikvClient.BatchPut(keys, values)
	if err != nil {
		log.Warnf("put topic to TiKV failed, error: %v", err)
		return err
	}
	mdm.topicMap[t.topicId] = t
	log.Info(fmt.Sprintf("topic: %s insert to map", t.topicName))
	for k, _ := range  mdm.topicMap{
		log.Info(fmt.Printf("topic: %d", k))
	}
	return nil
}

func (mdm *metadataManager) getQueue(queueId uint64) *Queue {
	mdm.queueMapMutex.RLock()
	defer mdm.queueMapMutex.RUnlock()

	q, exist := mdm.queueMap[queueId]
	if !exist {
		log.Infof("QueueId: %d doesn't exist.", queueId)
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
	mdm.queueMap[q.queueId] = q
	log.Info(fmt.Sprintf("add queue: %d", q.queueId))
	return mdm.tikvClient.Put([]byte(q.path), encodeQueue(q))
}

func encodeTopic(t *topic) []byte {
	qStr := ""
	for i := 0; i < len(t.queues)-1; i++ {
		qStr += fmt.Sprint(t.queues[i].queueId) + ","
	}
	qStr += fmt.Sprint(t.queues[len(t.queues)-1].queueId)
	str := fmt.Sprintf("%d;%s;%s;%s;%v", t.topicId, t.path, t.topicName, qStr, t.autoScaling)

	return []byte(str)
}

func decodeTopic(data []byte) *topic {
	if data == nil || len(data) == 0{
		return nil
	}
	t := &topic{}
	t.status = &topicStatus{producers: make(map[uint64]*ProducerInfo), consumers: make(map[uint64]*ConsumerInfo)}
	values := strings.Split(string(data), ";")
	if len(values) != 5 {
		log.Errorf("decode topic failed, expect length is 5, origin values: %v", values)
		return nil
	}

	// topicId
	v, err := strconv.Atoi(values[0])
	if err != nil {
		log.Errorf("convert topicId error, value:%s", values[0])
		return nil
	}
	t.topicId = uint64(v)
	// topic store path
	t.path = values[1]
	// name of topic
	t.topicName = values[2]

	// queues of topic
	queues := strings.Split(values[3], ",")
	t.queues = make([]*Queue, len(queues))
	for i := 0; i < len(queues); i++ {
		v, err := strconv.Atoi(queues[i])
		if err != nil {
			log.Errorf("convert queueId error, value:%s", queues[i])
			return nil
		}
		t.queues[i] = &Queue{queueId: uint64(v)}
	}

	b, err := strconv.ParseBool(values[4])
	if err != nil {
		log.Errorf("convert autoScaling error")
		return nil
	}
	t.autoScaling = b
	return t
}

func encodeQueue(q *Queue) []byte {
	return []byte(fmt.Sprintf("%d,%s,%v", q.queueId, q.path, pb.Permission_name[int32(q.perm)]))
}

func decodeQueue(data []byte) *Queue {
	q := &Queue{}
	q.status = &QueueStatus{}
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
	return q
}

// TODO
func (mdm *metadataManager) scale(t *topic) {

}
