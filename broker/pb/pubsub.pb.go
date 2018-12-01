// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pubsub.proto

package brokerpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PublishRequest struct {
	Msg                  []*Message `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PublishRequest) Reset()         { *m = PublishRequest{} }
func (m *PublishRequest) String() string { return proto.CompactTextString(m) }
func (*PublishRequest) ProtoMessage()    {}
func (*PublishRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{0}
}

func (m *PublishRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishRequest.Unmarshal(m, b)
}
func (m *PublishRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishRequest.Marshal(b, m, deterministic)
}
func (m *PublishRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishRequest.Merge(m, src)
}
func (m *PublishRequest) XXX_Size() int {
	return xxx_messageInfo_PublishRequest.Size(m)
}
func (m *PublishRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PublishRequest proto.InternalMessageInfo

func (m *PublishRequest) GetMsg() []*Message {
	if m != nil {
		return m.Msg
	}
	return nil
}

type PublishResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PublishResponse) Reset()         { *m = PublishResponse{} }
func (m *PublishResponse) String() string { return proto.CompactTextString(m) }
func (*PublishResponse) ProtoMessage()    {}
func (*PublishResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{1}
}

func (m *PublishResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublishResponse.Unmarshal(m, b)
}
func (m *PublishResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublishResponse.Marshal(b, m, deterministic)
}
func (m *PublishResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublishResponse.Merge(m, src)
}
func (m *PublishResponse) XXX_Size() int {
	return xxx_messageInfo_PublishResponse.Size(m)
}
func (m *PublishResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PublishResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PublishResponse proto.InternalMessageInfo

type SubscribeRequest struct {
	QueueId              uint64   `protobuf:"varint,1,opt,name=queueId,proto3" json:"queueId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SubscribeRequest) Reset()         { *m = SubscribeRequest{} }
func (m *SubscribeRequest) String() string { return proto.CompactTextString(m) }
func (*SubscribeRequest) ProtoMessage()    {}
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{2}
}

