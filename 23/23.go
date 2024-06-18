package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	i := 3
	fmt.Println("nums: ", nums)
	fmt.Println("i: ", i)

	nums = append(nums[:i], nums[i+1:]...) // append two parts of slice without ith index
	fmt.Println("Append: ", nums)

	nums = []int{1, 2, 3, 4, 5}
	nums[i], nums[len(nums)-1] = nums[len(nums)-1], nums[i] // swap ith and last elemnts
	nums = nums[:len(nums)-1]                               // make slice without last one
	fmt.Println("Swap with last element (Not preserving order): ", nums)
}
