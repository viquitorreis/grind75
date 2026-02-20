package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	fmt.Println(nums1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	arr := make([]int, 0)

	p1, p2 := 0, 0
	for p1 < m && p2 < n {
		if nums1[p1] <= nums2[p2] {
			arr = append(arr, nums1[p1])
			p1++
		} else {
			arr = append(arr, nums2[p2])
			p2++
		}
	}

	if p1 < m {
		arr = append(arr, nums1[p1:]...)
	}

	if p2 < n {
		arr = append(arr, nums2[p2:]...)
	}

	// nums1 = arr
	copy(nums1, arr)
}
