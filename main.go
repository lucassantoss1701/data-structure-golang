package main

import (
	"fmt"
	"lucassantos1701/data-structure-golang/queue"
)

func main() {
	queue := queue.NewCircularQueue(5)

	queue.Enqueue(5)
	queue.Enqueue(6)

	fmt.Println(queue.GetFirst())
	queue.Dequeue()
	fmt.Println(queue.GetFirst())
	queue.Dequeue()
	queue.Enqueue(4)
	queue.Enqueue(6)
	fmt.Println(queue.GetFirst())
	queue.Dequeue()
	fmt.Println(queue.GetFirst())
	queue.Enqueue(3)

	queue.Dequeue()
	fmt.Println(queue.GetFirst())
	queue.Dequeue()

}
