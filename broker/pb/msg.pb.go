// Code generated by protoc-gen-go. DO NOT EDIT.
// source: msg.proto

package brokerpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Message struct {
	MessageId            int64             `protobuf:"varint,1,opt,name=messageId,proto3" json:"messageId,omitempty"`
	Body                 string            `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
	QueueId              uint64            `protobuf:"varint,3,opt,name=queueId,proto3" json:"queueId,omitempty"`
	Properties           map[string]string `protobuf:"bytes,4,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06e4cca6c2cc899, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetMessageId() int64 {
	if m != nil {
		return m.MessageId
	}
	return 0
}

func (m *Message) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Message) GetQueueId() uint64 {
	if m != nil {
		return m.QueueId
	}
	return 0
}

func (m *Message) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "brokerpb.Message")
	proto.RegisterMapType((map[string]string)(nil), "brokerpb.Message.PropertiesEntry")
}

func init() { proto.RegisterFile("msg.proto", fileDescriptor_c06e4cca6c2cc899) }

var fileDescriptor_c06e4cca6c2cc899 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x2d, 0x4e, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x48, 0x2a, 0xca, 0xcf, 0x4e, 0x2d, 0x2a, 0x48, 0x52,
	0xba, 0xce, 0xc8, 0xc5, 0xee, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0x24, 0xc3, 0xc5, 0x99,
	0x0b, 0x61, 0x7a, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x21, 0x04, 0x84, 0x84, 0xb8,
	0x58, 0x92, 0xf2, 0x53, 0x2a, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0x6c, 0x21, 0x09,
	0x2e, 0xf6, 0xc2, 0xd2, 0xd4, 0x52, 0x90, 0x7a, 0x66, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x18, 0x57,
	0xc8, 0x91, 0x8b, 0xab, 0xa0, 0x28, 0xbf, 0x20, 0xb5, 0xa8, 0x24, 0x33, 0xb5, 0x58, 0x82, 0x45,
	0x81, 0x59, 0x83, 0xdb, 0x48, 0x51, 0x0f, 0x66, 0xad, 0x1e, 0xd4, 0x4a, 0xbd, 0x00, 0xb8, 0x1a,
	0xd7, 0xbc, 0x92, 0xa2, 0xca, 0x20, 0x24, 0x4d, 0x52, 0xb6, 0x5c, 0xfc, 0x68, 0xd2, 0x42, 0x02,
	0x5c, 0xcc, 0xd9, 0xa9, 0x95, 0x60, 0xb7, 0x71, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x65,
	0x89, 0x39, 0xa5, 0xa9, 0x50, 0x67, 0x41, 0x38, 0x56, 0x4c, 0x16, 0x8c, 0x49, 0x6c, 0x60, 0xaf,
	0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x12, 0xd2, 0xc0, 0xf7, 0x00, 0x00, 0x00,
}
