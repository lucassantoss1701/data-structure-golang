package bigo

import "fmt"

// 0(1) + O(5) + O(5) + O(N) + O(3)
// 0(9) + 0(2n) -> 0(n)
func Combination(n []int) {

	//O(1)
	println(n[0])

	//O(5)
	for i := 0; i < 5; i++ {
		println("teste ", i)
	}

	//O(N)
	for i := 0; i < len(n); i++ {
		println(n[i])
	}

	//O(N)
	for i := 0; i < len(n); i++ {
		println(n[i])
	}

	//O(3)
	fmt.Println("GO")
	fmt.Println("GO")
	fmt.Println("GO")

}
