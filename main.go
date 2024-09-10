package main

import "lucassantos1701/data-structure-golang/linkedlist"

func main() {

	myLinkedlist := linkedlist.NewLinkedList()
	myLinkedlist.Insert(4)
	myLinkedlist.Insert(5)
	myLinkedlist.Insert(2)
	myLinkedlist.Insert(6)
	myLinkedlist.Insert(7)
	myLinkedlist.Insert(8)

	myLinkedlist.Print()

	myLinkedlist.Reverse()

	myLinkedlist.Print()

}
