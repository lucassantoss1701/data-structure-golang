package stack

import "fmt"

type Stack struct {
	Capacity int
	Top      int
	Values   []int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		Capacity: capacity,
		Top:      -1,
		Values:   make([]int, capacity),
	}
}

func (s *Stack) Full() bool {
	return s.Top == s.Capacity-1
}

func (s *Stack) Empty() bool {
	return s.Top == -1
}

func (s *Stack) Push(value int) {
	if s.Full() {
		fmt.Println("the stack is full")
	} else {
		s.Top++
		s.Values[s.Top] = value
	}
}

func (s *Stack) Pop() {
	if s.Empty() {
		fmt.Println("the stack is empty")
	} else {
		s.Top--
	}
}

func (s *Stack) GetTop() int {
	if s.Top != -1 {
		return s.Values[s.Top]
	}
	return -1
}
