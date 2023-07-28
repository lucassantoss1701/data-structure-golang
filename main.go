package main

import (
	"fmt"
	unorderedvectors "lucassantos1701/data-structure-golang/unordered-vectors"
)

func main() {
	vector := unorderedvectors.NewUnorderedVector(3)
	vector.Insert(2)
	vector.Insert(3)
	vector.Insert(4)

	vector.Print()

	vector.Delete(4)

	fmt.Println("-------------")
	vector.Print()
	fmt.Println("-------------")

	fmt.Println(vector.Find(2))

}
