package queue

import "fmt"

type PriorityQueue struct {
	Capacity         int
	NumberOfElements int
	Values           []int
}

func NewPriorityQueue(capacity int) *PriorityQueue {
	return &PriorityQueue{
		Capacity:         capacity,
		NumberOfElements: 0,
		Values:           make([]int, capacity),
	}
}

func (p *PriorityQueue) Empty() bool {
	return p.NumberOfElements == 0
}

func (p *PriorityQueue) Full() bool {
	return p.NumberOfElements == p.Capacity
}

func (p *PriorityQueue) Enqueue(value int) {
	if p.Full() {
		fmt.Println("queue is full")
		return
	}

	if p.NumberOfElements == 0 {
		p.Values[p.NumberOfElements] = value
		p.NumberOfElements++
	} else {
		x := p.NumberOfElements - 1
		for x >= 0 {
			if value > p.Values[x] {
				p.Values[x+1] = p.Values[x]
			} else {
				break
			}
			x--
		}

		p.Values[x+1] = value
		p.NumberOfElements++
	}
}

func (p *PriorityQueue) Dequeue() int {

	if p.Empty() {
		fmt.Println("queue is empty")
		return -1
	}

	value := p.Values[p.NumberOfElements-1]
	p.NumberOfElements--
	return value

}

func (p *PriorityQueue) GetFirst() int {
	if p.Empty() {
		fmt.Println("queue is empty")
		return -1
	}

	return p.Values[p.NumberOfElements-1]
}
