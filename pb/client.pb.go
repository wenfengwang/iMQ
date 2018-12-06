// Code generated by protoc-gen-go. DO NOT EDIT.
// source: client.proto

package pb

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

type PublishResult int32

const (
	PublishResult_SUCCESS        PublishResult = 0
	PublishResult_FAILED         PublishResult = 1
	PublishResult_NO_QUEUE_ROUTE PublishResult = 2
)

var PublishResult_name = map[int32]string{
	0: "SUCCESS",
	1: "FAILED",
	2: "NO_QUEUE_ROUTE",
}

var PublishResult_value = map[string]int32{
	"SUCCESS":        0,
	"FAILED":         1,
	"NO_QUEUE_ROUTE": 2,
}

func (x PublishResult) String() string {
	return proto.EnumName(PublishResult_name, int32(x))
}

func (PublishResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{0}
}

type ConsumeResult int32

const (
	ConsumeResult_OK    ConsumeResult = 0
	ConsumeResult_RETRY ConsumeResult = 1
)

var ConsumeResult_name = map[int32]string{
	0: "OK",
	1: "RETRY",
}

var ConsumeResult_value = map[string]int32{
	"OK":    0,
	"RETRY": 1,
}

func (x ConsumeResult) String() string {
	return proto.EnumName(ConsumeResult_name, int32(x))
}

func (ConsumeResult) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_014de31d7ac8c57c, []int{1}
}

func init() {
	proto.RegisterEnum("pb.PublishResult", PublishResult_name, PublishResult_value)
	proto.RegisterEnum("pb.ConsumeResult", ConsumeResult_name, ConsumeResult_value)
}

func init() { proto.RegisterFile("client.proto", fileDescriptor_014de31d7ac8c57c) }

var fileDescriptor_014de31d7ac8c57c = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0xd2, 0xb2, 0xe1, 0xe2,
	0x0d, 0x28, 0x4d, 0xca, 0xc9, 0x2c, 0xce, 0x08, 0x4a, 0x2d, 0x2e, 0xcd, 0x29, 0x11, 0xe2, 0xe6,
	0x62, 0x0f, 0x0e, 0x75, 0x76, 0x76, 0x0d, 0x0e, 0x16, 0x60, 0x10, 0xe2, 0xe2, 0x62, 0x73, 0x73,
	0xf4, 0xf4, 0x71, 0x75, 0x11, 0x60, 0x14, 0x12, 0xe2, 0xe2, 0xf3, 0xf3, 0x8f, 0x0f, 0x0c, 0x75,
	0x0d, 0x75, 0x8d, 0x0f, 0xf2, 0x0f, 0x0d, 0x71, 0x15, 0x60, 0xd2, 0x52, 0xe2, 0xe2, 0x75, 0xce,
	0xcf, 0x2b, 0x2e, 0xcd, 0x4d, 0x85, 0xea, 0x66, 0xe3, 0x62, 0xf2, 0xf7, 0x16, 0x60, 0x10, 0xe2,
	0xe4, 0x62, 0x0d, 0x72, 0x0d, 0x09, 0x8a, 0x14, 0x60, 0x4c, 0x62, 0x03, 0x5b, 0x66, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x9d, 0xcf, 0x49, 0x56, 0x7c, 0x00, 0x00, 0x00,
}