package baton

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestTopic_Encode(t *testing.T) {
	topic := &topic{
		topicId: 12345678910,
		path:        TopicMetaPathPrefix + fmt.Sprint(12345678910),
		topicName:   "unitTestTopic",
		autoScaling: true,
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
	assert.Equal(t, topic.topicId, uint64(12345678910))
	assert.Equal(t, topic.path, "/imq/metadata/topic/12345678910")
	assert.Equal(t, topic.topicName, "unitTestTopic")
	assert.Equal(t, topic.autoScaling, true)
	assert.Equal(t, 4, len(topic.queues))
	assert.Equal(t, uint64(10000000000), topic.queues[0].queueId)
	assert.Equal(t, uint64(10000000001), topic.queues[1].queueId)
	assert.Equal(t, uint64(10000000002), topic.queues[2].queueId)
	assert.Equal(t, uint64(10000000003), topic.queues[3].queueId)
}

func TestQueue_Encode(t *testing.T) {
	
}


func TestQueue_Decode(t *testing.T) {

}
