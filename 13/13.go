package main

import "fmt"

func main() {
	a, b := 2, 4
	fmt.Println("Using tuple:")
	fmt.Println(a, b)
	fmt.Println(usingTuple(a, b))
	fmt.Println("Using return statement:")
	fmt.Println(usingReturn(a, b))
	fmt.Println("Using Math ops:")
	fmt.Println(usingMath(a, b))
	fmt.Println("Using XORs:")
	fmt.Println(usingXOR(a, b))
}

func usingTuple(a, b int) (int, int) {
	a, b = b, a
	return a, b
}

func usingReturn(a, b int) (int, int) {
	return b, a
}

func usingMath(a, b int) (int, int) { // for 2 and 4
	a = a + b // 6
	b = a - b // 2
	a = a - b // 4
	return a, b
}

func usingXOR(a, b int) (int, int) { // for 10 and 100 in binary
	a = a ^ b // 110
	b = a ^ b // 10
	a = a ^ b // 100
	return a, b
}
