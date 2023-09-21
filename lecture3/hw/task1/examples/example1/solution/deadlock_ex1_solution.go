package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		ch1 <- 100 // Сначала отправляем значение в ch1
		value := <-ch2
		fmt.Println(value)
		wg.Done()
	}()

	go func() {
		value := <-ch1
		ch2 <- value // Затем отправляем значение в ch2
		wg.Done()
	}()

	wg.Wait() // Ожидание завершения обеих горутин
}
