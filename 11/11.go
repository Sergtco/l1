package main

import "fmt"

type Set map[int]bool

func (s Set) Intersection(other Set) Set {
	inter := Set{}
	for k := range s {
		if _, ok := other[k]; ok {
			inter[k] = true
		}
	}
	return inter
}

func main() { // Intersection means values that present in A and B.
	first := Set{1: true, 2: true, 3: true, 4: true}
	second := Set{3: true, 4: true, 5: true, 6: true}
	fmt.Println(first.Intersection(second))
}
