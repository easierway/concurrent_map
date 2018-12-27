package concurrent_map

type Int64Key struct {
	value int64
}

// PartitionID is created by string's hash
func (i *Int64Key) PartitionKey() int64 {
	return i.value
}

// Value is the raw string
func (i *Int64Key) Value() interface{} {
	return i.value
}

// StrKey is to convert a string to StringKey
func I64Key(key int64) *Int64Key {
	return &Int64Key{key}
}