func (m *SubscribeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscribeRequest.Unmarshal(m, b)
}
func (m *SubscribeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscribeRequest.Marshal(b, m, deterministic)
}
func (m *SubscribeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscribeRequest.Merge(m, src)
}
func (m *SubscribeRequest) XXX_Size() int {
	return xxx_messageInfo_SubscribeRequest.Size(m)
}
func (m *SubscribeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscribeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SubscribeRequest proto.InternalMessageInfo

func (m *SubscribeRequest) GetQueueId() uint64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

type PullMessageRequest struct {
	QueueId              uint64   `protobuf:"varint,1,opt,name=queueId,proto3" json:"queueId,omitempty"`
	Numbers              int32    `protobuf:"varint,2,opt,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PullMessageRequest) Reset()         { *m = PullMessageRequest{} }
func (m *PullMessageRequest) String() string { return proto.CompactTextString(m) }
func (*PullMessageRequest) ProtoMessage()    {}
func (*PullMessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{3}
}

func (m *PullMessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMessageRequest.Unmarshal(m, b)
}
func (m *PullMessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMessageRequest.Marshal(b, m, deterministic)
}
func (m *PullMessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMessageRequest.Merge(m, src)
}
func (m *PullMessageRequest) XXX_Size() int {
	return xxx_messageInfo_PullMessageRequest.Size(m)
}
func (m *PullMessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PullMessageRequest proto.InternalMessageInfo

func (m *PullMessageRequest) GetQueueId() uint64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *PullMessageRequest) GetNumbers() int32 {
	if m != nil {
		return m.Numbers
	}
	return 0
}

type PullMessageResponse struct {
	Msg                  []*Message `protobuf:"bytes,1,rep,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *PullMessageResponse) Reset()         { *m = PullMessageResponse{} }
func (m *PullMessageResponse) String() string { return proto.CompactTextString(m) }
func (*PullMessageResponse) ProtoMessage()    {}
func (*PullMessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_91df006b05e20cf7, []int{4}
}

func (m *PullMessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PullMessageResponse.Unmarshal(m, b)
}
func (m *PullMessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PullMessageResponse.Marshal(b, m, deterministic)
}
func (m *PullMessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PullMessageResponse.Merge(m, src)
}
func (m *PullMessageResponse) XXX_Size() int {
	return xxx_messageInfo_PullMessageResponse.Size(m)
}
func (m *PullMessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PullMessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PullMessageResponse proto.InternalMessageInfo

func (m *PullMessageResponse) GetMsg() []*Message {
	if m != nil {
		return m.Msg
	}
	return nil
}

func init() {
	proto.RegisterType((*PublishRequest)(nil), "brokerpb.PublishRequest")
	proto.RegisterType((*PublishResponse)(nil), "brokerpb.PublishResponse")
	proto.RegisterType((*SubscribeRequest)(nil), "brokerpb.SubscribeRequest")
	proto.RegisterType((*PullMessageRequest)(nil), "brokerpb.PullMessageRequest")
	proto.RegisterType((*PullMessageResponse)(nil), "brokerpb.PullMessageResponse")
}

func init() { proto.RegisterFile("pubsub.proto", fileDescriptor_91df006b05e20cf7) }

var fileDescriptor_91df006b05e20cf7 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x51, 0xcb, 0x4e, 0xc3, 0x30,
	0x10, 0x94, 0x29, 0xb4, 0x74, 0x8b, 0x80, 0x9a, 0x8b, 0xb1, 0x40, 0x8a, 0xcc, 0x25, 0x07, 0x14,
	0x55, 0x45, 0x5c, 0x10, 0xe2, 0x0c, 0x07, 0x50, 0x94, 0x7e, 0x01, 0x0b, 0xab, 0x50, 0x91, 0x34,
	0xa9, 0xb7, 0xfb, 0xab, 0x7c, 0x0f, 0x2a, 0xa9, 0xfb, 0x02, 0x04, 0xc7, 0xf5, 0xcc, 0xce, 0xce,
	0x8c, 0xe1, 0xa0, 0x16, 0x64, 0xc1, 0xa4, 0xf6, 0xd5, 0xac, 0xd2, 0xfb, 0xe8, 0xab, 0x77, 0xf2,
	0x35, 0xda, 0x6e, 0xc9, 0x79, 0xf3, 0xe8, 0xae, 0xe1, 0x30, 0x15, 0x2c, 0xc6, 0xfc, 0x96, 0xd1,
	0x54, 0x88, 0x67, 0xfa, 0x02, 0x5a, 0x25, 0xe7, 0x46, 0x45, 0xad, 0xb8, 0x37, 0xec, 0x27, 0x61,
	0x29, 0x79, 0x24, 0xe6, 0xe7, 0x9c, 0xb2, 0x39, 0xea, 0xfa, 0x70, 0xb4, 0x5c, 0xe3, 0xba, 0x9a,
	0x30, 0xb9, 0x4b, 0x38, 0x1e, 0x09, 0xf2, 0x8b, 0x1f, 0x23, 0x05, 0x2d, 0x03, 0x9d, 0xa9, 0x90,
	0xd0, 0xc3, 0xab, 0x51, 0x91, 0x8a, 0x77, 0xb3, 0x30, 0xba, 0x7b, 0xd0, 0xa9, 0x14, 0x45, 0x10,
	0xfd, 0x8b, 0x3f, 0x47, 0x26, 0x52, 0x22, 0x79, 0x36, 0x3b, 0x91, 0x8a, 0xf7, 0xb2, 0x30, 0xba,
	0x1b, 0x38, 0xd9, 0x50, 0x6a, 0xec, 0xfc, 0x2b, 0xc6, 0xf0, 0x43, 0x41, 0x3b, 0x15, 0x1c, 0x09,
	0xea, 0x3b, 0xe8, 0x2c, 0x12, 0x69, 0xb3, 0x62, 0x6f, 0x76, 0x63, 0x4f, 0x7f, 0x40, 0x16, 0xf7,
	0x6e, 0xa1, 0xbb, 0x8c, 0xaf, 0xed, 0x8a, 0xb7, 0xdd, 0x89, 0xfd, 0xee, 0x65, 0xa0, 0xf4, 0x13,
	0xf4, 0xd6, 0x42, 0xe8, 0xb3, 0xf5, 0x3b, 0xdb, 0x2d, 0xd9, 0xf3, 0x5f, 0xd0, 0xc6, 0x49, 0xac,
	0x06, 0x0a, 0xdb, 0x5f, 0xbf, 0x7b, 0xf5, 0x19, 0x00, 0x00, 0xff, 0xff, 0x34, 0x9f, 0xfd, 0x19,
	0x02, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubSubClient is the client API for PubSub service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubSubClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSub_SubscribeClient, error)
	PullMessage(ctx context.Context, opts ...grpc.CallOption) (PubSub_PullMessageClient, error)
}

type pubSubClient struct {
	cc *grpc.ClientConn
}

func NewPubSubClient(cc *grpc.ClientConn) PubSubClient {
	return &pubSubClient{cc}
}

func (c *pubSubClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishResponse, error) {
	out := new(PublishResponse)
	err := c.cc.Invoke(ctx, "/brokerpb.PubSub/Publish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubSubClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (PubSub_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubSub_serviceDesc.Streams[0], "/brokerpb.PubSub/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubSub_SubscribeClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type pubSubSubscribeClient struct {
	grpc.ClientStream
}

func (x *pubSubSubscribeClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubSubClient) PullMessage(ctx context.Context, opts ...grpc.CallOption) (PubSub_PullMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubSub_serviceDesc.Streams[1], "/brokerpb.PubSub/PullMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubSubPullMessageClient{stream}
	return x, nil
}

type PubSub_PullMessageClient interface {
	Send(*PullMessageRequest) error
	Recv() (*PullMessageResponse, error)
	grpc.ClientStream
}

type pubSubPullMessageClient struct {
	grpc.ClientStream
}

func (x *pubSubPullMessageClient) Send(m *PullMessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pubSubPullMessageClient) Recv() (*PullMessageResponse, error) {
	m := new(PullMessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PubSubServer is the server API for PubSub service.
type PubSubServer interface {
	Publish(context.Context, *PublishRequest) (*PublishResponse, error)
	Subscribe(*SubscribeRequest, PubSub_SubscribeServer) error
	PullMessage(PubSub_PullMessageServer) error
}

func RegisterPubSubServer(s *grpc.Server, srv PubSubServer) {
	s.RegisterService(&_PubSub_serviceDesc, srv)
}

func _PubSub_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubSubServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/brokerpb.PubSub/Publish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubSubServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubSub_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubSubServer).Subscribe(m, &pubSubSubscribeServer{stream})
}

type PubSub_SubscribeServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type pubSubSubscribeServer struct {
	grpc.ServerStream
}

func (x *pubSubSubscribeServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _PubSub_PullMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PubSubServer).PullMessage(&pubSubPullMessageServer{stream})
}

type PubSub_PullMessageServer interface {
	Send(*PullMessageResponse) error
	Recv() (*PullMessageRequest, error)
	grpc.ServerStream
}

type pubSubPullMessageServer struct {
	grpc.ServerStream
}

func (x *pubSubPullMessageServer) Send(m *PullMessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pubSubPullMessageServer) Recv() (*PullMessageRequest, error) {
	m := new(PullMessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PubSub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "brokerpb.PubSub",
	HandlerType: (*PubSubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _PubSub_Publish_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _PubSub_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "PullMessage",
			Handler:       _PubSub_PullMessage_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pubsub.proto",
}