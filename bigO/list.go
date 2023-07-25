package bigo

// O(n)
func List1() []int {

	list := []int{}

	for i := 1; i <= 1000; i++ {
		list = append(list, i)
	}
	return list
}
