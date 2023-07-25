package bigo

import "fmt"

// O(n^2)
func Quadratic(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Println(i, j)
		}
	}
}
