package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		value := <-ch1
		ch2 <- value
	}()

	go func() {
		value := <-ch2
		ch1 <- value
	}()

	ch1 <- 100

	fmt.Println(<-ch2)
}

// deadlock из-за того, что обе горутины ожидают данных из каналов ch1 и ch2, но никто не отправляет данные в эти каналы.
