package main

import (
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	var list2 *ListNode
	mergedList := mergeTwoLists(list1, list2)
	expected := list1
	if !isSameList(mergedList, expected) {
		t.Errorf("Ожидаемый объединенный список: %v, но получено: %v", expected, mergedList)
	}

	var list3, list4 *ListNode
	mergedList = mergeTwoLists(list3, list4)
	if mergedList != nil {
		t.Errorf("Ожидаемый объединенный список будет нулевым, но получил: %v", mergedList)
	}

	list5 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	list6 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 6}}}
	mergedList = mergeTwoLists(list5, list6)
	expected = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}
	if !isSameList(mergedList, expected) {
		t.Errorf("Ожидаемый объединенный список: %v, но получено:%v", expected, mergedList)
	}
}

func isSameList(list1 *ListNode, list2 *ListNode) bool {
	for list1 != nil && list2 != nil {
		if list1.Val != list2.Val {
			return false
		}
		list1 = list1.Next
		list2 = list2.Next
	}
	return list1 == nil && list2 == nil
}
