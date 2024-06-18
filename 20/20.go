package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	fmt.Println("Please, enter a line: ")

	in := bufio.NewReader(os.Stdin) // bufferized reader
	var line string
	line, err := in.ReadString('\n') // reading whole string till end of line
	if err != nil {
		panic(err)
	}

	line = line[:len(line)-1]             // line includes \n, we need to remove it
	words := strings.Split(line, " ")     // split words by space
	slices.Reverse(words)                 // revresing words like in previous exercise
	fmt.Println(strings.Join(words, " ")) // join words back to one line
}
