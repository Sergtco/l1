package main

import "fmt"

type Set map[string]bool

func New(data []string) Set {
	out := Set{}
	for _, val := range data {
		out[val] = true
	}
	return out
}

func (s Set) getSelfSubset() []Set {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	out := []Set{}
	// The solution is use bitmask, because 000 -> 111 is all combinations
	for i := 1; i < (1<<len(keys))-1; i++ { // Number of elements in subset
		subset := Set{}
		for j := 0; j <= len(keys); j++ { // iterate over each bit 1 is presence of value
			if i>>j&1 == 1 {
				subset[keys[j]] = true
			}
		}
		out = append(out, subset)
	}
	return out
}

func main() {
	seq := New([]string{"cat", "cat", "dog", "cat", "tree"})
	for _, val := range seq.getSelfSubset() {
		for k := range val {
			fmt.Printf("%s ", k)
		}
		fmt.Println()
	}
}
