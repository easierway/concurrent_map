package concurrent_map

import (
	"sync"
	"unsafe"
)

// ConcurrentMap is a thread safe map collection with better performance.
// The backend map enties are separated into the different partitions.
// Threads can access the different partitions safely without lock.
type ConcurrentMap struct {
	partitions    []unsafe.Pointer
	numOfBlockets int
}

// Partitionable is the interface which should be implemented by key type.
// It is to define how to partition the entries.
type Partitionable interface {
	Value() interface{}
	PartitionID() int64
}

type innerMap struct {
	m    map[interface{}]interface{}
	lock sync.RWMutex
}

func createInnerMap() *innerMap {
	return &innerMap{
		m: make(map[interface{}]interface{}),
	}
}

func (im *innerMap) get(key Partitionable) (interface{}, bool) {
	keyVal := key.Value()
	im.lock.RLock()
	v, ok := im.m[keyVal]
	im.lock.RUnlock()
	return v, ok
}

func (im *innerMap) set(key Partitionable, v interface{}) {
	keyVal := key.Value()
	im.lock.Lock()
	im.m[keyVal] = v
	im.lock.Unlock()
}

func (im *innerMap) del(key Partitionable) {
	keyVal := key.Value()
	im.lock.Lock()
	delete(im.m, keyVal)
	im.lock.Unlock()
}

// CreateConcurrentMap is to create a ConcurrentMap with the setting number of the partitions
func CreateConcurrentMap(numOfPartitions int) *ConcurrentMap {
	var partitions []unsafe.Pointer
	for i := 0; i < numOfPartitions; i++ {
		partitions = append(partitions, unsafe.Pointer(createInnerMap()))
	}
	return &ConcurrentMap{partitions, numOfPartitions}
}

func (m *ConcurrentMap) getPartition(key Partitionable) *innerMap {
	partitionID := key.PartitionID() % int64(m.numOfBlockets)
	return (*innerMap)(m.partitions[partitionID])
}

// Get is to get the value by the key
func (m *ConcurrentMap) Get(key Partitionable) (interface{}, bool) {
	return m.getPartition(key).get(key)
}

// Set is to store the KV entry to the map
func (m *ConcurrentMap) Set(key Partitionable, v interface{}) {
	im := m.getPartition(key)
	im.set(key, v)
}

// Del is to delete the entries by the key
func (m *ConcurrentMap) Del(key Partitionable) {
	im := m.getPartition(key)
	im.del(key)
}
