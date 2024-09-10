package tree

import "fmt"

type Node struct {
	Left  *Node
	Right *Node
	Value int
}

func NewNode(value int) *Node {
	return &Node{
		Left:  nil,
		Right: nil,
		Value: value,
	}
}

func (n *Node) Print() {
	fmt.Println(n.Value)
}
