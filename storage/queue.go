package storage

type Queue interface {
	Put([][]byte) error
	Get(index uint64, num int32) ([][]byte, error)
}

type queue struct {
	dfile DistributedFile
}

func NewQueue(file DistributedFile) Queue {
	return &queue{dfile: file}
}

func (q *queue) Put([][]byte) error {
	return nil
}

func (q *queue) Get(index uint64, num int32) ([][]byte, error) {
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
