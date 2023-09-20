package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	previous := &ListNode{}
	current := previous
	p1, p2 := list1, list2

	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			current.Next = p1
			p1 = p1.Next
		} else {
			current.Next = p2
			p2 = p2.Next
		}
		current = current.Next
	}

	if p1 != nil {
		current.Next = p1
	} else {
		current.Next = p2
	}

	return previous.Next
}

func main() {

	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	mergedList := mergeTwoLists(list1, list2)
	printList(mergedList)
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " ")
		node = node.Next
	}
	fmt.Println("nil")
}
