package linkedlist

type NodeFull struct {
	val  int
	prev *NodeFull
	next *NodeFull
}

func NewNodeFull(val int, prev, next *NodeFull) *NodeFull {
	return &NodeFull{
		val:  val,
		prev: prev,
		next: next,
	}
}

type MyLinkedList struct {
	head *NodeFull
	tail *NodeFull
}

func Constructor() MyLinkedList {
	head := NewNodeFull(-1, nil, nil)
	tail := NewNodeFull(-1, head, nil)
	head.next = tail

	return MyLinkedList{
		head: head,
		tail: tail,
	}
}

func (this *MyLinkedList) Get(index int) int {
	cur := this.head.next
	for cur != nil && index > 0 {
		cur = cur.next
		index--
	}

	if cur != nil && cur != this.tail && index == 0 {
		return cur.val
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	NodeFull := NewNodeFull(val, this.head, this.head.next)
	this.head.next.prev = NodeFull
	this.head.next = NodeFull
}

func (this *MyLinkedList) AddAtTail(val int) {
	NodeFull := NewNodeFull(val, this.tail.prev, this.tail)
	this.tail.prev.next = NodeFull
	this.tail.prev = NodeFull
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	cur := this.head.next
	for cur.next != nil && index > 0 {
		cur = cur.next
		index--
	}

	if index == 0 {
		NodeFull := NewNodeFull(val, cur.prev, cur)
		cur.prev.next = NodeFull
		cur.prev = NodeFull
	}
	return
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	cur := this.head.next

	for cur.next != nil && index > 0 {
		cur = cur.next
		index--
	}

	if index == 0 && cur != nil && cur != this.tail {
		prev, next := cur.prev, cur.next
		prev.next, next.prev = next, prev
	}
	return
}
