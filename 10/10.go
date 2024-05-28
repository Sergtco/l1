package main

import (
	"fmt"
	"math"
	"sort"
)

func groupOf(a float32) int {
	return int(math.Floor(float64(a)/10) * 10)
}

func groupTemp(temps []float32) map[int][]float32 {
	res := map[int][]float32{}
	for _, num := range temps {
		group := groupOf(num)
		res[group] = append(res[group], num)
	}
	return res
}
func main() { // have fixed edge case: -9 and 9 shouldn't be in one group, because diff is 18.
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -6, 6, 5, -2, -9, 9}
	group := groupTemp(temp)

	keys := make([]int, 0, len(group))
	for k := range group {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		fmt.Printf("%d: %v\n", k, group[k])
	}
}
