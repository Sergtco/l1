package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isUnique(line string) bool {
	line = strings.ToLower(line) // all to lower
	occs := map[rune]bool{}      // set for occuresnces
	for _, c := range line {
		if occs[c] == true { // if occured return false
			return false
		}
		occs[c] = true
	}
	return true // if no falses returned return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = line[:len(line)-1]
	fmt.Println(isUnique(line))
}
