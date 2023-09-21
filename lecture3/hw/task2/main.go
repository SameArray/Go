package main

import (
	"fmt"
	"sync"
)

func main() {
	slice1 := []int{1, 2, 3, 4}
	slice2 := []int{5, 6}
	slice3 := []int{7, 8, 9}

	allslice := make(chan []int)

	var wg sync.WaitGroup

	wg.Add(3)

	// Горутина для объединения slice1
	go func() {
		allslice <- slice1
		wg.Done()
	}()

	// Горутина для объединения slice2
	go func() {
		allslice <- slice2
		wg.Done()
	}()

	// Горутина для объединения slice3
	go func() {
		allslice <- slice3
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(allslice)
	}()

	var mergedSlice []int

	// Собираем объединенные данные из канала
	for slice := range allslice {
		mergedSlice = append(mergedSlice, slice...)
	}

	fmt.Println(mergedSlice)
}
