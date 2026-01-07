package main

func main() {
	nums := []int{1, 2, 3, 1}
	println(containsDuplicate(nums))
	nums1 := []int{1, 2, 3, 4}
	println(containsDuplicate(nums1))
}

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, 0)
	for _, num := range nums {
		saw, _ := seen[num]
		if saw {
			return true
		}
		seen[num] = true
	}

	return false
}
