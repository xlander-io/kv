package kv_interface

type KVDB interface {
	Get(key []byte) (value []byte, err error)
	Put(key, value []byte, sync bool) error
	Delete(key []byte, sync bool) error
	WriteBatch(batch *Batch, sync bool) error
	Close() error
	NewIterator(start []byte, limit []byte) Iterator
}
