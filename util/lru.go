package util

type Object interface {
	Key() string
}

type Node struct {
	index string
	data  Object
	Prev  *Node
	Next  *Node
}

func NewNode(object Object) *Node {
	return &Node{data: object}
}

type LRUCache struct {
	dList struct {
		start *Node //头部，least recently used
		end   *Node //插入尾部
	}
	limit      uint //最大容量
	size       uint //当前大小
	storageMap map[string]*Node
}

// create a LRU
func NewLRUCache(capacity uint) *LRUCache {
	lru := &LRUCache{
		limit:      capacity,
		storageMap: make(map[string]*Node, capacity),
	}
	return lru
}
// Size returns current size of lru
func (c *LRUCache) Size() uint {
	return c.size
}
func (c *LRUCache) GetStorageMap() map[string]*Node {
	return c.storageMap
}

func (c *LRUCache) Add(object Object) {
	if c.Get(object.Key()) != nil {
		return
	}

	// new node
	node := NewNode(object)

	// full?
	if c.size >= c.limit {
		c.PopFront()
	}

	c.append(node)
	c.size++

	c.storageMap[object.Key()] = node

}

func (c *LRUCache) append(node *Node) {
	node.Prev = c.dList.end
	if c.size == 0 {
		// c.dList.start == nil
		// c.dList.end==nil
		c.dList.start = node

	} else {
		c.dList.end.Next = node
		node.Prev = c.dList.end
	}
	node.Next = nil
	c.dList.end = node
}

// PopFront allows user to pop item manually
func (c *LRUCache) PopFront() Object {
	if c.size > 0 {
		//delete start node
		nodeToPop := c.dList.start
		c.dList.start = nodeToPop.Next
		if c.dList.start != nil {
			c.dList.start.Prev = nil
		}
		delete(c.storageMap, nodeToPop.data.Key())
		c.size--
		return nodeToPop.data
	}

	return nil
}
func (c *LRUCache) moveBackward(node *Node) {
	if node == nil {
		panic("node is nil!!!")
	}
	if node == c.dList.end {
		return
	}
	prevNode := node.Prev
	nextNode := node.Next

	if prevNode != nil {
		prevNode.Next = nextNode
	}

	if nextNode != nil {
		nextNode.Prev = prevNode
		if node == c.dList.start {
			c.dList.start = nextNode
		}
		c.append(node)
		return
	}
	return
}

func (c *LRUCache) Get(key string) Object {
	node, ok := c.storageMap[key]
	if !ok {
		return nil
	}
	c.moveBackward(node)

	return node.data
}

func (c *LRUCache) ClearWithFunc(closeFunc func(object Object) bool) {
	c.size = 0
	for k, node := range c.storageMap {
		delete(c.storageMap, k)
		if closeFunc != nil {
			closeFunc(node.data)
		}
	}
	c.destroyDlink()
}

func (c *LRUCache) destroyDlink() {
	c.dList.start = nil
	c.dList.end = nil
}
