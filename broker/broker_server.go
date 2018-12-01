package broker

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/pingcap/tidb/config"
	"github.com/pingcap/tidb/store/tikv"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/wenfengwang/iMQ/baton/pb"
	pb "github.com/wenfengwang/iMQ/broker/pb"
	"google.golang.org/grpc"
	"io"
)

var (
	ErrRouteNotFound = errors.New("ErrRouteNotFound")
)

type BrokerConfig struct {
	PD           []string
	Address      string
	BatonAddress string
}

func NewPubSubServer(cfg BrokerConfig) pb.PubSubServer {
	tikvClient, err := tikv.NewRawKVClient(cfg.PD, config.Security{})
	if err != nil {
		panic(fmt.Sprintf("create TiKV client error: %v", err))
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial(cfg.BatonAddress, opts...)
	if err != nil {
		panic(fmt.Sprintf("Dial to Baton: %s error: %v", cfg.Address, err))
	}

	return &broker{qm: &QueueManager{tikvClient: tikvClient}, cfg: cfg, tikvClient: tikvClient,
		batonClient: batonpb.NewBatonClient(conn)}
}

type broker struct {
	brokerId    uint64
	qm          *QueueManager
	cfg         BrokerConfig
	tikvClient  *tikv.RawKVClient
	batonClient batonpb.BatonClient
}

func (b *broker) register() {
	response, err := b.batonClient.RegisterBroker(context.Background(),
		&batonpb.RegisterBrokerRequest{Addr: b.cfg.Address})
	if err != nil {
		panic(fmt.Sprintf("start broker error: %v", err))
	}
	b.brokerId = response.Id
}

func (b *broker) Publish(ctx context.Context, request *pb.PublishRequest) (*pb.PublishResponse, error) {
	var q Queue
	values := make([][]byte, len(request.Msg))
	for index, msg := range request.Msg {
		if q == nil {
			q = b.qm.getQueue(msg.QueueId)
		}
		//if q == nil {
		//	return
		//}
		log.Debugf("received message: %s", msg.Body)
		values[index], _ = proto.Marshal(msg)
	}

	err := q.Put(values)
	return &pb.PublishResponse{}, err
}

func (b *broker) Subscribe(request *pb.SubscribeRequest, stream pb.PubSub_SubscribeServer) error {
	q := b.qm.getQueue(request.QueueId)
	if q == nil {
		return ErrRouteNotFound
	}

	ch := make(chan error)
	go func() {
		for {
			values, _ := q.Get(16)
			for _, v := range values {
				msg := &pb.Message{}
				proto.Unmarshal(v, msg)
				stream.Send(msg)
			}
		}
	}()
	return <-ch
}

func (b *broker) PullMessage(stream pb.PubSub_PullMessageServer) error {
	log.Info("pull message...")
	for {
		request, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if request == nil {
			 return nil
		}
		q := b.qm.getQueue(request.QueueId)

		values, _ :=  q.Get(int(request.Numbers))
		msgs := make([]*pb.Message, request.Numbers)
		for index, v := range values {
			msg := &pb.Message{}
			proto.Unmarshal(v, msg)
			msgs[index] = msg
		}
		stream.Send(&pb.PullMessageResponse{Msg: msgs})
	}

	//stream.
	return nil
}
