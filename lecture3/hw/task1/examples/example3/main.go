package main

import (
	"sync"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		// Горутина 1 ждет значения из ch1 и пытается отправить его в ch2.
		value := <-ch1
		ch2 <- value
		wg.Done()
	}()

	go func() {
		// Горутина 2 ждет значения из ch2 и пытается отправить его в ch3.
		value := <-ch2
		ch3 <- value
		wg.Done()
	}()

	go func() {
		// Горутина 3 ждет значения из ch3 и пытается отправить его в ch1.
		value := <-ch3
		ch1 <- value
		wg.Done()
	}()

	ch1 <- 42 // Отправляем начальное значение в ch1

	wg.Wait() // Ожидание завершения всех горутин
}
