package lru

type Object interface {
	Key() string
}

type Node struct {
	Prev *Node
	Next *Node
	data Object
}

func NewNode(obj Object) *Node{
	return &Node{
		data:obj,
	}
}