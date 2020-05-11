// Package cache defined the interface of cache
// and implemented and LRU Cache
package cache

// ICache is the interface of a Cache with operations like
// CRUD, Size, Reset, etc.
type ICache interface {
	Add(key, val interface{})
	Remove(key interface{}) interface{}
	Get(key interface{}) interface{}
	Size() uint
	Reset()
	ResetWithFunc(closeFunc CloseFunc) bool
}

// CloseFunc is a func
type CloseFunc func(key, val interface{}) bool
