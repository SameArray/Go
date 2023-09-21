package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}

	wg.Add(3)

	go func() {
		defer wg.Done()
		time.Sleep(100 * time.Millisecond) //имитация бурной деятельности
		fmt.Println("first goroutine finished")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(210 * time.Millisecond)
		fmt.Println("second goroutine finished")
	}()

	go func() {
		defer wg.Done()
		time.Sleep(250 * time.Millisecond)
		fmt.Println("third goroutine finished")
	}()

	wg.Wait()

	fmt.Println("All goroutines are finsihed their work. ")
}
