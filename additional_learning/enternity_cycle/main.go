package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var count = 10

	for count > 0 {
		fmt.Println(count)
		time.Sleep(1 * time.Second)
		if rand.Intn(100) == 0 {
			break
		}
		count--

	}
	if count == 0 {
		fmt.Println("Запуск!")
	} else {
		fmt.Println("Запуск отменяется.")
	}
}
