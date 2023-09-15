package main

import "fmt"

func twoSum(nums []int, target int) []int {
	s := make(map[int]int)

	for idx, num := range nums {
		if pos, ok := s[target-num]; ok {
			return []int{pos, idx}
		}
		s[num] = idx
	}
	return []int{}
}

func main() {
	nums := []int{2, 11, 7, 15, 8, 1}
	target := 9
	result := twoSum(nums, target)

	if result != nil {
		fmt.Println("Output:", result)
	} else {
		fmt.Println("No solution found.")
	}
}
