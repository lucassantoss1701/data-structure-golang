package linkedlist

func (l *LinkedList) Reverse() {
	current := l.Head
	var previous *Node
	l.Tail = l.Head

	for current != nil {
		next := current.Next
		current.Next = previous
		previous = current
		current = next
		if current == nil {
			l.Head = previous
		}
	}
}
