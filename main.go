package main

import (
	"fmt"
	orderedvectors "lucassantos1701/data-structure-golang/ordered-vectors"
)

func main() {
	vector := orderedvectors.NewUnorderedVector(10)
	vector.Insert(3)
	vector.Insert(7)
	vector.Insert(5)
	vector.Insert(3)

	vector.Print()

	fmt.Println("--------")
	fmt.Println(vector.Find(7))
	fmt.Println("--------")

	vector.Delete(3)

	vector.Print()
	fmt.Println("--------")

	fmt.Println(vector.Find(7))

}
