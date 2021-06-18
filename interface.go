package gorocksdb

type DBClient interface {
	Get(opts *ReadOptions, key []byte) (*Slice, error)
	GetCF(opts *ReadOptions, cf *ColumnFamilyHandle, key []byte) (*Slice, error)
	Put(opts *WriteOptions, key, value []byte) error
	PutCF(opts *WriteOptions, cf *ColumnFamilyHandle, key, value []byte) error
	Delete(opts *WriteOptions, key []byte) error
	DeleteCF(opts *WriteOptions, cf *ColumnFamilyHandle, key []byte) error
	NewIterator(opts *ReadOptions) *Iterator
	NewIteratorCF(opts *ReadOptions, cf *ColumnFamilyHandle) *Iterator
}
