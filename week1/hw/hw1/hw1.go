package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numIndices := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if index, found := numIndices[complement]; found {
			return []int{index, i}
		}
		numIndices[num] = i
	}

	return nil // No solution found
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
