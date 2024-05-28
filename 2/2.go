package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{} // Making WaitGroup for waiting for every goroutine to execute
	for _, val := range arr {
		wg.Add(1) // Showing WaitGroup that we expect +1 goroutine to be executed
		go func(num int) {
			defer wg.Done()        // Defer is executed always when function ends
			fmt.Println(num * num) // Defer worked and wg -1
		}(val)
	}
	wg.Wait() // Waiting for every goroutine to be executed
}

// 100 will be printed first, because of one-element stack in the end of goroutine queue
// 100, 4, 16, 36, 64 - first 4 elements go to FIFO queue past the FILO stack, then last element goes to the stack
// and gets executed first.
