package main

import (
	"fmt"
)

func main() {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(search(nums, 0))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) >> 1
		fmt.Printf("mid: %d - left: %d - right: %d\n", mid, left, right)
		if nums[mid] == target {
			return mid
		}

		if nums[right] >= nums[mid] {
			// buscamos na segunda porção do array - dentro do intervalo corretamente
			if target >= nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				// o numero a extrema direita é maior que o número do meio (nums[right] >= nums[mid])
				// porém o target ou é maior que o número a extrema direita ou é menor que o número do meio...
				// nesse caso buscamos na primeira porção
				right = mid - 1
			}
		} else {
			// intervalo correto, primeira porção do array
			if target <= nums[mid] && target >= nums[left] {
				right = mid - 1
			} else {
				// o target ou é maior que o elemento do meio ou é menor que o elemento na extrema esquerda
				// nesse caso buscamos na segunda porção
				left = mid + 1
			}
		}

	}

	return -1
}
