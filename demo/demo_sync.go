package demo

import "sync"

type safeResource struct {
	resource map[string]string
	lock     sync.RWMutex
}

// Add 访问 add 方法来保证 resource 线程安全
func (s *safeResource) Add(key string, value string) {
	s.lock.Lock()
	defer s.lock.RUnlock()
	s.resource[key] = value
}

type SafeMap[K comparable, V any] struct {
	Values map[K]V
	lock   sync.RWMutex
}

// LoadOrStore 已经有 key，返回对应的值，然后 loaded = true
// 没有，则放进去，返回 loaded = false
func (s *SafeMap[K, V]) LoadOrStore(key K, newValue V) (V, bool) {
	// 先加读锁
	s.lock.RLock()
	oldVal, ok := s.Values[key]
	s.lock.RUnlock()
	if ok {
		return oldVal, true
	}
	// 再检查一遍，防止相同 key 写入覆盖 value
	s.lock.Lock()
	defer s.lock.Unlock()
	oldVal, ok = s.Values[key]
	if ok {
		return oldVal, true
	}

	s.Values[key] = newValue
	return newValue, false
}
