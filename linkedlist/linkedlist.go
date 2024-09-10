package linkedlist

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Tail *Node
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		Head: nil,
		Tail: nil,
	}
}

func (l *LinkedList) Insert(value int) {
	newNode := &Node{
		Value: value,
		Next:  nil,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode

	}
}

func (l *LinkedList) Print() {
	actual := l.Head
	for actual != nil {
		print(actual.Value)
		actual = actual.Next
	}

	println()
}
