package main

import (
	"fmt"
	"sync"
)

var counter int
var wg sync.WaitGroup

func main() {
	wg.Add(200)
	for i := 0; i < 200; i++ {
		go func() {
			for i := 0; i < 10; i++ {
				counter++
			}
		}()
		wg.Done()
	}
	wg.Wait()
	fmt.Println("Expected counter value:", 2000)
	fmt.Println("Actual counter value:", counter)
}
