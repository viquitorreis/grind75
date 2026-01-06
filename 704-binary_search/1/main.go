package main

func main() {
	nums := []int{-1, 0, 5}
	println(search(nums, -1))
}

// return target index if exists, -1 otherwise
// must write algo with O(log n) runtime complexity - Binary Search
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		midd := (left + right) / 2

		// fmt.Printf("midd: %d - nums[midd]: %d - nums[left]: %d - nums[right]: %d - left: %d - right: %d\n", midd, nums[midd], nums[left], nums[right], left, right)

		if nums[midd] == target {
			return midd
		} else if nums[midd] < target {
			left = midd + 1
		} else {
			right = midd - 1
		}
	}

	return -1
}
