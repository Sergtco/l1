package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Atomic struct {
	val atomic.Int32
}

func (a *Atomic) Add(val int32) {
	a.val.Add(val)
}

func (a *Atomic) Show() int {
	return int(a.val.Load())
}

type Synced struct {
	mu  sync.Mutex
	val int
}

func (s *Synced) Add(val int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.val += val
}

func (s *Synced) Show() int {
	return s.val
}

//Atomic makes addition in one operation so it'll preven race condition
//Mutex locks struct for only one routine

func main() {
	a, s := Atomic{}, Synced{}
	wg := sync.WaitGroup{}
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.Add(1)
			s.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println("Synced:", s.Show(), "Atomic:", a.Show())
}
