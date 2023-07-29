package unorderedvectors

import (
	"fmt"
)

type UnorderedVector struct {
	Capacity     int
	LastPosition int
	Values       []int
}

func NewUnorderedVector(capacity int) *UnorderedVector {
	return &UnorderedVector{
		Capacity:     capacity,
		LastPosition: -1,
		Values:       make([]int, capacity),
	}
}

// O(n)
func (u *UnorderedVector) Print() {
	if u.LastPosition == -1 {
		fmt.Println("The vector is empty")
	} else {
		for i := 0; i < u.LastPosition+1; i++ {
			fmt.Printf("%d - %d\n", i, u.Values[i])
		}
	}
}

// O(1) - O(2)
func (u *UnorderedVector) Insert(value int) {
	if u.LastPosition == u.Capacity-1 {
		fmt.Println("maximum capacity reached")
	} else {
		u.LastPosition += 1
		u.Values[u.LastPosition] = value
	}
}

// O(log n)
func (u *UnorderedVector) FindUsingBinarySearch(value int) int {
	inferiorLimit := 0
	upperLimit := u.LastPosition
	for {
		actualPosition := int((inferiorLimit + upperLimit) / 2)

		if u.Values[actualPosition] == value {
			return actualPosition
		} else if inferiorLimit > upperLimit {
			return -1
		} else {
			if u.Values[actualPosition] < value {
				inferiorLimit = actualPosition + 1
			} else {
				upperLimit = inferiorLimit - 1
			}
		}
	}
}

// O(n)
func (u *UnorderedVector) Delete(value int) int {
	position := u.FindUsingBinarySearch(value)
	if position == -1 {
		return -1
	}

	for i := position; i < u.LastPosition; i++ {
		u.Values[i] = u.Values[i+1]
	}

	u.LastPosition -= 1

	return 0
}
