// Cache abstract
package cache

type Cache interface {
	Add(object Object)
	PopFront() Object
	Get(key string) Object
	Size() uint
	ClearWithFunc(closeFunc func(object Object) bool)
}
