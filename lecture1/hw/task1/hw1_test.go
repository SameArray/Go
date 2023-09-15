package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums1 := []int{2, 11, 7, 15, 8, 1}
	target1 := 9
	expected1 := []int{2, 0}
	result1 := twoSum(nums1, target1)
	if !areEqual(result1, expected1) {
		t.Errorf("Тестовый случай 1 не пройден: ожидалось %v, но получено %v", expected1, result1)
	}

	nums2 := []int{3, 6, 9, 12}
	target2 := 25
	expected2 := []int{}
	result2 := twoSum(nums2, target2)
	if !areEqual(result2, expected2) {
		t.Errorf("Тестовый случай 2 не пройден: ожидалось %v, но получено %v", expected2, result2)
	}

	nums3 := []int{5}
	target3 := 5
	expected3 := []int{}
	result3 := twoSum(nums3, target3)
	if !areEqual(result3, expected3) {
		t.Errorf("Тестовый случай 3 не пройден: ожидалось %v, но получено %v", expected3, result3)
	}

	nums4 := []int{0, 0, 0}
	target4 := 0
	expected4 := []int{0, 1}
	result4 := twoSum(nums4, target4)
	if !areEqual(result4, expected4) {
		t.Errorf("Тестовый случай 4 не пройден: ожидалось %v, но получено %v", expected4, result4)
	}
}

func areEqual(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
