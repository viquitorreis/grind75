package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))

	left := 1
	for i := 0; i < len(nums); i++ {
		res[i] = left
		left = left * nums[i]
	}

	right := 1
	for i := len(nums); i >= 0; i-- {
		res[i] = res[i] * right
		right = right * nums[i]
	}

	return res
}
