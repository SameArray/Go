package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 4}
	equal(a, b)
}

func equal(a, b []int) bool {

	if len(a) != len(b) {
		return false
	}

	for i, num := range a {
		if num != b[i] {
			fmt.Println("false")
			return false
		}
	}
	fmt.Println("true")
	return true
}

// for i:=0; i < len(a); i++{
// 	fmt.Println(i, a[i])
// }
