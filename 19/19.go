package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println("Enter line: ")
	var line string
	fmt.Scan(&line)

	rv := []rune(line)                         // we make rune slice out of our string because we cannot change values inside string
	slices.Reverse(rv)                         // reverse slice
	fmt.Println("Slices reverse:", string(rv)) // make string out of slice again

	rev := ""
	for _, c := range line {
		rev = string(c) + rev // string concatenation
	}
	fmt.Println("Concatenating to left:", rev)
}
