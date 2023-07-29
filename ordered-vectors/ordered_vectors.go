package orderedvectors

import "fmt"

type OrderedVector struct {
	Capacity     int
	LastPosition int
	Values       []int
}

func NewUnorderedVector(capacity int) *OrderedVector {
	return &OrderedVector{
		Capacity:     capacity,
		LastPosition: -1,
		Values:       make([]int, capacity),
	}
}

// O(n)
func (u *OrderedVector) Print() {
	if u.LastPosition == -1 {
		fmt.Println("The vector is empty")
	} else {
		for i := 0; i < u.LastPosition+1; i++ {
			fmt.Printf("%d - %d\n", i, u.Values[i])
		}
	}
}

// O(n)
func (u *OrderedVector) Insert(value int) {
	if u.LastPosition == u.Capacity-1 {
		fmt.Println("The vector is empty")
	}

	position := 0
	for i := 0; i < u.LastPosition+1; i++ {
		position = i
		if u.Values[i] >= value {
			break
		}
		if i == u.LastPosition {
			position = i + 1
		}
	}

	x := u.LastPosition
	for x >= position {
		u.Values[x+1] = u.Values[x]
		x -= 1
	}

	u.Values[position] = value
	u.LastPosition += 1
}

// O(n)
func (u *OrderedVector) Find(value int) int {
	for i := 0; i < u.LastPosition+1; i++ {
		if u.Values[i] > value {
			return -1
		}
		if u.Values[i] == value {
			return i
		}
	}

	return -1
}

// O(n)
func (u *OrderedVector) Delete(value int) int {

	position := u.Find(value)
	if position == -1 {
		return -1
	}

	for i := position; i < u.LastPosition; i++ {
		u.Values[i] = u.Values[i+1]
	}

	u.LastPosition -= 1

	return 0
}
