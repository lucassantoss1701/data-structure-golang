package main

import (
	"fmt"
	"lucassantos1701/data-structure-golang/queue"
)

func main() {
	queue := queue.NewPriorityQueue(5)

	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(2)
	queue.Enqueue(9)
	queue.Enqueue(1)

	fmt.Println(queue)

	queue.Dequeue()
	queue.Dequeue()
	fmt.Println(queue)

}
