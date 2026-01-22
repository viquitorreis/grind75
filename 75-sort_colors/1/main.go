package main

func main() {

}

func sortColors(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, min, max int) []int {
	if min < max {
		piv := partition(nums, min, max)
		quickSort(nums, min, piv-1)
		quickSort(nums, piv+1, max)
	}

	return nums
}

func partition(nums []int, min, max int) int {
	piv := nums[min]
	left, right := min+1, max

	for {
		for left <= max && nums[left] <= piv {
			left++
		}

		for right > min && nums[right] > piv {
			right--
		}

		if left >= right {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	nums[min] = nums[right]
	nums[right] = piv

	return right
}
