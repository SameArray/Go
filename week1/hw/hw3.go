package main

import (
	"reflect"
	"sort"
)

func compareSlices(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	// Сортируем
	sort.Ints(slice1)
	sort.Ints(slice2)

	// Сравниваем
	return reflect.DeepEqual(slice1, slice2)
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3, 2, 1}

	result := compareSlices(slice1, slice2)
	if result {
		println("True")
	} else {
		println("False")
	}
}
