package main

func main() {

}

func sortColors(nums []int) {
	quicksort(nums, 0, len(nums)-1)
}

func quicksort(nums []int, min, max int) {
	if min < max {
		piv := partition(nums, min, max)
		quicksort(nums, min, piv-1)
		quicksort(nums, piv+1, max)
	}
}

func partition(nums []int, min, max int) int {
	left, right := min+1, max
	piv := nums[min]

	for {
		// left procura elementos maiores que pivot
		for left <= max && nums[left] <= piv {
			left++
		}

		// right procura elementos menores ou iguais que pivot
		for right > min && nums[right] > piv {
			right--
		}

		if left >= right {
			break
		}

		nums[left], nums[right] = nums[right], nums[left]
	}

	// swap piv
	// nums[right] agora é menor que nums[left]
	// nums[right] vira o piv -> manda piv para o centro
	nums[min] = nums[right]
	nums[right] = piv

	return right
}
