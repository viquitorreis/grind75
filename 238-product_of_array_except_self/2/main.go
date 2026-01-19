package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	// 1. left
	// fazemos em todos elementos à left do elemento atual
	res := make([]int, len(nums))
	left := 1
	for i, num := range nums {
		res[i] = left
		left = left * num
		// fmt.Printf("res[i]: %d - left: %d\n", res[i], left)
	}

	// 2. right
	// fazemos em todos elementos à right do elemento atual
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * right
		right = nums[i] * right
		// fmt.Printf("res[i]: %d - right: %d\n", res[i], right)
	}

	return res
}
