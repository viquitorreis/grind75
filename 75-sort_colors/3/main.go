package main

import "fmt"

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, min, max int) {
	if min < max {
		piv := partition(nums, min, max)
		quickSort(nums, min, piv-1)
		quickSort(nums, piv+1, max)
	}
}

func partition(nums []int, min, max int) int {
	left, right := min+1, max
	piv := nums[min]

	for {
		// achando todos elementos à direita que são maior que o pivo
		for left <= max && nums[left] <= piv {
			left++
		}

		// achando todos elementos à esquerad que sao menroes que o pivo
		for right > min && nums[right] > piv {
			right--
		}

		if left >= right {
			break
		}

		// swap
		if nums[left] > nums[right] {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	// elemento na right é menor entre os 3...
	nums[min] = nums[right]
	nums[right] = piv

	return right
}
