package cache

import "sync"

// ThreadSafeLRUWrapper is simply a lru plus one mutex to make it thread safe
type ThreadSafeLRUWrapper struct {
	lru LRU
	m   *sync.RWMutex
}

// NewThreadSafeLRU init a lru
func NewThreadSafeLRU(lru LRU) *ThreadSafeLRUWrapper {
	return &ThreadSafeLRUWrapper{
		lru: lru,
		m:   &sync.RWMutex{},
	}
}

// Add an obj with key
func (l *ThreadSafeLRUWrapper) Add(key, obj interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	l.lru.Add(key, obj)
}

// Remove an entry by key
// return the obj it add
func (l *ThreadSafeLRUWrapper) Remove(key interface{}) interface{} {

	l.m.Lock()
	defer l.m.Unlock()
	return l.lru.Remove(key)
}

// Get an obj by key
// return nil if key is not in LRU
func (l *ThreadSafeLRUWrapper) Get(key interface{}) interface{} {
	l.m.Lock()
	defer l.m.Unlock()
	return l.lru.Get(key)
}

// Size returns the size of LRU
func (l *ThreadSafeLRUWrapper) Size() uint {
	l.m.RLock()
	defer l.m.RUnlock()
	return l.lru.Size()
}

// ResetWithFunc pops elements and calls closeFunc one by one
func (l *ThreadSafeLRUWrapper) ResetWithFunc(closeFunc CloseFunc) bool {
	l.m.Lock()
	defer l.m.Unlock()
	return l.lru.ResetWithFunc(closeFunc)
}

// Reset simply re-init the map and the Doubly list
func (l *ThreadSafeLRUWrapper) Reset() {
	l.m.Lock()
	l.lru.Reset()
	defer l.m.Unlock()
}

// PopOldest pops the oldest element in LRU
// return nil if LRU is empty
func (l *ThreadSafeLRUWrapper) PopOldest() interface{} {

	l.m.Lock()
	defer l.m.Unlock()
	return l.lru.PopOldest()
}
