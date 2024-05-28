package main

import (
	"fmt"
	"reflect"
)

func getType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	case string:
		fmt.Println("string")
	case chan interface{}:
		fmt.Println("chan")
	}
}
func withReflect(x interface{}) {
	fmt.Println(reflect.TypeOf(x))
}

func main() {
	x := 123
	getType(x)
	v := make(chan int)
	withReflect(v)
}
