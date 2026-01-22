package main

func main() {
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	left, right := 0, 0
	newNums := make([]int, len(nums1))
	last := 0
	for left < m && right < n {
		if nums1[left] < nums2[right] {
			newNums[last] = nums1[left]
			left++
		} else {
			newNums[last] = nums2[right]
			right++
		}
		last++
	}

	for left < m {
		newNums[last] = nums1[left]
		left++
		last++
	}
	for right < n {
		newNums[last] = nums2[right]
		right++
		last++
	}

	copy(nums1, newNums)
}
