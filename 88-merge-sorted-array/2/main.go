package main

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	left, right := 0, 0

	res := []int{}

	for left < m && right < n {
		if nums1[left] <= nums2[right] {
			res = append(res, nums1[left])
			left++
		} else {
			res = append(res, nums2[right])
			right++
		}
	}

	if left < m {
		res = append(res, nums1[left:]...)
	}

	if right < n {
		res = append(res, nums2[right:]...)
	}

	copy(nums1, res)
}
