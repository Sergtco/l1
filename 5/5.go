package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	deadTime := 2
	if len(os.Args) > 1 {
		argTime, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic(err)
		}
		deadTime = argTime
	}
	// to time out after duration
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(deadTime))
	defer cancel()
	wg := sync.WaitGroup{} // for waiting both goroutines to end

	ch := make(chan string)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			time.Sleep(time.Second) // little delay between pings
			select {
			case <-ctx.Done(): // if time is up
				fmt.Println("No ping")
				return
			case ch <- "Beep": // sending signal
				fmt.Println("Ping")
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("No pong")
				return
			case <-ch: // receiving signal
				fmt.Println("Pong")
			}
		}
	}()
	wg.Wait() // waiting for routines to be done
}
