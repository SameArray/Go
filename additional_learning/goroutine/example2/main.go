package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{4, 5, 6} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{7, 8, 9} {
			c <- num
		}
		close(c)
	}()

	for num := range joinChannels(a, b, c) {
		fmt.Println(num)
	}
}

func joinChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		//join channels
		wg := &sync.WaitGroup{}

		wg.Add(len(chs))

		for _, channel := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for num := range ch {
					mergedCh <- num
				}
			}(channel, wg)
		}

		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}
