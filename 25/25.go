package main

import (
	"time"
)

// Blocks current goroutine while time passed less than duration
func mySleep(dur time.Duration) {
	curr := time.Now()
	for time.Now().Sub(curr) < dur {
	}
}

func main() {
	mySleep(time.Second)
}
