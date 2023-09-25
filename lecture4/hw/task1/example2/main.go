package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			value, ok := m.Load("key")
			if ok {
				m.Store("key", value.(int)+1)
			} else {
				m.Store("key", 1)
			}
		}()

	}
	for i := 0; i < 100; i++ {

		wg.Add(1)

		go func() {
			defer wg.Done()
			value, ok := m.Load("key2")
			if ok {
				m.Store("key2", value.(int)+1)
			} else {
				m.Store("key2", 1)
			}
		}()

	}
	wg.Wait()
	fmt.Println(m.Load("key"))
	fmt.Println(m.Load("key2"))
}
