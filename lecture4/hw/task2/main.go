package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
	return c.value
}

func main() {
	var counters sync.Map
	wg := &sync.WaitGroup{}
	results := make(chan string, 1)
	mu := sync.Mutex{}

	for _, key := range []string{"key1", "key2", "key3"} {
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func(k string) {
				defer wg.Done()
				mu.Lock()
				counter, _ := counters.LoadOrStore(k, &Counter{})
				newValue := counter.(*Counter).Increment()
				results <- fmt.Sprintf("Key: %s, New Value: %d", k, newValue)
				mu.Unlock()
			}(key)

		}
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println(result)
	}
}
