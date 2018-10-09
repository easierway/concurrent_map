package concurrent_map

type ConcurrentMapBenchmarkAdapter struct {
	cm *ConcurrentMap
}

func (m *ConcurrentMapBenchmarkAdapter) Set(key interface{}, value interface{}) {
	m.cm.Set(StrKey(key.(string)), value)
}

func (m *ConcurrentMapBenchmarkAdapter) Get(key interface{}) (interface{}, bool) {
	return m.cm.Get(StrKey(key.(string)))
}

func (m *ConcurrentMapBenchmarkAdapter) Del(key interface{}) {
	m.cm.Del(StrKey(key.(string)))
}

func CreateConcurrentMapBenchmarkAdapter(numOfPartitions int) *ConcurrentMapBenchmarkAdapter {
	conMap := CreateConcurrentMap(numOfPartitions)
	return &ConcurrentMapBenchmarkAdapter{conMap}
}
