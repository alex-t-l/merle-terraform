package client

type Keyspace struct {
	Keyspace   string
	version    int64
	Assignment Assignment
}

type Assignment struct {
	Partitions int64
	Created    int64
	Slices     []Slice
}

type Slice struct {
	Id        int64
	Start     uint64
	End       uint64
	Partition int64
}
