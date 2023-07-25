package bigo

// O(n)
func Sum1(n int) int {
	total := 0
	for i := 1; i <= n; i++ {
		total += i
	}

	return total
}

// O(3)
func Sum2(n int) int {
	return (n * (n + 1)) / 2
}
