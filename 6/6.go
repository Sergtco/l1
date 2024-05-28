package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	go func() {
		defer wg.Done()
		byTimeout()
		fmt.Println("By timeout ended")
	}()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	go func() {
		defer wg.Done()
		byContext(ctx)
		fmt.Println("By context ended")
	}()
	ch := make(chan struct{})
	go func() {
		defer wg.Done()
		byChannel(ch)
		fmt.Println("By channel ended")
	}()
	ch <- struct{}{}
	close(ch)
	intCh := make(chan int)
	go func() {
		defer wg.Done()
		byChannelClosing(intCh)
		fmt.Println("By channel closing ended")
	}()
	intCh <- 5
	close(intCh)

	shared := true
	go func() {
		defer wg.Done()
		bySharedVariable(&shared)
		fmt.Println("By shared variable ended")
	}()
	shared = false
	wg.Wait()
}

func byTimeout() { // after 2 seconds goroutine ends
	time.Sleep(time.Second * 2)
	return
}

func byContext(ctx context.Context) { // Done can be caused by anything context implements
	select {
	case <-ctx.Done():
		return
	}
}

func byChannel(quit chan struct{}) { // interface{}{} weighs noting, pretty close to ctx.Done()
	select {
	case <-quit:
		return
	}
}

func byChannelClosing(ch chan int) { // If channel is closed, then ok will be false
	select {
	case data, ok := <-ch:
		if !ok {
			return
		} else {
			fmt.Println(data)
		}
	}
}

func bySharedVariable(quit *bool) { // if something changed variable of quit pointer, then it's time to close
	for *quit {
	}
	return
}
