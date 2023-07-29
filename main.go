package main

import (
	"fmt"
	"lucassantos1701/data-structure-golang/stack"
)

func main() {
	stack := stack.NewStack(10)

	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	stack.Pop()
	stack.Pop()

	fmt.Println(stack.GetTop())
}
