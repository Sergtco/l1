package main

import "fmt"

func setBit(num int64, val int, i int) int64 {
	if val == 0 {
		return num & ^(1 << i) // mask of 11111011111
	}
	return num | (1 << i) // mask of 10000000
}

func main() {
	fmt.Println("Enter number,  bit value, i'th index") // counting bits from least significatnt one.
	var number int64
	var ind, bit int
	fmt.Scan(&number, &bit, &ind)
	out := setBit(number, bit, ind)
	fmt.Println(out)
}
