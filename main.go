package main

import (
	"fmt"
	"lucassantos1701/data-structure-golang/lists"
)

func main() {
	list := lists.NewSimpleList()

	list.AddOnInit(3)

	list.Print()
	fmt.Println(list.First)

}
