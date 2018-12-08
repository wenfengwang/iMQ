package client

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/wenfengwang/iMQ/pb"
	"google.golang.org/grpc"
	"sync"
	"sync/atomic"
)

type BrokerHub struct {
	brokerMap sync.Map
}

func (bh *BrokerHub) getBroker(info *pb.BrokerInfo) *pubsub {
	b, exist := bh.brokerMap.Load(info.BrokerId)
	if !exist {
		pb := newPubSub(info.Address, info.BrokerId)
		var stored bool
		b, stored = bh.brokerMap.LoadOrStore(info.BrokerId, pb)
		if !stored {
			pb.start()
		}
	}
	return b.(*pubsub)
}

type pubsub struct {
	address         string
	brokerId        uint64
	client          pb.PubSubClient
	pullMsgClient   pb.PubSub_PullMessageClient
	pullExitCh      chan interface{}
	requestIdGen    uint64
	pullResultChMap sync.Map
	mutex sync.Mutex
}

func newPubSub(address string, brokerId uint64) *pubsub {
	return &pubsub{address: address, brokerId: brokerId, pullExitCh: make(chan interface{}), requestIdGen: 0}
}

func (b *pubsub) start() error {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	conn, err := grpc.Dial("localhost:23456", opts...)
	if err != nil {
		return errors.New(fmt.Sprintf("dial broker: %s error %v", b.address, err))
	}

	log.Info("Dial to %s success\n", b.address)
	b.client = pb.NewPubSubClient(conn)
	b.pullMsgClient, _ = b.client.PullMessage(context.Background())
	for i :=0; i < 4 ; i++  {
		go func() {
			for {
				select {
				case <-b.pullExitCh:
					break
				default:
					res, _ := b.pullMsgClient.Recv()
					if res == nil {
						break
					}
					ch, exist := b.pullResultChMap.Load(res.ResponseId)

					if !exist {
						continue
					}
					b.pullResultChMap.Delete(res.ResponseId)
					ch.(chan *pb.PullMessageResponse) <- res
				}
			}
		}()
	}


	return nil
}

func (b *pubsub) shutdown() error {
	b.pullExitCh <- "exit"
	return nil
}

func (b *pubsub) publish(request *pb.PublishRequest) (*pb.PublishResponse, error) {
	return b.client.Publish(context.Background(), request)
}

func (b *pubsub) subscribe(request *pb.SubscribeRequest) (pb.PubSub_SubscribeClient, error) {
	return b.client.Subscribe(context.Background(), request)
}

func (b *pubsub) pullMessage(request *pb.PullMessageRequest, ch chan *pb.PullMessageResponse) error {
	atomic.AddUint64(&b.requestIdGen, 1)
	b.pullMsgClient.Send(request)
	b.pullResultChMap.Store(request.RequestId, ch)
	return nil
}
