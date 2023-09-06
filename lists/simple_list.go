package lists

import "fmt"

type Node struct {
	Value int
	Next  any
}

func NewNode(value int, next any) *Node {
	return &Node{
		Value: value,
		Next:  next,
	}
}

func (n *Node) Print() {
	fmt.Println(n.Value)
}

type SimpleList struct {
	First *Node
}

func NewSimpleList() *SimpleList {
	return &SimpleList{}
}

func (s *SimpleList) AddOnInit(value int) {
	firstNode := NewNode(value, nil)
	firstNode.Next = s.First
	s.First = firstNode
}

func (s *SimpleList) Print() {
	actual := s.First
	for actual != nil {
		actual.Print()
		actual = actual.Next.(*Node)
	}
}
