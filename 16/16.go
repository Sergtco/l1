package main

import "fmt"

func quicksort(arr []int, left, right int) {
	if left < right {
		piv := parition(arr, left, right)
		quicksort(arr, left, piv-1)
		quicksort(arr, piv+1, right)
	}
}

func parition(arr []int, left, right int) int { // choose pivot and move lower values to the left and higher to right
	pivot := arr[right]

	greater := left - 1

	for i := left; i < right; i++ {
		if arr[i] < pivot {
			greater++ // if arr[i] < pivot then greater is somewhere to the right
			arr[greater], arr[i] = arr[i], arr[greater]
		}
	}
	arr[greater+1], arr[right] = arr[right], arr[greater+1]
	return greater + 1
}

func main() {
	arr := []int{2, 3, 1, 4, 2, 4, 5, 6, 19, 2, 3}
	fmt.Println(arr)
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
