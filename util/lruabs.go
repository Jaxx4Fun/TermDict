package util

type LRU interface {
	Add(object Object)
	PopFront() Object
	Get(key string) Object
	Size() uint
	GetContentByOrder() []Object
	ClearWithFunc(closeFunc func(object Object) bool)
}
