package main

import (
	"fmt"
)

type BubbleSort struct{}

func (bs *BubbleSort) Sort(arr []int) {
	n := len(arr)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swapped = true
			}
		}
		n--
	}
}

func main() {
	arr := []int{55, 25, 10, 7, 100}
	bs := &BubbleSort{}
	bs.Sort(arr)
	fmt.Println("Отсортированный слайс:", arr)
}
