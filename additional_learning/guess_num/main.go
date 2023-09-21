package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var number = 42

	for {
		var n = rand.Intn(100) + 1
		if n < number {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%v слишком маленькое число.\n", n)
		} else if n > number {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%v слишком большое число.\n", n)
		} else {
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("Угадал! %v\n", n)
			break
		}
	}
}
