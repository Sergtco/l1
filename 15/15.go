package main

import "strings"

func createHugeString(size int) string {
	return strings.Repeat("-", size)
}

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = v[:100] // shared memory with v
	//garbage collector won't clear all other characters untill justString alive
}

func firstApproach() {
	v := createHugeString(100) // create small string in first place
}
func secondApproach() {
	v := createHugeString(1 << 10)
	tmp := make([]byte, 100)
	copy(tmp, v[:100]) // copying in temporary array
	justString := string(tmp)
}

func main() {
	someFunc()
}
