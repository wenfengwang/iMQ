package storage

import (
	"encoding/binary"
	"github.com/pingcap/tidb/store/tikv"
	"sync/atomic"
)

type tikvDFile struct {
	client      *tikv.RawKVClient
	size        uint64
	id          uint64
	writeOffset uint64
	readOffset  uint64
}

func (tikv *tikvDFile) Size() uint64 {
	return tikv.size
}

func (tikv *tikvDFile) FileID() uint64 {
	return tikv.id
}

func (tikv *tikvDFile) WriteOffset() uint64 {
	return tikv.writeOffset
}

func (tikv *tikvDFile) ReadOffset() uint64 {
	return tikv.readOffset
}

func (tikv *tikvDFile) Write(bytes []byte) error {
	offset := atomic.AddUint64(&tikv.writeOffset, 1)
	u64bytes := make([]byte, 8)
	binary.BigEndian.PutUint64(u64bytes, offset)
	return tikv.client.Put(bytes, u64bytes)
}

func (tikv *tikvDFile) Read(offset uint64, length int32) ([]byte, error) {
	//tikv.client.
	return nil, nil
}

type DistributedFile interface {
	Size() uint64
	FileID() uint64
	WriteOffset() uint64
	ReadOffset() uint64
	Write(bytes []byte) error
	Read(offset uint64, length int32) ([]byte, error)
}
