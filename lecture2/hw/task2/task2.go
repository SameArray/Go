package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	p1, p2 := list1, list2

	// While neither of the pointers is nil
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

	// If one of the lists becomes empty, append the remaining list
	if p1 != nil {
		current.Next = p1
	} else {
		current.Next = p2
	}

	return dummy.Next // Return the next of dummy since dummy itself is an empty placeholder.
}

func main() {
	// Create list1: 1 -> 2 -> 4
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}

	// Create list2: 1 -> 3 -> 4
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}

	// Merge the two lists
	mergedList := mergeTwoLists(list1, list2)

	// Print the merged list
	printList(mergedList)
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Print(node.Val, " -> ")
		node = node.Next
	}
	fmt.Println("nil")
}
