// LRU缓存
// 暂时还没有用
// 减少访问次数，减少调用API的次数，当然现在只是爬网页的；
// 但是毕竟拿人手短，so...

package cache

import (
	"container/list"
)

type LRU interface {
	ICache
	PopOldest() interface{}
	// PopNewest() interface{}
}

// entry is key,val pair structure in list
// to make remove element from hash map easily
type entry struct {
	key interface{}
	obj interface{}
}

// LRUCache is Lastest Recently Used Cache
type LRUCache struct {
	timeList *list.List
	capacity uint //capacity of cache
	hashMap  map[interface{}]*list.Element
}

//NewLRUCache create a cache and return the pointer of it
func NewLRUCache(capacity uint) *LRUCache {
	lru := &LRUCache{
		capacity: capacity,
		hashMap:  make(map[interface{}]*list.Element, capacity),
	}

	lru.timeList = list.New()

	return lru
}

// Add a element into Cache
// combile key,obj as entry, insert entry as the list.Element.Value
func (c *LRUCache) Add(key interface{}, obj interface{}) {
	if node, ok := c.hashMap[key]; ok {
		node.Value = entry{key, obj}
		c.timeList.MoveToBack(node)
	} else {
		entry := entry{key, obj}
		node = c.timeList.PushBack(entry)
		c.hashMap[key] = node

		if uint(c.timeList.Len()) > c.capacity {
			c.popOldest()
		}
	}
}

// Remove an entry by key
// return the obj it add
func (c *LRUCache) Remove(key interface{}) interface{} {
	if node, ok := c.hashMap[key]; ok {
		entry := c.timeList.Remove(node).(entry)
		delete(c.hashMap, entry.key)
		return entry.obj
	}
	return nil
}

// Get the obj by key
func (c *LRUCache) Get(key interface{}) interface{} {
	if node, ok := c.hashMap[key]; ok {
		c.timeList.MoveToBack(node)
		return node.Value.(entry).obj
	}
	return nil

}

// Size function returns the num of elements exist in LRU
func (c *LRUCache) Size() uint {
	return uint(c.timeList.Len())
}

// Reset removes all elements in LRU Cache
func (c *LRUCache) Reset() {
	c.timeList.Init()
	c.hashMap = make(map[interface{}]*list.Element, c.capacity)
}

// ResetWithFunc delete elements in LRU one by one
// from the oldest to the lastest
// and calls the closeFunc
// if one close func returns false then ResetWithFunc returns false
func (c *LRUCache) ResetWithFunc(closeFunc CloseFunc) bool {
	ans := true
	for c.timeList.Len() > 0 {
		key, obj := c.popOldest()
		ans = closeFunc(key, obj) && ans
	}
	return ans
}

func (c *LRUCache) popOldest() (key, obj interface{}) {
	// delete oldest element from both hash map and list
	toPop := c.timeList.Remove(c.timeList.Front())
	delete(c.hashMap, toPop.(entry).key)
	return toPop.(entry).key, toPop.(entry).obj
}

// PopOldest pops the oldest element in LRU
// return nil if LRU is empty
func (c *LRUCache) PopOldest() interface{} {
	if c.timeList.Len() == 0 {
		return nil
	}

	toPop := c.timeList.Remove(c.timeList.Front())
	delete(c.hashMap, toPop.(entry).key)
	return toPop.(entry).key
}

// PopNewest pops the Newest element in LRU
// return nil if LRU is empty
func (c *LRUCache) PopNewest() interface{} {
	if c.timeList.Len() == 0 {
		return nil
	}

	toPop := c.timeList.Remove(c.timeList.Back())
	delete(c.hashMap, toPop.(entry).key)
	return toPop.(entry).key
}
