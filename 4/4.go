package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
)

// Worker gets context by which he understands when it's time to end
func worker(ctx context.Context, id int, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d finished\n", id)
			return
		case val := <-ch:
			fmt.Printf("Worker %d predicted: %d\n", id, val)
		}
	}
}

// Same as worker but uses channel for pushing values
func Generator(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Generator finished\n")
			return
		case ch <- rand.Int():
		}
	}
}

func main() {
	nWorkers := 5
	if len(os.Args) > 1 {
		n, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		nWorkers = n
	}
	fmt.Printf("Amount of worksers is set to %d.\n", nWorkers)

	ch := make(chan int)
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt) // NotifyContext is for signal Interrupt
	defer cancel()
	wg := sync.WaitGroup{}
	for i := range nWorkers {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(ctx, id, ch)
		}(i)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		Generator(ctx, ch)
	}()
	wg.Wait()
}

// Сделал так, что по сигналу Interrupt в терминале одновременно завер шились все горутины через ctx.Done()
// За счет небуферизованного канала все данные из него прочитаются воркерами
