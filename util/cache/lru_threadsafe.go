package cache

import "sync"

type ThreadSafeLRUWrapper struct {
	lru LRU
	m   sync.RWMutex
}

func NewThreadSafeLRU(lru LRU) *ThreadSafeLRUWrapper {
	return &ThreadSafeLRUWrapper{
		lru: lru,
		m:   sync.RWMutex{},
	}
}
func (l *ThreadSafeLRUWrapper) Add(object Object) {
	l.m.Lock()
	defer l.m.Unlock()
	l.lru.Add(object)
}

func (l *ThreadSafeLRUWrapper) PopFront() Object {
	l.m.Lock()
	defer l.m.Unlock()
	return l.lru.PopFront()
}

func (l *ThreadSafeLRUWrapper) Get(key string) Object {
	l.m.Lock()
	defer l.m.Unlock()
	return l.lru.Get(key)
}

func (l *ThreadSafeLRUWrapper) Size() uint {
	l.m.RLock()
	defer l.m.RUnlock()
	return l.lru.Size()
}

func (l *ThreadSafeLRUWrapper) GetContentByOrder() []Object {
	l.m.RLock()
	defer l.m.RUnlock()
	return l.lru.GetContentByOrder()
}

func (l *ThreadSafeLRUWrapper) ClearWithFunc(closeFunc func(object Object) bool) {
	l.m.Lock()
	defer l.m.Unlock()
	l.lru.ClearWithFunc(closeFunc)
}
