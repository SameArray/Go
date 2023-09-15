package main

import (
	"fmt"
)

func CompareSlice(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	countMap := make(map[int]int)

	for _, num := range slice1 {
		countMap[num]++
	}
	for _, num := range slice2 {
		count, exists := countMap[num]
		if !exists || count == 0 {
			return false
		}
		countMap[num]--
	}

	return true
}

func main() {
	slice1 := []int{9, 5, 4, 1, 7}
	slice2 := []int{4, 5, 1, 7, 9}

	result := CompareSlice(slice1, slice2)
	if result {
		fmt.Println("Слайсы совпадают")
	} else {
		fmt.Println("Слайсы не совпадают")
	}
}
