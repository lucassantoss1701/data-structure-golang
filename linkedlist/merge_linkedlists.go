package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) Print() {
	fmt.Println(l.Val)
	if l.Next != nil {
		l.Next.Print()
	}
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	firstList, secondList := l1, l2

	resulta := &ListNode{}
	result := resulta

	for firstList != nil && secondList != nil {
		if firstList.Val < secondList.Val {
			result.Next = firstList
			result = result.Next
			firstList = firstList.Next
		} else {
			result.Next = secondList
			result = result.Next
			secondList = secondList.Next
		}
	}

	for firstList != nil {
		result.Next = firstList
		result = result.Next
		firstList = firstList.Next
	}

	for secondList != nil {
		result.Next = secondList
		result = result.Next
		secondList = secondList.Next
	}

	resulta = resulta.Next
	return resulta

}
