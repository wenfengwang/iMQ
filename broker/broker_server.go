package broker

import (
	"github.com/golang/protobuf/proto"
	"github.com/pingcap/tidb/store/tikv"
	pb "github.com/wenfengwang/iMQ/broker/pb"
	"io"
	"sync"
)

func NewPubSubServer() pb.PubSubServer {
	return &broker{}
}

type broker struct {
	queueMap   sync.Map
	tikvClient *tikv.RawKVClient
}

func (b *broker) Publish(stream pb.PubSub_PublishServer) error {
	ch := make(chan error)
	var q Queue
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			ch <- nil
			break
		}
		values := make([][]byte, len(req.Msg))
		for index, msg := range req.Msg {
			if q == nil {
				v, exist := b.queueMap.Load(msg.QueueId)
				if exist {
					q = v.(Queue)
				} else {
					q = NewQueue(msg.QueueId, b.tikvClient)
				}
			}
			values[index], _ = proto.Marshal(msg)
		}
		q.Put(values)
	}

	return <-ch
}

func (b *broker) Subscribe(request *pb.SubscribeRequest, stream pb.PubSub_SubscribeServer) error {

	//stream.
	return nil
}

func (b *broker) PullMessage(stream pb.PubSub_PullMessageServer) error {

	//stream.
	return nil
}
