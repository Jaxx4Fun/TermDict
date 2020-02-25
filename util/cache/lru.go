// LRU缓存
// 暂时还没有用
// 减少访问次数，减少调用API的次数，当然现在只是爬网页的；
// 但是毕竟拿人手短，so...

package cache

type LRU interface {
	Cache
	GetContentByOrder() []Object
}

type LRUCache struct {
	// TODO: replace my defined doubly linked list with container/list
	dList struct {
		start *Node
		end   *Node
	}
	cap        uint //capacity of cache
	len        uint //current len of cache
	storageMap map[string]*Node
}

//NewLRUCache create a cache and return the pointer of it
func NewLRUCache(capacity uint) *LRUCache {
	lru := &LRUCache{
		cap:        capacity,
		storageMap: make(map[string]*Node, capacity),
	}
	node := NewNode(nil)
	lru.dList.start = node
	lru.dList.end = node

	return lru
}

func (c *LRUCache) Add(object Object) {
	node := NewNode(object)
	if c.len >= c.cap {
		c.PopFront() // ++ --
	} else {
		c.len++
	}
	c.append(node)

	c.storageMap[node.data.Key()] = node
}

func (c *LRUCache) PopFront() Object {
	if c.len > 0 {
		nodeToPop := c.dList.start.Next
		c.dList.start.Next = nodeToPop.Next

		if nodeToPop.Next != nil {
			nodeToPop.Next.Prev = c.dList.start
		}
		obj := nodeToPop.data
		delete(c.storageMap, obj.Key())
		return obj
	} else {
		return nil
	}
}

func (c *LRUCache) append(node *Node) {
	node.Next = c.dList.end.Next
	c.dList.end.Next = node
	node.Prev = c.dList.end
	c.dList.end = node
}

func (c *LRUCache) Get(key string) Object {
	if node, ok := c.storageMap[key]; !ok {
		return nil
	} else {
		if node != c.dList.end {
			prevNode := node.Prev
			nextNode := node.Next

			prevNode.Next = nextNode
			nextNode.Prev = prevNode

			c.append(node)
		}
		return node.data
	}

}

func (c *LRUCache) Size() uint {
	return c.len
}

func (c *LRUCache) GetContentByOrder() []Object {
	objSlice := make([]Object, 0)
	for node := c.dList.start.Next; node != nil; node = node.Next {
		objSlice = append(objSlice, node.data)
	}
	return objSlice
}

func (c *LRUCache) ClearWithFunc(closeFunc func(object Object) bool) {
	c.len = 0
	for k, node := range c.storageMap {
		delete(c.storageMap, k)
		if closeFunc != nil {
			closeFunc(node.data)
		}
	}
	c.dList.start.Next = nil
	c.dList.end = c.dList.start
}
