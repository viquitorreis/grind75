package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	n := 1
	res := make([]int, len(nums))
	for i := range nums {
		res[i] += n
		n = nums[i] * n
	}

	n = 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * n
		n = nums[i] * n
	}

	return res
}
