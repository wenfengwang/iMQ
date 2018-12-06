package storage

import "github.com/pingcap/tidb/store/tikv"

type StoreEngine int

const (
	TiKV StoreEngine = 0
)

// Awesome Distributed Storage
type ADS interface {
	Put(key, value []byte) error
	Get(key []byte) ([]byte, error)
	Delete(key []byte) ([]byte, error)

	BatchPut(keys, values [][]byte) error
	BatchGet(keys [][]byte) ([][]byte, error)
	BatchDelete(keys [][]byte) ([][]byte, error)
}

type tikvADS struct {
	client      *tikv.RawKVClient
}

func (tds *tikvADS) Put(key, value []byte) error {
	return nil
}

func (tds *tikvADS) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (tds *tikvADS) Delete(key []byte) ([]byte, error) {
	return nil, nil
}

func (tds *tikvADS) BatchPut(keys, values [][]byte) error {
	return nil
}

func (tds *tikvADS) BatchGet(keys [][]byte) ([][]byte, error) {
	return nil, nil
}

func (tds *tikvADS) BatchDelete(keys [][]byte) ([][]byte, error) {
	return nil, nil
}