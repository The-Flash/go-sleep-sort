package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numbers := []int{3, 2, 3, 5, 6, 12, 7, 2, 3, 10}
	sorted := sleepSort(numbers)
	fmt.Println(sorted) // [2 2 3 3 5 6 7 10 12]
}

// sleepSort sorts the given slice of integers using the sleep sort algorithm.
// Each goroutine sleeps for an amount of time which is proportional
// to the corresponding array element.
func sleepSort(numbers []int) (result []int) {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for _, number := range numbers {
		wg.Add(1)
		go func(n int, wg *sync.WaitGroup) {
			defer wg.Done()
			time.Sleep(time.Duration(n) * time.Millisecond)
			// Prevent race conditions
			mutex.Lock()
			defer mutex.Unlock()
			result = append(result, n)
		}(number, &wg)
	}
	wg.Wait()
	return result
}
