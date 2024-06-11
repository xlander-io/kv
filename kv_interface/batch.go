package kv_interface

type Batch struct {
	data map[string][]byte
}

func NewBatch() *Batch {
	return &Batch{
		data: make(map[string][]byte),
	}
}

func (b *Batch) Put(key, value []byte) {
	b.data[string(key)] = value
}

func (b *Batch) Delete(key []byte) {
	delete(b.data, string(key))
}

func (b *Batch) Loop(callback func(key []byte, value []byte)) {
	for key_str, val_byte := range b.data {
		callback([]byte(key_str), val_byte)
	}
}

func (b *Batch) Exist(key, value []byte) ([]byte, bool) {
	result, ok := b.data[string(key)]
	if ok {
		return result, true
	} else {
		return nil, false
	}
}
