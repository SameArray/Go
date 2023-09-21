package main

import (
	"fmt"
	"sync"
)

func main() {
	numJobs := 100
	numWorker := 5
	var wg sync.WaitGroup
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for i := 1; i <= numWorker; i++ {
		wg.Add(1)
		go worker(square, jobs, results, i, &wg)
	}

	for i := 1; i <= numJobs; i += 2 {
		jobs <- i
	}

	close(jobs)
	wg.Wait()
	close(results)

	for result := range results {
		fmt.Println(result)
	}
}

func worker(f func(int) int, jobs <-chan int, results chan<- int, id int, wg *sync.WaitGroup) {
	for num := range jobs {
		results <- f(num)
	}
	defer wg.Done()
}

func square(num int) int {
	return num * num
}
