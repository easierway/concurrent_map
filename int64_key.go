package concurrent_map

type Int64Key struct {
	value int64
}

// PartitionKey is value of itself
func (i *Int64Key) PartitionKey() int64 {
	return i.value
}

// Value is the int64
func (i *Int64Key) Value() interface{} {
	return i.value
}

// I64Key is to convert a int64 to Int64Key
func I64Key(key int64) *Int64Key {
	return &Int64Key{key}
}
