package main

func main() {
	nums := []int{1, 2, 3, 4}
	println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {
	appears := make(map[int]int, len(nums))
	for slow := 0; slow < len(nums); slow++ {
		_, exists := appears[nums[slow]]
		if exists {
			return true
		}
		appears[nums[slow]]++
	}

	return false
}
