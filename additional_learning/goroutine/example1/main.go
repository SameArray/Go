package main

import "fmt"

func main() {
	nums := make(chan int)
	squares := make(chan int)

	go func() {
		// кладем значения в канал nums
		for _, num := range []int{1, 2, 3, 4, 5, 6, 7, 11} {
			nums <- num
		}
		close(nums)
	}()

	go func() {
		defer close(squares)
		// считываем данные с канала nums, возводим в квадрат и передаем в канал squares
		for num := range nums {
			squares <- num * num
		}
	}()

	for num := range squares {
		fmt.Println(num)
	}
}
