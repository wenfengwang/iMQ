package baton

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/wenfengwang/iMQ/pb"
)

func TestTopic_Encode(t *testing.T) {
	topic := &topic{
		topicId: 12345678910,
		path:        TopicMetaPathPrefix + fmt.Sprint(12345678910),
		topicName:   "unitTestTopic",
		autoScaling: true,
		status: &topicStatus{},
	}
	topic.queues = make([]*Queue, 4)
	for i := 0; i < 4 ; i++ {
		topic.queues[i] = &Queue{queueId: uint64(10e9 + i)}
	}

	expect := "12345678910;/imq/metadata/topic/12345678910;unitTestTopic;" +
		"10000000000,10000000001,10000000002,10000000003;true"
	assert.Equal(t, expect, string(encodeTopic(topic)))
}

func TestTopic_Decode(t *testing.T) {
	value := "12345678910;/imq/metadata/topic/12345678910;unitTestTopic;" +
		"10000000000,10000000001,10000000002,10000000003;true"
	topic := decodeTopic([]byte(value))
	assert.NotNil(t, topic)
	assert.Equal(t, topic.topicId, uint64(12345678910))
	assert.Equal(t, topic.path, "/imq/metadata/topic/12345678910")
	assert.Equal(t, topic.topicName, "unitTestTopic")
	assert.Equal(t, topic.autoScaling, true)
	assert.Equal(t, 4, len(topic.queues))
	assert.Equal(t, uint64(10000000000), topic.queues[0].queueId)
	assert.Equal(t, uint64(10000000001), topic.queues[1].queueId)
	assert.Equal(t, uint64(10000000002), topic.queues[2].queueId)
	assert.Equal(t, uint64(10000000003), topic.queues[3].queueId)
	assert.NotNil(t, topic.status)
	assert.NotNil(t, topic.status.producers)
	assert.NotNil(t, topic.status.consumers)
}

func TestQueue_Encode(t *testing.T) {
	q := &Queue{
		queueId: 123450987654,
		path:QueueMetaPathPrefix + fmt.Sprint(123450987654),
		perm: pb.Permission_READ_WRITE,
		status: &QueueStatus{},
	}

	expect := "123450987654,/imq/metadata/queue/123450987654,READ_WRITE"
	assert.Equal(t, expect, string(encodeQueue(q)))
}


func TestQueue_Decode(t *testing.T) {
	value := []byte("123450987654,/imq/metadata/queue/123450987654,READ_WRITE")
	q := decodeQueue(value)
	assert.NotNil(t, q)
	assert.Equal(t, q.queueId, uint64(123450987654))
	assert.Equal(t, q.path, "/imq/metadata/queue/123450987654")
	assert.Equal(t, q.perm, pb.Permission_READ_WRITE)
	assert.NotNil(t, q.status)
	assert.Equal(t, uint64(0), q.status.curPID)
	assert.Equal(t, uint64(0), q.status.nextPID)
	assert.Equal(t, int64(0), q.status.pExpired)
	assert.Equal(t, uint64(0), q.status.curCID)
	assert.Equal(t, uint64(0), q.status.nextCID)
	assert.Equal(t, int64(0), q.status.cExpired)
	assert.Nil(t, q.status.brokerForPub)
	assert.Nil(t, q.status.brokerForSub)
}
