package main

import "fmt"

func multiplyByTwo(rx chan int, tx chan int) {
	for val, ok := <-rx; ok; val, ok = <-rx { // Wait for channel to be closed
		tx <- val * 2
	}
	close(tx)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	rx := make(chan int, len(arr)) // receiever
	tx := make(chan int, len(arr)) // tranceiver
	go func() {
		multiplyByTwo(rx, tx)
	}()
	for _, num := range arr {
		rx <- num
	}
	close(rx)

	for val, ok := <-tx; ok; val, ok = <-tx { // also waiting for channel to be closed
		fmt.Println(val)
	}
} // goroutine will be executed because main waiting for tx to be closed by this goroutine
