package queue

import "fmt"

type CircularQueue struct {
	Capacity         int
	Init             int
	Final            int
	NumberOfElements int
	Values           []int
}

func NewCircularQueue(capacity int) *CircularQueue {
	return &CircularQueue{
		Capacity: capacity,
		Init:     0,
		Final:    -1,
		Values:   make([]int, capacity),
	}
}

func (c *CircularQueue) Empty() bool {
	return c.NumberOfElements == 0
}

func (c *CircularQueue) Full() bool {
	return c.NumberOfElements == c.Capacity
}

func (c *CircularQueue) Enqueue(value int) {
	if c.Full() {
		fmt.Println("queue is full")
		return
	}

	if c.Final == c.Capacity-1 {
		c.Final--
	}

	c.Final++
	c.Values[c.Final] = value
	c.NumberOfElements++
}

func (c *CircularQueue) Dequeue() int {
	if c.Empty() {
		fmt.Println("queue is empty")
		return -1
	}

	temp := c.Values[c.Init]

	c.Init++
	if c.Init == c.Capacity {
		c.Init = 0
	}

	c.NumberOfElements--

	return temp
}

func (c *CircularQueue) GetFirst() int {
	if c.Empty() {
		fmt.Println("queue is empty")
		return -1
	}

	return c.Values[c.Init]
}
