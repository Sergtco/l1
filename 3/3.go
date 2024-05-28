package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("With buffered channel:")
	first()
	fmt.Println("With atomic operation:")
	second()
}

// channel is safe for concurrency, because it locks, while some goroutine is writing to it
func first() {
	arr := []int{2, 4, 6, 8, 10}
	// make channel for piping data back to main goroutine
	ch := make(chan int, len(arr))
	defer close(ch)
	for _, val := range arr {
		go func(num int, ch chan int) {
			ch <- num * num // passing squares to channel
		}(val, ch)
	}
	// Don't need WaitGroup, because we have channel wich waits for len(arr) elements.
	res := 0
	for i := 0; i < len(arr); i++ {
		res += <-ch
	}
	fmt.Println(res)
}

// atomic operations are safe for concurrency because these operations are made in one assembly operation, that makes race condition almost impossible
func second() {
	arr := []int{2, 4, 6, 8, 10}
	var res int32
	wg := sync.WaitGroup{} // using WaitGroup so all goroutines could be executed
	for _, val := range arr {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			sq := int32(num * num)
			atomic.AddInt32(&res, sq) // adding square to result
		}(val)
	}
	wg.Wait()
	fmt.Println(res)
}
