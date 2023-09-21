package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		value := <-ch2
		value = value + 25
		fmt.Println("Было добавлено 25")
		ch3 <- value
		wg.Done()
	}()

	go func() {
		value := <-ch1
		value = value + 50
		fmt.Println("Было добавлено 50")
		ch2 <- value
		wg.Done()

	}()

	go func() {
		value := <-ch3
		value = value + 10
		fmt.Println("Было добавлено 10")
		ch1 <- value
		wg.Done()
	}()

	ch1 <- 100

	go func() {
		value := <-ch1
		fmt.Println("Полученная число после обработки горутинами:", value)
		wg.Done()
	}()

	wg.Wait() // Ожидание завершения всех горутин

}
