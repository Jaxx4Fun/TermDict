package util

type LRUCache2 struct {
	dList struct {
		start *Node
		end   *Node
	}
	limit      uint //最大容量
	size       uint //当前大小
	storageMap map[string]*Node
}

func NewLRUCache2(capacity uint) *LRUCache2 {
	lru := &LRUCache2{
		limit:      capacity,
		storageMap: make(map[string]*Node, capacity),
	}
	node := NewNode(nil)
	lru.dList.start = node
	lru.dList.end = node

	return lru
}

func (c *LRUCache2) Add(object Object) {
	node := NewNode(object)
	if c.size >= c.limit {
		c.PopFront() // ++ --
	} else {
		c.size++
	}
	c.append(node)
}

func (c *LRUCache2) PopFront() Object {
	if c.size > 0 {
		nodeToPop := c.dList.start.Next
		c.dList.start.Next = nodeToPop.Next

		if nodeToPop.Next != nil {
			nodeToPop.Next.Prev = c.dList.start
		}
		return nodeToPop.data
	} else {
		return nil
	}
}

func (c *LRUCache2) append(node *Node) {
	c.dList.end.Next = node
	node.Prev = c.dList.end
	c.dList.end = node
}

func (c *LRUCache2) Get(key string) Object {
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

func (c *LRUCache2) Size() uint {
	return c.size
}

func (c *LRUCache2) GetContentByOrder()  []Object  {
	objSlice := make([]Object,c.size)
	for node := c.dList.start.Next; node != nil; node = node.Next {
		objSlice = append(objSlice, node.data)
	}
	return objSlice
}

func (c *LRUCache2) ClearWithFunc(closeFunc func(object Object)) {
	c.size = 0
	for k, node := range c.storageMap {
		delete(c.storageMap, k)
		if closeFunc != nil {
			closeFunc(node.data)
		}
	}
}
