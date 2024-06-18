package main

import "fmt"

func binSearch(arr []int, val int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := (l + r) / 2  // find middle element
		if arr[mid] > val { // if middle is bigger than search value then search value is to the left
			r = mid - 1
		} else if arr[mid] < val { // same
			l = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 5, 6, 7, 9}
	fmt.Println(binSearch(arr, 5))
	fmt.Println(binSearch(arr, 4))
	fmt.Println(binSearch(arr, 6))
}
