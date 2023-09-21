package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		//Не отправляем значение
	}()
	// Не откуда получать значение
	fmt.Println(<-ch)
}
