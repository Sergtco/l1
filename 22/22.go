package main

import (
	"fmt"
	"math/big"
)

// Value of numeric values is more 2^20?
// It can be 2^100 so we need BigInt
func main() {
	fmt.Println("Enter a, b:")
	a, b := new(big.Int), new(big.Int)
	fmt.Scan(a, b)
	var c big.Int
	fmt.Println("a + b =", c.Add(b, a))
	fmt.Println("a - b =", c.Sub(b, a))
	fmt.Println("a * b =", c.Mul(b, a))
	fmt.Println("a / b =", c.Div(b, a))
}
